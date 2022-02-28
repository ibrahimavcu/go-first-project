package v1

import (
	"github.com/gin-gonic/gin"
	"ibrahimavcu.com/go-first-project/routers/api/v1/choice"
)

func RegisterRouter(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		choice.RegisterRouter(v1.Group("/choices"))
	}
}
