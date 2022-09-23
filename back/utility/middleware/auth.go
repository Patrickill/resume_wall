package middleware

import (
	"back/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Authorization(c *gin.Context) { //验证登陆的token
	token := c.GetHeader("Authorization")
	if token == "" {
		utility.Response(http.StatusUnauthorized, "未收到token", nil, c)
		c.Abort()
		return
	}
	id, err := utility.ParseToken(token)
	if err != nil {
		utility.Response(http.StatusUnauthorized, "Token验证失败", nil, c)
		c.Abort()
		return
	}
	//fmt.Println(id)
	idInt, ok := strconv.Atoi(id)
	fmt.Println("ID是", idInt)
	if ok != nil {
		utility.Response(http.StatusInternalServerError, "用户ID获取失败", nil, c)
		log.Println(ok)
		c.Abort()
	}
	c.Set("user_id", uint64(idInt))
	c.Next()
}
