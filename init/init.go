package globalInit

import (
	"blog/init/cache"
	"blog/init/config"
	"blog/init/db"
	"blog/model"
)

func Now() {

	// 配置初始化
	config.Init()

	// 数据库初始化
	db.Init()
	model.Migrate()

	// 缓存初始化
	cache.Init()

}