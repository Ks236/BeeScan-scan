package util

import (
	"github.com/projectdiscovery/mapcidr"
	"strings"
)

/*
创建人员：云深不知处
创建时间：2022/1/6
程序功能：目标处理
*/
//https://www.hebut.edu.cn

// TargetsHandle 目标处理函数
func TargetsHandle(target string) []string {
	var targets []string
	if strings.Contains(target, "https") {
		ip := target[12:]
		targets = append(targets, ip)

	} else if strings.Contains(target, "http") {
		ip := target[11:]
		targets = append(targets, ip)
	} else if (strings.Contains(target, "com") || strings.Contains(target, "cn")) && strings.Contains(target, "www.") {
		ip := target[4:]
		targets = append(targets, ip)
	} else if strings.Contains(target, "/") {
		ips, _ := mapcidr.IPAddresses(target)
		targets = ips
	} else {
		targets = append(targets, target)
	}
	return targets
}
