package routes

import (
	"tugas7/echo-rest/controllers"

	"github.com/labstack/echo"
)

//SuppliersRoute ...
func SuppliersRoute(g *echo.Group) {

	g.POST("/list", controllers.FetchAllSuppliers)

	g.POST("/add", controllers.AddSuppliers)

	g.POST("/update", controllers.UpdateSuppliers)

	g.POST("/delete", controllers.DeleteSuppliers)

}
