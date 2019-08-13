package main

import (
	"fmt"
	"go-storage/commands/application"
	"os"
)

func main() {
	err := application.Run()
	if err != nil {
		fmt.Println("")
		os.Exit(-1)
	}
}
