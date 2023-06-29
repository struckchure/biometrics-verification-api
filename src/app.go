package app

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/qinains/fastergoding"
)

var (
	app     *fiber.App
	appOnce sync.Once
)

func GetApp() *fiber.App {
	appOnce.Do(func() {
		fastergoding.Run()

		app = fiber.New()
		app.Use(logger.New())
	})

	return app
}
