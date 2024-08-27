package main

import "fmt"

func main(){
	fmt.Println("I am main function")

	app := fiber.New()

	app.Listen("localhost:9099")
}