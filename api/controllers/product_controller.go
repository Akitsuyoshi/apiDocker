package controllers

import (
	"os"

	"github.com/Akitsuyoshi/apiDocker/api/models"
	"github.com/martini-contrib/render"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type (
	ProductController struct {
		session *mgo.Session
	}
)

func NewProductController(s *mgo.Session) *ProductController {
	return &ProductController(s)
}

func (pc *ProductController) GetAllProducts(r render.Render) {
	products := []models.Product{}
	session := pc.session.DB(os.Getenv("DB_NAME")).C("products")

	err := session.Find(nil).Limit(100).All(&products)

	if err != nil {
		panic(err)
	}

	r.JSON(200, products)
}

func (pc *ProductController) PostProduct(product models.Product, r render.Render) {
	session := pc.session.DB(os.Getenv("DB_NAME")).C("products")

	product.Id = bson.NewObjectId()
	product.Title = product.Title
	product.Description = product.Description
	product.Price = product.Price
	session.Insert(product)

	r.JSON(201, product)
}
