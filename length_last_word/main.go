package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
}

func lengthOfLastWord(s string) int {
	wordCounter := 0
	s = strings.TrimRight(s, " ")
	sizeString := len(s)
	for i := 0; i < sizeString; i++ {
		if s[i] == ' ' {
			wordCounter = 0
		} else {
			wordCounter++
		}
	}
	return wordCounter
}

func lengthOfLastWord2(s string) int {
	s = strings.TrimSpace(s)
	words := strings.Split(s, " ")
	return len(words[len(words)-1])
}
