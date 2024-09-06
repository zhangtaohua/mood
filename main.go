package main

import (
	"embed"
	"mood/bootstrap"
	btsConig "mood/config"
	"mood/pkg/config"
)

//go:embed database/*
var DatabaseMigrationFS embed.FS

func init() {
	// 加载 config 目录下的配置信息
	btsConig.Initialize()
}

func main() {
	// 配置初始化，依赖命令行 --env 参数
	config.InitConfig()

	bootstrap.Bootstrap(DatabaseMigrationFS)

}
