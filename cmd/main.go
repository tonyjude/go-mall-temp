package main

import (
	"go-mall-temp/conf"
	"go-mall-temp/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
