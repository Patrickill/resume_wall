package utility

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response http响应
/*gin.H本质上是一种通用的Map<String>接口，即map[string]interface
然后gin将map渲染为Json返回给服务端
*/
func Response(code int, msg string, data gin.H, c *gin.Context) {
	if data != nil {
		c.JSON(code, gin.H{
			"msg":  msg,
			"data": data,
		})
	} else {
		c.JSON(code, gin.H{ //返回JSOM
			"msg": msg,
		})
	}
}

func ResponseBadRequest(c *gin.Context) {
	Response(http.StatusBadRequest, "Bad request", nil, c)
}

func ResponseInternalServerError(c *gin.Context) {
	Response(http.StatusInternalServerError, "Internal server error", nil, c)
}

func ResponseOK(c *gin.Context, data gin.H) {
	Response(http.StatusOK, "OK", data, c)
}
