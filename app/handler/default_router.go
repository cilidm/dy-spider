package controller

import (
	"github.com/cilidm/dy-spider/app/router"
)

func init() {
	r := router.New("system", "/system")
	r.GET("pear_config", PearConfig)
	r.GET("menu_config", MenuConfig)
}
