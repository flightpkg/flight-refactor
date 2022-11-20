package main

import (
	funcs "flightpkg/func"
	"fmt"
	"os"
)

func main() {
	cli := os.Args[1:]
	if len(cli) == 0 {
		fmt.Println("flight <command> [arguments]")
	} else {
		if cli[0] == "-h" || cli[0] == "--help" {
			funcs.Help()
		} else if cli[0] == "install" || cli[0] == "i" || cli[0] == "add" || cli[0] == "a" || cli[0] == "get" || cli[0] == "g" {
			funcs.Install(cli[1:])
		} else if cli[0] == "uninstall" || cli[0] == "u" || cli[0] == "remove" || cli[0] == "r" || cli[0] == "ui" {
			funcs.Uninstall(cli[1:])
		} else if cli[0] == "status" || cli[0] == "s" || cli[0] == "st" {
			funcs.Status()
		} else if cli[0] == "update" || cli[0] == "up" || cli[0] == "u" {
			funcs.Update()
		} else {
			fmt.Println("flight <command> [arguments]")
		}
	}
}
