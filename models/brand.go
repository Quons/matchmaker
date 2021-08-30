package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Brand struct {
	ID      int64     `gorm:"primary_key;column:id" json:"id"`
	Name    string    `gorm:"column:name" json:"name"`
	AddTime time.Time `gorm:"column:addTime" json:"addTime"`
}

func GetBrandList() ([]Brand, error) {
	var brandList []Brand
	err := readDB().Find(&brandList).Error
	if err == gorm.ErrRecordNotFound {
		return brandList, nil
	}
	return brandList, err
}

func AddBrand(brand *Brand) error {
	brand.AddTime = time.Now()
	if err := WriteDB().Create(brand).Error; err != nil {
		return err
	}
	return nil
}

func DeleteBrand(brand *Brand) error {
	if err := WriteDB().Where("id = ?", brand.ID).Delete(&Brand{}).Error; err != nil {
		return err
	}
	return nil
}
