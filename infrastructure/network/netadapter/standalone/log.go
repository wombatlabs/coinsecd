package standalone

import (
	"github.com/coinsec/coinsecd/infrastructure/logger"
	"github.com/coinsec/coinsecd/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
