package main

import (
	"math"
	"sort"
)

func main() {
	strs := []string{"f1lowers", "flow", "flight"}
	// first, last := strs[0], strs[len(strs)-1]

	// println(int(math.Min(float64(len(first)), float64(len(last)))))

	println(findCommonPrefix(strs))
	println(findCommonPrefix([]string{"dog", "racecar", "car"}))

}

func findCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	sort.Strings(strs)

	first, last := strs[0], strs[len(strs)-1]
	result := ""

	for i := 0; i < int(math.Min(float64(len(first)), float64(len(last)))); i++ {
		if first[i] == last[i] {
			result += string(first[i])
		} else {
			break
		}
	}

	return result
}
