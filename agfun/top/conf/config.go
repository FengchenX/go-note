package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type TopConf struct {
	TopCreateTime string `yaml:"top_create_time"`
}

var conf *TopConf

func TopConfInst() *TopConf {
	if conf == nil {
		data, _ := ioutil.ReadFile("D:/myPro/go-note/src/agfun/top/conf/config.yml")
		fmt.Println(string(data))
		t := TopConf{}
		//把yaml形式的字符串解析成struct类型
		yaml.Unmarshal(data, &t)
		fmt.Println("初始数据", t)
		if t.TopCreateTime == "" {
			fmt.Println("配置文件设置错误")
			return nil
		}
		conf = &t
	}
	return conf
}
