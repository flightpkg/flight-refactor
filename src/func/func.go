package flightpkg

import (
	"bufio"
	"fmt"
	"os"
)

func Figlet() {
	// Open the file.
	f, _ := os.Open("..\\misc\\flight.txt")
	// Create a new Scanner for the file.
	scanner := bufio.NewScanner(f)
	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	fmt.Println()
}

func Help() {
	fmt.Println(`flight <command> [arguments]

flight help
flight version
flight install <pkg>
flight uninstall <pkg>
flight update`)
}
