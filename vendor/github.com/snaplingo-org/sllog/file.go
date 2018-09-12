package sllog

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

// init needed config
type FileLogConfig struct {
	Path     string
	Filename string
	MaxLines int
	Maxsize  int
	Daily    bool
	MaxDays  int64
	Rotate   bool
	Level    int
}

// an *os.File writer with locker.
type muxWriter struct {
	sync.Mutex
	fd *os.File
}

// FileLogWriter implements LoggerInterface.
// It writes messages by lines limit, file size limit, or time frequency.
type fileLogWriter struct {
	FileLogConfig
	*log.Logger
	mw               *muxWriter
	maxLinesCurLines int
	maxsizeCurSize   int
	dailyOpenDate    int
	startLock        sync.Mutex // Only one log can write to the file
}

// write to os.File.
func (l *muxWriter) Write(b []byte) (int, error) {
	l.Lock()
	defer l.Unlock()
	return l.fd.Write(b)
}

// set os.File in writer.
func (l *muxWriter) SetFd(fd *os.File) {
	if l.fd != nil {
		l.fd.Close()
	}
	l.fd = fd
}

// create a FileLogWriter returning as LoggerInterface.
func newFileWriter() *fileLogWriter {
	w := &fileLogWriter{
		FileLogConfig: FileLogConfig{
			Path:     "",
			Filename: "",
			MaxLines: 1000000,
			Maxsize:  1 << 28, //256 MB
			Daily:    true,
			MaxDays:  7,
			Rotate:   true,
			Level:    DebugLevel,
		},
	}
	// use MuxWriter instead direct use os.File for lock write when rotate
	w.mw = new(muxWriter)
	// set MuxWriter as Logger's io.Writer
	w.Logger = log.New(w.mw, "", 0)
	return w
}

// Init file logger.
func (w *fileLogWriter) Init(conf FileLogConfig) (err error) {
	w.FileLogConfig = conf

	if len(w.Filename) == 0 {
		return errors.New("jsonconfig must have filename")
	}

	return w.startLogger()
}

// start file logger. create log file and set to locker-inside file writer.
func (w *fileLogWriter) startLogger() error {
	fd, err := w.createLogFile()
	if err != nil {
		return err
	}
	w.mw.SetFd(fd)
	err = w.initFd()

	return err
}

func (w *fileLogWriter) docheck(size int) {
	w.startLock.Lock()
	defer w.startLock.Unlock()
	if w.Rotate && ((w.MaxLines > 0 && w.maxLinesCurLines >= w.MaxLines) ||
		(w.Maxsize > 0 && w.maxsizeCurSize >= w.Maxsize) ||
		(w.Daily && time.Now().Day() != w.dailyOpenDate)) {
		if err := w.DoRotate(); err != nil {
			fmt.Fprintf(os.Stderr, "FileLogWriter(%q): %s\n", w.Filename, err)
			return
		}
	}
	w.maxLinesCurLines++
	w.maxsizeCurSize += size
}

// write logger message into file.
func (w *fileLogWriter) WriteMsg(msg string, level int) error {
	if level > w.Level {
		return nil
	}
	n := 24 + len(msg) // 24 stand for the length "2013/06/23 21:00:22 [T] "
	w.docheck(n)
	w.Logger.Print(msg)
	return nil
}

func (w *fileLogWriter) createLogFile() (*os.File, error) {
	// Open the log file
	fd, err := os.OpenFile(w.Filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0660)
	return fd, err
}

func (w *fileLogWriter) initFd() error {
	fd := w.mw.fd
	finfo, err := fd.Stat()
	if err != nil {
		return fmt.Errorf("get stat err: %s\n", err)
	}
	w.maxsizeCurSize = int(finfo.Size())
	w.dailyOpenDate = time.Now().Day()
	if finfo.Size() > 0 {
		content, err := ioutil.ReadFile(w.Filename)
		if err != nil {
			return err
		}
		w.maxLinesCurLines = len(strings.Split(string(content), "\n"))
	} else {
		w.maxLinesCurLines = 0
	}
	return nil
}

// DoRotate means it need to write file in new file.
// new file name like xx.log.2013-01-01.2
func (w *fileLogWriter) DoRotate() error {
	_, err := os.Lstat(w.Filename)
	if err == nil { // file exists
		// Find the next available number
		num := 1
		fname := ""
		for ; err == nil && num <= 999; num++ {
			fname = w.Filename + fmt.Sprintf(".%s.%03d", time.Now().Format("2006-01-02"), num)
			_, err = os.Lstat(fname)
		}
		// return error if the last file checked still existed
		if err == nil {
			return fmt.Errorf("Rotate: Cannot find free log number to rename %s\n", w.Filename)
		}

		// block Logger's io.Writer
		w.mw.Lock()
		defer w.mw.Unlock()

		fd := w.mw.fd
		fd.Close()

		// close fd before rename
		// Rename the file to its newfound home
		err = os.Rename(w.Filename, fname)
		if err != nil {
			return fmt.Errorf("Rotate: %s\n", err)
		}

		// re-start logger
		err = w.startLogger()
		if err != nil {
			return fmt.Errorf("Rotate StartLogger: %s\n", err)
		}

		go w.deleteOldLog()
	}

	return nil
}

func (w *fileLogWriter) deleteOldLog() {
	dir := filepath.Dir(w.Filename)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) (returnErr error) {
		defer func() {
			if r := recover(); r != nil {
				returnErr = fmt.Errorf("Unable to delete old log '%s', error: %+v", path, r)
			}
		}()

		if !info.IsDir() && info.ModTime().Unix() < (time.Now().Unix()-60*60*24*w.MaxDays) {
			if strings.HasPrefix(filepath.Base(path), filepath.Base(w.Filename)) {
				os.Remove(path)
			}
		}
		return
	})
}

// destroy file logger, close file writer.
func (w *fileLogWriter) Destroy() {
	w.mw.fd.Close()
}

// flush file logger.
// there are no buffering messages in file logger in memory.
// flush file means sync file from disk.
func (w *fileLogWriter) Flush() {
	w.mw.fd.Sync()
}
