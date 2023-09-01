package inits

import (
	"os"

	"github.com/sirupsen/logrus"
)

func InitLogrus() {

	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000 Z07:00",
	})
	logrus.SetReportCaller(true)
}
