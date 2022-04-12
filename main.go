package main

import (
	"fmt"
	"github.com/go-unit-test-tutorial/cmd"
	"time"

	"github.com/go-unit-test-tutorial/model"
)

func main() {

	resume := model.Resume{
		FullName:              "Muhammed Esad Kaya",
		YearsOfExperience:     5,
		IsDoneMilitaryService: true,
		Birthday:              time.Date(1993, time.Month(8), 26, 1, 10, 30, 0, time.UTC),
		Talent:                []string{"C#", "Golang", "MongoDB", "PostgreSQL", "ReactJS"},
	}

	var resumeApplication = cmd.NewResumeApplication(resume)

	result := resumeApplication.CheckResume()
	fmt.Println(result)
}
