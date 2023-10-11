package god

type Job interface {
	Enter()
	StartWorking()
	ContinueToWork()
	Finish()
}
