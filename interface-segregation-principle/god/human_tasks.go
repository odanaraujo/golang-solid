package god

type HumanTasks interface {
	Job
	CoffeBreak()
	Lunch()
}

type Human struct {
	Name       string
	Age        int
	Profession string
}

func (h *Human) Enter() {
	// Enter the office
}

func (h *Human) StartWorking() {
	// Start working
}

func (h *Human) CoffeBreak() {
	// Have a coffee break
}

func (h *Human) CheckTheOli() {
	// Check the oil level ?
}

func (h *Human) Lunch() {
	// Have a lunch
}

func (h *Human) Finish() {
	// Finish the work
}
