package main

import "testing"

func TestFinAllIpsByService(t *testing.T) {
	ips, err := findAllIpsByService("google.com")
	if err != nil {
		t.Error(err)
	} else {
		for _, ip := range ips {
			t.Log(ip)
		}
	}
}
