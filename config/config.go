package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"self-game/utils"
	"self-game/utils/async"
	"time"
)

func watchAndUpdateConfig() {
	viper.SetConfigName("conf")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	var configuration Conf

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
		return
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatal(err)
		return
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println(in.Op.String())
		if in.Op == fsnotify.Create || in.Op == fsnotify.Write {
			if err := viper.Unmarshal(&configuration); err != nil {
				log.Fatal(err)
			}
			fmt.Println(configuration.Development.Cfg.Port)
		}
	})
}
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

	go async.Do(func() {
		watchAndUpdateConfig()
	})
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

// redis configItem
type RedisConfig struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// Postres configItem
type PgConfig struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	Dbname          string        `yaml:"dbname"`
	Username        string        `yaml:"username"`
	Password        string        `yaml:"password"`
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
}

type CodeConfig struct {
	RuntimeRootPath string `yaml:"runtimeRootPath"`
	PrefixUrl       string `yaml:"prefixUrl"`
	QrCodeSavePath  string `yaml:"qrCodeSavePath"`
}

// 项目配置文件
type CfgConfig struct {
	Token        string `yaml:"token"`
	Port         int    `yaml:"port"`
	TimeZone     string `yaml:"timeZone"`
	TimeModelStr string `yaml:"timeModelStr"`
}

// env 配置文件
type EnvConfig struct {
	ENV string `yaml:"env"`
}

type WechatConfig struct {
	AppID  string `yaml:"appid"`
	Secret string `yaml:"secret"`
	Token  string `json:"token"`
}

type ConfigItem struct {
	Mysql MysqlConfig `yaml:"mysql"`
	Env   EnvConfig   `yaml:"env"`
	Cfg   CfgConfig   `yaml:"cfg"`
	Redis RedisConfig `yaml:"redis"`
	//Pgsql  PgConfig     `yaml:"pgsql"`
	Code   CodeConfig   `yaml:"app"`
	Wechat WechatConfig `yaml:"wechat"`
}

type Conf struct {
	Development ConfigItem `yaml:"development"`
	Test        ConfigItem `yaml:"test"`
	Production  ConfigItem `yaml:"production"`
}
