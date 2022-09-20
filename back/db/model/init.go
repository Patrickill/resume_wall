package model

import (
	"back/db"
	"log"
)

func InitModel() {
	ok := db.DB.AutoMigrate(&User{}, &Class{}) //自动补全键数值！
	if ok != nil {
		log.Panicln("s数据库错误: ", ok)
	}
}
