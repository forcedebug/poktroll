//go:build e2e

package e2e

import (
	"bytes"
	"fmt"
	"net"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"time"
)

// TODO_TECHDEBT(https://github.com/ignite/cli/issues/3737): We're using a combination
// of `pocketd` (legacy) and `poktrolld` (current) because of an issue of how ignite works.
var (
	// defaultRPCURL used by pocketdBin to run remote commands
	defaultRPCURL = os.Getenv("POCKET_NODE")
	// defaultRPCPort is the default RPC port that pocketd listens on
	defaultRPCPort = 26657
	// defaultRPCHost is the default RPC host that pocketd listens on
	defaultRPCHost = "127.0.0.1"
	// defaultHome is the default home directory for pocketd
	defaultHome = os.Getenv("POKTROLLD_HOME")
	// defaultPathURL used by curl commands to send relay requests
	defaultPathURL = os.Getenv("PATH_URL")
	// defaultPathHostOverride overrides the host in the URL used to send requests
	// Since the current DevNet infrastructure does not support arbitrary subdomains,
	// this is used to specify the host to connect to and the full host (with the service as a subdomain)
	// will be sent in the "Host" request header.
	defaultPathHostOverride = os.Getenv("PATH_HOST_OVERRIDE")
	// defaultDebugOutput provides verbose output on manipulations with binaries (cli command, stdout, stderr)
	defaultDebugOutput = os.Getenv("E2E_DEBUG_OUTPUT")
)

func isVerbose() bool {
	return defaultDebugOutput == "true"
}

func init() {
	if defaultRPCURL == "" {
		defaultRPCURL = fmt.Sprintf("tcp://%s:%d", defaultRPCHost, defaultRPCPort)
	}
	if defaultHome == "" {
		defaultHome = "../../localnet/poktrolld"
	}
}

// commandResult combines the stdout, stderr, and err of an operation
type commandResult struct {
	Command string // the command that was executed
	Stdout  string // standard output
	Stderr  string // standard error
	Err     error  // execution error, if any
}

// PocketClient is a single function interface for interacting with a node
type PocketClient interface {
	RunCommand(args ...string) (*commandResult, error)
	RunCommandOnHost(rpcUrl string, args ...string) (*commandResult, error)
	RunCurl(rpcUrl, service, method, path, appAddr, data string, args ...string) (*commandResult, error)
}

// Ensure that pocketdBin struct fulfills PocketClient
var _ PocketClient = (*pocketdBin)(nil)

// pocketdBin holds the reults of the last command that was run
type pocketdBin struct {
	result *commandResult // stores the result of the last command that was run
}

// RunCommand runs a command on the local machine using the pocketd binary
func (p *pocketdBin) RunCommand(args ...string) (*commandResult, error) {
	return p.runPocketCmd(args...)
}

// RunCommandOnHost runs a command on specified host with the given args.
// If rpcUrl is an empty string, the defaultRPCURL is used.
// If rpcUrl is "local", the command is run on the local machine and the `--node` flag is omitted.
func (p *pocketdBin) RunCommandOnHost(rpcUrl string, args ...string) (*commandResult, error) {
	if rpcUrl == "" {
		rpcUrl = defaultRPCURL
	}
	if rpcUrl != "local" {
		args = append(args, "--node", rpcUrl)
	}
	return p.runPocketCmd(args...)
}

// RunCommandOnHostWithRetry is the same as RunCommandOnHost but retries the
// command given the number of retries provided.
func (p *pocketdBin) RunCommandOnHostWithRetry(rpcUrl string, numRetries uint8, args ...string) (*commandResult, error) {
	if numRetries <= 0 {
		return p.RunCommandOnHost(rpcUrl, args...)
	}
	res, err := p.RunCommandOnHost(rpcUrl, args...)
	if err == nil {
		return res, nil
	}
	// TODO_HACK: Figure out a better solution for retries. A parameter? Exponential backoff? What else?
	time.Sleep(5 * time.Second)
	return p.RunCommandOnHostWithRetry(rpcUrl, numRetries-1, args...)
}

// RunCurl runs a curl command on the local machine
func (p *pocketdBin) RunCurl(rpcUrl, service, method, path, appAddr, data string, args ...string) (*commandResult, error) {
	if rpcUrl == "" {
		rpcUrl = defaultPathURL
	}
	return p.runCurlCmd(rpcUrl, service, method, path, appAddr, data, args...)
}

// RunCurlWithRetry runs a curl command on the local machine with multiple retries.
// It also accounts for an ephemeral error that may occur due to DNS resolution such as "no such host".
func (p *pocketdBin) RunCurlWithRetry(rpcUrl, service, method, path, appAddr, data string, numRetries uint8, args ...string) (*commandResult, error) {
	if service == "" {
		err := fmt.Errorf("Missing service name for curl request with url: %s", rpcUrl)
		return nil, err
	}

	// No more retries left
	if numRetries <= 0 {
		return p.RunCurl(rpcUrl, service, method, path, appAddr, data, args...)
	}
	// Run the curl command
	res, err := p.RunCurl(rpcUrl, service, method, path, appAddr, data, args...)
	if err != nil {
		return p.RunCurlWithRetry(rpcUrl, service, method, path, appAddr, data, numRetries-1, args...)
	}

	// TODO_HACK: This is a list of common flaky / ephemeral errors that can occur
	// during end-to-end tests. If any of them are hit, we retry the command.
	ephemeralEndToEndErrors := []string{
		"no such host",
		"internal error: upstream error",
	}
	for _, ephemeralError := range ephemeralEndToEndErrors {
		if strings.Contains(res.Stdout, ephemeralError) {
			if isVerbose() {
				fmt.Println("Retrying due to ephemeral error:", res.Stdout)
			}
			time.Sleep(10 * time.Millisecond)
			return p.RunCurlWithRetry(rpcUrl, service, method, path, appAddr, data, numRetries-1, args...)
		}
	}

	// Return a successful result
	return res, nil
}

