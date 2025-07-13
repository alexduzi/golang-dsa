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
	text = strings.ToLower(strings.TrimSpace(text))
	for i, j := 0, len(text)-1; i < j; i, j = i+1, j-1 {
		if text[i] != text[j] {
			return false
		}
	}
	return true
}

func isPalindrome2(text string) bool {
	start := 0
	end := len(text) - 1

	for start < end {
		letterStart := fmt.Sprintf("%c", text[start])
		letterEnd := fmt.Sprintf("%c", text[end])

		if letterStart != letterEnd {
			return false
		}
		start++
		end--
	}
	return true
}
