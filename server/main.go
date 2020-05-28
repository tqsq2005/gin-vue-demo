package main

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/tqsq2005/gin-vue/model"
)

func main() {
	r := gin.Default()
	r.GET("/api/user/register", func(c *gin.Context) {
		//接受数据
		user := model.User{}
		err := c.Bind(&user)
		if err != nil {
			log.Warnln("参数绑定失败,err:", err)
		}
		log.Infoln(user)
		//数据验证
		result, err := govalidator.ValidateStruct(user)
		if err != nil {
			log.Warnln("数据验证失败,err:", err)
			c.JSON(400, gin.H{
				"err": err,
			})
			return
		}
		log.Infoln(result)
		c.JSON(200, user)
	})
	log.Fatalln(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}
