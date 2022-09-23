package utility

import (
	"back/config"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"time"
)

type JwtData struct {
	ID string `json:"id"` //包含用户名称
	jwt.StandardClaims
}

func GenerateStandardJwt(jwtData *JwtData) string {

	claims := jwtData
	fmt.Println("传入的ID是：", claims.ID)

	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(time.Duration(config.Config.Jwt.Expires) * time.Hour)).Unix(), //token过期时间 单位：小时
		Issuer:    config.Config.Jwt.Issuer,                                                                   //签发人是谁                                                                   //指定token发行人
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //利用sha256加密生成token

	//用于获取完整、已签名的token，根据config的secret密钥生成签名字符串
	// 使用指定的secret签名并获得完整的编码后的字符串token
	token, err := tokenClaims.SignedString([]byte(config.Config.Jwt.Secret))
	if err != nil { //报错抛出
		log.Fatalln("Jwt错误", err)
		panic(err)
	}
	return token
}

// ParseToken 根据传入的token值获取到Claims对象信息，从而获取其中的用户名和密码）
func ParseToken(token string) (string, error) {
	//传入配置文件中的Secret密钥
	jwtSecret := []byte(config.Config.Jwt.Secret)
	// 解析token（传参分别为：字符串token，将解析结果保存至指定的结构体，返回生成token时所使用的secret从而用于解签名的方法）
	tokenClaims, err := jwt.ParseWithClaims(token, &JwtData{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	}) //解析token方法

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间
		if claims, ok := tokenClaims.Claims.(*JwtData); ok && tokenClaims.Valid {
			fmt.Println("解析出的ID是：", claims.ID)
			return claims.ID, err
		}
	}
	return "", err //BUT为什么要返回一个空值？？？
}
