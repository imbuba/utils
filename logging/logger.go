package logging

import (
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/finnan444/utils/time/cron"
)

// LoggersConsts
const (
	ErrorLogger = "errorLogger"
	StdLogger   = "logger"
)

var (
	logLocker     sync.RWMutex
	files         = make(map[string]*rotater)
	customLoggers = make(map[string]*log.Logger)
)

type rotater struct {
	sync.RWMutex
	filename string
	file     *os.File
}

// Write satisfies the io.Writer interface.
func (w *rotater) Write(output []byte) (int, error) {
	w.RLock()
	defer w.RUnlock()
	return w.file.Write(output)
}

// Perform the actual act of rotating and reopening file.
func (w *rotater) Rotate() (err error) {
	w.Lock()
	defer w.Unlock()

	// Close existing file if open
	if w.file != nil {
		err = w.file.Close()
		w.file = nil
		if err != nil {
			return
		}
	}

	// Rename dest file if it already exists
	_, err = os.Stat(w.filename)
	if err == nil {
		err = os.Rename(w.filename, w.filename+"."+time.Now().UTC().Format("2006-01-02-15-04"))
		if err != nil {
			return
		}
	}

	// Create a file.
	w.file, err = os.Create(w.filename)
	return
}

func (w *rotater) rotate() {
	w.Rotate()
}

func newRotater(filename string, period time.Duration) *rotater {
	result := &rotater{filename: filename}
	if err := result.Rotate(); err != nil {
		return nil
	}
	cron.Add(period, time.UTC, result.rotate)
	return result
}

// GetErrorLogger returns logger with stderr
func GetErrorLogger(name string, filenames ...string) (result *log.Logger) {
	var ok bool
	logLocker.RLock()
	if result, ok = customLoggers[name]; ok {
		logLocker.RUnlock()
		return
	}
	logLocker.RUnlock()
	var writers []io.Writer
	writers = append(writers, os.Stderr)
	var file *rotater
	logLocker.Lock()
	for _, fn := range filenames {
		if file, ok = files[fn]; !ok {
			file = newRotater(fn, 12*time.Hour)
			files[fn] = file
		}
		writers = append(writers, file)
	}
	result = log.New(io.MultiWriter(writers...), "", log.LstdFlags)
	customLoggers[name] = result
	logLocker.Unlock()
	return
}

// GetLogger returns logger with stdout
func GetLogger(name string, filenames ...string) (result *log.Logger) {
	var ok bool
	logLocker.RLock()
	if result, ok = customLoggers[name]; ok {
		logLocker.RUnlock()
		return
	}
	logLocker.RUnlock()
	var writers []io.Writer
	writers = append(writers, os.Stdout)
	var file *rotater
	logLocker.Lock()
	for _, fn := range filenames {
		if file, ok = files[fn]; !ok {
			file = newRotater(fn, 12*time.Hour)
			files[fn] = file
		}
		writers = append(writers, file)
	}
	result = log.New(io.MultiWriter(writers...), "", log.LstdFlags)
	customLoggers[name] = result
	logLocker.Unlock()
	return
}
