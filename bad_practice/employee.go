package bad_practice

import "reflect"

type EmployeeCLT struct {
	Name   string
	Age    int
	Salary float64
}

type Trainee struct {
	Name   string
	Age    int
	Salary float64
}

//caso queira colocar outros tipos, o código só tenderá a crescer e precisaremos mexer no código que já estava funcionando.

func Payroll[TypeEmployee EmployeeCLT | Trainee](employee TypeEmployee) float64 {
	if reflect.TypeOf(EmployeeCLT{}) == reflect.TypeOf(employee) {
		// todo implement calc payroll + benefits
		return 0.0
	}
	// todo implement calc treinee payroll
	return 0.0
}
