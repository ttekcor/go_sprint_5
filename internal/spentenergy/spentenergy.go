package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверка входных параметров
	if steps <= 0 {
		return 0, fmt.Errorf("неверное количество шагов")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("неверный вес")
	}
	if height <= 0 {
		return 0, fmt.Errorf("неверный рост")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("неверная продолжительность")
	}

	// Рассчитываем среднюю скорость
	speed := MeanSpeed(steps, height, duration)
	
	// Рассчитываем калории
	// (weight * meanSpeed * durationInMinutes) / minInH
	calories := weight * speed * duration.Minutes() / float64(minInH)
	
	return calories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// Проверка входных параметров
	if steps <= 0 {
		return 0, fmt.Errorf("неверное количество шагов")
	}
	if weight <= 0 {
		return 0, fmt.Errorf("неверный вес")
	}
	if height <= 0 {
		return 0, fmt.Errorf("неверный рост")
	}
	if duration <= 0 {
		return 0, fmt.Errorf("неверная продолжительность")
	}

	// Рассчитываем среднюю скорость
	speed := MeanSpeed(steps, height, duration)
	
	// Рассчитываем калории как для бега
	calories := weight * speed * duration.Minutes() / float64(minInH)
	
	// Применяем коэффициент для ходьбы
	walkingCalories := calories * walkingCaloriesCoefficient
	
	return walkingCalories, nil
}

func Distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	distanceStep := height * stepLengthCoefficient
	resultDistance := (distanceStep * float64(steps)) / mInKm
	return resultDistance
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration > 0 {
		distanceRes := Distance(steps, height)
		result := distanceRes / duration.Hours()
		return result
	} else {
		return 0
	}

}