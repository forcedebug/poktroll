// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poktroll/supplier/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/pokt-network/poktroll/x/shared/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8c18c53656bd0d, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8c18c53656bd0d, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetSupplierRequest struct {
	OperatorAddress string `protobuf:"bytes,1,opt,name=operator_address,json=operatorAddress,proto3" json:"operator_address,omitempty"`
}

func (m *QueryGetSupplierRequest) Reset()         { *m = QueryGetSupplierRequest{} }
func (m *QueryGetSupplierRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetSupplierRequest) ProtoMessage()    {}
func (*QueryGetSupplierRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8c18c53656bd0d, []int{2}
}
func (m *QueryGetSupplierRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetSupplierRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetSupplierRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetSupplierRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetSupplierRequest.Merge(m, src)
}
func (m *QueryGetSupplierRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetSupplierRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetSupplierRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetSupplierRequest proto.InternalMessageInfo

func (m *QueryGetSupplierRequest) GetOperatorAddress() string {
	if m != nil {
		return m.OperatorAddress
	}
	return ""
}

type QueryGetSupplierResponse struct {
	Supplier types.Supplier `protobuf:"bytes,1,opt,name=supplier,proto3" json:"supplier"`
}

func (m *QueryGetSupplierResponse) Reset()         { *m = QueryGetSupplierResponse{} }
func (m *QueryGetSupplierResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetSupplierResponse) ProtoMessage()    {}
func (*QueryGetSupplierResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8c18c53656bd0d, []int{3}
}
func (m *QueryGetSupplierResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetSupplierResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetSupplierResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetSupplierResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetSupplierResponse.Merge(m, src)
}
func (m *QueryGetSupplierResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetSupplierResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetSupplierResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetSupplierResponse proto.InternalMessageInfo

func (m *QueryGetSupplierResponse) GetSupplier() types.Supplier {
	if m != nil {
		return m.Supplier
	}
	return types.Supplier{}
}

type QueryAllSuppliersRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllSuppliersRequest) Reset()         { *m = QueryAllSuppliersRequest{} }
func (m *QueryAllSuppliersRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllSuppliersRequest) ProtoMessage()    {}
func (*QueryAllSuppliersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8c18c53656bd0d, []int{4}
}
func (m *QueryAllSuppliersRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllSuppliersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllSuppliersRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllSuppliersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllSuppliersRequest.Merge(m, src)
}
func (m *QueryAllSuppliersRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllSuppliersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllSuppliersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllSuppliersRequest proto.InternalMessageInfo

func (m *QueryAllSuppliersRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllSuppliersResponse struct {
	Supplier   []types.Supplier    `protobuf:"bytes,1,rep,name=supplier,proto3" json:"supplier"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllSuppliersResponse) Reset()         { *m = QueryAllSuppliersResponse{} }
func (m *QueryAllSuppliersResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllSuppliersResponse) ProtoMessage()    {}
func (*QueryAllSuppliersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8c18c53656bd0d, []int{5}
}
func (m *QueryAllSuppliersResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllSuppliersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllSuppliersResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllSuppliersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllSuppliersResponse.Merge(m, src)
}
func (m *QueryAllSuppliersResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllSuppliersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllSuppliersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllSuppliersResponse proto.InternalMessageInfo

func (m *QueryAllSuppliersResponse) GetSupplier() []types.Supplier {
	if m != nil {
		return m.Supplier
	}
	return nil
}

func (m *QueryAllSuppliersResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "poktroll.supplier.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "poktroll.supplier.QueryParamsResponse")
	proto.RegisterType((*QueryGetSupplierRequest)(nil), "poktroll.supplier.QueryGetSupplierRequest")
	proto.RegisterType((*QueryGetSupplierResponse)(nil), "poktroll.supplier.QueryGetSupplierResponse")
	proto.RegisterType((*QueryAllSuppliersRequest)(nil), "poktroll.supplier.QueryAllSuppliersRequest")
	proto.RegisterType((*QueryAllSuppliersResponse)(nil), "poktroll.supplier.QueryAllSuppliersResponse")
}

func init() { proto.RegisterFile("poktroll/supplier/query.proto", fileDescriptor_7a8c18c53656bd0d) }

var fileDescriptor_7a8c18c53656bd0d = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x31, 0x6f, 0x13, 0x31,
	0x14, 0xc7, 0xe3, 0x16, 0xaa, 0xd6, 0x20, 0x41, 0x4d, 0x24, 0x92, 0x08, 0x0e, 0x74, 0x12, 0x21,
	0x0a, 0xd4, 0x26, 0x65, 0x2c, 0x0c, 0x4d, 0x25, 0x3a, 0xb0, 0x94, 0x64, 0x40, 0x62, 0xa0, 0x72,
	0x12, 0xeb, 0x7a, 0xea, 0xe5, 0xec, 0xda, 0x0e, 0x50, 0x21, 0x16, 0x16, 0x56, 0x24, 0x46, 0x26,
	0xb6, 0x8e, 0x0c, 0x7c, 0x88, 0x8e, 0x15, 0x2c, 0x88, 0x01, 0xa1, 0x04, 0x89, 0xaf, 0x81, 0xce,
	0xf6, 0xa5, 0x29, 0x77, 0x51, 0xc2, 0x12, 0xf9, 0xfc, 0xfe, 0xff, 0xf7, 0x7e, 0x7e, 0xef, 0x29,
	0xf0, 0xba, 0xe0, 0xfb, 0x5a, 0xf2, 0x28, 0x22, 0x6a, 0x20, 0x44, 0x14, 0x32, 0x49, 0x0e, 0x06,
	0x4c, 0x1e, 0x62, 0x21, 0xb9, 0xe6, 0x68, 0x35, 0x0d, 0xe3, 0x34, 0x5c, 0x59, 0xa5, 0xfd, 0x30,
	0xe6, 0xc4, 0xfc, 0x5a, 0x55, 0xa5, 0x18, 0xf0, 0x80, 0x9b, 0x23, 0x49, 0x4e, 0xee, 0xf6, 0x5a,
	0xc0, 0x79, 0x10, 0x31, 0x42, 0x45, 0x48, 0x68, 0x1c, 0x73, 0x4d, 0x75, 0xc8, 0x63, 0xe5, 0xa2,
	0xe5, 0x2e, 0x57, 0x7d, 0xae, 0x76, 0xad, 0xcd, 0x7e, 0xb8, 0x50, 0xdd, 0x7e, 0x91, 0x0e, 0x55,
	0xcc, 0xd2, 0x90, 0x17, 0x8d, 0x0e, 0xd3, 0xb4, 0x41, 0x04, 0x0d, 0xc2, 0xd8, 0xe4, 0x71, 0x5a,
	0x6f, 0x52, 0x9b, 0xaa, 0xba, 0x3c, 0x1c, 0xc7, 0xb3, 0xef, 0x13, 0x54, 0xd2, 0xbe, 0xca, 0xc6,
	0xf7, 0xa8, 0x64, 0xbd, 0xb1, 0xcc, 0xc6, 0xfd, 0x22, 0x44, 0x4f, 0x12, 0x82, 0x1d, 0x63, 0x6a,
	0xb1, 0x83, 0x01, 0x53, 0xda, 0x6f, 0xc3, 0x2b, 0x67, 0x6e, 0x95, 0xe0, 0xb1, 0x62, 0xe8, 0x01,
	0x5c, 0xb2, 0xc9, 0x4b, 0xe0, 0x26, 0xa8, 0x5d, 0x58, 0x2f, 0xe3, 0x4c, 0xfb, 0xb0, 0xb5, 0x34,
	0x57, 0x8e, 0x7f, 0xde, 0x28, 0x1c, 0xfd, 0xf9, 0x5c, 0x07, 0x2d, 0xe7, 0xf1, 0x9f, 0xc3, 0xab,
	0x26, 0xe9, 0x36, 0xd3, 0x6d, 0xa7, 0x76, 0xf5, 0xd0, 0x16, 0xbc, 0xcc, 0x05, 0x93, 0x54, 0x73,
	0xb9, 0x4b, 0x7b, 0x3d, 0xc9, 0x94, 0x2d, 0xb1, 0xd2, 0x2c, 0x7d, 0xfd, 0xb2, 0x56, 0x74, 0xdd,
	0xdb, 0xb4, 0x91, 0xb6, 0x96, 0x61, 0x1c, 0xb4, 0x2e, 0xa5, 0x0e, 0x77, 0xed, 0x3f, 0x85, 0xa5,
	0x6c, 0x7e, 0x47, 0xbe, 0x01, 0x97, 0x53, 0xc2, 0x1c, 0x76, 0xd3, 0x19, 0x9c, 0x9a, 0x9a, 0xe7,
	0x12, 0xf6, 0xd6, 0xd8, 0xe0, 0x77, 0x5c, 0xe2, 0xcd, 0x28, 0x4a, 0x35, 0x69, 0xa7, 0xd0, 0x23,
	0x08, 0x4f, 0x67, 0xe6, 0x52, 0x57, 0xb1, 0x03, 0x4e, 0x86, 0x86, 0xed, 0xba, 0xb9, 0xd1, 0xe1,
	0x1d, 0x1a, 0x30, 0xe7, 0x6d, 0x4d, 0x38, 0xfd, 0x4f, 0x00, 0x96, 0x73, 0x8a, 0xe4, 0xe2, 0x2f,
	0xfe, 0x17, 0x3e, 0xda, 0x3e, 0x83, 0xb8, 0x60, 0x10, 0x6f, 0xcf, 0x44, 0xb4, 0x95, 0x27, 0x19,
	0xd7, 0x7f, 0x2c, 0xc2, 0xf3, 0x86, 0x11, 0xbd, 0x03, 0x70, 0xc9, 0x0e, 0x1a, 0xdd, 0xca, 0xd9,
	0x81, 0xec, 0x46, 0x55, 0xaa, 0xb3, 0x64, 0xb6, 0x9e, 0x8f, 0xdf, 0x7e, 0xfb, 0xfd, 0x61, 0xa1,
	0x86, 0xaa, 0x24, 0xd1, 0xaf, 0xc5, 0x4c, 0xbf, 0xe4, 0x72, 0x9f, 0x4c, 0xdb, 0x72, 0x74, 0x04,
	0xe0, 0x72, 0xfa, 0x72, 0x54, 0x9f, 0x56, 0x24, 0xbb, 0x72, 0x95, 0x3b, 0x73, 0x69, 0x1d, 0xd5,
	0x96, 0xa1, 0x7a, 0x88, 0x36, 0x66, 0x51, 0x8d, 0x0f, 0xaf, 0xff, 0xdd, 0xe7, 0x37, 0xe8, 0x23,
	0x80, 0x17, 0x27, 0xa7, 0x8b, 0xa6, 0x22, 0xe4, 0x2c, 0x5a, 0xe5, 0xee, 0x7c, 0x62, 0x07, 0x7c,
	0xcf, 0x00, 0xd7, 0x51, 0x6d, 0x5e, 0xe0, 0xe6, 0xe3, 0xe3, 0xa1, 0x07, 0x4e, 0x86, 0x1e, 0xf8,
	0x35, 0xf4, 0xc0, 0xfb, 0x91, 0x57, 0x38, 0x19, 0x79, 0x85, 0xef, 0x23, 0xaf, 0xf0, 0xac, 0x11,
	0x84, 0x7a, 0x6f, 0xd0, 0xc1, 0x5d, 0xde, 0x9f, 0x92, 0xed, 0xd5, 0x69, 0x3e, 0x7d, 0x28, 0x98,
	0xea, 0x2c, 0x99, 0x3f, 0x97, 0xfb, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x20, 0x1b, 0x4b, 0x4a,
	0x7e, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a list of Supplier items.
	Supplier(ctx context.Context, in *QueryGetSupplierRequest, opts ...grpc.CallOption) (*QueryGetSupplierResponse, error)
	AllSuppliers(ctx context.Context, in *QueryAllSuppliersRequest, opts ...grpc.CallOption) (*QueryAllSuppliersResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/poktroll.supplier.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Supplier(ctx context.Context, in *QueryGetSupplierRequest, opts ...grpc.CallOption) (*QueryGetSupplierResponse, error) {
	out := new(QueryGetSupplierResponse)
	err := c.cc.Invoke(ctx, "/poktroll.supplier.Query/Supplier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllSuppliers(ctx context.Context, in *QueryAllSuppliersRequest, opts ...grpc.CallOption) (*QueryAllSuppliersResponse, error) {
	out := new(QueryAllSuppliersResponse)
	err := c.cc.Invoke(ctx, "/poktroll.supplier.Query/AllSuppliers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a list of Supplier items.
	Supplier(context.Context, *QueryGetSupplierRequest) (*QueryGetSupplierResponse, error)
	AllSuppliers(context.Context, *QueryAllSuppliersRequest) (*QueryAllSuppliersResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) Supplier(ctx context.Context, req *QueryGetSupplierRequest) (*QueryGetSupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Supplier not implemented")
}
func (*UnimplementedQueryServer) AllSuppliers(ctx context.Context, req *QueryAllSuppliersRequest) (*QueryAllSuppliersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllSuppliers not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poktroll.supplier.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Supplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetSupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Supplier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poktroll.supplier.Query/Supplier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Supplier(ctx, req.(*QueryGetSupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllSuppliers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllSuppliersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllSuppliers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poktroll.supplier.Query/AllSuppliers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllSuppliers(ctx, req.(*QueryAllSuppliersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "poktroll.supplier.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "Supplier",
			Handler:    _Query_Supplier_Handler,
		},
		{
			MethodName: "AllSuppliers",
			Handler:    _Query_AllSuppliers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "poktroll/supplier/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetSupplierRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetSupplierRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetSupplierRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OperatorAddress) > 0 {
		i -= len(m.OperatorAddress)
		copy(dAtA[i:], m.OperatorAddress)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.OperatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetSupplierResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetSupplierResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetSupplierResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Supplier.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllSuppliersRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllSuppliersRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllSuppliersRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllSuppliersResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllSuppliersResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllSuppliersResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Supplier) > 0 {
		for iNdEx := len(m.Supplier) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Supplier[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetSupplierRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetSupplierResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Supplier.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllSuppliersRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllSuppliersResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Supplier) > 0 {
		for _, e := range m.Supplier {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetSupplierRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetSupplierRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetSupplierRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetSupplierResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetSupplierResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetSupplierResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supplier", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Supplier.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllSuppliersRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllSuppliersRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllSuppliersRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllSuppliersResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllSuppliersResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllSuppliersResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supplier", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Supplier = append(m.Supplier, types.Supplier{})
			if err := m.Supplier[len(m.Supplier)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
