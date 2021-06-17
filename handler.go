package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Handler struct {
	storage Storage
}

func NewHandler(storage Storage) *Handler {
	return &Handler{
		storage: storage,
	}
}

func (h *Handler) CreateUser(c echo.Context) error {
	var user User
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		fmt.Printf("Failed to bind user: %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	h.storage.Insert(&user)

	return c.JSON(http.StatusOK, "Success")
}

func (h Handler) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("Failed to convert id param to int: %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := h.storage.Get(uint(id))
	if err != nil {
		fmt.Printf("Failed to get user: %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h Handler) GetUsers(c echo.Context) error {
	users, err := h.storage.GetAll()
	if err != nil {
		fmt.Printf("Failed to get all users: %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if len(users) < 1 {
		fmt.Println("Failed to get all users: slice is empty")
		return c.JSON(http.StatusBadRequest, "No users")
	}
	return c.JSON(http.StatusOK, users)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("Failed to convert id param to int: %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var user User
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		fmt.Printf("Failed to bind user: %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.storage.Update(uint(id), user)
	if err != nil {
		fmt.Printf("Failed to update user: %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Success")
}

func (h *Handler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("Failed to convert id param to int: %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = h.storage.Delete(uint(id))
	if err != nil {
		fmt.Printf("Failed to delete user: %s\n", err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, "Success")
}
