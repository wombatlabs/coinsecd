package rpc

import (
	"github.com/wombatlabs/coinsecd/infrastructure/logger"
	"github.com/wombatlabs/coinsecd/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
