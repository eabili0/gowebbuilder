package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Started running...")
	fmt.Println("Hello World!")
	logrus.Info("Done!")
}
