package repository

import (
	"sushiApi/internal/db/model"

	"gorm.io/gorm"
)

type MockSushiData struct {
}

func (m MockSushiData) First(i interface{}, _ ...interface{}) *gorm.DB {
	i = model.SushiData{
		ID:    99,
		Name:  "mock",
		Price: 100,
		Sozai: `["納豆","のり","しゃり"]`,
	}
	return new(gorm.DB)
}

func (m MockSushiData) Create(i interface{}) *gorm.DB {
	i = model.SushiData{
		ID:    99,
		Name:  "mock",
		Price: 100,
		Sozai: "mock",
	}
	return new(gorm.DB)
}

func (m MockSushiData) Finds(_ string, _ int, i interface{}) error {
	i = []model.SushiData{
		{
			ID:    99,
			Name:  "mock",
			Price: 100,
			Sozai: "mock",
		},
	}
	return nil
}

func NewMockSushiData() SushiDataInterface {
	return &MockSushiData{}
}
