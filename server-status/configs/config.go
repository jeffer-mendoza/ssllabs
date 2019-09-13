package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

//data structure to database configuration
type Conf struct {
	Host       string
	User       string
	Port       string
	Name       string
	Password   string
	ServerHost string
	ServerPort string
}

func (config *Conf) GetConf() *Conf {

	yamlFile, err := ioutil.ReadFile("conf.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return config
}

func (config *Conf) GetUrlConnection(conector string) string {
	if conector == "postgres" {
		return "host=" + config.Host + " port=" + config.Port + " user=" + config.User + " password=" + config.Password + " dbname=" + config.Name
	}
	return ""
}
