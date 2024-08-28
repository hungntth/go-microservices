package main

import "fmt"

func main(){
	fmt.Println("I am main function test change")

	app := fiber.New()

	app.Listen("localhost:9099")
}