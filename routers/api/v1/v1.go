package v1

import (
	"berkantbegdilili.com/first-project/routers/api/v1/choice"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{
		choice.RegisterRouter(v1.Group("/choices"))
	}
}
