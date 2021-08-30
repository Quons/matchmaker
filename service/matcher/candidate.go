package matcher

import (
	"github.com/Quons/matchmaker/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Candidate struct {
	Email  string        `form:"email" json:"email" binding:"required"`
	Gender models.Gender `form:"gender" json:"gender" binding:"required"`
}

// 添加候选人
func AddCandidate(candidate Candidate) error {
	gender := candidate.Gender
	if gender != models.GenderMale && gender != models.GenderFemale {
		logrus.Errorf("invalid_gender:%v", gender)
		return errors.New("invalid_gender")
	}
	dbCandidate := models.Candidate{
		Gender: candidate.Gender,
		Email:  candidate.Email,
	}
	err := models.AddCandidate(dbCandidate)
	if err != nil {
		logrus.Errorf("add_candidate_err:%v", err)
		return err
	}
	return nil
}
