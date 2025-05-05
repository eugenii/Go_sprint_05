package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/eugenii/Go_sprint_05/internal/personaldata"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("invalid data format")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("invalid steps value")
	}
	t.Steps = steps
	t.TrainingType = parts[1]
	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("invalid duration value")
	}
	t.Duration = duration
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
}
