package protowire

import (
	"github.com/coinsec/coinsecd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *CoinsecdMessage_StopNotifyingUtxosChangedRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "CoinsecdMessage_StopNotifyingUtxosChangedRequest is nil")
	}
	return x.StopNotifyingUtxosChangedRequest.toAppMessage()
}

func (x *CoinsecdMessage_StopNotifyingUtxosChangedRequest) fromAppMessage(message *appmessage.StopNotifyingUTXOsChangedRequestMessage) error {
	x.StopNotifyingUtxosChangedRequest = &StopNotifyingUtxosChangedRequestMessage{
		Addresses: message.Addresses,
	}
	return nil
}

func (x *StopNotifyingUtxosChangedRequestMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "StopNotifyingUtxosChangedRequestMessage is nil")
	}
	return &appmessage.StopNotifyingUTXOsChangedRequestMessage{
		Addresses: x.Addresses,
	}, nil
}

func (x *CoinsecdMessage_StopNotifyingUtxosChangedResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "CoinsecdMessage_StopNotifyingUtxosChangedResponse is nil")
	}
	return x.StopNotifyingUtxosChangedResponse.toAppMessage()
}

func (x *CoinsecdMessage_StopNotifyingUtxosChangedResponse) fromAppMessage(message *appmessage.StopNotifyingUTXOsChangedResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.StopNotifyingUtxosChangedResponse = &StopNotifyingUtxosChangedResponseMessage{
		Error: err,
	}
	return nil
}

func (x *StopNotifyingUtxosChangedResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "StopNotifyingUtxosChangedResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}
	return &appmessage.StopNotifyingUTXOsChangedResponseMessage{
		Error: rpcErr,
	}, nil
}
