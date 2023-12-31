// Code generated by hertz generator.

package main

import (
	"douyin/cmd/api/biz/rpc"
	"douyin/pkg/global"
	"douyin/pkg/initialize"
	"douyin/pkg/viper"
	"flag"
)

func Init() {
	flag.StringVar(&global.ConfigPath, "c", "./pkg/config/config.yml", "config file path")
	flag.Parse()
	viper.InitConfig()
	rpc.Init()
	initialize.Hertz()
}

func main() {
	Init()
}
