package repository

import "gorm.io/gorm"

type baseInterface interface {
	Create(interface{}) *gorm.DB
	Finds(string, int, interface{}) error
	First(interface{}, ...interface{}) *gorm.DB
}
