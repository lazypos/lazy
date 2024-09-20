package lazy

import (
	"fmt"
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

var DefaultLog = &LogUtil{}

type LogUtil struct {
	log.Logger
}

func (l *LogUtil) DefaultLogInfo(fileName string) {
	l.SetLogInfo(fileName, 20, 3, 0)
}
func (l *LogUtil) SetLogInfo(fileName string, maxMSize, saveCounts, flags int) {

	l.SetOutput(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxMSize,
		MaxBackups: saveCounts,
		MaxAge:     30,
		LocalTime:  true,
		Compress:   true,
	})

	if flags == 0 {
		l.SetFlags(log.LstdFlags)
	} else {
		l.SetFlags(flags)
	}
}

func (l *LogUtil) LogPrint(v ...any) {
	fmt.Println(v...)
	l.Println(v...)
}

func (l *LogUtil) Log(v ...any) {
	l.Println(v...)
}

func (l *LogUtil) Printf(s string, v ...interface{}) {
	l.Println(fmt.Sprintf(s, v...))
}
