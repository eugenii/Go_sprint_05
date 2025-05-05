package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"internal/personaldata"
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
	ds.Duration, err = time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("invalid duration format")
	}
	return nil

}

func (ds DaySteps) ActionInfo() (string, error) {
	// Количество шагов: 792.
	// Дистанция составила 0.51 км.
	// Вы сожгли 221.33 ккал.
	distance := personaldata.Distance(ds.Steps, ds.Personal.Height)
	calories := personaldata.Calories(ds.Steps, ds.Personal.Weight, ds.Duration)
	return fmt.Sprintf("Количество шагов: %d\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", ds.Steps, distance, calories), nil
}
