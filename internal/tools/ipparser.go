package tools

import (
	"net"
)

func ValidIP(IP string) bool {
	if IP == "127.0.0.1" {
		return false
	}
	for i := 0; i < len(IP); i++ {
		if IP[i] == '.' {
			if ip := net.ParseIP(IP); ip != nil {
				return true
			}
		}
	}
	return false
}