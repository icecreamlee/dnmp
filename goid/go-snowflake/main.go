package main

import (
	"encoding/json"
	"fmt"
	"github.com/IcecreamLee/goutils"
	"os"
)

var id *goutils.IDGenServ

func main() {
	loadConfig()
	id.Run()
}

// 加载配置文件
func loadConfig() {
	id = goutils.IDServSingleton()
	file, _ := os.Open(goutils.GetCurrentPath() + "conf.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	err := decoder.Decode(id)
	if err != nil {
		fmt.Println("Load Config Error:", err)
	}
}
