package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
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
	return nil
}

func (t Training) ActionInfo() (string, error) {
	var (
		calories float64
		err      error
	)
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	if t.TrainingType == "Бег" {
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	} else if t.TrainingType == "Ходьба" {
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	} else {
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	if err != nil {
		return "", err
	}
	hours := t.Duration.Hours()
	formattedDuration := fmt.Sprintf("%.2f ч.", hours)
	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %s\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, formattedDuration, distance, speed, calories), nil
}
