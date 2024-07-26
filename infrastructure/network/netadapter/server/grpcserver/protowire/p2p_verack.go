package protowire

import (
	"github.com/coinsec/coinsecd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *CoinsecdMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "CoinsecdMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *CoinsecdMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
