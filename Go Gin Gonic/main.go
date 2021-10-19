package main

import (
	"goGin/actions"
	"goGin/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	//key := os.Getenv("SEC_KEY_GETAWAY")

	r := gin.Default()

	//r.Use(middleware.LogIn, middleware.LogOut)

	//store := cookie.NewStore([]byte(key))
	//r.Use(sessions.Sessions("mysession", store))

	v1 := r.Group("/product")
	{
		v1.GET("/all", actions.ProductsList)
		v1.GET("/product/:id", actions.ShowProduct)
		v1.POST("/bind", actions.Bind)
		v1.POST("/productAdd", actions.ProductAdd)
	}

	v2 := r.Group("/user")
	{
		v2.GET("/userInfo", actions.GetUserByNameAndPhone)
		v2.POST("/login", middleware.LogIn)
		v2.GET("/logOut", middleware.LogOut)
	}

	r.GET("/filter", actions.FilterExample)

	r.Run(":8080")
}
