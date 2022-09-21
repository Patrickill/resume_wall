package model

import "back/db"

type LoveMessage struct {
	ID      int    `gorm:"primaryKey;UNIQUE_INDEX;AUTO_INCREMENT" json:"id"`
	Email   string `gorm:"not null" json:"email"`    //发送者邮箱
	Name    string `gorm:"not null" json:"name" `    //匿名显示
	Message string `gorm:"not null" json:"message" ` //消息内容
	Time    string `gorm:"not null" json:"time" `    //发送时间
	UserID  uint64 `gorm:"default:1" json:"user_id"`
}

func AddMessage(message *LoveMessage) error {
	return db.DB.Create(message).Error
}

func GetMessageByID(id uint64) (*APIUser, error) {
	var user APIUser
	err := db.DB.Model(&User{}).Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}
