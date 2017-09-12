package config

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port                 string `envconfig:"port" default:"9111"`
	MysqlHost            string `envconfig:"mysql_host" default:"devbox-carneles.dev.svc.cluster.local"`
	MysqlUsername        string `envconfig:"mysql_username" default:"root"`
	MysqlPassword        string `envconfig:"mysql_password" default:"root-is-not-used"`
	MysqlConnectionLimit int    `envconfig:"mysql_connection_limit" default:"40"`
	MysqlDatabase        string `envconfig:"mysql_database" default:"test"`
	TimeZoneHelper       string `envconfig:"timezone_helper" default:"Asia/Jakarta"`
	PaginationSize       int    `envconfig:"pagination_size" default:"20"`
}

var conf Config
var once sync.Once

func GetConfig() Config {
	once.Do(func() {
		err := envconfig.Process("", &conf)
		if err != nil {
			log.Fatal("Can't load config: ", err)
		}
	})

	return conf
}
