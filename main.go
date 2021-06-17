package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
)

func main() {
	file, err := os.OpenFile("test.json", os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Can't open file")
		return
	}
	defer file.Close()

	userstorage := NewUserStorage(file)
	handler := NewHandler(userstorage)

	e := echo.New()

	e.POST("/user", handler.CreateUser)
	e.GET("/user/", handler.GetUsers)
	e.GET("/user/:id", handler.GetUser)
	e.PUT("/user/:id", handler.UpdateUser)
	e.DELETE("/user/:id", handler.DeleteUser)

	e.Logger.Fatal(e.Start(":8080"))
}
