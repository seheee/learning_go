package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	//Get gin's default engine
	r := gin.Default()
	h,_ := NewHandler()

	r.GET("/products", h.GetProducts)

	r.GET("/promos", h.GetPromos)

	/*
	r.POST("/users/signin", h.SignIn)

	r.POST("/users", h.AddUser)

	r.POST("/user/:id/signout", h.SignOut)

	r.GET("/user/:id/orders", h.GetOrders)

	r.POST("/users/charge", h.Charge)
	*/

	// grouping routes
	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id/signout", h.SignOut)
		userGroup.GET("/:id/orders", h.GetOrders)
	}
	usersGroup := r.Group("/users")
	{
		usersGroup.POST("/charge", h.Charge)
		usersGroup.POST("/signin", h.SignIn)
		userGroup.POST("", h.AddUser)
	}

	// start server
	return r.Run(address)
} 