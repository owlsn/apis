package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var Conf *Config

type Config struct {
	Env     string `yaml:"Env"`     // 环境：prod、dev
	BaseUrl string `yaml:"BaseUrl"` // base url
	Port    string `yaml:"Port"`    // 端口
	LogFile string `yaml:"LogFile"` // 日志文件

	MySqlUrl     string `yaml:"MySqlUrl"` // 数据库连接地址
	MaxIdleConns int    `yaml:"MaxIdleConns"`
	MaxOpenConns int    `yaml:"MaxOpenConns"`
	ShowSql      bool   `yaml:"ShowSql"` // 是否显示日志
}

func InitConfig(filename string) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		logrus.Error(err)
		return
	}

	Conf = &Config{}
	err = yaml.Unmarshal(yamlFile, Conf)
	if err != nil {
		logrus.Error(err)
	}
}
