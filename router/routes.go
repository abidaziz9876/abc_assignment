package router

import (
	"demo/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	//this is my router package where all my different routes or endpoint is defined
	router.POST("/classes", controllers.CreateClassesController())
	router.GET("/bookings", controllers.CreateBooking())
	router.GET("/getclasses", controllers.GetAllClassesController())
}
