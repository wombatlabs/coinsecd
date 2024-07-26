package rpchandlers

import (
	"github.com/coinsec/coinsecd/app/appmessage"
	"github.com/coinsec/coinsecd/app/rpc/rpccontext"
	"github.com/coinsec/coinsecd/infrastructure/network/netadapter/router"
)

// HandleGetSelectedTipHash handles the respectively named RPC command
func HandleGetSelectedTipHash(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	selectedTip, err := context.Domain.Consensus().GetVirtualSelectedParent()
	if err != nil {
		return nil, err
	}

	response := appmessage.NewGetSelectedTipHashResponseMessage(selectedTip.String())

	return response, nil
}
