package standalone

import (
	"github.com/wombatlabs/coinsecd/infrastructure/logger"
	"github.com/wombatlabs/coinsecd/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
