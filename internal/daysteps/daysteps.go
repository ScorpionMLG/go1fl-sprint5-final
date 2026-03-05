package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps содержит данные суточной шаговой активности и персональные данные.
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse разбирает строку "шаги,длительность".
func (ds *DaySteps) Parse(datastring string) (err error) {
	splitData := strings.Split(datastring, ",")
	if len(splitData) != 2 {
		return fmt.Errorf("incorrect data: %q, transferred: %d", datastring, len(splitData))
	}

	steps, err := strconv.Atoi(splitData[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return fmt.Errorf("incorrect data: %q, steps: %d", datastring, steps)
	}

	ds.Steps = steps

	duration, err := time.ParseDuration(splitData[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return fmt.Errorf("incorrect data: %q, duration: %d", datastring, duration)
	}

	ds.Duration = duration

	return nil
}

// ActionInfo возвращает статистику ходьбы или ошибку при некорректных данных.
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Steps <= 0 {
		return "", fmt.Errorf("number of steps must be greater than 0, steps: %d", ds.Steps)
	}
	if ds.Weight <= 0 {
		return "", fmt.Errorf("weight must be greater than 0, steps: %.2f", ds.Weight)
	}
	if ds.Height <= 0 {
		return "", fmt.Errorf("height must be greater than 0, steps: %.2f", ds.Height)
	}
	if ds.Duration <= 0 {
		return "", fmt.Errorf("duration must be greater than 0, steps: %.2f", ds.Duration.Hours())
	}
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	return fmt.Sprintf(`Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
`, ds.Steps, spentenergy.Distance(ds.Steps, ds.Height), calories), err
}
