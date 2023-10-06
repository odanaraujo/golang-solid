package good_practice

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

type EmployeePJ struct {
	Name   string
	Age    int
	Salary float64
}

type Employee interface {
	Payroll() float64
}

func (e *EmployeeCLT) Payroll() float64 {
	return e.Salary + 1000
}

func (e *Trainee) Payroll() float64 {
	return e.Salary + 100
}

func (e *EmployeePJ) Payroll() float64 {
	return e.Salary + 2000
}
