package scheduler

// Scheduler does something which I am not entirely sure of
// TODO: Understand what the Scheduler *really* does and why is it an interface
type Scheduler interface {
	SelectCandidateNodes()
	Score()
	Pick()
}
