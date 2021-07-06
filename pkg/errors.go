package pkg

import (
	"encoding/json"
	"net/http"
	"sushiApi/internal/http/gen"

	"github.com/labstack/echo/v4"
)

const (
	BadRequestCanNotBind      string = "400_CanNotBind"
	BadRequestInvalidParam           = "400_InvalidParam"
	BadRequestDatabaseTrouble        = "400_DatabaseTrouble"
	NotFoundNoData                   = "404_NoData"
)

type HttpError struct {
	HttpStatus int
	Code       string
	Message    string
}

func NewHttpError(httpStatus int, code string, err error) HttpError {
	return HttpError{HttpStatus: httpStatus, Code: code, Message: err.Error()}
}

func NewBindError(err error) HttpError {
	return HttpError{HttpStatus: http.StatusBadRequest, Code: BadRequestCanNotBind, Message: err.Error()}
}

func (r HttpError) Error() string {
	v, _ := json.Marshal(r)
	return string(v)
}

func SendError(ctx echo.Context, err error) error {
	e := err.(HttpError)
	return ctx.JSON(e.HttpStatus, gen.Error{
		Code:    int32(e.HttpStatus),
		Message: e.Error(),
	})
}
