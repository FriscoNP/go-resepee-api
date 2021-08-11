package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
}

func main() {
	fmt.Println("Hello World")
}
