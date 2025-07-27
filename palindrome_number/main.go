package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(1221))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reverse := 0
	xcopy := x

	for x > 0 {
		reverse = (reverse * 10) + (x % 10)
		x = x / 10
	}
	return reverse == xcopy
}

// reverse digits technique
// You need mod 10 and divide by 10 because we're working in base 10 (decimal)
// Number: 1234
// Step 1: Take 4 (1234 % 10), remove it (1234/10 = 123), reverse = 4
// Step 2: Take 3 (123 % 10), remove it (123/10 = 12), reverse = 4*10 + 3 = 43
// Step 3: Take 2 (12 % 10), remove it (12/10 = 1), reverse = 43*10 + 2 = 432
// Step 4: Take 1 (1 % 10), remove it (1/10 = 0), reverse = 432*10 + 1 = 4321
