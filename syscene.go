package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func main() {
	f, _ := ioutil.ReadFile("/proc/loadavg")
	s := strings.Split(string(f), " ")
	fmt.Printf("load avg %s", s[0])
	s1 := exec.Command("/bin/df", "-h", ".")
	r, _ := s1.Output()
	r1 := strings.Split(string(r), "/")
	r2 := strings.Split(r1[2], "  ")
	fmt.Printf("\n%s of %s used", r2[4], r2[1])
}
