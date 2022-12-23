package logs

import (
	"github.com/ian-kent/go-log/log"
	"github.com/ian-kent/go-log/logger"
)

var Log logger.Logger

func InitLog() {
	Log = log.Logger()
}
