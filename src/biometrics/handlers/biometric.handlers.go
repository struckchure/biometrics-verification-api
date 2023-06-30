package biomoterics_handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	biometric_service "struckchure.bv_api/src/biometrics/services"
)

// TODO: add function to utility
func getBasePath() string {
	f, _ := os.Getwd()
	projectPath := fmt.Sprintf("%s/%s/", filepath.Dir(f), filepath.Base(f))

	return projectPath
}

// TODO: raise errors where necessary
func FingerPrintMatchHandler(c *fiber.Ctx) error {
	target, _ := c.FormFile("target")
	sample, _ := c.FormFile("sample")

	if (target == nil) || (sample == nil) {
		c.SendStatus(400)

		return c.JSON(fiber.Map{"error": "sample and target are required"})
	}

	projectPath := getBasePath()

	// TODO: add to file upload utility
	os.Mkdir(fmt.Sprintf("%s/%s", projectPath, "/tmp/"), 0755)

	// TODO: write utilty to upload file
	targetName, _ := uuid.NewUUID()
	targetPath := fmt.Sprintf(
		"%s/tmp/%s.%s",
		projectPath,
		targetName.String(),
		strings.Split(target.Filename, ".")[1],
	)
	c.SaveFile(target, targetPath)

	// TODO: write utilty to upload file
	sampleName, _ := uuid.NewUUID()
	samplePath := fmt.Sprintf(
		"%s/tmp/%s.%s",
		projectPath,
		sampleName.String(),
		strings.Split(sample.Filename, ".")[1],
	)
	c.SaveFile(sample, samplePath)

	matchScore := biometric_service.FingerPrintMatch(targetPath, samplePath)

	// TODO: add to util for save file to remove after 2 minutes (*)
	os.Remove(targetPath)
	os.Remove(samplePath)

	return c.JSON(fiber.Map{"matchScore": matchScore})
}
