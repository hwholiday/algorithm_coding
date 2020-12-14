package leet_code

import (
	"strings"
	"testing"
)

//山羊拉丁文
func toGoatLatin(S string) string {
	srtArray := strings.Split(S, " ")
	var result []string
	for i, v := range srtArray {
		var d string
		if isVowel(string(v[0])) {
			d = v + "ma"
		} else {
			d = v[1:] + string(v[0]) + "ma"
		}
		var a string
		for j := 0; j <= i; j++ {
			a += "a"
		}
		d += a
		result = append(result, d)
	}
	return strings.Join(result, " ")
}

func isVowel(s string) bool {
	s = strings.ToLower(s)
	var vowel = []string{"a", "e", "i", "o", "u"}
	for _, v := range vowel {
		if v == s {
			return true
		}
	}
	return false
}

func Test_toGoatLatin(t *testing.T) {
	t.Log(toGoatLatin("I speak Goat Latin"))
}
