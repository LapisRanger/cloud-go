package service

import (
	"github.com/go-martini/martini"
)

// NewServer configures and returns a Server.
func NewServer() *martini.ClassicMartini {
	//return instance
	m := martini.Classic()

	//route
	m.Get("/", func(params martini.Params) string {
		return "hello world"
	})

	return m
}
