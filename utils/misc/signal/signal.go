package signal

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/glog"
)

func HandleSignal(closeF func()) {
	osCh := make(chan os.Signal, 1)
	glog.Infoln("Start Signal Hooker!")
	signal.Notify(osCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT) // , syscall.SIGSTOP) cannot compile on windows
	glog.Infof("Got a signal [%s]", <-osCh)
	glog.Flush()
	closeF()
}
