package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Host missing\n")
		os.Exit(1)
	}
	domain := os.Args[1]
	daysleft, err := checkssl(domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(2)
	}
	fmt.Println(daysleft)
}

func checkssl(domain string) (int64, error) {
	conn, err := tls.Dial("tcp", domain+":443", nil)
	if err != nil {
		return -1, err
	}
	defer conn.Close()
	now := time.Now().Unix()
	var min int64 = 1000000
	for _, chain := range conn.ConnectionState().VerifiedChains {
		for _, cert := range chain {
			for _, name := range cert.DNSNames {
				if s, e := filepath.Match("/"+name, "/"+domain); s && e == nil {
					left := cert.NotAfter.Unix() - now
					daysleft := (left / 60 / 60 / 24)
					if daysleft < min {
						min = daysleft
					}
				}
			}
		}
	}
	if min >= 1000000 {
		return -2, errors.New("No valid Expiration found\n")
	} else {
		return min, nil
	}
}
