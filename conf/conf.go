package conf

import (
	"os"
	"encoding/json"
)

type mysql struct {
	User     string
	Password string
	Database string
}

type redis struct {
	User     string
	Password string
}

type prod struct {
	Host string
	Port string
}

type dev struct {
	Host string
	Port string
}

type Conf struct {
	Mysql mysql
	Redis redis
	Prod  prod
	Dev   dev
}

func LoadConf() *Conf {
	file, _ := os.Open("conf/conf.json")
	decoder := json.NewDecoder(file)
	conf := Conf{}
	err := decoder.Decode(&conf)
	if err != nil {
		panic("Load Config File Failed")
	}
	return &conf
}
