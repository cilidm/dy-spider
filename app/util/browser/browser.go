package browser

import (
	"fmt"
	"github.com/cilidm/dy-spider/app/core"
	"github.com/cilidm/dy-spider/app/global"
	"net/http"
	"os/exec"
	"time"
)

func Open()  {
	go func() {
		for {
			time.Sleep(time.Second)
			resp, err := http.Get(fmt.Sprintf("http://localhost:%d/system/index", core.Conf.App.HttpPort))
			if err != nil {
				global.ZapLog.Error(err.Error())
				continue
			}
			resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				continue
			}

			cmd := exec.Command("cmd", fmt.Sprintf("/c start http://localhost:%d/system/index", core.Conf.App.HttpPort))
			err = cmd.Start()
			if err != nil {
				fmt.Println(fmt.Sprintf("程序已运行,前台网址http://localhost:%d/system/index", core.Conf.App.HttpPort))
			}
			break
		}
	}()
}
