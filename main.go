package main

import (
	"fmt"

	"github.com/yurichandra/gunners/cmd"
)

func main() {
	fmt.Println("Welcome to gunners: A Livescore API")

	cmd.Serve()
}
