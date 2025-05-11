package models

import (
	"fmt"
	"github.com/qiangge2020/go-gin-blog/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// 数据库连接
var Db *gorm.DB

func Setup() {
	var err error
	Db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

}
