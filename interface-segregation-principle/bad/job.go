package bad

type Job interface {
	Enter()
	StartWorking()
	CoffeBreak()
	CheckTheOli()
	Lunch()
	ChargeTheBattery()
	ContinueToWork()
	Finish()
}

type Human struct {
	Name       string
	Age        int
	Profession string
}

type Robot struct {
	Model string
	Year  int
	Power int
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

func (h *Human) ChargeTheBattery() {
	// Charge the battery ?
}

func (h *Human) ContinueToWork() {
	// Continue to work
}

func (h *Human) Finish() {
	// Finish the work
}

func (r *Robot) Enter() {
	// Enter the office
}

func (r *Robot) StartWorking() {
	// Start working
}

func (r *Robot) CoffeBreak() {
	// Have a coffee break ?
}

func (r *Robot) CheckTheOli() {
	// Check the oil level
}

func (r *Robot) Lunch() {
	// Have a lunch ?
}

func (r *Robot) ChargeTheBattery() {
	// Charge the battery
}

func (r *Robot) ContinueToWork() {
	// Continue to work
}

func (r *Robot) Finish() {
	// Finish the work
}
