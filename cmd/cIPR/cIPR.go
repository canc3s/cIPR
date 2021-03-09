package main

import (
	"fmt"
	"github.com/canc3s/cIPR/internal/app"
	"github.com/canc3s/cIPR/internal/tools"
	"net"
	"sort"
	"strings"
)


func main() {
	options := app.ParseOptions()
	scanner := tools.FileScaner(options.InputFile)
	var (
		ipRecords []net.IP
	)
	for scanner.Scan() {
		domain := strings.TrimSpace(scanner.Text())
		ipRecord, _ :=net.LookupIP(domain)
		fmt.Println(domain,"\t",ipRecord)
		ipRecords = append(ipRecords,ipRecord...)
	}
	ips := tools.RemoveDuplicateIP(ipRecords)

	app.InitIPDB(options.DatPath)
	ress := make(map[string]int)
	for _, ip := range ips {
		result := app.ParseIP(ip)
		flag := true
		for _, keyWord := range options.BlackList {
			if strings.Contains(result, keyWord) {
				flag = false
				break
			}
		}
		if flag && tools.ValidIP(ip) {
			buf := strings.Split(ip,".")
			ipd := strings.Join(buf[:3] , ".")
			ipd += ".0/24"
			ress[ipd]++
		}
	}

	type IPR struct {
		ip		string
		count	int
	}

	var lstIPR []IPR

	for k, v := range ress {
		lstIPR = append(lstIPR, IPR {k, v})
	}

	sort.Slice(lstIPR, func(i, j int) bool {
		return lstIPR[i].count > lstIPR[j].count  // 降序
	})

	for _, i := range lstIPR {
		fmt.Printf("%-18s||%4d\n",i.ip, i.count)
	}
}
