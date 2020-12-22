package controllers

import (
	"net/http"
	"tugas7/echo-rest/models"

	"github.com/labstack/echo"
)

//FetchAllSuppliers ...
func FetchAllSuppliers(c echo.Context) (err error) {

	result, err := models.FetchSuppliers()

	return c.JSON(http.StatusOK, result)
}

//AddSuppliers ...
func AddSuppliers(c echo.Context) (err error) {

	result, err := models.AddSuppliers(c)

	return c.JSON(http.StatusOK, result)
}

//UpdateSuppliers ...
func UpdateSuppliers(c echo.Context) (err error) {

	result, err := models.UpdateSuppliers(c)

	return c.JSON(http.StatusOK, result)
}

//DeleteSuppliers ...
func DeleteSuppliers(c echo.Context) (err error) {

	result, err := models.DeleteSuppliers(c)

	return c.JSON(http.StatusOK, result)
}
