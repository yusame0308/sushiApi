package usecase

import (
	"encoding/json"
	"net/http"
	"sushiApi/internal/db/model"
	"sushiApi/internal/db/repository"
	"sushiApi/internal/http/gen"
	"sushiApi/pkg"

	"golang.org/x/xerrors"
)

type Sushi struct {
	sd repository.SushiDataInterface
}

func NewSushi(db repository.SushiDataInterface) *Sushi {
	return &Sushi{sd: db}
}

func (p *Sushi) AddSushi(param *gen.NewSushi) (*gen.Sushi, error) {
	// Array to String
	sozaiString, err := arrayToString(param.Sozai)
	if err != nil {
		return nil, pkg.NewHttpError(http.StatusBadRequest, pkg.BadRequestInvalidParam, err)
	}
	// Create
	sushiData := &model.SushiData{
		Name:  param.Name,
		Price: param.Price,
		Sozai: sozaiString,
	}
	if tx := p.sd.Create(sushiData); tx.Error != nil {
		return nil, pkg.NewHttpError(http.StatusBadRequest, pkg.BadRequestDatabaseTrouble, tx.Error)
	}
	return &gen.Sushi{
		Id:       sushiData.ID,
		NewSushi: *param,
	}, nil
}

func (p *Sushi) FindSushiById(id int64) (*gen.Sushi, error) {
	// データを取得
	m := new(model.SushiData)
	if tx := p.sd.First(m, id); tx.Error != nil {
		return nil, pkg.NewHttpError(http.StatusNotFound, pkg.NotFoundNoData, tx.Error)
	}
	// String to Array
	sozaiArray, err := stringToArray(m.Sozai)
	if err != nil {
		return nil, pkg.NewHttpError(http.StatusBadRequest, pkg.BadRequestInvalidParam, err)
	}

	sushi := &gen.Sushi{
		Id: id,
		NewSushi: gen.NewSushi{
			Name:  m.Name,
			Price: m.Price,
			Sozai: sozaiArray,
		},
	}
	return sushi, nil
}

func (p *Sushi) FindSushis(params gen.FindSushisParams) (*[]gen.Sushi, error) {
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
	if err := p.sd.Finds(order, limit, m); err != nil {
		return nil, pkg.NewHttpError(http.StatusBadRequest, pkg.BadRequestDatabaseTrouble, err)
	}

	var sushis []gen.Sushi
	for _, sushiData := range *m {
		// String to Array
		sozaiArray, err := stringToArray(sushiData.Sozai)
		if err != nil {
			return nil, pkg.NewHttpError(http.StatusBadRequest, pkg.BadRequestInvalidParam, err)
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
	return &sushis, nil
}

func arrayToString(array []string) (string, error) {
	b, err := json.Marshal(array)
	if err != nil {
		return "", xerrors.Errorf("Marshal failed: %+v", array)
	}
	return string(b), nil
}

func stringToArray(str string) ([]string, error) {
	b := []byte(str)
	sl := make([]string, 0)
	err := json.Unmarshal(b, &sl)
	if err != nil {
		return nil, xerrors.Errorf("Unmarshal failed: %s", str)
	}
	return sl, nil
}
