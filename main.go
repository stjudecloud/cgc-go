package main

import (
	"github.com/sirupsen/logrus"
	"github.com/stjudecloud/cgc-go/cmd"
)

var cgcURL string = ""

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	cmd.Execute()
}
