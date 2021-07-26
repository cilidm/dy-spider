package controller

import (
	"github.com/cilidm/dy-spider/app/router"
)

func init() {
	r := router.New("system", "/system")
	r.GET("/", Index)
	r.GET("index", Index)
	r.GET("main", FrameIndex)
}
