package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func Test_Logger(t *testing.T) {
	fname := filepath.Base(os.Args[0])
	fmt.Println(fname)

	log := NewLogger()
	defer log.Close()
	log.InitFileLog(filepath.Dir(os.Args[0]), fname, "", true)
	log.FileLogMode(true, true, true)

	log.Info("This is INFO-1")
	log.Infof("This is INFO-%d", 2)
	log.Infom("Method-1", "This is INFO-%d", 2)

	time.Sleep(time.Second * 15)

	log.Debug("This is INFO-1")
	log.Debugf("This is INFO-%d", 2)

	time.Sleep(time.Second * 15)

	log.Error("This is INFO-1")
	log.Errorf("This is INFO-%d", 2)

	time.Sleep(time.Second * 30)
}

func Test_LoggerVar(t *testing.T) {
	log := NewLogger()
	defer log.Close()
	err := log.InitFileLog("/var/log", "logtest", "", true)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.FileLogMode(true, true, true)

	log.Info("This is INFO-1")
	log.Infof("This is INFO-%d", 2)
	log.Infom("Method-1", "This is INFO-%d", 2)

	log.Debug("This is INFO-1")
	log.Debugf("This is INFO-%d", 2)

	log.Error("This is INFO-1")
	log.Errorf("This is INFO-%d", 2)

}
