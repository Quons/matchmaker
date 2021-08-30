package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Attitude int

const (
	AttitudeUnknown Attitude = 0
	AttitudeAgree   Attitude = 1
	AttitudeUnAgree Attitude = 2
)

type ConfirmStatus int

const (
	ConfirmStatusUnknown ConfirmStatus = 0
	ConfirmStatusOk      ConfirmStatus = 1
)

type SendStatus int

const (
	SendStatusNo SendStatus = 0
	SendStatusOK SendStatus = 1
)

type MatchRecord struct {
	ID             int64         `gorm:"primary_key;column:id" json:"id"`
	MaleID         int64         `gorm:"column:male_id" json:"male_id"`
	FemaleID       int64         `gorm:"column:female_id" json:"female_id"`
	MaleEmail      string        `gorm:"column:male_email" json:"male_email"`
	FemaleEmail    string        `gorm:"column:female_email" json:"female_email"`
	CreateTime     int64         `gorm:"column:create_time" json:"create_time"`
	SendStatus     SendStatus    `gorm:"column:send_status" json:"send_status"`
	ConfirmStatus  ConfirmStatus `gorm:"column:confirm_status" json:"confirm_status"` // TODO 这里男女生都需要有个确认状态
	MatchResult    Attitude      `gorm:"column:match_result" json:"match_result"`
	MaleAttitude   Attitude      `gorm:"column:male_attitude" json:"male_attitude"`
	FemaleAttitude Attitude      `gorm:"column:female_attitude" json:"female_attitude"`
	MaleNote       string        `gorm:"column:male_note" json:"male_note"`
	FemaleNote     string        `gorm:"column:female_note" json:"female_note"`
}

func GetMatchRecordList() ([]MatchRecord, error) {
	var recordList []MatchRecord
	err := readDB().Find(&recordList).Error
	if err == gorm.ErrRecordNotFound {
		return recordList, nil
	}
	return recordList, err
}

func AddMatchRecord(record *MatchRecord) error {
	record.CreateTime = time.Now().Unix()
	if err := WriteDB().Create(record).Error; err != nil {
		return err
	}
	return nil
}

func UpdateMatchRecord(id int64) error {
	if err := WriteDB().Table("match_record").Where("id=?", id).Update("send_status", SendStatusOK).Error; err != nil {
		return err
	}
	return nil
}
