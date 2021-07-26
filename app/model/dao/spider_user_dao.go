package dao

import (
	"github.com/cilidm/dy-spider/app/global"
	"github.com/cilidm/dy-spider/app/model"
	"github.com/cilidm/toolbox/gconv"
	"github.com/jinzhu/gorm"
	"strings"
)

type SpiderUserDao interface {
	InsertUser(su model.SpiderUserInfo) (id uint, err error)
	FindUserByUid(uid string) (su *model.SpiderUserInfo, err error)
	GetScreenByPage(page, limit int, filters ...interface{}) (su []model.SpiderUserInfo, count int, err error)
}

func NewSpiderUserDaoImpl() SpiderUserDao {
	sc := new(SpiderUserDaoImpl)
	return sc
}

type SpiderUserDaoImpl struct {
}

func (b SpiderUserDaoImpl) InsertUser(su model.SpiderUserInfo) (id uint, err error) {
	err = global.DBConn.Create(&su).Error
	id = su.ID
	return id, err
}

func (b SpiderUserDaoImpl) FindUserByUid(uid string) (*model.SpiderUserInfo, error) {
	var su = &model.SpiderUserInfo{}
	err := global.DBConn.Model(model.SpiderUserInfo{}).Where("uid = ?", uid).First(&su).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return su, err
}


func (b SpiderUserDaoImpl) GetScreenByPage(page, limit int, filters ...interface{}) (su []model.SpiderUserInfo, count int, err error) {
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
	err = query.Where(strings.Join(queryArr, " AND "), values...).Order("id desc").Limit(pageSize).Offset(offset).Find(&su).Error
	if err != nil {
		return nil, 0, err
	}
	return
}
