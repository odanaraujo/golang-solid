package main

import (
	"fmt"
	"github.com/odanaraujo/golang/solid/good_practice"
)

func main() {
	pj := good_practice.EmployeePJ{
		Name:   "Dan",
		Age:    26,
		Salary: 1000,
	}

	pj.Salary = pj.Payroll()

	fmt.Println(pj.PrintEmployee())
}
