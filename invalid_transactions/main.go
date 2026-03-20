package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/invalid-transactions/description/

// A transaction is possibly invalid if:

// the amount exceeds $1000, or;
// if it occurs within (and including) 60 minutes of another transaction with the same name in a different city.
// You are given an array of strings transaction where transactions[i] consists of comma-separated values representing the name,
// time (in minutes), amount, and city of the transaction.

// Return a list of transactions that are possibly invalid. You may return the answer in any order.

// Example 1:

// Input: transactions = ["alice,20,800,mtv","alice,50,100,beijing"]
// Output: ["alice,20,800,mtv","alice,50,100,beijing"]
// Explanation: The first transaction is invalid because the second transaction occurs within a difference of 60 minutes,
// have the same name and is in a different city. Similarly the second one is invalid too.
// Example 2:

// Input: transactions = ["alice,20,800,mtv","alice,50,1200,mtv"]
// Output: ["alice,50,1200,mtv"]
// Example 3:

// Input: transactions = ["alice,20,800,mtv","bob,50,1200,mtv"]
// Output: ["bob,50,1200,mtv"]

type Transaction struct {
	Name  string
	Time  int
	Money float64
	City  string
}

func NewTransaction(transaction string) Transaction {
	parts := strings.Split(transaction, ",")
	time, _ := strconv.Atoi(parts[1])
	money, _ := strconv.ParseFloat(parts[2], 64)
	return Transaction{
		Name:  parts[0],
		Time:  time,
		Money: money,
		City:  parts[3],
	}
}

func main() {
	transactions := []string{
		"alice,20,800,mtv",
		"alice,50,1200,beijing",
		"bob,60,300,sp",
		"ana,70,1100,sp",
	}
	invalid := invalidTransactions(transactions)
	fmt.Printf("%+v\n", invalid)
}

// example using only arrays to track invalid transactions
func invalidTransactions(transactions []string) []string {
	invalidTrans := make([]bool, len(transactions))
	result := make([]string, 0)

	for i := 0; i < len(transactions); i++ {
		ti := NewTransaction(transactions[i])
		fmt.Printf("%+v\n", ti)
		if ti.Money > 1000 {
			invalidTrans[i] = true
		}
		for j := i + 1; j < len(transactions); j++ {
			tj := NewTransaction(transactions[j])
			if ti.Name == tj.Name && int(math.Abs(float64(ti.Time)-float64(tj.Time))) <= 60 && ti.City != tj.City {
				invalidTrans[i] = true
				invalidTrans[j] = true
			}
		}
	}

	for i := 0; i < len(transactions); i++ {
		if invalidTrans[i] {
			result = append(result, transactions[i])
		}
	}

	return result
}
