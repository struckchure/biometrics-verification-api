package main

import (
	app "struckchure.bv_api/src"
	"struckchure.bv_api/src/biometrics"
)

func main() {
	app := app.GetApp()

	biometrics.BiometricsModule()

	app.Listen(":3000")
}
