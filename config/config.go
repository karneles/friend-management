package config

import (
	"log"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
/*	Port                 string `envconfig:"port" default:"9111"`
	MysqlHost            string `envconfig:"mysql_host" default:"devbox-carneles.dev.svc.cluster.local"`
	MysqlUsername        string `envconfig:"mysql_username" default:"ssi"`
	MysqlPassword        string `envconfig:"mysql_password" default:"teramakuro"`
	MysqlConnectionLimit int    `envconfig:"mysql_connection_limit" default:"40"`
	MysqlDatabase        string `envconfig:"mysql_database" default:"test"`
*/
	Port                 string `envconfig:"port" default:"9111"`
	MysqlHost            string `envconfig:"mysql_host" default:"us-cdbr-iron-east-05.cleardb.net"`
	MysqlUsername        string `envconfig:"mysql_username" default:"ba4b5d33f28a01"`
	MysqlPassword        string `envconfig:"mysql_password" default:"0aa6bf4b"`
	MysqlConnectionLimit int    `envconfig:"mysql_connection_limit" default:"40"`
	MysqlDatabase        string `envconfig:"mysql_database" default:"heroku_3a05363a2d674d8"`
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
