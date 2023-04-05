package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

var mac string

func main() {
	flag.StringVar(&mac, "mac", "xx:xx:xx:xx", "please enter your mac")
	flag.Parse()

	formattedMac, err := FormatMacAddress(mac)
	if err != nil {
		log.Fatalf("Invalid MAC address format: %s. The correct format is xx:xx:xx:xx:xx:xx or xxxxxxxxxxxx", mac)
	}
	log.Println(formattedMac)
}

func FormatMacAddress(mac string) (string, error) {
	mac = strings.ToLower(mac)
	mac = strings.ReplaceAll(mac, "-", "")
	mac = strings.ReplaceAll(mac, ":", "")
	if len(mac) != 12 {
		return "", fmt.Errorf("Invalid MAC address length: %d", len(mac))
	}

	var b strings.Builder
	for i := 0; i < 12; i += 2 {
		if i > 0 {
			b.WriteByte(':')
		}
		b.WriteString(mac[i : i+2])
	}

	return b.String(), nil
}
