// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package grpcserver

import (
	"github.com/wombatlabs/coinsecd/infrastructure/logger"
	"github.com/wombatlabs/coinsecd/util/panics"
)

var log = logger.RegisterSubSystem("TXMP")
var spawn = panics.GoroutineWrapperFunc(log)
