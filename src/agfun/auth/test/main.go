package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	// force colors on for TextFormatter
	formatter := &logrus.TextFormatter{
		ForceColors: true,
	}
	logrus.SetFormatter(formatter)
	// then wrap the log output with it
	//logrus.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
	logrus.Println("hello")
	logrus.WithFields(logrus.Fields{"Animal": "bird", "Size": 10}).Infoln("todo")
}
