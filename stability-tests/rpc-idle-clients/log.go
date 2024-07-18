package main

import (
	"github.com/wombatlabs/coinsecd/infrastructure/logger"
	"github.com/wombatlabs/coinsecd/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("RPIC")
	spawn      = panics.GoroutineWrapperFunc(log)
)
