package app

import (
	"github.com/canc3s/cIPR/internal/tools"
	"github.com/canc3s/cIPR/pkg/qqwry"
	"path/filepath"
)

type IPDB interface {
	Find(ip string) string
}

var (
	db    IPDB
	qqip  qqwry.QQwry
)

// init ip db content
func InitIPDB(datPath string) {
	db = qqwry.NewQQwry(filepath.Join(datPath))
}

// parse several ips
func ParseIP(ip string) string {
	db := db

	if tools.ValidIP(ip) {
		result := db.Find(ip)
		return result
	}else{
		return ""
	}
}