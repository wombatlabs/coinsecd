package connmanager

import (
	"github.com/wombatlabs/coinsecd/infrastructure/logger"
	"github.com/wombatlabs/coinsecd/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
