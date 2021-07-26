package dao

import (
	"github.com/cilidm/dy-spider/app/global"
	"github.com/cilidm/dy-spider/app/model"
	"github.com/cilidm/toolbox/gconv"
	"github.com/jinzhu/gorm"
	"strings"
)

type SpiderDao interface {
	InsertSpider(model.Spider) (id uint, err error)
	FindSpiderByUrl(url string) (ws []model.Spider, err error)
	FindSpiderBySavePath(savePath string) (s *model.Spider, err error)
	FindSpiderByAwemeId(aid string) (s *model.Spider, err error)
	GetSpiderByPage(page, limit int, filters ...interface{}) (s []model.Spider, count int, err error)
	Update(id uint, attr map[string]interface{}) (err error)
}

func NewSpiderDaoImpl() SpiderDao {
	sc := new(SpiderDaoImpl)
	return sc
}

type SpiderDaoImpl struct {
}

func (b SpiderDaoImpl) Update(id uint, attr map[string]interface{}) (err error) {
	err = global.DBConn.Model(&model.Spider{}).Where("id = ?", id).Updates(attr).Error
	return err
}

func (b SpiderDaoImpl) InsertSpider(ws model.Spider) (id uint, err error) {
	err = global.DBConn.Create(&ws).Error
	id = ws.ID
	return id, err
}

func (b SpiderDaoImpl) FindSpiderByUrl(url string) (ws []model.Spider, err error) {
	err = global.DBConn.Model(model.Spider{}).Where("url = ?", url).Find(&ws).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return ws, err
}

func (b SpiderDaoImpl) FindSpiderByAwemeId(aid string) (*model.Spider, error) {
	s := new(model.Spider)
	err := global.DBConn.Model(model.Spider{}).Where("aweme_id = ?", aid).First(&s).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return s, err
}

func (b SpiderDaoImpl) FindSpiderBySavePath(savePath string) (*model.Spider, error) {
	s := new(model.Spider)
	err := global.DBConn.Model(model.Spider{}).Where("save_path = ?", savePath).First(&s).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return s, err
}

func (b SpiderDaoImpl) GetSpiderByPage(page, limit int, filters ...interface{}) (spider []model.Spider, count int, err error) {
	pageSize := limit
	offset := (page - 1) * pageSize
	var (
		queryArr []string
		values   []interface{}
	)
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			queryArr = append(queryArr, gconv.String(filters[k]))
			values = append(values, filters[k+1])
		}
	}

	query := global.DBConn.Model(model.Spider{})
	err = query.Where(strings.Join(queryArr, " AND "), values...).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Where(strings.Join(queryArr, " AND "), values...).Order("id desc").Limit(pageSize).Offset(offset).Find(&spider).Error
	if err != nil {
		return nil, 0, err
	}
	return
}
