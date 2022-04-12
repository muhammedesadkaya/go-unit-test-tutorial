package model

import "time"

type Resume struct {
	FullName              string
	Talent                []string
	YearsOfExperience     uint8
	IsDoneMilitaryService bool
	Birthday              time.Time
}
