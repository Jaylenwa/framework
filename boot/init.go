package boot

import (
	"framework/global"
	"framework/infra/config"
	"framework/infra/db/mysql"
)

// 初始化
func init() {
	global.GConfig = config.NewConfig() // 初始化全局配置
	global.GDB = mysql.NewDB()          // 初始化全局DB
}
