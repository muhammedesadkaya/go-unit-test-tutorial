package test

import (
	"github.com/go-unit-test-tutorial/cmd"
	. "github.com/go-unit-test-tutorial/enum"
	"github.com/go-unit-test-tutorial/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_UnderAge_AutoRejectExpected(t *testing.T) {

	resume := model.Resume{
		Birthday: time.Date(2021, time.Month(8), 26, 1, 10, 30, 0, time.UTC),
	}

	var resumeApplication = cmd.NewResumeApplication(resume)

	result := resumeApplication.CheckResume()

	assert.NotNil(t, result)

	assert.Equal(t, result, AutoReject)
}

func Test_NotDoneMilitaryService_AutoRejectExpected(t *testing.T) {

	resume := model.Resume{
		Birthday:              time.Date(1993, time.Month(8), 26, 1, 10, 30, 0, time.UTC),
		IsDoneMilitaryService: false,
	}

	var resumeApplication = cmd.NewResumeApplication(resume)

	result := resumeApplication.CheckResume()

	assert.NotNil(t, result)

	assert.Equal(t, result, AutoReject)
}

func Test_OneYearExperience_AutoRejectExpected(t *testing.T) {

	resume := model.Resume{
		Birthday:              time.Date(1993, time.Month(8), 26, 1, 10, 30, 0, time.UTC),
		IsDoneMilitaryService: false,
		YearsOfExperience:     1,
	}

	var resumeApplication = cmd.NewResumeApplication(resume)

	result := resumeApplication.CheckResume()

	assert.NotNil(t, result)

	assert.Equal(t, result, AutoReject)
}

func Test_OneTalent_AutoRejectExpected(t *testing.T) {

	resume := model.Resume{
		Birthday:              time.Date(1993, time.Month(8), 26, 1, 10, 30, 0, time.UTC),
		IsDoneMilitaryService: false,
		Talent:                []string{"C#"},
	}

	var resumeApplication = cmd.NewResumeApplication(resume)

	result := resumeApplication.CheckResume()

	assert.NotNil(t, result)

	assert.Equal(t, result, AutoReject)
}

func Test_TwoTalentsAndThreeYearsOfExperience_SendMailToHRExpected(t *testing.T) {

	resume := model.Resume{
		Birthday:              time.Date(1993, time.Month(8), 26, 1, 10, 30, 0, time.UTC),
		IsDoneMilitaryService: true,
		YearsOfExperience:     3,
		Talent:                []string{"C#", "Golang"},
	}

	var resumeApplication = cmd.NewResumeApplication(resume)

	result := resumeApplication.CheckResume()

	assert.NotNil(t, result)

	assert.Equal(t, result, SendMailToHR)
}

func Test_FiveTalentsAndFiveYearsOfExperience_AutoAcceptExpected(t *testing.T) {

	resume := model.Resume{
		YearsOfExperience:     5,
		IsDoneMilitaryService: true,
		Birthday:              time.Date(1993, time.Month(8), 26, 1, 10, 30, 0, time.UTC),
		Talent:                []string{"C#", "Golang", "MongoDB", "PostgreSQL", "ReactJS"},
	}

	var resumeApplication = cmd.NewResumeApplication(resume)

	result := resumeApplication.CheckResume()

	assert.NotNil(t, result)

	assert.Equal(t, result, SendMailToTeamLead)
}

func Test_SixTalentsAndSixYearsOfExperience_AutoAcceptExpected(t *testing.T) {

	resume := model.Resume{
		YearsOfExperience:     6,
		IsDoneMilitaryService: true,
		Birthday:              time.Date(1993, time.Month(8), 26, 1, 10, 30, 0, time.UTC),
		Talent:                []string{"C#", "Golang", "MongoDB", "PostgreSQL", "ReactJS", "MSSQL"},
	}

	var resumeApplication = cmd.NewResumeApplication(resume)

	result := resumeApplication.CheckResume()

	assert.NotNil(t, result)

	assert.Equal(t, result, AutoAccept)
}
