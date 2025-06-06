package trainings

import (
	"fmt"
	"log"
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
	Weight float64
	Height float64
}

func (ds *Training) Parse(data string) (err error) {
	log.Print("???",data)
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return fmt.Errorf("неверный формат входных данных")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("неверный формат количества шагов: %v", err)
	}
	if steps <= 0 {
		return fmt.Errorf("отрицательные шаги")
	}
	ds.Steps = steps

	ds.TrainingType = parts[1]

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		log.Print("!!!!",err)
		return fmt.Errorf("неверный формат длительности")
	}
	if duration <= 0 {
		return fmt.Errorf("отрицательное время")
	}
	ds.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.Steps == 0 || t.Duration <= 0 {
		return "", fmt.Errorf("недостаточно данных для подсчёта")
	}

	distance := float64(t.Steps) * t.Personal.Height * 0.45 / 1000
	speed := distance / t.Duration.Hours()
	
	var calories float64
	var err error
	
	// Выбираем функцию в зависимости от типа тренировки
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
	
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil
}

func (t Training) Print() {
	t.Personal.Print()
}
