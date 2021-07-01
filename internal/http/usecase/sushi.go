package usecase

import (
	"encoding/json"
	"net/http"
	"sushiApi/internal/http/gen"
	"sushiApi/internal/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Sushi struct {
	db *gorm.DB
}

func NewSushi(db *gorm.DB) *Sushi {
	return &Sushi{
		db: db,
	}
}

func (p *Sushi) AddSushi(c echo.Context) error {
	// リクエストを取得
	sushi := new(gen.Sushi)
	if err := c.Bind(sushi); err != nil {
		return sendError(c, http.StatusBadRequest, "Invalid format")
	}
	// Array to String
	sozaiString, err := arrayToString(c, sushi.Sozai)
	if err != nil {
		return err
	}
	// Create
	p.db.Create(&repository.SushiData{
		Name:  sushi.Name,
		Price: sushi.Price,
		Sozai: sozaiString,
	})
	return c.JSON(http.StatusCreated, sushi)
}

func (p *Sushi) FindSushiById(c echo.Context, id int64) error {
	// データを取得
	m := new(repository.SushiData)
	if tx := p.db.First(m, id); tx.Error != nil {
		return sendError(c, http.StatusNotFound, tx.Error.Error())
	}
	// String to Array
	sozaiArray, err := stringToArray(c, m.Sozai)
	if err != nil {
		return err
	}
	sushi := &gen.Sushi{
		Name:  m.Name,
		Price: m.Price,
		Sozai: sozaiArray,
	}
	return c.JSON(http.StatusOK, sushi)
}

func arrayToString(c echo.Context, array []string) (string, error) {
	b, err := json.Marshal(array)
	if err != nil {
		return "", sendError(c, http.StatusBadRequest, "Invalid format")
	}
	return string(b), nil
}

func stringToArray(c echo.Context, str string) ([]string, error) {
	b := []byte(str)
	sl := make([]string, 0)
	err := json.Unmarshal(b, &sl)
	if err != nil {
		return nil, sendError(c, http.StatusBadRequest, "Invalid format")
	}
	return sl, nil
}
