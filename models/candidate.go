package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Gender int

const (
	GenderMale   Gender = 1
	GenderFemale Gender = 2
)

type CandidateStatus int

const (
	CandidateStatusOk   CandidateStatus = 1
	CandidateStatusStop CandidateStatus = 2
)

type Candidate struct {
	ID         int64           `gorm:"primary_key;column:id" json:"id"`
	Gender     Gender          `gorm:"column:gender" json:"gender"`
	Email      string          `gorm:"column:email" json:"email"`
	Age        int64           `gorm:"column:age" json:"age"`
	Status     CandidateStatus `gorm:"column:status" json:"status"`
	UpdateTime int64           `gorm:"column:update_time" json:"update_time"`
	CreateTime int64           `gorm:"column:create_time" json:"create_time"`
}

func GetCandidateList(gender Gender) ([]Candidate, error) {
	var CandidateList []Candidate
	err := readDB().Where("gender = ?", gender).Where("status=?", CandidateStatusOk).Find(&CandidateList).Error
	if err == gorm.ErrRecordNotFound {
		return CandidateList, nil
	}
	return CandidateList, err
}

func GetCandidateByEmail(email string) (candidate Candidate, err error) {
	err = readDB().Where("email = ?", email).Find(&candidate).Error
	if err == gorm.ErrRecordNotFound {
		return candidate, nil
	}
	return candidate, err
}

func AddCandidate(candidate Candidate) error {
	candidate.CreateTime = time.Now().Unix()
	if err := WriteDB().Create(&candidate).Error; err != nil {
		return err
	}
	return nil
}

func UpdateCandidate(candidate Candidate) error {
	candidate.UpdateTime = time.Now().Unix()
	candidate.ID = 0
	if err := WriteDB().Table("candidate").Where("email=?", candidate.Email).Update(&candidate).Error; err != nil {
		return err
	}
	return nil
}
