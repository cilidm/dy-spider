package controller

import (
	"github.com/cilidm/dy-spider/app/middleware"
	"github.com/cilidm/dy-spider/app/router"
)

func init() {
	r := router.New("system", "/")
	r.GET("/", middleware.CheckDefaultPage)
}

