package helper

import (
	"time"
)

func CheckAgeAdult(birthday time.Time) int {
	t := time.Now()
	result := t.Year() - birthday.Year()
	return result
}
