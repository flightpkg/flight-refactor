package flightpkg

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	grab "github.com/cavaliergopher/grab/v3"
	targz "github.com/walle/targz"
)

func Install(args []string) {
	registry := "https://registry.yarnpkg.com/"
	// registry2 := "https://registry.npmmirror.com/"
	if len(args) != 1 {
		fmt.Println("flight install <pkg>")
	} else {

		resp, err := http.Get(fmt.Sprintf(registry+"%v", args[0]))
		if err != nil {
			fmt.Println(err)
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println(err)
		}

		/* Getting the version of the package that is being installed. */
		var data map[string]interface{}
		json.Unmarshal(body, &data)

		latest := data["dist-tags"].(map[string]interface{})["latest"].(string)
		tarball := data["versions"].(map[string]interface{})[latest].(map[string]interface{})["dist"].(map[string]interface{})["tarball"].(string)

		os.Mkdir(".flight", 0777)
		os.Mkdir(".flight\\"+args[0], 0777)
		_, err = grab.Get(".flight\\"+args[0], tarball)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Installed " + args[0] + " @ " + latest)

		targz.Extract(fmt.Sprintf("./.flight/%v/%v-%v.tgz", args[0], args[0], latest), "./node_modules")
		os.RemoveAll(".flight")
		os.Rename(fmt.Sprintf("./node_modules/package"), fmt.Sprintf("./node_modules/%v", args[0]))
		fmt.Println("Extracted " + args[0] + " @ " + latest)
	}
}

func Uninstall(args []string) {
	if len(args) != 1 {
		fmt.Println("flight uninstall <pkg>")
	} else {
		os.RemoveAll(fmt.Sprintf("./node_modules/%v", args[0]))
		fmt.Println("Uninstalled " + args[0])

		files, _ := ioutil.ReadDir("./node_modules")
		if len(files) == 0 {
			os.Remove("./node_modules")
		}
	}
}

func Status() {
	version := "2.0.3"
	version_dotless, _ := strconv.ParseInt(strings.ReplaceAll(version, ".", ""), 10, 64)
	latest := "https://api.github.com/repos/flightpkg/flight-v2/releases/latest"
	fmt.Printf("Version: %v\n", version)
	resp, err := http.Get(latest)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	/* Getting the version of the package that is being installed. */
	var data map[string]interface{}
	json.Unmarshal(body, &data)

	latest_tag := data["tag_name"].(string)
	latest_tag_dotless, _ := strconv.ParseInt(strings.Replace(latest_tag, ".", "", -1), 10, 64)

	if int(version_dotless) <= int(latest_tag_dotless) {
		fmt.Printf("Update Available:\n %v -> %v", version, latest_tag)
	} else {
		fmt.Println("Up to date!")
	}
}

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
flight update
flight status`)
}
