package cmd

import (
	"github.com/go-unit-test-tutorial/pkg/helper"

	. "github.com/go-unit-test-tutorial/enum"
	"github.com/go-unit-test-tutorial/model"
)

type ResumeService interface {
	CheckResume() string
}

type resumeService struct {
	resume model.Resume
}

func (self *resumeService) CheckResume() string {

	ageYear := helper.CheckAgeAdult(self.resume.Birthday)

	if ageYear < 18 ||
		!self.resume.IsDoneMilitaryService ||
		len(self.resume.Talent) <= 1 ||
		self.resume.YearsOfExperience <= 1 {
		return AutoReject
	}

	if len(self.resume.Talent) <= 2 {
		return SendMailToHR
	}

	if (len(self.resume.Talent) >= 3 && len(self.resume.Talent) <= 5) && (self.resume.YearsOfExperience >= 3 && self.resume.YearsOfExperience <= 5) {
		return SendMailToTeamLead
	}

	return AutoAccept
}

func NewResumeApplication(resume model.Resume) ResumeService {
	return &resumeService{resume: resume}
}
