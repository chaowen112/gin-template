package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ImplementMe(c *gin.Context) {
	c.JSON(http.StatusOK, "implement me")
}
