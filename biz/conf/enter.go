package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
	Jwt    Jwt    `yaml:"jwt"`
	Redis  Redis  `yaml:"redis"`
	EasyAR EasyAR `yaml:"easyAR"`
	Qiniu  Qiniu  `yaml:"qi_niu"`
}

var Conf = new(Config)

const ConfigFile = "conf/production.yaml"

// Init 读取yaml文件的配置
func Init() {
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, Conf)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success.")
}

// SetYaml 设置yaml文件
func SetYaml() error {
	byteData, err := yaml.Marshal(Conf)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ConfigFile, byteData, 0644)

	if err != nil {
		return err
	}
	log.Println("配置文件修改成功")
	return nil
}
