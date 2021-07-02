//+build wireinject

package main

import (
	"sushiApi/internal/db/repository"
	api "sushiApi/internal/http"
	"sushiApi/internal/usecase"

	"github.com/google/wire"
)

func InitServer() *api.Server {
	wire.Build(
		// ↓DBのコネクションを作る
		repository.NewBaseRepository,
		// ↓DBにコネクションでリポジトリを作る
		repositories,
		//　↓リポジトリを使ってusecaseを作る
		usecases,
		//　↓usecaseを使ってAPIを作る
		api.NewApi,
		//　↓APIを使ってserverを作る
		api.NewServer,
	)

	// 最終的にserverを返す
	return &api.Server{}
}

var repositories = wire.NewSet(
	repository.NewSushiData,
)
var usecases = wire.NewSet(
	usecase.NewSushi,
)

//// mock作るときはこれを使う
//func InitMockServer() *internal.Server {
//	wire.Build()
//
//	return &internal.Server{}
//}
