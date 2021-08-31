package matcher

import (
	"github.com/Quons/matchmaker/models"
	"testing"
)

func TestAddCandidate(t *testing.T) {
	cList := []Candidate{
		{
			Email:  "cugses@qq.com",
			Gender: models.GenderMale,
			Age:    1993,
		},
		{
			Email:  "cugjyb@163.com",
			Gender: models.GenderFemale,
			Age:    1993,
		},
	}
	for _, candidate := range cList {
		err := AddCandidate(candidate)
		if err != nil {
			t.Error(err)
			return
		}
	}
}
