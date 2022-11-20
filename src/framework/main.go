package framework

import (
	"fmt"
	"go/importer"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func Main() {
	runCommand("d", [2]string{"hi", "hello"})
}

func runCommand(cmd string, args [2]string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(fmt.Errorf("%v directory does not exist", err))
	}
	commandFiles, matchErr := filepath.Glob(fmt.Sprintf("%s/commands/*.go", dir))

	if matchErr != nil {
		fmt.Println("Failed to load commands.")
	}

	println(fmt.Sprint(commandFiles), dir)

	for _, file := range commandFiles {
		println(file)
		imported, err := importer.Default().Import("flightpkg/commands") // does not work

		if err != nil {
			fmt.Println("Failed to load a command.", err)
		}

		println(imported.Imports())
	}
}

func getCommandName(filePath string) string {
	if runtime.GOOS == "windows" {
		s := strings.Split(filePath, "\\")

		return strings.Split(s[len(s)-1], ".go")[0]
	}

	return "UNKNOWN"
}
