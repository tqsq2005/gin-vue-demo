package util

import . "github.com/tqsq2005/gin-vue/pkg/setting"

func InitUtil()  {
	jwtSecret = []byte(Config.App.JwtSecret)
}
