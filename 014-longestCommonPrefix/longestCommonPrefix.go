package longestCommonPrefix

import (
	"sort"
)

func longestCommonPrefix(strs []string) string {
	result := ""
	for i := 0; ; i++ {
		for _, str := range strs {
			if i == len(str) || strs[0][i] != str[i] {
				return result
			}
		}
		result += string(strs[0][i])
	}
}

func longestCommonPrefixBrute(strs []string) string {
	result := ""
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	for i := 1; i <= len(strs[0]); i++ {
		counter := 0
		for _, str := range strs {
			if len(str) == 0 {
				return result
			}
			if strs[0][0:i] == str[0:i] {
				counter++
			}
		}
		if counter == len(strs) {
			result = strs[0][0:i]
		}
	}
	return result
}
