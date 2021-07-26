package controller

import (
	"encoding/json"
	"github.com/cilidm/dy-spider/app/global/api/response"
	"github.com/cilidm/dy-spider/app/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

var pearConfig = `{
	"logo": {
		"title": "Base Spider",
		"image": "/static/admin/images/logo.png"
	},
	"menu": {
		"data": "/system/menu_config",
		"accordion": true,
		"method": "GET",
		"control": false,
		"select": "10"
	},
	"tab": {
		"muiltTab": true,
		"keepState": true,
		"tabMax": 30,
		"index": {
			"id": "10",
			"href": "/system/spider/list",
			"title": "数据列表"
		}
	},
	"theme": {
		"defaultColor": "2",
		"defaultMenu": "dark-theme",
		"allowCustom": true
	},
	"colors": [{
			"id": "1",
			"color": "#2d8cf0"
		},
		{
			"id": "2",
			"color": "#5FB878"
		},
		{
			"id": "3",
			"color": "#1E9FFF"
		}, {
			"id": "4",
			"color": "#FFB800"
		}, {
			"id": "5",
			"color": "darkgray"
		}
	],
	"links": [
		{
			"icon": "layui-icon layui-icon-auz",
			"title": "cli版本开源地址",
			"href": "https://gitee.com/cilidm/go-spider"
		}
	],
	"other": {
		"keepLoad": 100,
		"autoHead": false
	},
	"header": {
		"message": "/static/admin/data/message.json"
	}
}
`

var menuConfig = `[
	{
		"id": 1,
		"title": "工作空间",
		"type": 0,
		"icon": "layui-icon layui-icon-console",
		"href": "",
		"children": [ {
			"id": 10,
			"title": "下载数据",
			"icon": "layui-icon layui-icon-console",
			"type": 1,
			"openType": "_iframe",
			"href": "/system/spider/list"
		}]
	}
]
`
func PearConfig(c *gin.Context) {
	var configData model.PearConfigForm
	err := json.Unmarshal([]byte(pearConfig), &configData)
	if err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	c.JSON(http.StatusOK, configData)
}

func MenuConfig(c *gin.Context) {
	var menuData model.MenuData
	err := json.Unmarshal([]byte(menuConfig), &menuData)
	if err != nil {
		response.ErrorResp(c).SetMsg(err.Error()).WriteJsonExit()
		return
	}
	c.JSON(http.StatusOK, menuData)
}

