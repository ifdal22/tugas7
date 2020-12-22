package controllers

import (
	"net/http"
	"tugas7/echo-rest/models"

	"github.com/labstack/echo"
)

//AddUser ...
func FetchAllUsers(c echo.Context) (err error) {

	result, err := models.FetchUsers()

	return c.JSON(http.StatusOK, result)
}

//AddSuppliers ...
func StoreUser(c echo.Context) (err error) {

	result, err := models.StoreUser(c)

	return c.JSON(http.StatusOK, result)
}

//UpdateUser ...
func UpdateUser(c echo.Context) (err error) {

	result, err := models.UpdateUser(c)

	return c.JSON(http.StatusOK, result)
}

//DeleteUser ...
func DeleteUser(c echo.Context) (err error) {

	result, err := models.DeleteUser(c)

	return c.JSON(http.StatusOK, result)
}
