package main

import (
	"fmt"
	"github.com/wannanbigpig/gin-layout/boot"
	"github.com/wannanbigpig/gin-layout/config"
	"github.com/wannanbigpig/gin-layout/internal/command"
	"github.com/wannanbigpig/gin-layout/internal/routers"
	"strings"
)

func main() {
	run()
}

func run() {
	script := strings.Split(boot.Run, ":")
	switch script[0] {
	case "http":
		r := routers.SetRouters()

		err := r.Run(fmt.Sprintf("%s:%d", config.Config.Server.Host, config.Config.Server.Port))
		if err != nil {
			panic(err)
		}
	case "command":
		if len(script) != 2 {
			panic("命令错误，缺少重要参数")
		}
		command.Run(script[1])
	default:
		panic("执行脚本错误")
	}
}
