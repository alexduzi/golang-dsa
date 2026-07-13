package main

import (
	"fmt"
	"math"
)

func minJumpsValue(values []int) int {
	n := len(values)
	memo := make([]int, n+1)
	for i := 1; i < n; i++ {
		memo[i] = -1
	}

	return minJumpsValueUtil(n-1, values, memo)
}

func minJumpsValueUtil(n int, values []int, memo []int) int {
	if n == 0 {
		return 0
	}

	jump1Cost := minJumpsValueUtil(n-1, values, memo) + int(math.Abs(float64(values[n]-values[n-1])))

	jump2Cost := math.MaxInt
	if n-2 >= 0 {
		jump2Cost = minJumpsValueUtil(n-2, values, memo) + int(math.Abs(float64(values[n]-values[n-2])))
	}

	return min(jump1Cost, jump2Cost)
}

// Problema "frog_jumps"
// Suponha que temos uma sequência de pedras numeradas de 1 até N, onde cada pedra possui uma altura dada por height[i].
// Um sapo está inicialmente na pedra 1 e deseja alcançar a pedra N.
// Para isso, ele pode pular para a próxima pedra (i+1) ou para a seguinte (i+2).
// No entanto, cada salto tem um custo |height[i] - height[j]| (valor absoluto), onde j é a pedra onde ele pousará.
// Determine o custo mínimo total para que o sapo alcance a pedra N.
func main() {
	fmt.Println(minJumpsValue([]int{10, 30, 40, 20}))
	fmt.Println(minJumpsValue([]int{10, 10}))
	fmt.Println(minJumpsValue([]int{30, 10, 60, 10, 60, 50}))
}
