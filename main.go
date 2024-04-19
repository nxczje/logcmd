package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/nxczje/logcmd/log"
)

func main() {
	pty := os.Args[1]
	cmds := os.Args[2:]
	log.Banner()
	arrcmd := []string{
		"bash",
		"-c",
	}
	arrcmd = append(arrcmd, strings.Join(cmds, " "))
	log.Log(log.LogData{
		Level: "info",
		Func:  "Pending",
		Data:  "Command : " + strings.Join(cmds, " "),
	})
	file, _ := os.OpenFile(pty, os.O_WRONLY, os.ModeAppend)
	commmandRun := exec.Command(arrcmd[0], arrcmd[1:]...)
	fmt.Println(commmandRun.String())
	commmandRun.Stdout = file
	commmandRun.Stderr = file
	if err := commmandRun.Start(); err != nil {
		log.Log(log.LogData{
			Level: "error",
			Func:  "Running",
			Data:  "Running command failed",
		})
		return
	}
	if err := commmandRun.Wait(); err != nil {
		log.Log(log.LogData{
			Level: "error",
			Func:  "Running",
			Data:  "Running command failed",
		})
		return
	}
	log.Log(log.LogData{
		Level: "done",
		Func:  "Done",
		Data:  "OK",
	})
}
