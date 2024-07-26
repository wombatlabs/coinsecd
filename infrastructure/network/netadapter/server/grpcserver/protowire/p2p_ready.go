package protowire

import (
	"github.com/coinsec/coinsecd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *CoinsecdMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "CoinsecdMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *CoinsecdMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
