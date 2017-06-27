package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

func main() {

	if len(os.Args) == 1 {
		panic("must pass search target as arg")
	}

	target := os.Args[1]

	if rand.Intn(100) == 100 {
		resp, _ := http.Get("http://go-search.org/api?action=packages")

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		_ = ioutil.WriteFile("packages.data", body, 0644)
	}

	packages, _ := ioutil.ReadFile("packages.data")
	data := strings.Split(string(packages), ",")

	for _, item := range data {
		if strings.Contains(item, target) {
			fmt.Println(item)
		}
	}

}
