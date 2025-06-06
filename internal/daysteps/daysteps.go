package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	Type     string
	Personal personaldata.Personal
	Weight   float64
	Height   float64
}

func (ds *DaySteps) Parse(data string) (err error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
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

	if parts[1] == "time.Hour" {
		
		ds.Duration = time.Hour
		return nil
	}
	
	duration, err := time.ParseDuration(parts[1])

	if err != nil {
		log.Print("!!!!",err)
		return fmt.Errorf("неверный формат длительности")
	}
	
	if duration.Hours()*60+duration.Minutes() <= 0 {
		return fmt.Errorf("отрицательное время")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 || ds.Duration <= 0 || ds.Personal.Height <= 0 || ds.Personal.Weight<=0 {
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

func (ds DaySteps) Print() {
	fmt.Printf("Пользователь: %s, Вес: %.2f, Рост: %.2f\n", ds.Personal.Name, ds.Personal.Weight, ds.Personal.Height)
}
