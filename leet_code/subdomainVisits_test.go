package leet_code

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

//子域名访问计数
func subdomainVisits(cpdomains []string) []string {
	var cpdomainsMap = make(map[string]int)
	for _, v := range cpdomains {
		t := strings.Split(v, " ")
		num, _ := strconv.Atoi(t[0])
		splitT := strings.Split(t[1], ".")
		for i := 0; i < len(splitT); i++ {
			key := strings.Join(splitT[i:], ".")
			cpdomainsMap[key] = cpdomainsMap[key] + num
		}
	}
	var data []string
	for k, v := range cpdomainsMap {
		data = append(data, fmt.Sprintf("%d %s", v, k))
	}
	return data
}

func Test_subdomainVisits(t *testing.T) {
	t.Log(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}
