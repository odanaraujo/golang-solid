package good_practice

import "fmt"

type PrinterDataEmployee interface {
	PrintEmployee() string
}

func (e *EmployeeCLT) PrintEmployee() string {
	return fmt.Sprintf("Name: %s, Age: %d, Salary: %f", e.Name, e.Age, e.Salary)
}

func (e *Trainee) PrintEmployee() string {
	return fmt.Sprintf("Name: %s, Age: %d, Salary: %f", e.Name, e.Age, e.Salary)
}

func (e *EmployeePJ) PrintEmployee() string {
	return fmt.Sprintf("Name: %s, Age: %d, Salary: %0.2f", e.Name, e.Age, e.Salary)
}
