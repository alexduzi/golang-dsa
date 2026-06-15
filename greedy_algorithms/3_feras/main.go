package main

import (
	"fmt"
	"math"
	"sort"
)

// Suponha que existem C (1 ≤ C ≤ 5) jaulas que podem abrigar 0, 1 ou 2 feras,
// S (1 ≤ S ≤ 2C) feras e uma lista M das massas das S feras. Determine qual jaula deve
// conter cada fera para minimizar o ‘desbalanceamento’.
func main() {
	fmt.Println(minimumImbalance(3, 4, []int{5, 1, 2, 7}))
}

func minimumImbalance(c, s int, m []int) int {
	// calcular a média
	a := 0
	for j := 0; j < s; j++ {
		a += m[j]
	}
	a /= c

	// adicionar extras
	if s < 2*c {
		dummies := 2*c - s
		for i := 0; i < dummies; i++ {
			m = append(m, 0)
		}
	}

	// ordenar
	sort.Ints(m)

	chambers := make([][]int, c)
	for i := 0; i < c; i++ {
		chambers[i] = make([]int, 0)
	}

	// fazemos o pareamento
	for i := 0; i < c; i++ {
		chambers[i] = append(chambers[i], m[i])
		chambers[i] = append(chambers[i], m[2*c-1-i])
	}

	// percorrer jaulas
	imbalance := 0
	for i := 0; i < c; i++ {
		xi := 0
		// massa total das feras na jaula i
		for j := 0; j < len(chambers[i]); j++ {
			xi += chambers[i][j]
		}
		imbalance += int(math.Abs(float64(xi) - float64(a)))
	}

	return imbalance
}
