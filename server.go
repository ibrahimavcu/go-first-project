package server

import (
	"os"

	"berkantbegdilili.com/routers"
)

func main() {
	r := routers.InitRouter()
	r.Run(":" + os.Getenv("SERVER_PORT"))
}
