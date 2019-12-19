package main

import (
	"github.com/kakaisaname/infra"
	"github.com/kakaisaname/infra/base"
	"github.com/kakaisaname/props/ini"
	"github.com/kakaisaname/props/kvs"
	_ "goRed-ui"
)

func main() {
	//获取程序运行文件所在的路径
	file := kvs.GetCurrentFilePath("config.ini", 1)
	//加载和解析配置文件
	conf := ini.NewIniFileCompositeConfigSource(file)
	base.InitLog(conf)
	app := infra.New(conf)
	app.Start()
}
