package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {

	localAddresses()

}

func localAddresses() {
	ifaces, err := net.Interfaces()
	if err != nil {
		log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
		return
	}
	for _, i := range ifaces {

		if i.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}

		addrs, err := i.Addrs()

		if err != nil {
			log.Print(fmt.Errorf("localAddresses: %v\n", err.Error()))
			continue
		}

		for _, a := range addrs {
			log.Printf("Scanning %+v\n", a)

			hosts, _ := Hosts(a.String())

			for _, ip := range hosts {

				ip += ":22"

				conn, err := net.Dial("tcp", ip)
				if err != nil {
					log.Print(fmt.Errorf("%v", err.Error()))
					continue
				}

				log.Printf("%v ", ip)

				res, _ := bufio.NewReader(conn).ReadString('\n')

				fmt.Print(res)

			}
		}
	}
}

func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
