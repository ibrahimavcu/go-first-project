package choice

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.RouterGroup) {
	r.POST("/register", register)
}
