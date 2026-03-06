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

// Training содержит данные тренировки и персональные параметры.
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает строку "шаги,тип тренировки, длительность".
func (t *Training) Parse(datastring string) (err error) {
	splitData := strings.Split(datastring, ",")
	if len(splitData) != 3 {
		return fmt.Errorf("incorrect data: %q, transferred: %d", datastring, len(splitData))
	}

	steps, err := strconv.Atoi(splitData[0])
	if err != nil {
		return fmt.Errorf("invalid steps format: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("incorrect data: %q, steps: %d", datastring, steps)
	}

	t.Steps = steps
	t.TrainingType = splitData[1]

	duration, err := time.ParseDuration(splitData[2])
	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("incorrect data: %q, duration: %d", datastring, duration)
	}

	t.Duration = duration

	return nil
}

// ActionInfo возвращает статистику тренировки.
func (t Training) ActionInfo() (string, error) {
	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		err = fmt.Errorf("неизвестный тип тренировки")
		log.Println(err)
		return "", err
	}
	return fmt.Sprintf(`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`, t.TrainingType, t.Duration.Hours(), spentenergy.Distance(t.Steps, t.Height), spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration), calories), err
}
