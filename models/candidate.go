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

type Candidate struct {
	ID         int64  `gorm:"primary_key;column:id" json:"id"`
	Gender     Gender `gorm:"column:gender" json:"gender"`
	Email      string `gorm:"column:email" json:"email"`
	Statue     int    `gorm:"column:statue" json:"statue"`
	UpdateTime int64  `gorm:"column:update_time" json:"update_time"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
}

func GetCandidateList(gender Gender) ([]Candidate, error) {
	var CandidateList []Candidate
	err := readDB().Where("gender = ?", gender).Find(&CandidateList).Error
	if err == gorm.ErrRecordNotFound {
		return CandidateList, nil
	}
	return CandidateList, err
}

func AddCandidate(candidate Candidate) error {
	candidate.CreateTime = time.Now().Unix()
	if err := WriteDB().Create(&candidate).Error; err != nil {
		return err
	}
	return nil
}
