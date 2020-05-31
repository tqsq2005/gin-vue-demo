package main

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/tqsq2005/gin-vue/model"
	"github.com/tqsq2005/gin-vue/pkg/logging"
	. "github.com/tqsq2005/gin-vue/pkg/setting"
	"time"
)

func init() {
	//step1：读取配置文件 通过导入读取
	//step2：连接数据库
	//model.InitDB()
	//step3：启动日志
	logging.InitLog()
	//step4：启动redis

	//step5：其他
	//util.InitUtil()
}

func main() {
	//logging.Logger.Warnln("something is here,", time.Now())
	logging.AppLog.Warnln("something is here,", time.Now())
	logging.SQLLog.Warnln("something is here,", time.Now())
}

func runServer()  {
	r := gin.Default()
	r.POST("/api/user/register", func(c *gin.Context) {
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
	log.Fatalln(r.Run(Config.Server.Port)) // 监听并在 0.0.0.0:8080 上启动服务
}