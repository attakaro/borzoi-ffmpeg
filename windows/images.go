package windows

import (
	"bytes"
	"embed"
	"image"

	"fyne.io/fyne/v2/canvas"
)

//go:embed images/*.jpg
var Images embed.FS

func GetImage(imgPath string) *canvas.Image {
	imageBytes, err := Images.ReadFile(imgPath)
	if err != nil {
		panic(err)
	}

	image, _, err := image.Decode(bytes.NewBuffer(imageBytes))
	if err != nil {
		panic(err)
	}

	mainMenuCanvas := canvas.NewImageFromImage(image)

	return mainMenuCanvas
}
