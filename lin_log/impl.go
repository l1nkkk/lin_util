package lin_log

import (
	"io"
	"log"
	"os"
	"runtime"
	"sync"
)

var (
	once sync.Once
	ins  Logger
)

func Init(maxLevel Level, outList []io.Writer, flag *int) {
	once.Do(func() {
		callLevel := 3
		ins = NewLogger(maxLevel, outList, flag, &callLevel)
	})
}

func Debug(format string, a ...any) {
	ins.Debug(format, a...)
}
func Info(format string, a ...any) {
	ins.Info(format, a...)
}
func Warn(format string, a ...any) {
	ins.Warn(format, a...)
}
func Error(format string, a ...any) {
	ins.Error(format, a...)
}
func Fatal(format string, a ...any) {
	ins.Fatal(format, a...)
}

var (
	defaultFlag        = log.LstdFlags
	defaultOutput      = []io.Writer{os.Stderr}
	defaultCallerLevel = 2
)

func NewLogger(maxLevel Level, outList []io.Writer, flag *int, callerLevel *int) Logger {
	if len(outList) == 0 {
		outList = defaultOutput
	}
	if flag == nil {
		flag = &defaultFlag
	}
	if callerLevel == nil {
		callerLevel = &defaultCallerLevel
	}

	rtn := &logger{
		maxLevel:    maxLevel,
		callerLevel: *callerLevel,
	}
	rtn.levelToLogList = make(map[Level][]*log.Logger)
	for _, out := range outList {
		rtn.levelToLogList[DebugLevel] = append(rtn.levelToLogList[DebugLevel], log.New(out, DebugLevel.String(), *flag))
		rtn.levelToLogList[InfoLevel] = append(rtn.levelToLogList[InfoLevel], log.New(out, InfoLevel.String(), *flag))
		rtn.levelToLogList[WarnLevel] = append(rtn.levelToLogList[WarnLevel], log.New(out, WarnLevel.String(), *flag))
		rtn.levelToLogList[ErrorLevel] = append(rtn.levelToLogList[ErrorLevel], log.New(out, ErrorLevel.String(), *flag))
		rtn.levelToLogList[FatalLevel] = append(rtn.levelToLogList[FatalLevel], log.New(out, FatalLevel.String(), *flag))
	}
	return rtn
}

type logger struct {
	maxLevel       Level
	levelToLogList map[Level][]*log.Logger
	callerLevel    int
}

func (impl *logger) Debug(format string, a ...any) {
	impl.logf(DebugLevel, format, a...)
}
func (impl *logger) Info(format string, a ...any) {
	impl.logf(InfoLevel, format, a...)
}
func (impl *logger) Warn(format string, a ...any) {
	impl.logf(WarnLevel, format, a...)
}
func (impl *logger) Error(format string, a ...any) {
	impl.logf(ErrorLevel, format, a...)
}
func (impl *logger) Fatal(format string, a ...any) {
	impl.logf(FatalLevel, format, a...)
}

func (impl *logger) logf(l Level, format string, a ...any) {
	if l < impl.maxLevel {
		return
	}
	_, file, line, _ := runtime.Caller(impl.callerLevel)
	switch l {
	case DebugLevel, InfoLevel, WarnLevel, ErrorLevel:
		for _, log := range impl.levelToLogList[l] {
			log.Printf("%s:%d - "+format, append([]interface{}{file, line}, a...)...)
		}
	case FatalLevel:
		for _, log := range impl.levelToLogList[l] {
			log.Fatalf("%s:%d - "+format, append([]interface{}{file, line}, a...)...)
		}
	}
}

type (
	Level int32
)

const (
	DebugLevel Level = 1 // log.Println()
	InfoLevel  Level = 2 // log.Println()
	WarnLevel  Level = 3 // log.Println()
	ErrorLevel Level = 4 // log.Println()
	FatalLevel Level = 5 // log.Fatalln()
)

var levelString = map[Level]string{
	DebugLevel: "[Debug]\t",
	InfoLevel:  "[Info]\t",
	WarnLevel:  "[Warn]\t",
	ErrorLevel: "[Error]\t",
	FatalLevel: "[Fatal]\t",
}

func (l Level) String() string {
	return levelString[l]
}
