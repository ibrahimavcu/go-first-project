package routers

import (
	"github.com/gin-gonic/gin"
	v1 "ibrahimavcu.com/go-first-project/routers/api/v1"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	v1.RegisterRouter(api)
	return r
}
