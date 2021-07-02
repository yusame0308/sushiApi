package repository

type SushiData struct {
	ID    int64  `gorm:"primaryKey"`
	Name  string `gorm:"unique"`
	Price int
	Sozai string
}
