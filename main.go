package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/nxczje/logcmd/log"
)

func main() {
	pty := os.Args[1]
	cmds := os.Args[2:]
	pid := 0
	go func() {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGINT)
		<-interrupt
		fmt.Println("\n")
		log.Log(log.LogData{
			Level: "warn",
			Func:  "Interrupt",
			Data:  "Cancel command",
		})
		if pid != 0 {
			syscall.Kill(pid, syscall.SIGINT)
		}
	}()
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
	commmandRun.Stdout = file
	commmandRun.Stderr = file
	if err := commmandRun.Start(); err != nil {
		log.Log(log.LogData{
			Level: "error",
			Func:  "Start",
			Data:  "Start command failed : " + err.Error(),
		})
		return
	}
	pid = commmandRun.Process.Pid
	if err := commmandRun.Wait(); err != nil {
		log.Log(log.LogData{
			Level: "error",
			Func:  "Running",
			Data:  "Running command failed : " + err.Error(),
		})
		return
	}
	log.Log(log.LogData{
		Level: "done",
		Func:  "Done",
		Data:  "OK",
	})
}
