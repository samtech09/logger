package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

//Logger is object responsible for logging
type Logger struct {
	fileLogger   *log.Logger
	stdOutLogger *log.Logger
	logToFile    bool

	logInfoStd  bool
	logDebugStd bool
	logErrorStd bool

	logInfoFile  bool
	logDebugFile bool
	logErrorFile bool

	file *os.File
}

//NewLogger create new instance of Logger
func NewLogger() Logger {
	l := Logger{}
	l.logInfoStd = true
	l.logDebugStd = false
	l.logErrorStd = true

	l.logInfoFile = false
	l.logDebugFile = false
	l.logErrorFile = true
	l.stdOutLogger = log.New(os.Stdout, "", log.LstdFlags)
	return l
}

//StdOutLogMode sets what to log on screen
//Default is Info=true, debug=false, error=true
func (l *Logger) StdOutLogMode(loginfo, logdebug, logerror bool) {
	l.logInfoStd = loginfo
	l.logDebugStd = logdebug
	l.logErrorStd = logerror
}

//FileLogMode sets what to log in file
//Default is Info=false, debug=false, error=true
func (l *Logger) FileLogMode(loginfo, logdebug, logerror bool) {
	l.logInfoFile = loginfo
	l.logDebugFile = logdebug
	l.logErrorFile = logerror
}

