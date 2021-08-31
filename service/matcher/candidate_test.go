package matcher

import (
	"github.com/Quons/matchmaker/models"
	"testing"
)

func TestAddCandidate(t *testing.T) {
	cList := []Candidate{
		{
			Email:  "11",
			Gender: models.GenderMale,
			Age:    1993,
		},
		{
			Email:  "22",
			Gender: models.GenderMale,
			Age:    1993,
		},
		{
			Email:  "33",
			Gender: models.GenderMale,
			Age:    1993,
		},
		{
			Email:  "44",
			Gender: models.GenderMale,
			Age:    1993,
		},
		{
			Email:  "55",
			Gender: models.GenderMale,
			Age:    1993,
		},
		{
			Email:  "1",
			Gender: models.GenderFemale,
			Age:    1993,
		},
		{
			Email:  "2",
			Gender: models.GenderFemale,
			Age:    1993,
		},
		{
			Email:  "3",
			Gender: models.GenderFemale,
			Age:    1993,
		},
		{
			Email:  "4",
			Gender: models.GenderFemale,
			Age:    1993,
		},
		{
			Email:  "5",
			Gender: models.GenderFemale,
			Age:    1993,
		},
		{
			Email:  "6",
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
