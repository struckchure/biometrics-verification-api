package biometric_routes

import (
	app "struckchure.bv_api/src"
	biomoterics_handlers "struckchure.bv_api/src/biometrics/handlers"
)

func BiometricsRoutes() {
	router := app.GetApp()

	router.Post("/fingerprint/", biomoterics_handlers.FingerPrintMatchHandler)
}