//InitFileLog initialize logging by setting logfile
//It returns *os.File, it is responsibility of caller to close file
//using defer f.Close()
func (l *Logger) InitFileLog(pathToFolder, prefix string, suffix string, appendDate bool) error {
	filename := ""
	if appendDate {
		filename = "_" + time.Now().Local().Format("2006-01-02")
	}
	filename = prefix + strings.Replace(filename, ".", "", 1) + suffix + ".log"
	fullpath := filepath.Join(pathToFolder, filename)
	f, err := os.OpenFile(fullpath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	l.fileLogger = log.New(f, "", log.LstdFlags)
	l.logToFile = true
	l.file = f

	log.Printf("Logging to file: %s", fullpath)

	return nil
}

//Close closed the log file
func (l *Logger) Close() {
	// ensure buffers are written
	l.file.Sync()

	l.file.Close()
}

//Info log informational message
func (l *Logger) Info(msg ...string) {
	if l.logToFile && l.logInfoFile {
		l.fileLogger.Println("INFO: ", msg)
	}
	if l.logInfoStd {
		l.stdOutLogger.Println("INFO: ", msg)
	}
}

//Infof log informational message along with formating
func (l *Logger) Infof(format string, msg ...interface{}) {
	if l.logToFile && l.logInfoFile {
		l.fileLogger.Println("INFO: ", fmt.Sprintf(format, msg...))
	}
	if l.logInfoStd {
		l.stdOutLogger.Println("INFO: ", fmt.Sprintf(format, msg...))
	}
}

//Infom log informational message along with methodname
func (l *Logger) Infom(methodname, format string, msg ...interface{}) {
	if l.logToFile && l.logInfoFile {
		//l.fileLogger.Println("INFO: " + fmt.Sprintf("[%s] "+format, methodname, msg))
		l.fileLogger.Printf("INFO: [%s] [%s]\n", methodname, fmt.Sprintf(format, msg...))
	}
	if l.logInfoStd {
		//l.stdOutLogger.Println("INFO: " + fmt.Sprintf("[%s] "+format, methodname, msg))
		l.stdOutLogger.Printf("INFO: [%s] [%s]\n", methodname, fmt.Sprintf(format, msg...))
	}
}

//Debug log debug message
func (l *Logger) Debug(msg ...string) {
	if l.logToFile && l.logDebugFile {
		l.fileLogger.Println("DEBUG: ", msg)
	}
	if l.logDebugStd {
		l.stdOutLogger.Println("DEBUG: ", msg)
	}
}

//Debugf log debug message along with formatting
func (l *Logger) Debugf(format string, msg ...interface{}) {
	if l.logToFile && l.logDebugFile {
		l.fileLogger.Println("DEBUG: ", fmt.Sprintf(format, msg...))
	}
	if l.logDebugStd {
		l.stdOutLogger.Println("DEBUG: ", fmt.Sprintf(format, msg...))
	}
}

//Debugm log debug message along with methodname
func (l *Logger) Debugm(methodname, format string, msg ...interface{}) {
	if l.logToFile && l.logDebugFile {
		//l.fileLogger.Println("DEBUG: " + fmt.Sprintf("[%s] "+format, methodname, msg))
		l.fileLogger.Printf("DEBUG: [%s] [%s]\n", methodname, fmt.Sprintf(format, msg...))
	}
	if l.logDebugStd {
		//l.stdOutLogger.Println("DEBUG: " + fmt.Sprintf("[%s] "+format, methodname, msg))
		l.stdOutLogger.Printf("DEBUG: [%s] [%s]\n", methodname, fmt.Sprintf(format, msg...))
	}
}

func (l *Logger) Error(err ...string) {
	if l.logToFile && l.logErrorFile {
		l.fileLogger.Println("ERROR: ", err)
	}
	if l.logErrorStd {
		l.stdOutLogger.Println("ERROR: ", err)
	}
}

//Errorf logs error message along with formatting
func (l *Logger) Errorf(format string, msg ...interface{}) {
	if l.logToFile && l.logErrorFile {
		l.fileLogger.Println("ERROR: ", fmt.Sprintf(format, msg...))
	}
	if l.logErrorStd {
		l.stdOutLogger.Println("ERROR: ", fmt.Sprintf(format, msg...))
	}
}

//Errorm log error message along with methodname
func (l *Logger) Errorm(methodname, format string, msg ...interface{}) {
	if l.logToFile && l.logErrorFile {
		//l.fileLogger.Println("ERROR: " + fmt.Sprintf("[%s] "+format, methodname, msg))
		l.fileLogger.Printf("ERROR: [%s] [%s]\n", methodname, fmt.Sprintf(format, msg...))
	}
	if l.logErrorStd {
		//l.stdOutLogger.Println("ERROR: " + fmt.Sprintf("[%s] "+format, methodname, msg))
		l.stdOutLogger.Printf("ERROR: [%s] [%s]\n", methodname, fmt.Sprintf(format, msg...))
	}
}

//Fatal log fatam message
func (l *Logger) Fatal(msg ...string) {
	if l.logToFile {
		l.fileLogger.Println("FATAL: ", msg)
	}
	l.stdOutLogger.Fatal("FATAL: ", msg)
}

//Fatalf log fatam message along with formatting
func (l *Logger) Fatalf(format string, msg ...interface{}) {
	if l.logToFile {
		l.fileLogger.Println("FATAL: ", fmt.Sprintf(format, msg...))
	}
	l.stdOutLogger.Fatal("FATAL: ", fmt.Sprintf(format, msg...))
}

//Fatalm log fatam message along with methodname
func (l *Logger) Fatalm(methodname, format string, msg ...interface{}) {
	if l.logToFile {
		l.fileLogger.Printf("FATAL: [%s] [%s]\n", methodname, fmt.Sprintf(format, msg...))
	}
	l.stdOutLogger.Fatalf("FATAL: [%s] [%s]\n", methodname, fmt.Sprintf(format, msg...))
}

//InitMsg logs messages that are out of category
// like when initializing app, it may call Direct to log app startup messages
func (l *Logger) InitMsg(msg ...string) {
	if l.logToFile {
		l.fileLogger.Println("INIT: ", msg)
	}
	l.stdOutLogger.Println("INIT: ", msg)
}

//InitMsgf - init messages with formatting
func (l *Logger) InitMsgf(format string, msg ...interface{}) {
	if l.logToFile {
		l.fileLogger.Println("INIT: ", fmt.Sprintf(format, msg...))
	}
	l.stdOutLogger.Println("INIT: ", fmt.Sprintf(format, msg...))
}