// runPocketCmd is a helper to run a command using the local pocketd binary with the flags provided
func (p *pocketdBin) runPocketCmd(args ...string) (*commandResult, error) {
	base := []string{"--home", defaultHome}
	args = append(base, args...)
	commandStr := "poktrolld " + strings.Join(args, " ") // Create a string representation of the command
	cmd := exec.Command("poktrolld", args...)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err := cmd.Run()
	r := &commandResult{
		Command: commandStr, // Set the command string
		Stdout:  stdoutBuf.String(),
		Stderr:  stderrBuf.String(),
		Err:     err,
	}
	p.result = r

	if err != nil {
		// Include the command executed in the error message for context
		err = fmt.Errorf("error running command [%s]: %v, stderr: %s", commandStr, err, stderrBuf.String())
	}

	if defaultDebugOutput == "true" {
		fmt.Printf("%#v\n", r)
	}

	return r, err
}

// runCurlCmd is a helper to run a command using the local pocketd binary with the flags provided
func (p *pocketdBin) runCurlCmd(rpcBaseURL, service, method, path, appAddr, data string, args ...string) (*commandResult, error) {
	rpcUrl, err := url.Parse(rpcBaseURL)
	if err != nil {
		return nil, err
	}

	// Get the virtual host that will be sent in the "Host" request header
	virtualHost := getVirtualHostFromUrlForService(rpcUrl, service)

	// TODO_HACK: As of PR #879, the DevNet infrastructure does not support routing
	// requests to arbitrary subdomains due to TLS certificate-related complexities.
	// In such environment, defaultPathHostOverride (which contains no subdomain)
	// is used as:
	//   1. The gateway's 'host:port' to connect to
	//   2. A base to which the service is added as a subdomain then set as the "Host" request header.
	//      (i.e. Host: <service>.<defaultPathHostOverride>)
	//
	// Override the actual connection address if the environment requires it.
	if defaultPathHostOverride != "" {
		rpcUrl.Host = defaultPathHostOverride
	}

	// Ensure that if a path is provided, it starts with a "/".
	// This is required for RESTful APIs that use a path to identify resources.
	// For JSON-RPC APIs, the resource path should be empty, so empty paths are allowed.
	if len(path) > 0 && path[0] != '/' {
		path = "/" + path
	}
	rpcUrl.Path = rpcUrl.Path + path

	// Ensure that the path also ends with a "/" if it only contains the version.
	// This is required because the server responds with a 301 redirect for "/v1"
	// and curl binaries on some platforms MAY NOT support re-sending POST data
	// while following a redirect (`-L` flag).
	if strings.HasSuffix(rpcUrl.Path, "/v1") {
		rpcUrl.Path = rpcUrl.Path + "/"
	}

	base := []string{
		"-v",                                   // verbose output
		"-sS",                                  // silent with error
		"-H", `Content-Type: application/json`, // HTTP headers
		"-H", fmt.Sprintf("Host: %s", virtualHost), // Add virtual host header
		"-H", fmt.Sprintf("X-App-Address: %s", appAddr),
		rpcUrl.String(),
	}

	if method == "POST" {
		base = append(base, "--data", data)
	} else if len(data) > 0 {
		fmt.Printf("WARN: data provided but not being included in the %s request because it is not of type POST", method)
	}
	args = append(base, args...)
	commandStr := "curl " + strings.Join(args, " ") // Create a string representation of the command
	cmd := exec.Command("curl", args...)
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	err = cmd.Run()
	r := &commandResult{
		Command: commandStr, // Set the command string
		Stdout:  stdoutBuf.String(),
		Stderr:  stderrBuf.String(),
		Err:     err,
	}
	p.result = r

	if defaultDebugOutput == "true" {
		fmt.Printf("%#v\n", r)
	}

	if err != nil {
		// Include the command executed in the error message for context
		err = fmt.Errorf("error running command [%s]: %v, stderr: %s", commandStr, err, stderrBuf.String())
	}

	return r, err
}

// formatURLString returns RESTful or JSON-RPC API endpoint URL depending
// on the parameters provided.
func formatURLString(serviceAlias, rpcUrl, path string) string {
	// For JSON-RPC APIs, the path should be empty
	if len(path) == 0 {
		return fmt.Sprintf("http://%s.%s/v1", serviceAlias, rpcUrl)
	}

	// For RESTful APIs, the path should not be empty.
	// We remove the leading / to make the format string below easier to read.
	if path[0] == '/' {
		path = path[1:]
	}
	return fmt.Sprintf("http://%s.%s/v1/%s", serviceAlias, rpcUrl, path)
}

// getVirtualHostFromUrlForService returns a virtual host taking into consideration
// the URL's host and the service if it's non-empty.
// Specifically, it:
//  1. Extract's the host from the rpcURL
//  2. Prefixes the service as a subdomain to (1) the given rpcUrl host stripped of the port
//
// TODO_HACK: This is needed as of PR #879 because the DevNet infrastructure does
// not support arbitrary subdomains due to TLS certificate-related complexities.
func getVirtualHostFromUrlForService(rpcUrl *url.URL, service string) string {
	// Strip port if it exists and add service prefix
	host, _, err := net.SplitHostPort(rpcUrl.Host)
	if err != nil {
		// err is non-nil if rpcUrl.Host does not have a port.
		// Use the entire host as is
		host = rpcUrl.Host
	}
	virtualHost := fmt.Sprintf("%s.%s", service, host)

	return virtualHost
}
