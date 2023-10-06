package bad_pratics

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

func Payroll[TypeEmployee EmployeeCLT | Trainee](employee TypeEmployee) float64 {
	if reflect.TypeOf(EmployeeCLT{}) == reflect.TypeOf(employee) {
		// todo implement calc payroll + benefits

		return 0.0
	}
	// todo implement calc treinee payroll
	return 0.0
}
