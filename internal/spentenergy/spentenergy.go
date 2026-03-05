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

// WalkingSpentCalories вычисляет калории при ходьбе.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("incorrect argument, steps: %d", steps)
	}

	if weight <= 0 {
		return 0, fmt.Errorf("incorrect argument, weight: %.2f", weight)
	}

	if height <= 0 {
		return 0, fmt.Errorf("incorrect argument, height: %.2f", height)
	}

	if duration <= 0 {
		return 0, fmt.Errorf("incorrect argument, duration: %d", duration)
	}

	avgSpeed := MeanSpeed(steps, height, duration)
	durationInMin := duration.Minutes()
	calories := (weight * durationInMin * avgSpeed) / minInH

	walkingCalories := calories * walkingCaloriesCoefficient

	return walkingCalories, nil
}

// RunningSpentCalories вычисляет калории при беге.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("incorrect argument, steps: %d", steps)
	}

	if weight <= 0 {
		return 0, fmt.Errorf("incorrect argument, weight: %.2f", weight)
	}

	if height <= 0 {
		return 0, fmt.Errorf("incorrect argument, height: %.2f", height)
	}

	if duration <= 0 {
		return 0, fmt.Errorf("incorrect argument, duration: %d", duration)
	}

	avgSpeed := MeanSpeed(steps, height, duration)
	durationInMin := duration.Minutes()
	calories := (weight * durationInMin * avgSpeed) / minInH

	return calories, nil
}

// MeanSpeed возвращает среднюю скорость (км/ч).
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	distance := Distance(steps, height)
	avgSpeed := distance / duration.Hours()

	return avgSpeed
}

// Distance вычисляет дистанцию (км) по шагам и росту.
func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanceInM := float64(steps) * stepLength
	distanceInKm := distanceInM / mInKm
	return distanceInKm
}
