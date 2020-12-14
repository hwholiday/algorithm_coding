package leet_code

import (
	"strings"
	"testing"
)

//翻转单词顺序
func reverseWords(s string) string {
	data := strings.Split(s, " ")
	var dataTrim []string
	for i := 0; i < len(data); i++ {
		if len(data[i]) == 0 {
			continue
		}
		dataTrim = append(dataTrim, data[i])
	}
	for i, j := 0, len(dataTrim)-1; i < j; i, j = i+1, j-1 {
		dataTrim[i], dataTrim[j] = dataTrim[j], dataTrim[i]
	}
	return strings.Join(dataTrim, " ")
}

func Test_reverseWords(t *testing.T) {
	t.Log(reverseWords("a good   example      "))
}
