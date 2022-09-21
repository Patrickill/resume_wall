package controller

import (
	"back/db"
	"back/db/model"
	"back/utility"
	"crypto/rand"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"math/big"
	"net/http"
	"time"
)

func rangeRand(min, max int64) int64 {
	if min > max {
		return min
	}
	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))

		return result.Int64() - i64Min
	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}

func GetRandomMessage(c *gin.Context) {
	var message model.LoveMessage
	var message2 model.LoveMessage
	var message3 model.LoveMessage
	db.DB.First(&message)
	first := message.ID
	db.DB.Last(&message2)
	last := message2.ID
	id := rangeRand(int64(first), int64(last))
	fmt.Println(id)
	fmt.Println(last)
	fmt.Println(first)
	err := db.DB.Find(&message3, id)
	fmt.Println(err)
	//if err != nil {
	//	log.Println(err)
	//	utility.ResponseInternalServerError(c)
	//	return
	//}
	utility.Response(http.StatusOK, "OK", gin.H{
		"info": message3,
	}, c)
}

func ClaimLoveMessage(c *gin.Context) {
	var data model.LoveMessage
	userID, _ := c.Get("user_id")
	user, err := model.GetUserByID(userID.(uint64))
	if err != nil {
		log.Println(err)
		utility.ResponseInternalServerError(c)
		return
	}
	data.Email = user.Email
	err = c.ShouldBindJSON(&data) //转换为JSON 必需包含name message user_id
	if err != nil {
		log.Println(err)
		utility.ResponseBadRequest(c)
		return
	}
	//处理时间
	t1 := time.Now().Year()   //年
	t2 := time.Now().Month()  //月
	t3 := time.Now().Day()    //日
	t4 := time.Now().Hour()   //小时
	t5 := time.Now().Minute() //分钟
	time := fmt.Sprintf("%d年%d月%d日%d时%d分",
		t1, t2, t3, t4, t5,
	)
	data.Time = time
	err = model.AddMessage(&data) //添加消息
	if err != nil {
		log.Println(err)
		utility.ResponseInternalServerError(c)
		return
	}
	utility.ResponseOK(c, nil)
}
