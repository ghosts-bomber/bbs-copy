package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct{

	LogFile 		string 	`yaml:"LogFile"` // 日志文件
	ShowSql			bool	`yaml:"ShowSql"` //是否显示数据库日志
	MySqlUrl 		string 	`yaml:"MySqlUrl"` //数据库地址
}

var Instance *Config

func Init(filename string) *Config{
	Instance = &Config{}
	if yamlFile,err:=ioutil.ReadFile(filename);err!=nil{
		logrus.Error(err)
	}else if err = yaml.Unmarshal(yamlFile,Instance);err!=nil{
		logrus.Error(err)
	}
	return Instance
}