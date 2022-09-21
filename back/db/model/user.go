package model

import "back/db"

type User struct { //后面的‘gorm’代表一些特殊属性
	//json:''为字段指定一个标记信息,这些标记信息通过反射接口可见，并参与结构体的类型标识，但在其他情况下被忽略。
	UserID   uint64 `gorm:"primaryKey" json:"user_id"`                //指定一个列作为主键
	Name     string `gorm:"not null" json:"name" validate:"required"` //validate表示必须要有
	Email    string `gorm:"not null" json:"email" validate:"required"`
	Phone    string `json:"phone"`
	Password string `gorm:"not null" json:"password" validate:"required"`
	School   string `json:"school"`
	StuID    string `json:"stu_id"`
}

type APIUser struct {
	Name        string `gorm:"primaryKey" json:"name"`
	Email       string `gorm:"primaryKey" json:"email"`
	Phone       string `json:"phone"`
	School      string `json:"school"`
	StuID       string `json:"stu_id"`
	ClassAmount uint64 `gorm:"not null; default:12" json:"class_amount"`
	WeekAmount  uint64 `gorm:"not null; default:16" json:"week_amount"`
}

// AddUser Omit仿佛是忽略指定键值
func AddUser(user *User) error {
	return db.DB.Omit("UserID", "Classes", "ClassAmount", "WeekAmount").Create(user).Error
}

func GetUserByEmail(email string) (*User, error) {
	var user *User
	err := db.DB.Where("email = ?", email).First(&user).Error //Where是限定条件，所以只有一条数据的时候第一个必然是这个
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func GetUserByID(id uint64) (*APIUser, error) {
	var user APIUser
	err := db.DB.Model(&User{}).Where("user_id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}

func UpdateUser(user *User) error {
	var u User
	err := db.DB.Where("user_id = ?", user.UserID).Take(&u).Error //利用UserID获取User对象
	if err != nil {                                               //看看是不是能查到，虽然不是很清楚为什么这一步
		return err
	}
	return db.DB.Omit("UserID", "Classes").Save(user).Error
}

func DeleteUserByID(id uint64) error {
	return db.DB.Where("user_id = ?", id).Delete(&User{}).Error
}
