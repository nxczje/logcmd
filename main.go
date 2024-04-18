package main

import (
	"os"
	"os/exec"

	"github.com/nxczje/logcmd/log"
)

func main() {
	cmd := os.Args[1]
	pty := os.Args[2]
	log.Banner()
	arrcmd := []string{
		"bash",
		"-c",
		cmd,
	}
	log.Log(log.LogData{
		Level: "info",
		Func:  "Pending",
		Data:  "Command : " + cmd,
	})
	file, _ := os.OpenFile(pty, os.O_WRONLY, os.ModeAppend)
	commmandRun := exec.Command(arrcmd[0], arrcmd[1:]...)
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
