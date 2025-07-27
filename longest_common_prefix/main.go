package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"flower", "flow", "flight"}
	sort.Strings(words)
	fmt.Println(words)
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	sort.Strings(strs)
	first, last := strs[0], strs[len(strs)-1]
	result := ""
	for i := 0; i < len(first); i++ {
		if i < len(last) && first[i] == last[i] {
			result += string(first[i])
		} else {
			break
		}
	}
	return result
}
