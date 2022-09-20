package db

import (
	"back/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB //数据库类型，所有的数据都在这里，可以通过他读取数据库！

func InitDB() {
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.DB.UserName,
		config.Config.DB.Password,
		config.Config.DB.Address,
		config.Config.DB.DBName,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Panicln("数据库链接出错: ", err)
	}
}
