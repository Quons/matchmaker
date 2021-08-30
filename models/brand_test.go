package models

import (
	"testing"
)

func TestGetBrandList(t *testing.T) {
	list, err := GetBrandList()
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%v", list)
}

func TestAddBrand(t *testing.T) {
	err := AddBrand(&Brand{Name: "别克"})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%v", "ok")
}
