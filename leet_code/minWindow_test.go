package leet_code

import (
	"testing"
)

//最小覆盖子串
func minWindow(s string, t string) string {
	var (
		tMap = make(map[rune]int, len(t))
	)
	for _, v := range t {
		tMap[v]++
	}
	check := func(d string) bool {
		var (
			a = make(map[rune]int, len(d))
		)
		for _, v := range d {
			a[v]++
		}
		for k, v := range tMap {
			val, ok := a[k]
			if !ok || val < v {
				return false
			}
		}
		return true
	}
	var (
		end    = s
		isFind = false
	)
	for i, _ := range s {
		for j, _ := range s {
			if j >= i {
				temp := s[i : j+1]
				if check(temp) {
					isFind = true
					if len(end) > len(temp) {
						end = temp
					}
				}
			}
		}
	}
	if !isFind {
		return ""
	}
	return end
}

func Test_minWindow(t *testing.T) {
	t.Log(minWindow("qdsvbyuipqxnhkbgqdgozelvapgcainsofnrfjzvawihgmpwpwnqcylcnufqcsiqzwhhhjchfmqmagkrexigytklnrdslmkniuurdwzikrwlxhcbgkjegwsvnvpzhamrwgjzekjobizbreexqqewmwubtjadlowhwhiarurvcsyvwcunsylgwhisrivezrmvzwwsqppuhnreqmtkkgrjozbhjwlkpzgqwejotylamcgeqzobihmwinduloecqmtoqcejcpmqusukulncsbabodxbtbeloxzgbesdveupyocuzryutyxjdulzvpklokspqkslqodqfhlgajatkxfntqorhzcxlwmdigoyxtrcccidnlyxidnevdveczbpwpugyjhveyxhcfkpqipboehjhcombrdzhyghjncnnzwpggezrvcfzjqjngvoyyqhwwohlsvarrpzavatrcasnqbazyrzxhivfydsqasjtjiteloxposdhtfgswhrfpomnteftyonjyiojxnznfeubjctijmnyaanwgsphieqhpgsoutbbxycjaxrklekogakpsbwdimkxvelpyosvmxgnuxzgejwmjgbehxhpmtohzbyxqsvepbrmzsufcqrnwttfscxgxlpxnpufirjxtdjuvfzzvqprlizelwmkjchwtcdbvpbosminsjpncehnmgtzegknkrmdvrhrgihywsoobdedhltvtmxzuhmeaakysrpybyzxqnouqszzfswahtzbanidoubilsgoqfnjubdmvclaxkaedbfeppj", "fjknwevk"))
}
