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
	projectPath := getBasePath()

	// TODO: add to file upload utility
	os.Mkdir(fmt.Sprintf("%s/%s", projectPath, "/tmp/"), 0755)

	// TODO: write utilty to upload file
	img1, _ := c.FormFile("target")
	img1Name, _ := uuid.NewUUID()
	img1Path := fmt.Sprintf(
		"%s/tmp/%s.%s",
		projectPath,
		img1Name.String(),
		strings.Split(img1.Filename, ".")[1],
	)
	c.SaveFile(img1, img1Path)

	// TODO: write utilty to upload file
	img2, _ := c.FormFile("sample")
	img2Name, _ := uuid.NewUUID()
	img2Path := fmt.Sprintf(
		"%s/tmp/%s.%s",
		projectPath,
		img2Name.String(),
		strings.Split(img2.Filename, ".")[1],
	)
	c.SaveFile(img2, img2Path)

	matchScore := biometric_service.FingerPrintMatch(img1Path, img2Path)

	// TODO: add to util for save file to remove after 2 minutes (*)
	os.Remove(img1Path)
	os.Remove(img2Path)

	return c.JSON(fiber.Map{"matchScore": matchScore})
}
