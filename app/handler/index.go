package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func FrameIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", gin.H{})
}
