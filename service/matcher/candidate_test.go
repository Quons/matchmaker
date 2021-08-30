package matcher

import (
	"github.com/Quons/matchmaker/models"
	"testing"
)

func TestAddCandidate(t *testing.T) {
	c := Candidate{
		Email:  "cugjyb@qq.com",
		Gender: models.GenderMale,
	}
	err := AddCandidate(c)
	if err != nil {
		t.Error(err)
		return
	}

}
