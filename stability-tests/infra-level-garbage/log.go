package main

import (
	"github.com/coinsec/coinsecd/infrastructure/logger"
	"github.com/coinsec/coinsecd/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("IFLG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
