package main

import (
	"fmt"
	"regexp"
)

// Use MustCompile: When the regex is a hardcoded constant that you've tested
// Use Compile: When the regex comes from user input, config files, or could be invalid

var (
	cepPattern      = regexp.MustCompile(`^\d{5}-?\d{3}$`)
	cpfPattern      = regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}-\d{2}$|^\d{11}$`)
	emailPattern    = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	brDomainPattern = regexp.MustCompile(`^[a-zA-Z0-9.-]+\.br$`)
)

func main() {
	testCases := []string{
		"12345-678",
		"12345678",
		"1234-678",
		"12345-67",
		"1234-5678",
	}

	for _, cep := range testCases {
		fmt.Println(cep, "  ", validateCEP(cep))
	}

	fmt.Println("        ")

	cpfTests := []string{
		"123.456.789-00", // valid format
		"12345678900",    // valid format (no formatting)
		"123.456.789",    // invalid
	}
	for _, cep := range cpfTests {
		fmt.Println(cep, "  ", validateCPF(cep))
	}

	fmt.Println("        ")

	emailTests := []string{
		"user@example.com",         // valid
		"user.name+tag@example.br", // valid
		"invalid@",                 // invalid
		"@example.com",             // invalid
	}
	for _, cep := range emailTests {
		fmt.Println(cep, "  ", validateEmail(cep))
	}

	fmt.Println("        ")

	domainTests := []string{
		"google.com.br", // valid
		"uol.com.br",    // valid
		"example.br",    // valid
		"google.com",    // invalid
		".br",           // invalid
	}
	for _, cep := range domainTests {
		fmt.Println(cep, "  ", validateDomains(cep))
	}
}

func validateCEP(cep string) bool {
	return cepPattern.Match([]byte(cep))
}

func validateCPF(cpf string) bool {
	return cpfPattern.Match([]byte(cpf))
}

func validateEmail(email string) bool {
	return emailPattern.Match([]byte(email))
}

func validateDomains(domain string) bool {
	return brDomainPattern.Match([]byte(domain))
}
