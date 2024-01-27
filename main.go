package main

import (
	windows "borzoiffmpeg/windows"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	initApp := app.New()
	w := windows.Windows{App: initApp}

	mainWindow := initApp.NewWindow("borzoi ffmpeg")

	resizeButton := widget.NewButton("RESIZE VIDEO OR IMAGE", func() {
		w.ResizeWindow()
	})

	trimButton := widget.NewButton("TRIM VIDEO", func() {
		w.TrimWindow()
	})

	videoToGifButton := widget.NewButton("VIDEO TO GIF", func() {
		w.VideoToGifWindow()
	})

	extractAudioButton := widget.NewButton("EXTRACT AUDIO FROM VIDEO", func() {
		w.ExtractAudioWindow()
	})

	convertVideoButton := widget.NewButton("CONVERT VIDEO", func() {
		w.ConvertVideoWindow()
	})

	resizeButton.Move(fyne.NewPos(0, 0))
	trimButton.Move(fyne.NewPos(490, 0))
	videoToGifButton.Move(fyne.NewPos(980, 0))
	extractAudioButton.Move(fyne.NewPos(0, 860))
	convertVideoButton.Move(fyne.NewPos(980, 860))

	buttonSize := fyne.NewSize(300, 100)

	resizeButton.Resize(buttonSize)
	trimButton.Resize(buttonSize)
	videoToGifButton.Resize(buttonSize)
	extractAudioButton.Resize(buttonSize)
	convertVideoButton.Resize(buttonSize)

	backgroundImage := windows.GetImage("images/gaming.jpg")
	backgroundImage.FillMode = canvas.ImageFillOriginal
	backgroundImageContainer := container.NewHBox(backgroundImage)

	buttonContainer := container.NewWithoutLayout(
		resizeButton,
		trimButton,
		videoToGifButton,
		extractAudioButton,
		convertVideoButton,
	)

	content := container.NewStack(
		backgroundImageContainer,
		buttonContainer,
	)

	mainWindow.SetMaster()
	mainWindow.SetContent(content)
	mainWindow.Resize(mainWindow.Canvas().Size())
	mainWindow.SetFixedSize(true)
	mainWindow.ShowAndRun()
}
