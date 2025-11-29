package main

import (
	"fmt"
	"strings"
)

func main() {
	var sb strings.Builder

	var name string = "Alex"
	var job string = "works with I.T"
	var path string = "C:\\Program Files\\Windows32\\"

	var text string = `
		huge
		text
		with
		a
		lot
		paragraphs
	`

	var text2 string = "car \"123\""

	var interpolation string = fmt.Sprintf("%s has a salary of %.2f\n", name, 9000000.00)

	sb.WriteString(name)
	sb.WriteRune(' ')
	sb.WriteString(job)

	fmt.Println(name + " " + job)
	fmt.Println(sb.String())
	fmt.Println(path)
	fmt.Println(text)
	fmt.Println(text2)
	fmt.Println(interpolation)
	fmt.Println()

	var hello string = "Hello, world!"
	fmt.Println(hello[2:5])
	fmt.Println("            ")

	// Basic tests
	fmt.Println(countVowels("hello")) // 2 (e, o)
	fmt.Println(countVowels("world")) // 1 (o)
	fmt.Println(countVowels("aeiou")) // 5 (all vowels)
	fmt.Println(countVowels("AEIOU")) // 5 (uppercase vowels)
	fmt.Println(countVowels("AeIoU")) // 5 (mixed case)

	// Edge cases
	fmt.Println(countVowels(""))       // 0 (empty string)
	fmt.Println(countVowels("bcdfg"))  // 0 (no vowels)
	fmt.Println(countVowels("rhythm")) // 0 (no vowels)
	fmt.Println(countVowels("a"))      // 1 (single vowel)
	fmt.Println(countVowels("b"))      // 0 (single consonant)

	// Sentences and phrases
	fmt.Println(countVowels("Hello World"))                                 // 3 (e, o, o)
	fmt.Println(countVowels("The quick brown fox jumps over the lazy dog")) // 11
	fmt.Println(countVowels("Programming in Go"))                           // 5

	// Special characters and numbers
	fmt.Println(countVowels("123456"))         // 0 (numbers)
	fmt.Println(countVowels("!@#$%^&*()"))     // 0 (special chars)
	fmt.Println(countVowels("hello123world!")) // 3 (e, o, o)
	fmt.Println(countVowels("a e i o u"))      // 5 (with spaces)

	// Longer strings
	fmt.Println(countVowels("Beautiful")) // 5 (e, a, u, i, u)
	fmt.Println(countVowels("Education")) // 5 (E, u, a, i, o)
}

var (
	vowels = map[rune]rune{'a': 'a', 'e': 'e', 'i': 'i', 'o': 'o', 'u': 'u', 'A': 'A', 'E': 'E', 'I': 'I', 'O': 'O', 'U': 'U'}
)

func countVowels(text string) (counter int) {
	for _, char := range text {
		if _, ok := vowels[char]; ok {
			counter++
		}
	}
	return
}
