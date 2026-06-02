package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"sync"
	"time"
)

// Fazer uma função que receba uma lista de registros de log de acesso a um website,
// e retorne a quantidade de usuários únicos que acessaram o site.
// Cada registro de log está no formato CSV (valores separados por vírgula) que contém: nome de usuário, momento de acesso e URL acessada.
func main() {
	visitors := downloadVisitorsJSON()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go measureExecution(&wg, "total", visitors, total)

	wg.Add(1)
	go measureExecution(&wg, "totalNaive", visitors, totalNaive)

	wg.Wait()
}

func downloadVisitorsJSON() []string {
	visitorsJsonURL := "https://raw.githubusercontent.com/devsuperior/curso-eda/main/conjuntos-dicionarios/java/visitantes/visitantes-input.json"

	req, err := http.NewRequest("GET", visitorsJsonURL, nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var result []string
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		panic(err)
	}
	return result
}

func measureExecution(wg *sync.WaitGroup, funcName string, visitors []string, exec func([]string) int) {
	defer wg.Done()

	start := time.Now()

	result := exec(visitors)

	end := time.Now()

	fmt.Printf("%s result is: %d executed in: %.2fs\n", funcName, result, end.Sub(start).Seconds())
}

// solução correta implementada com Set (conjunto)
func total(visitors []string) int {
	set := NewSet()
	for _, line := range visitors {
		info := strings.Split(line, ",")
		name := info[0]

		// O(1)
		set.Add(name)
	}
	return set.Size()
}

// solução mais lenta, implementada com lista
func totalNaive(visitors []string) int {
	setSlice := []string{}

	for _, line := range visitors {
		info := strings.Split(line, ",")
		name := info[0]

		// O(n × n)
		if !slices.Contains(setSlice, name) {
			setSlice = append(setSlice, name)
		}
	}

	return len(setSlice)
}

// Set representa um conjunto de elementos únicos usando um mapa internamente
type Set struct {
	elements map[any]struct{}
}

// NewSet cria e retorna um novo conjunto vazio
func NewSet() Set {
	return Set{
		elements: make(map[any]struct{}),
	}
}

// Add insere um elemento no conjunto (ignorado se já existir)
func (set *Set) Add(elem any) {
	set.elements[elem] = struct{}{}
}

// Remove exclui um elemento do conjunto
func (set *Set) Remove(elem any) {
	delete(set.elements, elem)
}

// IsEmpty retorna true se o conjunto não contiver nenhum elemento
func (set *Set) IsEmpty(elem any) bool {
	return set.Size() == 0
}

// Size retorna a quantidade de elementos no conjunto
func (set *Set) Size() int {
	return len(set.elements)
}

// List retorna todos os elementos do conjunto em um slice
func (set *Set) List() (list []any) {
	for k := range set.elements {
		list = append(list, k)
	}
	return
}
