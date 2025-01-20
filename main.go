package main

import (
	"demo/router"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


//I tried to structure the project and followed some best practices to showcase how can we just stared with..

func main() {
	fmt.Println("Welcome to Studio!")
	r := gin.Default()
	router.Routes(r)
	fmt.Println("Server is listening on port 8000")
	err:=http.ListenAndServe(":8000",r)
	if err!=nil{
		log.Fatal("could not start server",err)
	}
}
