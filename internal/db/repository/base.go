package repository

import (
	"sushiApi/internal/db/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BaseRepository struct {
	*gorm.DB
}

func NewBaseRepository() *BaseRepository {
	// mysql connection
	dsn := "docker:docker@tcp(127.0.0.1:3306)/sushi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	// Migrate the schema
	if err := db.AutoMigrate(&model.SushiData{}); err != nil {
		panic(err.Error())
	}

	return &BaseRepository{DB: db}
}

func (r *BaseRepository) Finds(order string, limit int, dst interface{}) error {
	tx := r.DB
	if len(order) > 0 {
		tx = tx.Order(order)
	}
	if limit > 0 {
		tx = tx.Limit(limit)
	}
	if tx := tx.Find(dst); tx.Error != nil {
		return tx.Error
	}

	return nil

}
