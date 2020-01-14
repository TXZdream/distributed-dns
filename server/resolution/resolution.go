package resolution

import (
	"strconv"
	"strings"
)

// 将域名解析为对应的ipv4地址
func FindIpAddrOfDomain(domainName string) (error, string) {
	return nil, "xxx.xxx.xxx.xxx"
}

// 已经过测试，可用
func Ipv4ToArr(ipAddr string) [4]int {
	ipAddr = strings.Split(ipAddr, "@")[0]
	var ipArr [4]int
	s := strings.Split(ipAddr, ".")
	for i := 0; i < 4; i++ {
		tmp, _ := strconv.Atoi(s[i])
		ipArr[i] = tmp
	}
	return ipArr
}
