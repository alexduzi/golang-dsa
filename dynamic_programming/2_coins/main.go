package main

import (
	"fmt"
	"math"
)

func init() {
	memo = make(map[int]int)
}

// problema de retornar o melhor troco possível
// onde v = valor
// c = array de moedas que temos para dar o troco
// retorna a quantidade de moedas que podemos dar o troco
// no caso de um valor de 14 o melhor troco seria de duas moedas de 7, o resultado esperado é 2
// essa solução é lenta, quanto maior for o valor, mais lento vai ser, é um problema exponencial
func minCoins(v int, c []int) int {
	if v == 0 {
		return 0
	}

	result := math.MaxInt

	for _, coin := range c {
		if (v - coin) >= 0 {
			result = min(result, minCoins(v-coin, c)+1)
		}
	}

	return result
}

var memo map[int]int

// solução de programação dinamica com memoização topdown, recursiva
// se o input for o valor 100 e temos apenas as moedas [10, 7, 1], essa solução é muito mais rápida
func minCoins2(v int, c []int) int {
	if val, ok := memo[v]; ok {
		return val
	}

	if v == 0 {
		return 0
	}

	result := math.MaxInt

	for _, coin := range c {
		if (v - coin) >= 0 {
			result = min(result, minCoins2(v-coin, c)+1)
		}
	}

	memo[v] = result

	return memo[v]
}

// solução bottomup interativa
// aqui não tem recursão, construímos a tabela memo do zero até v
func minCoins3(v int, c []int) int {
	memo := make([]int, v+1)
	for i := 1; i <= v; i++ {
		memo[i] = math.MaxInt
	}

	for i := 1; i <= v; i++ {
		for _, coin := range c {
			if (i - coin) >= 0 {
				memo[i] = min(memo[i], memo[i-coin]+1)
			}
		}
	}

	return memo[v]
}

func main() {
	// fmt.Println(minCoins(14, []int{10, 7, 1}))
	fmt.Println(minCoins2(100, []int{10, 7, 1}))
	fmt.Println(minCoins3(100, []int{10, 7, 1}))
}
