package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/pokt-network/poktroll/x/supplier/keeper"
	"github.com/pokt-network/poktroll/x/supplier/types"
)

func SimulateMsgStakeSupplier(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgStakeSupplier{
			OperatorAddress: simAccount.Address.String(),
			// TODO_TECHDEBT: Update all stake message fields
		}

		// TODO_TECHDEBT: Handling the StakeSupplier simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "StakeSupplier simulation not implemented"), nil, nil
	}
}
