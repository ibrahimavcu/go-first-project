package main

import (
	"berkantbegdilili.com/first-project/config"
	"berkantbegdilili.com/first-project/routers"
)

func main() {
	config.InitConfig()
	r := routers.InitRouters()
	r.Run(":3000")
}
