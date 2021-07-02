package usecase

import (
	"encoding/json"
	"net/http"
	"sushiApi/internal/db/model"
	"sushiApi/internal/db/repository"
	"sushiApi/internal/http/gen"

	"github.com/labstack/echo/v4"
)

type Sushi struct {
	db *repository.SushiData
}

func NewSushi(db *repository.SushiData) *Sushi {
	return &Sushi{db: db}
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
	sushiData := &model.SushiData{
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
	m := new(model.SushiData)
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

func (p *Sushi) FindSushis(c echo.Context, params gen.FindSushisParams) error {
	// データを取得
	var (
		order string
		limit int
	)
	if params.Asc != nil {
		if *params.Asc {
			order = "id ASC"
		} else {
			order = "id DESC"
		}
	}
	if params.Limit != nil {
		limit = int(*params.Limit)
	}
	m := new([]model.SushiData)
	if err := p.db.Finds(order, limit, m); err != nil {
		return sendError(c, http.StatusNotFound, err.Error())
	}

	var sushis []gen.Sushi
	for _, sushiData := range *m {
		// String to Array
		sozaiArray, err := stringToArray(c, sushiData.Sozai)
		if err != nil {
			return sendError(c, http.StatusBadRequest, err.Error())
		}
		newSushi := gen.Sushi{
			Id: sushiData.ID,
			NewSushi: gen.NewSushi{
				Name:  sushiData.Name,
				Price: sushiData.Price,
				Sozai: sozaiArray,
			},
		}
		sushis = append(sushis, newSushi)
	}
	return c.JSON(http.StatusOK, sushis)
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
