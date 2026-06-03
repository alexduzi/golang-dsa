package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	entry1 := []string{
		"57,Vendas,8032,Meire Silva,8000.0,57",
		"32,Estoque,4368,Dom Dias,7000.0,32",
		"57,Vendas,3298,Pedro Neto,8500.0,57",
		"57,Vendas,8639,Carol Souza,9000.0,57",
		"18,Marketing,6421,Davi Souto,7500.0,18",
		"32,Estoque,7523,Lara Matos,8000.0,32",
		"18,Marketing,2732,Bob Costa,6500.0,18",
	}

	departments1 := convertRecords(entry1)

	printDepartment(departments1)

	println()

	entry2 := []string{
		"57,Vendas,8032,Meire Silva,8000.0,57",
		"18,Marketing,6421,Davi Souto,7500.0,18",
		"18,Marketing,2732,Bob Costa,6500.0,18",
	}

	departments2 := convertRecords(entry2)

	printDepartment(departments2)
}

func convertRecords(records []string) []Department {
	depControl := make(map[int64]Department)

	for _, csv := range records {
		data := strings.Split(csv, ",")

		depId, _ := strconv.ParseInt(data[0], 0, 64)
		depName := data[1]
		empId, _ := strconv.ParseInt(data[2], 0, 64)
		empName := data[3]
		empSalary, _ := strconv.ParseFloat(data[4], 64)

		var department Department
		employee := NewEmployee(empId, empName, empSalary, depId)

		department, ok := depControl[depId]

		if !ok {
			department = NewDepartment(depId, depName)
		}

		department.addEmployee(employee)
		depControl[depId] = department
	}

	result := make([]Department, 0)
	for _, dep := range depControl {
		result = append(result, dep)
	}

	sort.Sort(ByDepartment(result))

	return result
}

func printDepartment(department []Department) {
	for _, dep := range department {
		fmt.Printf("%s:\n", dep.name)
		for _, emp := range dep.employees {
			fmt.Printf("\t%d: %s, $ %.2f\n", emp.id, emp.name, emp.salary)
		}
	}
}

type ByDepartment []Department

func (a ByDepartment) Len() int           { return len(a) }
func (a ByDepartment) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDepartment) Less(i, j int) bool { return a[i].name < a[j].name }

type Department struct {
	id        int64
	name      string
	employees []Employee
}

func (dep *Department) addEmployee(employee Employee) {
	dep.employees = append(dep.employees, employee)
}

func NewDepartment(id int64, name string) Department {
	return Department{
		id:        id,
		name:      name,
		employees: make([]Employee, 0),
	}
}

type Employee struct {
	id           int64
	name         string
	salary       float64
	departmentId int64
}

func NewEmployee(id int64, name string, salary float64, departmentId int64) Employee {
	return Employee{
		id:           id,
		name:         name,
		salary:       salary,
		departmentId: departmentId,
	}
}
