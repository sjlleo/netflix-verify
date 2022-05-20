package util

import (
	"errors"
	"net"
)

func DnsResolver(addrType int) (string, error) {
	dns := "www.netflix.com"

	if ns, err := net.LookupHost(dns); err != nil {
		return "", err
	} else {
		switch {
		case len(ns) != 0:
			for _, n := range ns {
				if ParseIP(n) == 4 && addrType == 4 {
					return n, nil
				}
				if ParseIP(n) == 6 && addrType == 6 {
					return "[" + n + "]", nil
				}
			}
		}
	}
	return "", errors.New("no IP Address")
}
