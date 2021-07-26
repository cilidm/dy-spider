package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckDefaultPage(c *gin.Context) {
	c.Redirect(http.StatusFound, "/system/index")
}
