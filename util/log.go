package util

import (
	"log"
)

var (
	Info *log.Logger
	Warning *log.Logger
	Error *log.Logger
)

func init() {
	// logfile, err := os.OpenFile(".log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	// if err!=nil {
	// 	log.Fatalln("打开日志文件失败：",err)
	// }
	// Info = log.New(logfile,"Info:",log.LstdFlags|log.Llongfile)
	// Warning = log.New(logfile,"Warning:",log.LstdFlags|log.Llongfile)
	// Error = log.New(logfile,"Error:",log.LstdFlags|log.Llongfile)
}