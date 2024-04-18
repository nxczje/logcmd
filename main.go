package main

import (
	"os"
	"os/exec"

	"github.com/nxczje/logcmd/log"
)

func main() {
	cmd := os.Args[1]
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
	commmandRun := exec.Command(arrcmd[0], arrcmd[1:]...)
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
