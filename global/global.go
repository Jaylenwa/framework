package global

import (
	"framework/infra/config"
	"gorm.io/gorm"
)

var (
	GConfig *config.Config // 全局配置
	GDB     *gorm.DB       // 全局 DB
)
