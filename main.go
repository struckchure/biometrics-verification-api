package main

import (
	app "pkg.struckchure.com/bv_api/src"
	"pkg.struckchure.com/bv_api/src/biometrics"
)

func main() {
	app := app.GetApp()

	biometrics.BiometricsModule()

	app.Listen(":3000")
}
