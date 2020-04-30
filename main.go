package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jeandias/monitoring/cmd"
)

const turns = 2
const delay = 2

func main() {
	for {
		option, err := cmd.GetOption(os.Stdin)
		throwError(err)

		switch option {
		case 1:
			startMonitoring()
		case 2:
			printLog()
		case 0:
			fmt.Println("Exit program")
			os.Exit(0)
		default:
			fmt.Println("Option not found")
			os.Exit(-1)
		}
	}
}

func readFile() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	throwError(err)

	reader := bufio.NewReader(file)
	for {
		row, err := reader.ReadString('\n')
		row = strings.TrimSpace(row)
		sites = append(sites, row)

		if err == io.EOF {
			break
		}
	}

	file.Close()
	return sites
}

func writeLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	throwError(err)

	file.WriteString(time.Now().Format("2006-01-02 15:04:05") + " | " + site + " | online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func testSite(site string) {
	resp, err := http.Get(site)
	throwError(err)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was successfully loaded!")
		writeLog(site, true)
	} else {
		fmt.Println("Site:", site, "it has a problem. Status Code:", resp.StatusCode)
		writeLog(site, false)
	}
}

func startMonitoring() {
	sites := readFile()
	for t := 0; t < turns; t++ {
		for i, site := range sites {
			fmt.Println("Row", i, "test:", site)
			testSite(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func printLog() {
	file, err := ioutil.ReadFile("log.txt")
	throwError(err)
	fmt.Println(string(file))
}

func throwError(err error) {
	if err != nil {
		fmt.Printf("Something went wrong: %v", err)
	}
}
