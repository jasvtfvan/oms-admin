package core

import (
	"flag"
	"fmt"
	"os"

	"github.com/jasvtfvan/oms-admin/server/utils"
)

// 将pid写入到 .vscode/pid.txt 中，debug结束后，根据pid进行kill操作
func WritePIDToFile() {
	var env string
	flag.StringVar(&env, "env", "", "go main.go -env debug")
	flag.Parse()
	if env == "debug" {
		wd, err := os.Getwd() // ${workspaceFolder}/server
		if err != nil {
			fmt.Println("[Golang] vscode debug 启动，尝试pid写入文件，项目绝对路径获取失败")
			return
		}
		pidFile := wd[:len(wd)-6] + ".vscode/pid.txt"
		pid := os.Getpid()
		// 将PID写入pid.txt文件
		err = os.WriteFile(pidFile, []byte(fmt.Sprintf("%d", pid)), 0644)
		if err != nil {
			fmt.Printf("[Golang] 写入PID到文件 ${workspaceFolder}/.vscode/pid.txt 失败: %v\n", err)
			return
		}
		fmt.Printf("[Golang] PID %d 已写入文件 %s\n", pid, pidFile)
		fmt.Println(utils.GetStringWithTime("====== [Golang] vscode debug 启动 ======"))
	} else {
		fmt.Println(utils.GetStringWithTime("====== [Golang] 非 vscode debug 启动 ======"))
	}
}
