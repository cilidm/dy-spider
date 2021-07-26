package controller

import (
	"github.com/cilidm/dy-spider/app/global/api/request"
	"github.com/cilidm/dy-spider/app/global/api/response"
	"github.com/cilidm/dy-spider/app/service"
	"github.com/cilidm/dy-spider/app/util/spider"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SpiderList(c *gin.Context) {
	c.HTML(http.StatusOK, "dy_list.html", nil)
}

func SpiderJson(c *gin.Context) {
	var f request.SpiderListForm
	if err := c.ShouldBind(&f); err != nil {
		response.SuccessResp(c).SetCode(0).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	data, count, err := service.SpiderJsonService(f)
	if err != nil {
		response.SuccessResp(c).SetCode(0).SetMsg(err.Error()).SetCount(count).WriteJsonExit()
		return
	}
	response.SuccessResp(c).SetCode(0).SetCount(count).SetData(data).WriteJsonExit()
}

func SpiderAddPage(c *gin.Context) {
	c.HTML(http.StatusOK, "dy_add.html", nil)
}

func SpiderAdd(c *gin.Context) {
	var f request.SpiderAddForm
	if err := c.ShouldBind(&f); err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	err := service.SpiderAddService(f)
	if err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	response.SuccessResp(c).WriteJsonExit()
}

func SpiderDownAll(c *gin.Context){
	url := c.PostForm("url")
	go spider.SpiderDY(false,[]string{url},1)
	response.SuccessResp(c).SetMsg("请求已执行，请稍后刷新页面").WriteJsonExit()
}