package controller

import (
	"back/db"
	"back/db/model"
	"back/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.ResponseBadRequest(c)
		return
	}
	var user *model.User
	user, err = model.GetUserByEmail(data.Email) //利用邮箱登录获取对应USER
	if user == nil {
		utility.Response(http.StatusNotFound, "User not found", nil, c)
		return
	}
	if !utility.PasswordVerify(data.Password, user.Password) {
		utility.Response(http.StatusBadRequest, "Wrong Password", nil, c)
		return
	}
	fmt.Println("Login: ", user.UserID)
	token := utility.GenerateStandardJwt(&utility.JwtData{
		ID: strconv.Itoa(int(user.UserID)),
	})
	utility.Response(http.StatusOK, "OK", gin.H{"token": token}, c)
}

func Register(c *gin.Context) {
	var data model.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utility.ResponseBadRequest(c)
		return
	}
	data.Password, err = utility.PasswordHash(data.Password)
	if err != nil {
		log.Println(err)
		utility.ResponseInternalServerError(c)
		return
	}
	var data2 model.User
	var data3 model.User
	err = db.DB.Where("name = ?", data.Name).Find(&data2).Error
	err = db.DB.Where("email = ?", data.Email).Find(&data3).Error
	if data2.UserID != 0 {
		utility.Response(http.StatusBadRequest, "用户名已经被注册", nil, c)
		return
	}
	if data3.UserID != 0 {
		utility.Response(http.StatusBadRequest, "邮箱已经被注册", nil, c)
		return
	}
	err = model.AddUser(&data)
	if err != nil {
		log.Println(err)
		utility.ResponseInternalServerError(c)
		return
	}
	utility.Response(http.StatusOK, "OK", nil, c)
}

// NameIsRegister 名字是否被占用
func NameIsRegister(c *gin.Context) {
	var data model.User
	var data2 model.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utility.ResponseBadRequest(c)
		return
	}
	err = db.DB.Where("name = ?", data.Name).Find(&data2).Error
	if data2.UserID != 0 {
		utility.Response(http.StatusBadRequest, "用户名已经被注册", nil, c)
		return
	}
	utility.Response(http.StatusOK, "名称可以使用", nil, c)
	return
}
func EmailIsRegister(c *gin.Context) {
	var data model.User
	var data2 model.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utility.ResponseBadRequest(c)
		return
	}
	err = db.DB.Where("email = ?", data.Email).Find(&data2).Error
	if data2.UserID != 0 {
		utility.Response(http.StatusBadRequest, "邮箱已经被注册", nil, c)
		return
	}
	utility.Response(http.StatusOK, "邮箱可以使用", nil, c)
	return
}
func UpdateUser(c *gin.Context) {
	var data model.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utility.ResponseBadRequest(c)
		return
	}
	data.UserID = c.GetUint64("user_id")
	err = model.UpdateUser(&data)
}

func GetUserInfo(c *gin.Context) {
	userID, _ := c.Get("user_id")
	user, err := model.GetUserByID(userID.(uint64))
	if err != nil {
		log.Println(err)
		utility.ResponseInternalServerError(c)
		return
	}
	utility.Response(http.StatusOK, "OK", gin.H{
		"info": user,
	}, c)
}
