package main

import "github.com/wombatlabs/coinsecd/cmd/coinsecwallet/daemon/server"

func startDaemon(conf *startDaemonConfig) error {
	return server.Start(conf.NetParams(), conf.Listen, conf.RPCServer, conf.KeysFile, conf.Profile, conf.Timeout)
}
