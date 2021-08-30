package brand_service

import (
	"github.com/Quons/matchmaker/models"
	log "github.com/sirupsen/logrus"
)

type Brand struct {
	ID   int64  `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
}

func (b *Brand) GetBrandList() ([]models.Brand, error) {
	list, err := models.GetBrandList()
	if err != nil {
		log.Error(err)
		return list, err
	}
	return list, nil
}

func (b *Brand) AddBrand() error {
	err := models.AddBrand(&models.Brand{Name: b.Name})
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (b *Brand) DeleteBrand() error {
	err := models.DeleteBrand(&models.Brand{ID: b.ID})
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
