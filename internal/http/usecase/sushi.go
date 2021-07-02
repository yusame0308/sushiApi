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
	newSushi := new(gen.NewSushi)
	if err := c.Bind(newSushi); err != nil {
		return sendError(c, http.StatusBadRequest, "Invalid format")
	}
	// Array to String
	sozaiString, err := arrayToString(c, newSushi.Sozai)
	if err != nil {
		return sendError(c, http.StatusBadRequest, err.Error())
	}
	// Create
	sushiData := &repository.SushiData{
		Name:  newSushi.Name,
		Price: newSushi.Price,
		Sozai: sozaiString,
	}
	if tx := p.db.Create(sushiData); tx.Error != nil {
		return sendError(c, http.StatusBadRequest, tx.Error.Error())
	}
	return c.JSON(http.StatusCreated, gen.Sushi{
		Id:       sushiData.ID,
		NewSushi: *newSushi,
	})
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
		return sendError(c, http.StatusBadRequest, err.Error())
	}
	sushi := &gen.Sushi{
		Id: id,
		NewSushi: gen.NewSushi{
			Name:  m.Name,
			Price: m.Price,
			Sozai: sozaiArray,
		},
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
