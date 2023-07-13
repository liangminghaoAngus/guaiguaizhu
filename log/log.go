package log

import (
	"github.com/robfig/cron/v3"
	"io"
	"io/fs"
	"liangminghaoangus/guaiguaizhu/config"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	debug         *log.Logger
	info          *log.Logger
	warn          *log.Logger
	errorLog      *log.Logger
	dayChangeLock sync.RWMutex
)

const (
	debugLevel = iota //iota=0
	infoLevel
	warnLevel
	errorLevel
)

func createLogFile() {
	dayChangeLock.Lock()
	defer dayChangeLock.Unlock()
	now := time.Now()
	postFix := now.Format("2006-01-02")
	logFile := postFix + ".log"
	logFilePath := path.Join("log/files", logFile)
	if _, err := os.Stat("log/files"); err != nil && !os.IsExist(err) {
		err = os.MkdirAll("log/files", fs.ModePerm)
	}
	logOut, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	} else {
		multiWriter := io.MultiWriter(os.Stdout, logOut)
		debug = log.New(multiWriter, "[DEBUG] ", log.Ldate|log.Ltime)
		info = log.New(multiWriter, "[INFO] ", log.Ldate|log.Ltime)
		warn = log.New(multiWriter, "[WARN] ", log.Ldate|log.Ltime)
		errorLog = log.New(multiWriter, "[ERROR] ", log.Ldate|log.Ltime)
	}
}

func Debug(format string, v ...any) {
	if config.GetConfig().LogLevel <= debugLevel {
		debug.Printf(getLineNo()+format, v...)
	}
}

func Info(format string, v ...any) {
	if config.GetConfig().LogLevel <= infoLevel {
		info.Printf(getLineNo()+format, v...)
	}
}

func Warn(format string, v ...any) {
	if config.GetConfig().LogLevel <= warnLevel {
		warn.Printf(getLineNo()+format, v...)
	}
}

func Error(format string, v ...any) {
	if config.GetConfig().LogLevel <= errorLevel {
		errorLog.Printf(getLineNo()+format, v...)
	}
}

func getLineNo() string {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		split := strings.Split(file, "/")
		file = split[len(split)-1]
		fileLine := file + ":" + strconv.Itoa(line) + " "
		return fileLine
	}
	return ""
}

func logJob() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("@daily", func() {
		Info("执行log定时任务")
		//now := time.Now()
		createLogFile()
		//closeYesterdayLogFile := fmt.Sprintf("plume_log_%s.log", now.Add(-24*time.Hour).Format("20060102"))
		//file, _ := os.Open(closeYesterdayLogFile)
		//_ = file.Sync()
		//_ = file.Close()
	})
	c.Start()
}

func Init() {
	dayChangeLock = sync.RWMutex{}
	createLogFile()
	go logJob()
}
