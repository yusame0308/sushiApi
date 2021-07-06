// +build !mock

package main

import api "sushiApi/internal/http"

func NewServer() *api.Server {
	return InitServer()
}
