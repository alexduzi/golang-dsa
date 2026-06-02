package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"golang.org/x/text/collate"
	"golang.org/x/text/language"
)

// Você deve fazer uma função que recebe um texto, e retorna o número de ocorrências de cada palavra no texto.
// Cada ocorrência deve ser representada por um objeto do tipo Rank, com os campos word e count.
// A saída da função deve ser uma lista de objetos tipo Rank. Você deve mostrar o resultado em letras minúsculas,
// ordenado da palavra mais frequente para a menos frequente, e as palavras com mesma frequência devem aparecer ordenadas alfabeticamente.

// Antes de analisar as frequências das palavras, você deve aplicar a seguinte função normalize ao texto,
// para remover símbolos de pontuação, remover espaços adicionais e converter tudo para minúsculo
//
// function normalize(text) {
//   const words = text.replace(/[^\p{L}\p{N}\s]/gu, " ");
//   return words.replace(/\s+/g, " ").trim().toLowerCase();
// }

func main() {
	text := `
		O vento sussurra sons entre as árvores, 
		sons que fazem animais correrem. 
		A floresta e a natureza vibram com segredos e sons.
	`

	result := wordCount(text)
	for _, rank := range result {
		fmt.Printf("%s: %d\n", rank.word, rank.count)
	}
}

type Rank struct {
	word  string
	count int
}

func NewRank(w string, c int) Rank {
	return Rank{
		word:  w,
		count: c,
	}
}

var cl = collate.New(language.Portuguese)

type ByRank []Rank

func (a ByRank) Len() int { return len(a) }
func (a ByRank) Less(i, j int) bool {
	if a[i].count != a[j].count {
		return a[i].count > a[j].count
	}
	return cl.CompareString(a[i].word, a[j].word) < 0
}
func (a ByRank) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func wordCount(text string) []Rank {
	m := make(map[string]int)

	text = normalize(text)
	textArr := strings.Split(text, " ")

	for _, word := range textArr {
		m[word] += 1
	}

	result := make([]Rank, 0, len(m))
	for k, v := range m {
		result = append(result, NewRank(k, v))
	}

	sort.Sort(ByRank(result))

	return result
}

func normalize(text string) string {
	re1 := regexp.MustCompile(`[^\pL\pN\s]`)
	re2 := regexp.MustCompile(`\s+`)

	result := re1.ReplaceAllString(text, "")
	result = re2.ReplaceAllString(result, " ")
	result = strings.TrimSpace(result)
	result = strings.ToLower(result)
	return result
}
