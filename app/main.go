package main

import (
	"b2/controller"
	"b2/model"
	"fmt"
)

func main() {
	model.Init()
	controller.Start()

	fmt.Println("Iniciou")

}
