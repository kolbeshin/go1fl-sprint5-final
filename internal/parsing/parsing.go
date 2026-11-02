package parsing

import (
	"fmt"
	"strconv"
	"time"
)

// ParseSteps - парсит количество шагов
func ParseSteps(value string) (int, error) {
	steps, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть больше 0")
	}

	return steps, nil
}

// ParseDuration - парсит продолжительность
func ParseDuration(value string) (time.Duration, error) {
	duration, err := time.ParseDuration(value)
	if err != nil {
		return 0, err
	}

	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть больше 0")
	}

	return duration, nil
}
