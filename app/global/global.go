package global

import (
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

var (
	DBConn    *gorm.DB
	ZapLog    *zap.Logger
)
