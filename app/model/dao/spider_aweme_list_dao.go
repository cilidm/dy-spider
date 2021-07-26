package dao

import (
	"github.com/cilidm/dy-spider/app/global"
	"github.com/cilidm/dy-spider/app/model"
	"github.com/cilidm/toolbox/gconv"
	"github.com/jinzhu/gorm"
	"strings"
)

type SpiderAwemeListDao interface {
	InsertAwemeList(al model.AwemeList) (id uint, err error)
	FindByAwemeId(aid string) (*model.AwemeList, error)
	GetScreenByPage(page, limit int, filters ...interface{}) (al []model.AwemeList, count int, err error)
}

func NewSpiderAwemeListDaoImpl() SpiderAwemeListDao {
	sc := new(SpiderAwemeListDaoImpl)
	return sc
}

type SpiderAwemeListDaoImpl struct {
}

func (b SpiderAwemeListDaoImpl) InsertAwemeList(al model.AwemeList) (id uint, err error) {
	err = global.DBConn.Create(&al).Error
	id = al.ID
	return id, err
}

func (b SpiderAwemeListDaoImpl) FindByAwemeId(aid string) (*model.AwemeList, error) {
	var su = &model.AwemeList{}
	err := global.DBConn.Model(model.AwemeList{}).Where("aweme_id = ?", aid).First(&su).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return su, err
}

func (b SpiderAwemeListDaoImpl) GetScreenByPage(page, limit int, filters ...interface{}) (al []model.AwemeList, count int, err error) {
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

	query := global.DBConn.Model(model.SpiderUserInfo{})
	err = query.Where(strings.Join(queryArr, " AND "), values...).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	err = query.Where(strings.Join(queryArr, " AND "), values...).Order("id desc").Limit(pageSize).Offset(offset).Find(&al).Error
	if err != nil {
		return nil, 0, err
	}
	return
}
