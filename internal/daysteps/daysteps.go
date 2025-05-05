package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// "678,0h50m"
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("invalid data format")
	}
	ds.Steps, err = strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("invalid steps format")
	}
	if ds.Steps <= 0 {
		return fmt.Errorf("invalid steps value")
	}
	ds.Duration, err = time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("invalid duration format")
	}
	if ds.Duration <= 0 {
		return fmt.Errorf("invalid duration value")
	}
	return nil

}

func (ds DaySteps) ActionInfo() (string, error) {
	// Количество шагов: 792.
	// Дистанция составила 0.51 км.
	// Вы сожгли 221.33 ккал.
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("invalid data format")
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
