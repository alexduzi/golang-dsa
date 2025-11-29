package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	// create a map with all symbols and respective values
	romanMap := make(map[byte]int)
	romanMap['I'] = 1
	romanMap['V'] = 5
	romanMap['X'] = 10
	romanMap['L'] = 50
	romanMap['C'] = 100
	romanMap['D'] = 500
	romanMap['M'] = 1000

	result := 0
	size := len(s)

	for i := 0; i < size; i++ {
		// if the current value is less than the next
		// subtract current from net and add to result
		if i+1 < size && romanMap[s[i]] < romanMap[s[i+1]] {
			result += romanMap[s[i+1]] - romanMap[s[i]]

			// skip to the next symbol
			i++
		} else {
			// if not, add the current value to result
			result += romanMap[s[i]]
		}
	}
	return result
}
