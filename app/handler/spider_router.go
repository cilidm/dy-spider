package controller

import (
	"github.com/cilidm/dy-spider/app/router"
)

func init() {
	r := router.New("system", "/system")
	r.GET("spider/list", SpiderList)
	r.GET("spider/json", SpiderJson)

	r.GET("spider/add", SpiderAddPage)
	r.POST("spider/add", SpiderAdd)

	r.POST("spider/down_all", SpiderDownAll)
}
