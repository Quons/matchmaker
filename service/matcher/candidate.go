package matcher

import (
	"github.com/Quons/matchmaker/models"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Candidate struct {
	Email  string        `form:"email" json:"email" binding:"required"`
	Gender models.Gender `form:"gender" json:"gender" binding:"required"`
	Age    int64         `form:"age" json:"age" binding:"required"`
}

// 添加候选人
func AddCandidate(candidate Candidate) error {
	gender := candidate.Gender
	if gender != models.GenderMale && gender != models.GenderFemale {
		logrus.Errorf("invalid_gender:%v", gender)
		return errors.New("invalid_gender")
	}
	if candidate.Email == "" {
		logrus.Errorf("empty_email")
		return errors.New("empty_email")
	}
	if candidate.Age <= 14 || candidate.Age >= 50 {
		logrus.Errorf("invalid_age:%v", candidate.Age)
		return errors.New("invalid_age")
	}
	// 先查询该邮箱是否已经存在
	existCandidate, err := models.GetCandidateByEmail(candidate.Email)
	if err != nil {
		logrus.Errorf("get_candidate_by_email_err:%v", err)
		return err
	}
	dbCandidate := models.Candidate{
		Gender: candidate.Gender,
		Email:  candidate.Email,
		Age:    candidate.Age,
		Status: models.CandidateStatusOk,
	}

	// 已存在，就更新
	if existCandidate.Email != "" {
		err := models.UpdateCandidate(dbCandidate)
		if err != nil {
			logrus.Errorf("udpate_candidate_err:%v", err)
			return err
		}
		return nil
	}

	err = models.AddCandidate(dbCandidate)
	if err != nil {
		logrus.Errorf("add_candidate_err:%v", err)
		return err
	}
	return nil
}

type PushStatus struct {
	Email  string `form:"email" json:"email" binding:"required"`
	Status int64  `form:"status" json:"status" binding:"required"`
}

func UpdatePushStatus(pushStatus PushStatus) error {
	// 参数校验
	if pushStatus.Email == "" {
		logrus.Errorf("empty_email")
		return errors.New("empty_email")
	}
	// 查询邮件
	candidate, err := models.GetCandidateByEmail(pushStatus.Email)
	if err != nil {
		logrus.WithField("err", err).Info("get_candidate_by_email_err")
		return err
	}
	if candidate.Email == "" {
		logrus.WithField("email", pushStatus.Email).Warn("no_such_email")
		return errors.New("no_such_email")
	}
	if pushStatus.Status == 2 {
		// 修改成停止推送
		candidate.Status = models.CandidateStatusStop
	} else {
		// 修改成停止推送
		candidate.Status = models.CandidateStatusOk
	}
	err = models.UpdateCandidate(candidate)
	if err != nil {
		logrus.WithField("err", err).Error("update_candidate_err")
		return err
	}
	return err

}
