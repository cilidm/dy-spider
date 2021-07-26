package core

import (
	"fmt"
	"github.com/cilidm/dy-spider/app/constant"
	"github.com/cilidm/dy-spider/app/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
)

func InitConn() *gorm.DB {
	return GormSqlite()
}

var (
	db  *gorm.DB
	err error
)

func GormSqlite() *gorm.DB {
	dbFile := fmt.Sprintf("./%s.db", constant.DbName)
	db, err = gorm.Open("sqlite3", dbFile)
	if err != nil {
		if err := createDB(dbFile); err != nil {
			log.Fatal(err)
		}
	}
	db.SingularTable(true)
	db.LogMode(true)
	initTables()
	return db
}

func createDB(path string) error {
	fp, err := os.Create(path) // 如果文件已存在，会将文件清空。
	if err != nil {
		return err
	}
	defer fp.Close() //关闭文件，释放资源。
	return nil
}

func initTables() {
	checkTableData(&model.Spider{})
	checkTableData(&model.SpiderUserInfo{})
	checkTableData(&model.AwemeList{})
}

func checkTableData(tb interface{}) {
	if db.HasTable(tb) == false {
		if err := db.Debug().CreateTable(tb).Error; err != nil {
			log.Fatal("创建数据表失败", err.Error())
		}
	} else {
		// 已存在的表校验一下是否有新增字段
		if err := db.Debug().AutoMigrate(tb).Error; err != nil {
			log.Fatal("数据库初始化失败", err.Error())
		}
	}
}
