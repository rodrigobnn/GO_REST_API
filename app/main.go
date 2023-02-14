package main

import (
	"b2/controller"
	"b2/model"
	"fmt"

	_ "b2/docs"

	_ "github.com/swaggo/http-swagger" // http-swagger middleware
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.

func main() {
	model.Init()
	controller.Start()

	fmt.Println("Iniciou")

}
