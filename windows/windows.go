package windows

import (
	borzoi "borzoiffmpeg/ffmpeg"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Windows struct {
	App fyne.App
}

func (w *Windows) ResizeWindow() {
	resizeWindow := w.App.NewWindow("Resize video or image")
	resizeWindow.Resize(fyne.NewSize(500, 500))

	mainLabel := widget.NewLabel("Status: not running")

	selectFileLabel := widget.NewLabel("Select video or image:")

	inputFilePath := ""

	selectFileButton := widget.NewButton("Select file", func() {
		dialog.ShowFileOpen(func(response fyne.URIReadCloser, err error) {
			if err == nil && response != nil {
				inputFilePath = response.URI().String()
				fileName := getFileNameWithExt(inputFilePath)
				mainLabel.Text = fmt.Sprintf("Status: file selected(%s)", fileName)
				mainLabel.Refresh()
				response.Close()
			}
		}, resizeWindow)
	})

	widthLabel := widget.NewLabel("Width")
	widthEntry := widget.NewEntry()

	heightLabel := widget.NewLabel("Height")
	heightEntry := widget.NewEntry()

	runButton := widget.NewButton("Run", func() {
		switch {
		case inputFilePath == "":
			mainLabel.Text = "Status: file not selected"
			mainLabel.Refresh()
		case widthEntry.Text == "":
			mainLabel.Text = "Status: width not specified"
			mainLabel.Refresh()
		case heightEntry.Text == "":
			mainLabel.Text = "Status: height not specified"
			mainLabel.Refresh()
		default:
			outputFilePath := getOutputFilePath(inputFilePath, "_resized", "")

			mainLabel.Text = "Status: processing, please wait..."
			mainLabel.Refresh()

			err := borzoi.Resize(
				inputFilePath,
				outputFilePath,
				widthEntry.Text,
				heightEntry.Text,
			)
			if err != nil {
				mainLabel.Text = fmt.Sprintf("Status: %v", err)
				mainLabel.Refresh()
			} else {
				mainLabel.Text = "Status: file processed successfully"
				mainLabel.Refresh()
			}
		}
	})

	mainLabelContainer := container.NewHBox(
		layout.NewSpacer(),
		mainLabel,
		layout.NewSpacer(),
	)

	resizeWindow.SetContent(container.NewVBox(
		mainLabelContainer,
		selectFileLabel,
		selectFileButton,
		widthLabel,
		widthEntry,
		heightLabel,
		heightEntry,
		layout.NewSpacer(),
		runButton,
	))

	resizeWindow.Show()
}

func (w *Windows) TrimWindow() {
	trimWindow := w.App.NewWindow("Trim video")
	trimWindow.Resize(fyne.NewSize(500, 500))

	mainLabel := widget.NewLabel("Status: not running")

	selectFileLabel := widget.NewLabel("Select video:")

	inputFilePath := ""

	selectFileButton := widget.NewButton("Select file", func() {
		dialog.ShowFileOpen(func(response fyne.URIReadCloser, err error) {
			if err == nil && response != nil {
				inputFilePath = response.URI().String()
				fileName := getFileNameWithExt(inputFilePath)
				mainLabel.Text = fmt.Sprintf("Status: file selected(%s)", fileName)
				mainLabel.Refresh()
				response.Close()
			}
		}, trimWindow)
	})

	startTimeLabel := widget.NewLabel("Start time in 15:04:05 format:")
	startTimeEntry := widget.NewEntry()

	endTimeLabel := widget.NewLabel("End time in 15:04:05 format:")
	endTimeEntry := widget.NewEntry()

	runButton := widget.NewButton("Run", func() {
		switch {
		case inputFilePath == "":
			mainLabel.Text = "Status: file not selected"
			mainLabel.Refresh()
		case startTimeEntry.Text == "":
			mainLabel.Text = "Status: start time not specified"
			mainLabel.Refresh()
		case endTimeEntry.Text == "":
			mainLabel.Text = "Status: end time not specified"
			mainLabel.Refresh()
		default:
			outputFilePath := getOutputFilePath(inputFilePath, "_trimmed", "")

			mainLabel.Text = "Status: processing, please wait..."
			mainLabel.Refresh()

			err := borzoi.Trim(
				inputFilePath,
				outputFilePath,
				startTimeEntry.Text,
				endTimeEntry.Text,
			)
			if err != nil {
				mainLabel.Text = fmt.Sprintf("Status: %v", err)
				mainLabel.Refresh()
			} else {
				mainLabel.Text = "Status: file processed successfully"
				mainLabel.Refresh()
			}
		}
	})

	mainLabelContainer := container.NewHBox(
		layout.NewSpacer(),
		mainLabel,
		layout.NewSpacer(),
	)

	trimWindow.SetContent(container.NewVBox(
		mainLabelContainer,
		selectFileLabel,
		selectFileButton,
		startTimeLabel,
		startTimeEntry,
		endTimeLabel,
		endTimeEntry,
		layout.NewSpacer(),
		runButton,
	))

	trimWindow.Show()
}

func (w *Windows) VideoToGifWindow() {
	videoToGifWindow := w.App.NewWindow("Video to gif")
	videoToGifWindow.Resize(fyne.NewSize(500, 500))

	mainLabel := widget.NewLabel("Status: not running")

	selectFileLabel := widget.NewLabel("Select video:")

	inputFilePath := ""

	selectFileButton := widget.NewButton("Select file", func() {
		dialog.ShowFileOpen(func(response fyne.URIReadCloser, err error) {
			if err == nil && response != nil {
				inputFilePath = response.URI().String()
				fileName := getFileNameWithExt(inputFilePath)
				mainLabel.Text = fmt.Sprintf("Status: file selected(%s)", fileName)
				mainLabel.Refresh()
				response.Close()
			}
		}, videoToGifWindow)
	})

	gifSizeLabel := widget.NewLabel("Select gif size:")
	gifSize := ""
	gifSizeRadioGroup := widget.NewRadioGroup([]string{"320x240", "640x480", "1280x720"}, func(selected string) {
		gifSize = selected
	})

	gifFpsLabel := widget.NewLabel("Select gif fps:")
	gifFps := ""
	gifFpsRadioGroup := widget.NewRadioGroup([]string{"10", "15", "20", "25", "30"}, func(selected string) {
		gifFps = selected
	})

	startTimeLabel := widget.NewLabel("Start time in 15:04:05 format:")
	startTimeEntry := widget.NewEntry()

	endTimeLabel := widget.NewLabel("End time in 15:04:05 format:")
	endTimeEntry := widget.NewEntry()

	runButton := widget.NewButton("Run", func() {
		switch {
		case inputFilePath == "":
			mainLabel.Text = "Status: file not selected"
			mainLabel.Refresh()
		case gifSize == "":
			mainLabel.Text = "Status: gif size not selected"
			mainLabel.Refresh()
		case gifFps == "":
			mainLabel.Text = "Status: gif fps not selected"
			mainLabel.Refresh()
		case startTimeEntry.Text == "":
			mainLabel.Text = "Status: start time not specified"
			mainLabel.Refresh()
		case endTimeEntry.Text == "":
			mainLabel.Text = "Status: end time not specified"
			mainLabel.Refresh()
		default:
			outputFilePath := getOutputFilePath(inputFilePath, "", "gif")

			mainLabel.Text = "Status: processing, please wait..."
			mainLabel.Refresh()

			err := borzoi.VideoToGif(
				inputFilePath,
				outputFilePath,
				gifSize,
				startTimeEntry.Text,
				endTimeEntry.Text,
				gifFps,
			)
			if err != nil {
				mainLabel.Text = fmt.Sprintf("Status: %v", err)
				mainLabel.Refresh()
			} else {
				mainLabel.Text = "Status: file processed successfully"
				mainLabel.Refresh()
			}
		}
	})

	mainLabelContainer := container.NewHBox(
		layout.NewSpacer(),
		mainLabel,
		layout.NewSpacer(),
	)

	videoToGifWindow.SetContent(container.NewVBox(
		mainLabelContainer,
		selectFileLabel,
		selectFileButton,
		gifSizeLabel,
		gifSizeRadioGroup,
		gifFpsLabel,
		gifFpsRadioGroup,
		startTimeLabel,
		startTimeEntry,
		endTimeLabel,
		endTimeEntry,
		layout.NewSpacer(),
		runButton,
	))

	videoToGifWindow.Show()
}

func (w *Windows) ExtractAudioWindow() {
	extractAudioWindow := w.App.NewWindow("Extract audio from video")
	extractAudioWindow.Resize(fyne.NewSize(500, 500))

	mainLabel := widget.NewLabel("Status: not running")

	selectFileLabel := widget.NewLabel("Select video:")

	inputFilePath := ""

	selectFileButton := widget.NewButton("Select file", func() {
		dialog.ShowFileOpen(func(response fyne.URIReadCloser, err error) {
			if err == nil && response != nil {
				inputFilePath = response.URI().String()
				fileName := getFileNameWithExt(inputFilePath)
				mainLabel.Text = fmt.Sprintf("Status: file selected(%s)", fileName)
				mainLabel.Refresh()
				response.Close()
			}
		}, extractAudioWindow)
	})

	runButton := widget.NewButton("Run", func() {
		switch {
		case inputFilePath == "":
			mainLabel.Text = "Status: file not selected"
			mainLabel.Refresh()
		default:
			outputFilePath := getOutputFilePath(inputFilePath, "", "mp3")

			mainLabel.Text = "Status: processing, please wait..."
			mainLabel.Refresh()

			err := borzoi.ExtractAudio(
				inputFilePath,
				outputFilePath,
			)
			if err != nil {
				mainLabel.Text = fmt.Sprintf("Status: %v", err)
				mainLabel.Refresh()
			} else {
				mainLabel.Text = "Status: file processed successfully"
				mainLabel.Refresh()
			}
		}
	})

	mainLabelContainer := container.NewHBox(
		layout.NewSpacer(),
		mainLabel,
		layout.NewSpacer(),
	)

	extractAudioWindow.SetContent(container.NewVBox(
		mainLabelContainer,
		selectFileLabel,
		selectFileButton,
		layout.NewSpacer(),
		runButton,
	))

	extractAudioWindow.Show()
}

func (w *Windows) ConvertVideoWindow() {
	convertVideoWindow := w.App.NewWindow("Convert video")
	convertVideoWindow.Resize(fyne.NewSize(500, 500))

	mainLabel := widget.NewLabel("Status: not running")

	selectFileLabel := widget.NewLabel("Select video:")

	inputFilePath := ""

	selectFileButton := widget.NewButton("Select file", func() {
		dialog.ShowFileOpen(func(response fyne.URIReadCloser, err error) {
			if err == nil && response != nil {
				inputFilePath = response.URI().String()
				fileName := getFileNameWithExt(inputFilePath)
				mainLabel.Text = fmt.Sprintf("Status: file selected(%s)", fileName)
				mainLabel.Refresh()
				response.Close()
			}
		}, convertVideoWindow)
	})

	formatLabel := widget.NewLabel("Select video format:")
	format := ""
	formatRadioGroup := widget.NewRadioGroup([]string{"mp4", "avi", "wmv", "mkv", "webm", "flv", "3gp"}, func(selected string) {
		format = selected
	})

	runButton := widget.NewButton("Run", func() {
		switch {
		case inputFilePath == "":
			mainLabel.Text = "Status: file not selected"
			mainLabel.Refresh()
		case format == "":
			mainLabel.Text = "Status: video format not selected"
			mainLabel.Refresh()
		default:
			outputFilePath := getOutputFilePath(inputFilePath, "", format)

			mainLabel.Text = "Status: processing, please wait..."
			mainLabel.Refresh()

			err := borzoi.ConvertVideo(
				inputFilePath,
				outputFilePath,
			)
			if err != nil {
				mainLabel.Text = fmt.Sprintf("Status: %v", err)
				mainLabel.Refresh()
			} else {
				mainLabel.Text = "Status: file processed successfully"
				mainLabel.Refresh()
			}
		}
	})

	mainLabelContainer := container.NewHBox(
		layout.NewSpacer(),
		mainLabel,
		layout.NewSpacer(),
	)

	convertVideoWindow.SetContent(container.NewVBox(
		mainLabelContainer,
		selectFileLabel,
		selectFileButton,
		formatLabel,
		formatRadioGroup,
		layout.NewSpacer(),
		runButton,
	))

	convertVideoWindow.Show()
}
