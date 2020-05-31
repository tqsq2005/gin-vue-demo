package setting

import (
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"sync"
)

type App struct {
	Name            string
	Mode            string
	PageSize        int
	JwtSecret       string
	JwtTokenTimeout int
	PrefixUrl       string
	RuntimeRootPath string
}

type Log struct {
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

type Server struct {
	Port         string
	ReadTimeout  int
	WriteTimeout int
}

type DataSource struct {
	DriveName   string
	Host        string
	Database    string
	Username    string
	Password    string
	Port        string
	Charset     string
	Loc         string
	TablePrefix string
}

type Configuration struct {
	App        App
	Log        Log
	Server     Server
	DataSource DataSource
}

var Config *Configuration
var once sync.Once

func init() {
	workDir, _ := os.Getwd()
	configPath := workDir + "/config"
	viperConfig := viper.New()
	viperConfig.AddConfigPath(configPath)
	viperConfig.SetConfigName("application")
	viperConfig.SetConfigType("yml")
	err := viperConfig.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件信息出错：%v\n", err)
	}
	err = viperConfig.Unmarshal(&Config)
	if err != nil {
		log.Fatalf("解析配置文件信息出错：%v\n", err)
	}
	viperConfig.WatchConfig()
	viperConfig.OnConfigChange(func(in fsnotify.Event) {
		log.Infof("配置文件有更新，重新读取中...\n")
		err = viperConfig.Unmarshal(&Config)
		if err != nil {
			log.Fatalf("配置文件更新后解析配置文件信息出错：%v\n", err)
		}
	})
}
