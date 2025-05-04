package trainings

import (
	"time"
	//"internal/personaldata"
	personaldata "github.com/eugenii/Go_sprint_05/internal/personaldata"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	Person       personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
}

