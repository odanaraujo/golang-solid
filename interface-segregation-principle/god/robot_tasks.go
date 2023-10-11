package god

type JobRobot interface {
	Job
	CheckTheOli()
	ChargeTheBattery()
}

type Robot struct {
	Model string
	Year  int
	Power int
}

func (r *Robot) Enter() {
	// Enter the office
}

func (r *Robot) StartWorking() {
	// Start working
}

func (r *Robot) CheckTheOli() {
	// Check the oil level
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
