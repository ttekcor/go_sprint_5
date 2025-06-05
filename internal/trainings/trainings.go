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
	Steps int
	TrainingType string
	Duration time.Duration
	Personal personaldata.Personal
}

func (t *Training) Parse(data string) (err error) {
	parts := strings.Split(data, ";")
	if len(parts) != 3 {
		return fmt.Errorf("неверный формат входных данных")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("неверный формат количества шагов")
	}
	t.Steps = steps

	t.TrainingType = parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("неверный формат длительности")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.Steps == 0 || t.Duration == 0 {
		return "", fmt.Errorf("недостаточно данных для подсчёта")
	}

	distance := float64(t.Steps) * t.Personal.Height * 0.45 / 1000 // distance in km
	speed := distance / t.Duration.Hours()
	calories, err := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %v\nДистанция: %.2f км\nСкорость: %.2f км/ч\nПотрачено ккал: %.2f\n",
		t.TrainingType, t.Duration, distance, speed, calories), nil
}

func (t Training) Print() {
	t.Personal.Print()
}
