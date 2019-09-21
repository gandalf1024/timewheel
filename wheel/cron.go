package wheel

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/juju/utils"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
	"timewheel/util"
)

type TomlWheel struct {
	W Wheel
}

var client *http.Client

func init() {
	client = utils.GetNonValidatingHTTPClient()
}

func Exec() {
	//遍历所有轮子

	flag := true
	index := 1
	for flag {
		var env = "task_"
		env += strconv.Itoa(index)
		tomlAbsPath, err := filepath.Abs(fmt.Sprintf("./conf/%s.tml", env))
		if err != nil {
			flag = false
			break
		}
		if !util.FileExist(tomlAbsPath) {
			flag = false
			break
		}

		w := TomlWheel{}
		w = TOMLobj(tomlAbsPath, w)
		AllWheel = append(AllWheel, w.W)
		index++
	}

	for i, _ := range AllWheel {
		time := AllWheel[i].WheelExecTime
		if util.GetTime() >= time {
			if AllWheel[i].Mode == EXEC_ONE {
				ExecOne(&AllWheel[i])
			}

			if AllWheel[i].Mode == EXEC_FOR {
				ExecFor(&AllWheel[i])
			}
		}
	}

}

func ExecOne(w *Wheel) {
	tasks := w.Tasks
	exec(tasks)
}

func ExecFor(w *Wheel) {
	tasks := w.Tasks
	for {
		exec(tasks)
		time.Sleep(time.Second)
	}
}

func exec(tasks []Task) {
	for _, v := range tasks {
		client.Get(v.Url)
		fmt.Println("-------------------->>>")
	}
}

func TOMLobj(filename string, c TomlWheel) TomlWheel {
	tomlAbsPath, err := filepath.Abs(filename)
	if err != nil {
		return TomlWheel{}
	}

	data, err := ioutil.ReadFile(tomlAbsPath)
	if err != nil {
		return TomlWheel{}
	}

	if _, err := toml.Decode(string(data), &c); err != nil {
		return TomlWheel{}
	}
	return c
}
