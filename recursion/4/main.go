package main

import (
	"fmt"
	"slices"
)

func main() {
	list := []string{"azul", "verde", "preto", "rosa"}
	reversed := reverse2(list)
	fmt.Printf("%v\n", reversed)
}

func reverse(list []string) []string {
	reverse := make([]string, len(list))
	for i, j := len(list)-1, 0; i >= 0; i, j = i-1, j+1 {
		reverse[j] = list[i]
	}
	return reverse
}

func reverse2(list []string) []string {
	if len(list) <= 1 {
		return list
	}

	head := list[0]
	tail := list[1:]

	return slices.Concat(reverse(tail), []string{head})
}
