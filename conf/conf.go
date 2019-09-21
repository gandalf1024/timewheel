package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"timewheel/errors"
	"timewheel/golog"
	"timewheel/wheel"
)

var Config GlobalConf

type ServerConf struct {
	Name    string //项目名称
	Port    int32  //监听端口
	Addr    string //监听地址(0.0.0.0:Port)
	Log     string //日志路径
	GitHash string //当前运行hash值
}

type GlobalConf struct {
	Wheel  wheel.Wheel
	Server ServerConf
}

func init() {
	var env = "task"
	path := fmt.Sprintf("./%s.tml", env)
	Config = GlobalConf{}
	Config = TOMLobj(path, Config)
	Config.Server.Addr = fmt.Sprintf("0.0.0.0:%d", Config.Server.Port)
	golog.Infof("项目配置初始化成功:=(%v)", Config)
}

const (
	globalConfigurationKeyword = "~"
)

var errConfigurationDecode = errors.New("error while trying to decode configuration")

func TOMLobj(filename string, c GlobalConf) GlobalConf {

	if filename == globalConfigurationKeyword {
		filename = homeConfigurationFilename(".tml")
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			panic("default configuration file '" + filename + "' does not exist")
		}
	}

	tomlAbsPath, err := filepath.Abs(filename)
	if err != nil {
		panic(errConfigurationDecode.AppendErr(err))
	}

	data, err := ioutil.ReadFile(tomlAbsPath)
	if err != nil {
		panic(errConfigurationDecode.AppendErr(err))
	}

	if _, err := toml.Decode(string(data), &c); err != nil {
		panic(errConfigurationDecode.AppendErr(err))
	}
	return c
}

func homeConfigurationFilename(ext string) string {
	return filepath.Join(homeDir(), "iris"+ext)
}

func homeDir() (home string) {
	u, err := user.Current()
	if u != nil && err == nil {
		home = u.HomeDir
	}

	if home == "" {
		home = os.Getenv("HOME")
	}

	if home == "" {
		if runtime.GOOS == "plan9" {
			home = os.Getenv("home")
		} else if runtime.GOOS == "windows" {
			home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
			if home == "" {
				home = os.Getenv("USERPROFILE")
			}
		}
	}

	return
}
