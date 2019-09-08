package main

import (
	"github.com/alimtegar/nggading-car-rental-system/app"
)

func main() {
	app := &app.App{}

	app.Initialize()
	app.Run(":3001")
}
