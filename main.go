package main

import (
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	logrus.Infoln("I'm alive")
	time.Sleep(time.Hour)
}
