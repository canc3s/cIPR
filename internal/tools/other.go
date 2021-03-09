package tools

import (
	"net"
	"sort"
)

func RemoveDuplicateIP(ips []net.IP) []string {
	result := make([]string, 0, len(ips))
	temp := map[string]struct{}{}
	for _, ip := range ips {
		item := ip.String()
		if _, ok := temp[string(item)]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}


func in(target string, str_array []string) bool {
	sort.Strings(str_array)
	index := sort.SearchStrings(str_array, target)
	//index的取值：[0,len(str_array)]
	if index < len(str_array) && str_array[index] == target { //需要注意此处的判断，先判断 &&左侧的条件，如果不满足则结束此处判断，不会再进行右侧的判断
		return true
	}
	return false
}