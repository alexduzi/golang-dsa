package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//"transacoes" (ref: Leetcode invalid-transactions)
// A transaction is possibly invalid if:
// the amount exceeds $1000, or;
// if it occurs within (and including) 60 minutes of another transaction with the same name in a different city.

// You are given an array of strings transaction where transactions[i] consists of comma-separated values representing the name,
// time (in minutes), amount, and city of the transaction. Return a list of transactions that are possibly invalid.
// You may return the answer in any order.

// Constraints:
// transactions.length <= 1000
// Each transactions[i] takes the form "{name},{time},{amount},{city}"
// Each {name} and {city} consist of lowercase English letters, and have lengths between 1 and 10.
// Each {time} consist of digits, and represent an integer between 0 and 1000.
// Each {amount} consist of digits, and represent an integer between 0 and 2000.

func main() {
	fmt.Printf("%+v\n", invalidTransactions([]string{"alice,20,800,mtv", "alice,50,100,beijing"})) // result -> "alice,20,800,mtv","alice,50,100,beijing"
	fmt.Printf("%+v\n", invalidTransactions([]string{"alice,20,800,mtv", "alice,50,1200,mtv"}))    // result -> "alice,50,1200,mtv"
	fmt.Printf("%+v\n", invalidTransactions([]string{"alice,20,800,mtv", "bob,50,1200,mtv"}))      // result -> "bob,50,1200,mtv"
}

func invalidTransactions(transactions []string) []string {
	allTransactions := make([]Transaction, 0, len(transactions))
	m := make(map[string][]Transaction)
	result := []string{}

	for _, csv := range transactions {
		tran := NewTransaction(csv)
		allTransactions = append(allTransactions, tran)
		m[tran.name] = append(m[tran.name], tran)
	}

	for _, transaction := range allTransactions {
		if transaction.amount > 1000 {
			result = append(result, transaction.csv)
		} else {
			for _, transaction2 := range m[transaction.name] {
				if math.Abs(float64(transaction.time)-float64(transaction2.time)) <= 60 && transaction.city != transaction2.city {
					result = append(result, transaction.csv)
					break
				}
			}
		}
	}

	return result
}

type Transaction struct {
	name   string
	time   int
	amount int
	city   string
	csv    string
}

func NewTransaction(csv string) Transaction {
	data := strings.Split(csv, ",")
	name := data[0]
	time, _ := strconv.Atoi(data[1])
	amount, _ := strconv.Atoi(data[2])
	city := data[3]
	return Transaction{
		name:   name,
		time:   time,
		amount: amount,
		city:   city,
		csv:    csv,
	}
}
