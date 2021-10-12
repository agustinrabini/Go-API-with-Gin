package main

import (
	"goGin/actions"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	//r.GET("/products", actions.ProductsList)
	//r.GET("/product/:id", actions.ShowProduct)
	//r.POST("/productAdd", actions.ProductAdd)

	//r.GET("/user", actions.GetUserByNameAndPhone)

	v1 := r.Group("/v1")
	{
		v1.GET("/products", actions.ProductsList)
		v1.GET("/product/:id", actions.ShowProduct)
		v1.POST("/productAdd", actions.ProductAdd)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/user", actions.GetUserByNameAndPhone)
	}

	r.Run()
}
