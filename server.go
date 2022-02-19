package main

import (
	"berkantbegdilili.com/first-project/routers"
)

func main() {
	r := routers.InitRouters()
	r.Run(":3000")
}
