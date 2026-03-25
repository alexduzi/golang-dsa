package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
)

// Dado um array de vendedores, cada um representado por um objeto com o nome do vendedor name e o valor de suas vendas amount,
// crie uma função para encontrar e retornar o vendedor que obteve o maior valor de venda.

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	bytes, err := os.ReadFile(filepath.Join(dir, "seller.json"))
	if err != nil {
		panic(err)
	}

	var sellers []Seller
	err = json.Unmarshal(bytes, &sellers)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The best seller is: %s\n", findBestSeller(sellers))
	fmt.Printf("The best seller is: %s\n", findBestSeller2(sellers))
}

type Seller struct {
	Name   string
	Amount float64
}

type ByAmount []Seller

func (a ByAmount) Len() int           { return len(a) }
func (a ByAmount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAmount) Less(i, j int) bool { return a[i].Amount > a[j].Amount } // ordenar por asc (<) ou desc (>)

func findBestSeller(sellers []Seller) string {
	name := ""
	maxAmount := sellers[0].Amount
	for i := 0; i < len(sellers); i++ {
		if sellers[i].Amount > maxAmount {
			maxAmount = sellers[i].Amount
			name = sellers[i].Name
		}
	}
	return name
}

func findBestSeller2(sellers []Seller) string {
	sort.Sort(ByAmount(sellers))
	return sellers[0].Name
}
