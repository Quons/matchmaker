package models

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAddCandidate(t *testing.T) {

	c := Candidate{
		Gender:     1,
		Email:      "cugjyb@qq.com",
		Statue:     1,
		UpdateTime: 1,
		CreateTime: 1,
	}
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			c.Gender = GenderMale
			c.Email = fmt.Sprintf("m_%v", i/2)
		} else {
			c.Gender = GenderFemale
			c.Email = fmt.Sprintf("f_%v", i/2)
		}
		err := AddCandidate(c)
		if err != nil {
			t.Error(err)
			return
		}
	}

}

func TestGetCandidateList(t *testing.T) {
	got, err := GetCandidateList(GenderMale)
	if err != nil {
		t.Error(err)
		return
	}
	marshal, err := json.Marshal(got)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(marshal))
}

func TestGetCandidateByEmail(t *testing.T) {
	gotCandidate, err := GetCandidateByEmail("cugses@qq.com")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(gotCandidate)
}

func TestUpdateCandidate(t *testing.T) {
	c := Candidate{
		Gender:     2,
		Email:      "cugjyb@163.com",
		Statue:     1,
		UpdateTime: 1,
		CreateTime: 1211111,
	}
	err := UpdateCandidate(c)
	if err != nil {
		t.Error(err)
		return
	}

}
