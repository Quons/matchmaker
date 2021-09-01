package models

import (
	"testing"
)

func TestAddMatchRecord(t *testing.T) {
	record := &MatchRecord{
		MaleID:           1,
		FemaleID:         1,
		MaleEmail:        "cugjyb@163.com",
		FemaleEmail:      "cugses@qq.com",
		CreateTime:       1,
		MaleSendStatus:   1,
		FemaleSendStatus: 1,
		ConfirmStatus:    ConfirmStatusOk,
		MatchResult:      AttitudeUnAgree,
		MaleAttitude:     AttitudeUnAgree,
		FemaleAttitude:   AttitudeUnAgree,
		MaleNote:         "聊的挺不错的",
		FemaleNote:       "这人有点无聊",
	}
	err := AddMatchRecord(record)

	if err != nil {
		t.Error(err)
		return
	}

}

func TestGetMatchRecordList(t *testing.T) {

	got, err := GetMatchRecordList()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(got)
}

func TestUpdateMatchRecord(t *testing.T) {

	err := UpdateMatchRecord(3, GenderMale)
	if err != nil {
		t.Error(err)
		return
	}

}

func TestGetFailedMatchRecordList(t *testing.T) {

	got, err := GetFailedMatchRecordList()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(got)
}
