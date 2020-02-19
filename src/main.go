package main

import (
	"github.com/even-c/learnGo/router"
)

func main() {
	router := router.SetupRouter()
	router.Run()
}