package main

import (
	db "tugas7/echo-rest/db"
	routes "tugas7/echo-rest/routes"
)

func main() {

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":3000"))

}
