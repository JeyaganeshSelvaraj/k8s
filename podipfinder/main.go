package main

import (
	"flag"
	"log"
	"net"
	"os"
)

var serviceName string

func findAllIpsByService(service string) ([]string, error) {
	ips, err := net.LookupIP(service)
	if err != nil {
		return nil, err
	}
	var result []string
	for _, ip := range ips {
		result = append(result, ip.String())
	}
	return result, nil
}
func init() {
	flag.Parse()
	serviceName = flag.Arg(0)
}
func main() {
	if serviceName == "" {
		serviceName = os.Getenv("SERVICE_NAME")
	}
	if serviceName == "" {
		log.Fatal("Service name is empty")
	}
	if ips, err := findAllIpsByService(serviceName); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("IP Addresses for service: %s\n", serviceName)
		for _, ip := range ips {
			log.Println(ip)
		}
	}
	os.Exit(0)
}
