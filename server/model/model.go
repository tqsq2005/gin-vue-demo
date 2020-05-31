package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	. "github.com/tqsq2005/gin-vue/pkg/setting"
	"net/url"
	"time"
)

var db *gorm.DB

//连接到数据库
func InitDB() {
	var err error
	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=%s",
		Config.DataSource.Username,
		Config.DataSource.Password,
		Config.DataSource.Host,
		Config.DataSource.Port,
		Config.DataSource.Database,
		Config.DataSource.Charset,
		url.QueryEscape(Config.DataSource.Loc),
	)
	db, err = gorm.Open(Config.DataSource.DriveName, dns)
	if err != nil {
		log.Fatalf("连接数据库出错，连接字符串为：%q, 错误信息：%v\n", dns, err)
	}
	//设置表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return viper.GetString("dataSource.tablePrefix") + defaultTableName;
	}

	//表名是否用单数
	//db.SingularTable(true)

	//callback替换
	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.Callback().Create().Register("logging_create", logCreated)

	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(10)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(100)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(time.Hour)
	//migrate
	initMigrate()
}

func logCreated(scope *gorm.Scope) {
	sql := scope.SQL
	log.Infof("查询数据库的SQL语句：%v\n", sql)
}

func CloseDB()  {
	defer db.Close()
}

