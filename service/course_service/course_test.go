package course_service

import (
	"github.com/Quons/matchmaker/models"
	"github.com/Quons/matchmaker/pkg/setting"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func init() {
	setting.Setup("dev")
	models.Setup()
}

type MockedObj struct {
	mock.Mock
}

func (m *MockedObj) Speek(word string) (s string) {
	arg := m.Called(word)
	return arg.String(0)
}

func TestAddArticleAndTag(t *testing.T) {
	testObj := new(MockedObj)
	testObj.On("Speek", mock.Anything).Return("hhhhh")
	s := GetTest(testObj)
	t.Log(s)
	assert.Equal(t, "hhhhh", s)
	testObj.AssertCalled(t, "Speek", "hello")
}
