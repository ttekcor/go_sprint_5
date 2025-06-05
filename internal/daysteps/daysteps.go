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
	Type string
	Personal personaldata.Personal
}

func (ds *DaySteps) Parse(data string) (err error) {
	parts := strings.Split(data, ";")
	if len(parts) != 3 {
		return fmt.Errorf("неверный формат входных данных")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("неверный формат количества шагов")
	}
	ds.Steps = steps

	ds.Type = parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("неверный формат длительности")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps == 0 || ds.Duration == 0 {
		return "", fmt.Errorf("недостаточно данных для подсчёта")
	}

	distance := float64(ds.Steps) * ds.Personal.Height * 0.45 / 1000 // distance in km
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories), nil
}
