package sllog

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type fileHook struct {
	W *fileLogWriter
}

func newFileHook(conf FileLogConfig) (hook logrus.Hook, err error) {
	path := strings.Split(conf.Filename, "/")
	if len(path) > 1 {
		exec.Command("mkdir", path[0]).Run()
	}
	w := newFileWriter()
	if err = w.Init(conf); err != nil {
		return
	}
	hook = &fileHook{W: w}
	return
}

func (p *fileHook) Fire(entry *logrus.Entry) (err error) {
	message, err := getMessage(entry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}
	switch entry.Level {
	case logrus.PanicLevel:
		fallthrough
	case logrus.FatalLevel:
		fallthrough
	case logrus.ErrorLevel:
		return p.W.WriteMsg(string(message), ErrorLevel)
	case logrus.WarnLevel:
		return p.W.WriteMsg(string(message), WarnLevel)
	case logrus.InfoLevel:
		return p.W.WriteMsg(string(message), InfoLevel)
	case logrus.DebugLevel:
		return p.W.WriteMsg(string(message), DebugLevel)
	default:
		return nil
	}

	return
}

func (p *fileHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

func getMessage(entry *logrus.Entry) (message string, err error) {
	var formatter = logrus.JSONFormatter{}
	file, lineNumber := getCallerIgnoringLogMulti(2)
	if file != "" {
		sep := fmt.Sprintf("%s/src/", os.Getenv("GOPATH"))
		fileName := strings.Split(file, sep)
		if len(fileName) >= 2 {
			file = fileName[1]
		}
	}
	entry.Data["file"] = file
	entry.Data["lineNumber"] = lineNumber
	messageBytes, err := formatter.Format(entry)
	if err != nil {
		fmt.Println(err)
	}
	message = string(messageBytes)
	return
}

// get file and line number ignore
func getCallerIgnoringLogMulti(callDepth int) (file string, line int) {
	return getCaller(callDepth+1, "logging.go", "logrus/hooks.go", "logrus/entry.go", "logrus/logger.go", "logrus/exported.go", "proc.go", "asm_amd64.s")
}

// get file and line number.
func getCaller(callDepth int, suffixesToIgnore ...string) (file string, line int) {
	// bump by 1 to ignore the getCaller (this) stackframe
	callDepth++
outer:
	for {
		var ok bool
		_, file, line, ok = runtime.Caller(callDepth)
		if !ok {
			file = "???"
			line = 0
			break
		}

		for _, s := range suffixesToIgnore {
			if strings.HasSuffix(file, s) {
				callDepth++
				continue outer
			}
		}
		break
	}
	return
}
