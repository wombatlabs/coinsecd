package rpchandlers

import (
	"github.com/coinsec/coinsecd/app/appmessage"
	"github.com/coinsec/coinsecd/app/rpc/rpccontext"
	"github.com/coinsec/coinsecd/infrastructure/network/netadapter/router"
	"github.com/coinsec/coinsecd/version"
)

// HandleGetInfo handles the respectively named RPC command
func HandleGetInfo(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	isNearlySynced, err := context.Domain.Consensus().IsNearlySynced()
	if err != nil {
		return nil, err
	}

	response := appmessage.NewGetInfoResponseMessage(
		context.NetAdapter.ID().String(),
		uint64(context.Domain.MiningManager().TransactionCount(true, false)),
		version.Version(),
		context.Config.UTXOIndex,
		context.ProtocolManager.Context().HasPeers() && isNearlySynced,
	)

	return response, nil
}
