package router

import (
	"back/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine //声明路由Router
//Cors
/*
Cors-跨院与资源共享
以往的后端设计中使用了同源策略保证安全性，同源，即域名、协议、端口完全相同的地址，如https://a.com与http://a.com是非同源地址
而ajax请求只能同源，而CORS是一个W3C标准，它允许浏览器向跨源服务器发出XMLHttpRequest请求，从而突破了ajax只能使用同源服务的局限性。
*/

// InitRouter
/*
简单说明一下有关gin框架中路由存在的理解，实际上router才是服务端与客户端通信的核心
router将地址映射到指定函数凭此完成整个后端的调用————函数的返回值将作为http相应返回给客户端
而同时router也会处理来自客户端的所有http请求
总之 这个是一切的起点
*/
func InitRouter() {
	Router = gin.Default()             //创建路由
	corsConfig := cors.DefaultConfig() //cors资源共享
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	if config.Config.Dev {
		corsConfig.AllowAllOrigins = true //如果是开发模式，那么接受所有源的请求
	} else {
		corsConfig.AllowOrigins = config.Config.Server.AllowOrigins //如果不是，那么就仅允许指定源访问
	}
	Router.Use(cors.New(corsConfig))
	SetRouter()
}
