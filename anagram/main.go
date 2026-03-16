package main

import (
	"sort"
	"strings"
)

func main() {
	println(isAnagram("anagram", "nagaram"))
	println(isAnagram("rat", "cat"))

	println(isAnagram2("anagram", "nagaram"))
	println(isAnagram2("rat", "cat"))

	println(isAnagram3("anagram", "nagaram"))
	println(isAnagram3("rat", "cat"))
}

// anagram é um texto formado pelo rearanjo das letras em um texto diferente
// tipicamente utilizando todas as letras originais exatamente uma vez
func isAnagram(s, t string) bool {

	if len(s) != len(t) {
		return false
	}

	hash := make(map[byte]int, len(s))

	for i := 0; i < len(s); i++ {
		hash[s[i]] += 1
		hash[t[i]] -= 1
	}

	for _, v := range hash {
		if v != 0 {
			return false
		}
	}

	return true
}

// ordenando as duas palavras
func isAnagram2(s, t string) bool {

	if len(s) != len(t) {
		return false
	}
	sSplit := strings.Split(s, "")
	tSplit := strings.Split(t, "")

	sort.Strings(sSplit)
	sort.Strings(tSplit)

	sSorted := strings.Join(sSplit, "")
	tSorted := strings.Join(tSplit, "")

	return sSorted == tSorted
}

// tabela unicode
func isAnagram3(s, t string) bool {

	if len(s) != len(t) {
		return false
	}

	var table [256]int

	for i := 0; i < len(s); i++ {
		table[s[i]]++
		table[t[i]]--
	}

	for _, v := range table {
		if v != 0 {
			return false
		}
	}

	return true
}
