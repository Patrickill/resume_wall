package middleware

import (
	"back/utility"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Authorization(c *gin.Context) { //验证登陆的token
	token := c.GetHeader("Authorization")
	if token == "" {
		utility.Response(http.StatusUnauthorized, "No Token", nil, c)
		c.Abort()
		return
	}
	id, err := utility.ParseToken(token)
	if err != nil {
		utility.Response(http.StatusUnauthorized, "Bad Token", nil, c)
		c.Abort()
		return
	}
	id_int, ok := strconv.Atoi(id)
	if ok != nil {
		utility.Response(http.StatusInternalServerError, "Internal server error", nil, c)
		log.Println(ok)
		c.Abort()
	}
	c.Set("user_id", uint64(id_int))
	c.Next()
}
