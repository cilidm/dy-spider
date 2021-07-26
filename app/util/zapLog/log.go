package zapLog

import (
	"github.com/cilidm/dy-spider/app/core"
	"github.com/cilidm/dy-spider/app/global"
	"go.uber.org/zap"
)

func NewLog() *ZapLog {
	z := new(ZapLog)
	if global.ZapLog == nil {
		global.ZapLog = core.InitZap()
	}
	z.log = global.ZapLog
	return z
}

type ZapLog struct {
	log *zap.Logger
}

func (z *ZapLog) Info(msg, key string, infoMsg interface{}) {
	z.log.Info(msg, zap.Any(key, infoMsg))
}

func (z *ZapLog) Error(msg, key string, errMsg interface{}) {
	z.log.Error(msg, zap.Any(key, errMsg))
}

func (z *ZapLog) Warn(msg, key string, infoMsg interface{}) {
	z.log.Warn(msg, zap.Any(key, infoMsg))
}
