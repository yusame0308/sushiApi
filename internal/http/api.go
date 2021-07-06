package api

import (
	"net/http"
	"sushiApi/internal/http/gen"
	"sushiApi/internal/usecase"
	"sushiApi/pkg"

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
	res, err := p.sushi.FindSushis(params)
	if err != nil {
		return pkg.SendError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (p *Api) AddSushi(ctx echo.Context) error {
	// リクエストを取得
	newSushi := new(gen.NewSushi)
	if err := ctx.Bind(newSushi); err != nil {
		return pkg.SendError(ctx, pkg.NewBindError(err))
	}
	res, err := p.sushi.AddSushi(newSushi)
	if err != nil {
		return pkg.SendError(ctx, err)
	}

	return ctx.JSON(http.StatusOK, res)
}

func (p *Api) FindSushiById(ctx echo.Context, id int64) error {
	res, err := p.sushi.FindSushiById(id)
	if err != nil {
		return pkg.SendError(ctx, err)
	}
	return ctx.JSON(http.StatusOK, res)
}
