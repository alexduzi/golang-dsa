package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	file, err := os.Open("./palindrome/words.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	wordCounter := 0
	palindromeCounter := 0
	isPalin := false

	for scanner.Scan() {
		word := scanner.Text()
		wordCounter++
		isPalin = isPalindrome(word)
		if isPalin {
			palindromeCounter++
		}

		fmt.Printf("IsPalindrome:\t%s\t%v\n", string(word), isPalin)
	}
	end := time.Now()

	fmt.Printf("Finished in: %v palindrome words in file: %d from %d\n", end.Sub(start).Microseconds(), palindromeCounter, wordCounter)
}

func isPalindrome(text string) bool {
	if text == "" {
		return false
	}
	text = strings.ToLower(strings.TrimSpace(text))
	aux := len(text) - 1
	for i := 0; i < len(text); i++ {
		if i == aux && text[i] == text[aux] {
			break
		}
		if text[i] != text[aux] {
			return false
		}
		aux--
	}
	return true
}
