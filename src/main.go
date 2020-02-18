package main

import (
	"fmt"

	"github.com/even-c/learnGo/src/testlib"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println(testlib.Add(1, 2))
	log.Info("IThome Iron man")
}
