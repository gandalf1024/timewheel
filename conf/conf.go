package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"path/filepath"
	"timewheel/glog"
)

var Config GlobalConf

type GlobalConf struct {
	Log     string //日志路径
	GitHash string //当前运行hash值
}

func init() {
	var env = "conf"
	path := fmt.Sprintf("./conf/%s.tml", env)
	Config = GlobalConf{}
	Config = TOMLobj(path, Config)
	glog.Infof("项目配置初始化成功:=(%v)", Config)
}

func TOMLobj(filename string, c GlobalConf) GlobalConf {
	tomlAbsPath, err := filepath.Abs(filename)
	if err != nil {
		return GlobalConf{}
	}

	data, err := ioutil.ReadFile(tomlAbsPath)
	if err != nil {
		return GlobalConf{}
	}

	if _, err := toml.Decode(string(data), &c); err != nil {
		return GlobalConf{}
	}
	return c
}
