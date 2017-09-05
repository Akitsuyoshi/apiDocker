package main

import (
	"github.com/Akitsuyoshi/apiDocker/api/controllers"
	"github.com/Akitsuyoshi/apiDocker/api/models"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Map(models.Database())
	m.Use(render.Renderer())

	pc := controllers.NewProductController(models.Database())

	m.Get("/products", binding.Bind(models.Product{}), pc.GetAllProducts)
	m.Post("/products".binding.Bind(models.Product{}), pc.PostProduct)
	m.Run()
}
