package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	domain := flag.String("domain", "", "source domain to check")
	checks := flag.Int("checks", 10, "numbers of checks to perform")
	flag.Parse()
	if *domain == "" {
		fmt.Println("Can't run with an empty domain.")
	}
	log.Printf("performing lookup for %s, with %d checks...", *domain, *checks)
	ipForHost := make(map[string]int)
	for i := 0; i < *checks; i++ {
		ips, err := net.LookupIP(*domain)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, currentIP := range ips {
			ipForHost[currentIP.String()] = ipForHost[currentIP.String()] + 1
		}
	}
	fmt.Printf("found %d addresses:\n", len(ipForHost))
	fmt.Println(ipForHost)
}
