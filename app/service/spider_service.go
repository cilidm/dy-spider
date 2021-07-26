package service

import (
	"errors"
	"github.com/cilidm/dy-spider/app/global/api/request"
	"github.com/cilidm/dy-spider/app/model"
	"github.com/cilidm/dy-spider/app/model/dao"
	"github.com/cilidm/dy-spider/app/util/emoji"
	"github.com/cilidm/dy-spider/app/util/spider"
	"strings"
)

func SpiderAddService(f request.SpiderAddForm) error {
	urls := strings.Split(f.Ulr, "\r\n")
	if len(urls) < 1 {
		return errors.New("未检测到分享链接")
	}
	var urlsSpider []string
	for _, v := range urls {
		if v != "" {
			temp := strings.ReplaceAll(v, "在抖音，记录美好生活！", "")
			urlsSpider = append(urlsSpider, strings.TrimSpace(temp))
		}
	}
	var byYear = false
	if f.ByYear == 1 {
		byYear = true
	}
	spider.SpiderDY(byYear, urlsSpider, f.Down)
	return nil
}

func SpiderJsonService(f request.SpiderListForm) ([]model.Spider, int, error) {
	if f.Page == 0 {
		f.Page = 1
	}
	if f.Limit == 0 {
		f.Limit = 10
	}
	filters := make([]interface{}, 0)
	if f.UserName != "" {
		filters = append(filters, "user_name LIKE ?", "%"+f.UserName+"%")
	}
	if f.Url != "" {
		filters = append(filters, "url LIKE ?", "%"+f.Url+"%")
	}
	if f.Info != "" {
		filters = append(filters, "info LIKE ?", "%"+f.Info+"%")
	}
	list, count, err := dao.NewSpiderDaoImpl().GetSpiderByPage(f.Page, f.Limit, filters...)
	if err != nil {
		return nil, 0, err
	}
	var data []model.Spider
	for _, v := range list {
		temp := v
		temp.UserName = emoji.DeCode(v.UserName)
		data = append(data, temp)
	}
	return data, count, nil
}
