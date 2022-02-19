package choice

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func findById(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var ch choice
	ch.getOneById(uint(idInt))

	c.JSON(http.StatusOK, ch)
}
