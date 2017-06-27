package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	packages, _ := ioutil.ReadFile("packages.data")
	data := strings.Split(string(packages), ",")

	for _, item := range data {
		if strings.Contains(item, os.Args[1]) {
			fmt.Println(item)
		}
	}

}
