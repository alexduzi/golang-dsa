package main

import (
	"fmt"
	"time"
)

func init() {
	memo = make(map[int]int)
}

// recursão padrão de fibonacci
// demora muito para executar pois sempre calcula a mesma coisa várias vezes
// quanto maior o "n" maior vai ser a quantidade de rendundância
// essa redundância seriam os cálculos repetidos
func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

// resolvendo isso com programação dinamica
// podemos criar um "cache" do resultado de alguma chamada
// isso seria memoizar o resultado
// recursão + memoização = programação dinamica topdown
// essa abordagem é conhecida como topdown
var memo map[int]int

func fib2(n int) int {
	if val, ok := memo[n]; ok {
		return val
	}

	if n <= 2 {
		return 1
	}

	memo[n] = fib2(n-1) + fib2(n-2)

	return memo[n]
}

// essa abordagem é conhecida como bottomup
// pré calculamos todos os resultados antes para depois retornar o resultado esperado
// iteração + tabulação = programação dinamica bottomup
func fib3(n int) int {
	m := make(map[int]int)

	m[0] = 0
	m[1] = 1
	m[2] = 1

	for i := 3; i <= n; i++ {
		m[i] = m[i-1] + m[i-2]
	}

	return m[n]
}

type FibData struct {
	Name     string
	Result   int
	Duration time.Duration
}

func main() {
	const n = 100

	fibChan := make(chan FibData, 3)

	go func() {
		start := time.Now()
		result := fib(n)
		fibChan <- FibData{Name: "normal fib", Result: result, Duration: time.Since(start)}
	}()

	go func() {
		start := time.Now()
		result := fib2(n)
		fibChan <- FibData{Name: "topdown memo fib2", Result: result, Duration: time.Since(start)}
	}()

	go func() {
		start := time.Now()
		result := fib3(n)
		fibChan <- FibData{Name: "bottomup memo fib3", Result: result, Duration: time.Since(start)}
	}()

	for range 3 {
		data := <-fibChan
		fmt.Printf("%s(%d) = %d, took %v\n", data.Name, n, data.Result, data.Duration)
	}
}
