package main

import (
	"bbs-copy/config"
	"bbs-copy/controllers"
	"bbs-copy/model"
	"bbs-copy/pkg/simple"
	"flag"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
)

var configFile = flag.String("config","./bbs-go.yaml","配置文件路径")

func init(){
	flag.Parse()

	config := config.Init(*configFile)

	gormConf := &gorm.Config{}

	// 初始化日志
	if file,err:= os.OpenFile(config.LogFile,os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666);err!=nil{
		logrus.SetOutput(io.MultiWriter(os.Stdout,file))
		if config.ShowSql{
			gormConf.Logger = logger.New(log.New(file,"\r\n",log.LstdFlags),logger.Config{
				SlowThreshold: time.Second,
				Colorful: true,
				LogLevel: logger.Info,
			})
		}
	}else{
		logrus.SetOutput(os.Stdout)
		logrus.Error(err)
	}
	if err:= simple.OpenDB(config.MySqlUrl,gormConf,20,20,model.Models...);err!=nil{
		logrus.Error(err)
	}
}

func main() {
	controllers.Router()
}
