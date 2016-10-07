package webwx

import (
	"fmt"
	"github.com/cihub/seelog"
)

/*
 * global log object
 */
var log = &mLogger{nil, seelog.Off}

/*
 * Custom Logger
 *
 * 1. use "fmt.Sprintln" underline
 * 2. use a min-max dispatcher
 * 3. can adjust loglevel
 */
type mLogger struct {
	innerLogger seelog.LoggerInterface
	logLevel    seelog.LogLevel
}

func (l *mLogger) Trace(params ...interface{}) {
	if l.logLevel != seelog.Off && l.logLevel <= seelog.TraceLvl && l.innerLogger != nil {
		s := fmt.Sprintln(params...)
		l.innerLogger.Trace(s)
	}
}
func (l *mLogger) Debug(params ...interface{}) {
	if l.logLevel != seelog.Off && l.logLevel <= seelog.DebugLvl && l.innerLogger != nil {
		s := fmt.Sprintln(params...)
		l.innerLogger.Debug(s)
	}
}
func (l *mLogger) Info(params ...interface{}) {
	if l.logLevel != seelog.Off && l.logLevel <= seelog.InfoLvl && l.innerLogger != nil {
		s := fmt.Sprintln(params...)
		l.innerLogger.Info(s)
	}
}
func (l *mLogger) Warn(params ...interface{}) {
	if l.logLevel != seelog.Off && l.logLevel <= seelog.WarnLvl && l.innerLogger != nil {
		s := fmt.Sprintln(params...)
		l.innerLogger.Warn(s)
	}
}
func (l *mLogger) Error(params ...interface{}) {
	if l.logLevel != seelog.Off && l.logLevel <= seelog.ErrorLvl && l.innerLogger != nil {
		s := fmt.Sprintln(params...)
		l.innerLogger.Error(s)
	}
}
func (l *mLogger) Critical(params ...interface{}) {
	if l.logLevel != seelog.Off && l.logLevel <= seelog.CriticalLvl && l.innerLogger != nil {
		s := fmt.Sprintln(params...)
		l.innerLogger.Critical(s)
	}
}
func (l *mLogger) Tracef(format string, params ...interface{}) {
	if l.logLevel != seelog.Off && l.logLevel <= seelog.TraceLvl && l.innerLogger != nil {
		l.innerLogger.Tracef(format+"\n", params...)
	}
}
func (l *mLogger) Debugf(format string, params ...interface{}) {
	if l.logLevel != seelog.Off && l.logLevel <= seelog.DebugLvl && l.innerLogger != nil {
		l.innerLogger.Debugf(format+"\n", params...)
	}
}
func (l *mLogger) Infof(format string, params ...interface{}) {
	if l.logLevel != seelog.Off && l.logLevel <= seelog.InfoLvl && l.innerLogger != nil {
		l.innerLogger.Infof(format+"\n", params...)
	}
}
func (l *mLogger) Warnf(format string, params ...interface{}) {
	if l.logLevel != seelog.Off && l.logLevel <= seelog.WarnLvl && l.innerLogger != nil {
		l.innerLogger.Warnf(format+"\n", params...)
	}
}
func (l *mLogger) Errorf(format string, params ...interface{}) {
	if l.logLevel != seelog.Off && l.logLevel <= seelog.ErrorLvl && l.innerLogger != nil {
		l.innerLogger.Errorf(format+"\n", params...)
	}
}
func (l *mLogger) Criticalf(format string, params ...interface{}) {
	if l.logLevel != seelog.Off && l.logLevel <= seelog.CriticalLvl && l.innerLogger != nil {
		l.innerLogger.Criticalf(format+"\n", params...)
	}
}
func (l *mLogger) Flush() {
	if l.innerLogger != nil {
		l.innerLogger.Flush()
	}
}

func (l *mLogger) Close() {
	if l.innerLogger != nil {
		l.innerLogger.Close()
	}
}

func init() {
	formatter, err := seelog.NewFormatter(`[WebWx] %Date_%Time %LEVEL %Msg`)
	if err != nil {
		fmt.Println("[WebWx] init logger failed:", err)
		return
	}
	writer, _ := seelog.NewConsoleWriter()
	if err != nil {
		fmt.Println("[WebWx] init logger failed:", err)
		return
	}
	dispatcher, err := seelog.NewSplitDispatcher(formatter, []interface{}{writer})
	if err != nil {
		fmt.Println("[WebWx] init logger failed:", err)
		return
	}
	constrains, err := seelog.NewMinMaxConstraints(seelog.TraceLvl, seelog.CriticalLvl)
	if err != nil {
		fmt.Println("[WebWx] init logger failed:", err)
		return
	}
	config := seelog.NewLoggerConfig(
		constrains,
		make([]*seelog.LogLevelException, 0),
		dispatcher,
	)

	log.innerLogger = seelog.NewAsyncLoopLogger(config)
	log.logLevel = seelog.InfoLvl
}

func SetLogLevel(level string) bool {
	l, ok := seelog.LogLevelFromString(level)
	if !ok {
		fmt.Println("[WebWx] invalid log level:", level)
		return false
	}
	log.logLevel = l
	return true
}
