package api

import (
	"sushiApi/internal/http/gen"

	om "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	api *Api
}

func NewServer(api *Api) *Server {
	return &Server{api: api}
}

func (r *Server) Run() {
	e := echo.New()
	// リクエストIDの設定
	e.Use(middleware.RequestID())
	// loggerの設定
	e.Use(middleware.Logger())
	// recoverの設定
	e.Use(middleware.Recover())

	// validator
	spec, err := gen.GetSwagger()
	if err != nil {
		panic(err)
	}
	e.Use(om.OapiRequestValidator(spec))

	gen.RegisterHandlers(e, r.api)
	e.Logger.Fatal(e.Start(":1232"))
}
