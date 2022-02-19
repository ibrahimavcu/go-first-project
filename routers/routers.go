package routers

import (
	v1 "berkantbegdilili.com/first-project/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	v1.RegisterRouter(api)
	return r
}
