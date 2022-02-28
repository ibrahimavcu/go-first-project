package main

import (
	"ibrahimavcu.com/go-first-project/config"
	"ibrahimavcu.com/go-first-project/routers"
)

func main() {
	config.InitConfig()
	r := routers.InitRouters()
	r.Run(":3000")
}
