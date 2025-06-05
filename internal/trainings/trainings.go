package trainings

import "time"

type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
}
