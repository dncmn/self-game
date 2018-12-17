package bash

import (
	"bytes"
	"os/exec"
	"self-game/utils/logging"
)

var logs = logging.GetLogger()

// Bash bash执行命令，子进程方式
func Bash(cmdStr string) (outStr string, err error) {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)
	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	logs.Infof("bashBash start cmdStr:%s", cmdStr)
	err = cmd.Run()
	if err != nil {
		logs.Errorf("bashBash err stderr:%q err:%v cmdStr:%s", stderr.String(), err, cmdStr)
		return
	}
	defer cmd.Process.Release()
	defer cmd.Wait()
	logs.Infof("bashBash success stdout:%q cmdStr:%s", stdout.String(), cmdStr)
	return
}
