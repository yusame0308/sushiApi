package repository

import (
	"sushiApi/internal/db/model"

	"gorm.io/gorm"
)

type MockSushiData struct {
}

func (m MockSushiData) First(i interface{}, _ ...interface{}) *gorm.DB {
	p := i.(*model.SushiData)
	p.ID = 99
	p.Name = "mock"
	p.Price = 99
	p.Sozai = `["納豆","のり","しゃり"]`
	return new(gorm.DB)
}

func (m MockSushiData) Create(i interface{}) *gorm.DB {
	p := i.(*model.SushiData)
	p.ID = 99
	p.Name = "mock"
	p.Price = 99
	p.Sozai = `["納豆","のり","しゃり"]`
	return new(gorm.DB)
}

func (m MockSushiData) Finds(_ string, _ int, i interface{}) error {
	p := i.(*[]model.SushiData)
	*p = append(*p, model.SushiData{
		ID:    99,
		Name:  "mock",
		Price: 99,
		Sozai: `["納豆","のり","しゃり"]`,
	})
	return nil
}

func NewMockSushiData() SushiDataInterface {
	return &MockSushiData{}
}
