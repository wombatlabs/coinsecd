package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/coinsec/coinsecd/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.CoinsecdMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.CoinsecdMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.CoinsecdMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.CoinsecdMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.CoinsecdMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.CoinsecdMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.CoinsecdMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.CoinsecdMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.CoinsecdMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.CoinsecdMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.CoinsecdMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.CoinsecdMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.CoinsecdMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.CoinsecdMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.CoinsecdMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.CoinsecdMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.CoinsecdMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.CoinsecdMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.CoinsecdMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.CoinsecdMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.CoinsecdMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.CoinsecdMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.CoinsecdMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.CoinsecdMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.CoinsecdMessage_BanRequest{}),
	reflect.TypeOf(protowire.CoinsecdMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
