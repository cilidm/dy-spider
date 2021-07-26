package main

import (
	"embed"
	"fmt"
	"github.com/cilidm/dy-spider/app/core"
	"github.com/cilidm/dy-spider/app/global"
	_ "github.com/cilidm/dy-spider/app/handler"
	"github.com/cilidm/dy-spider/app/router"
	"github.com/cilidm/dy-spider/app/util/browser"
	"github.com/cilidm/toolbox/gconv"
	"net/http"
	"time"
)

//go:embed template
var templateFs embed.FS

//go:embed static
var staticFs embed.FS

func main() {
	global.ZapLog = core.InitZap()
	core.InitConfig()
	global.DBConn = core.InitConn()
	defer global.DBConn.Close()

	r := router.InitRouter(staticFs,templateFs)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", core.Conf.App.HttpPort),
		Handler:        r,
		ReadTimeout:    time.Duration(core.Conf.App.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(core.Conf.App.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf(`	欢迎使用 base-spider
	当前版本:V1.0.1
	默认前端运行地址:http://127.0.0.1:%s
`, gconv.String(core.Conf.App.HttpPort))
	browser.Open()
	global.ZapLog.Error(s.ListenAndServe().Error())
}
