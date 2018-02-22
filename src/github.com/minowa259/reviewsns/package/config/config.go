package config

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
)

/* 他クラスからもアクセス可能な構造体 */
type Data struct {
	Config dataCategory `yaml:"config"`
}

type dataCategory struct {
	Site     site     `yaml:"site"`
	Database database `yaml:"detabase"`
}

type site struct {
	Host string `yaml:"hosturl"`
	Sitename string `yaml:"sitename"`
}

type database struct {
	Dbname   string `yaml:"dbname"`
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     string `yaml:"port"`
}


func LoadConfig() Data {
	// yamlを読み込むstring
	buf, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	// structにUnmasrshal
	var d Data
	err = yaml.Unmarshal(buf, &d)
	if err != nil {
		panic(err)
	}
	return d

}
