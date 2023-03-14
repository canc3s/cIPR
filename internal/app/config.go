package app

import (
	"github.com/canc3s/cIPR/internal/tools"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

var ConfigFile string

type Config struct {
	DatPath		string				`yaml:"datPath"`
	BlackList	[]string			`yaml:"blackList"`
}

func init()  {
	ConfigFile = os.Getenv("CIPR_Conf")
}

func InitConf(option *Options) {
	if ConfigFile == "" {
		log.Println("未检测到环境变量 CIPR_Conf，默认读取./config.yml")
		ConfigFile = "./config.yml"
	}
	if !tools.FileExists(ConfigFile) {
		DefaultConf()
	}
	err := LoadConf(option)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func DefaultConf() {
	config := &Config{
		"./qqwry.dat",
		[]string{"阿里巴巴","CDN","局域网","美国","阿里","Azure","华为","腾讯","Microsoft","微软","网宿","Amazon"},
	}
	d, err := yaml.Marshal(config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	ioutil.WriteFile(ConfigFile, d, 0644)
}

func LoadConf(option *Options) error {
	yamlFile, err := ioutil.ReadFile(ConfigFile)
	config := &Config{}
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return err
	}
	option.DatPath = config.DatPath
	option.BlackList = config.BlackList
	return nil
}