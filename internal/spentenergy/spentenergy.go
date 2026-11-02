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

// WalkingSpentCalories возвращает затраченные калории при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	return calories * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories возвращает затраченные калории при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("некорректные входные параметры")
	}

	speed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := weight * speed * durationInMinutes / minInH

	return calories, nil
}

// MeanSpeed возвращает среднюю скорость в километрах в час.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration > 0 {
		avgSpeed := Distance(steps, height) / duration.Hours()

		return avgSpeed
	}
	return 0
}

// Distance возвращает пройденное расстояние в километрах.
func Distance(steps int, height float64) float64 {
	footLength := height * stepLengthCoefficient
	distance := footLength * float64(steps) / mInKm

	return distance
}
