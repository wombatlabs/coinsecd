package prefixmanager

import (
	"github.com/coinsec/coinsecd/infrastructure/logger"
	"github.com/coinsec/coinsecd/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
