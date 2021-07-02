package api

import (
	"sushiApi/internal/http/gen"
	"sushiApi/internal/usecase"

	"github.com/labstack/echo/v4"
)

type Api struct {
	sushi *usecase.Sushi
}

func NewApi(sushi *usecase.Sushi) *Api {
	return &Api{sushi: sushi}
}

var _ gen.ServerInterface = (*Api)(nil)

func (p *Api) FindSushis(ctx echo.Context, params gen.FindSushisParams) error {
	return p.sushi.FindSushis(ctx, params)
}

func (p *Api) AddSushi(ctx echo.Context) error {
	return p.sushi.AddSushi(ctx)
}

func (p *Api) FindSushiById(ctx echo.Context, id int64) error {
	return p.sushi.FindSushiById(ctx, id)
}
