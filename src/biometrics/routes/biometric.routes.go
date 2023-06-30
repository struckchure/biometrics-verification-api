package biometric_routes

import (
	app "pkg.struckchure.com/bv_api/src"
	biomoterics_handlers "pkg.struckchure.com/bv_api/src/biometrics/handlers"
)

func BiometricsRoutes() {
	router := app.GetApp()

	router.Post("/fingerprint/", biomoterics_handlers.FingerPrintMatchHandler)
}
