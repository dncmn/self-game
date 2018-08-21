package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"self_game/utils"
	"time"
)

func init() {
	allConf := new(Conf)

	filePath := ""
	switch {
	case utils.FileExist("./conf.yaml"):
		filePath = "./conf.yaml"
	case utils.FileExist("./config/conf.yaml"):
		filePath = "./config/conf.yaml"
	case utils.FileExist("../config/conf.yaml"):
		filePath = "../config/conf.yaml"
	default:
		log.Fatal("config file not found")
	}

	yamlBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yamlBytes, allConf)
	if err != nil {
		log.Fatal(err)
	}

	// 给全局变量赋值
	env := os.Getenv("GO_ENV")
	switch env {
	case "test", "":
		Config = allConf.Test
	case "production":
		Config = allConf.Production
	default:
		Config = allConf.Development
	}
}

var Config ConfigItem

// mysql 配置文件
type MysqlConfig struct {
	Host            string        `yaml:"host"`
	Dbname          string        `yaml:"dbname"`
	Username        string        `yaml:"username"`
	Password        string        `yaml:"password"`
	MaxOpenConns    int           `yam:"maxOpenConns"`
	GOMaxIdleConns  int           `yaml:"maxIdelConns"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
}

// 项目配置文件
type CfgConfig struct{
	Token string  `yaml:"token"`
	Port int `yaml:"port"`
}


// env 配置文件
type EnvConfig struct {
	ENV string `yaml:"env"`
}

type ConfigItem struct {
	Mysql MysqlConfig
	Env   EnvConfig
	Cfg CfgConfig
}

type Conf struct {
	Development ConfigItem `yaml:"development"`
	Test        ConfigItem `yaml:"test"`
	Production  ConfigItem `yaml:"production"`
}
