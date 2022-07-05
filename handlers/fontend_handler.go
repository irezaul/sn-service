package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FrontendService(c *gin.Context) {
	c.HTML(http.StatusOK, "index.gohtml", nil)
}
