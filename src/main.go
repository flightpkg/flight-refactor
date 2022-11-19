package main

import (
	funcs "flightpkg/func"
	"fmt"
	"os"
)

func main() {
	cli := os.Args[1:]
	funcs.Figlet()
	if len(cli) == 0 {
		fmt.Println("flight <command> [arguments]")
	} else {
		if cli[0] == "help" {
			funcs.Help()
		} else if cli[0] == "install" {
			funcs.Install(cli[1:])
		} else {
			fmt.Println("flight <command> [arguments]")
		}
	}
}
