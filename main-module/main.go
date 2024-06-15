package main

import (
	"fmt"

	"sample.com/exec-module/mini-exec"
	"sample.com/greet-module"
	"sample.com/main-module/one"
)

func main() {
	fmt.Println("Welcome to main modile")
	greet.Greets()
	mini.CheckMini()
	one.One()

}
