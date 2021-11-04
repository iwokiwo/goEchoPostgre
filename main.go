package main

import (
	"hotels/db"
	"hotels/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8085"))
}
