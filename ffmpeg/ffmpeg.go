package borzoiffmpeg

import (
	"errors"
	"fmt"
	"strings"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// resize video or image
func Resize(inputFilePath, outputFilePath, width, height string) error {
	scaleStr := fmt.Sprintf("scale=w=%s:h=%s", width, height)
	err := ffmpeg.Input(inputFilePath).
		Output(outputFilePath, ffmpeg.KwArgs{"vf": scaleStr}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		return err
	}

	return nil
}

// trim a videoclip
func Trim(inputFilePath, outputFilePath, start, end string) error {
	startSeconds, err := timeStringToSeconds(start)
	if err != nil {
		return errors.New("error parsing start time to seconds")
	}

	endSeconds, err := timeStringToSeconds(end)
	if err != nil {
		return errors.New("error parsing end time to seconds")
	}

	timeSeconds := endSeconds - startSeconds

	err = ffmpeg.Input(inputFilePath, ffmpeg.KwArgs{"ss": startSeconds}).
		Output(outputFilePath, ffmpeg.KwArgs{"t": timeSeconds}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		return err
	}

	return nil
}

// convert video to gif
func VideoToGif(inputFilePath, outputFilePath string, size, start, end, fps string) error {
	startSeconds, err := timeStringToSeconds(start)
	if err != nil {
		return errors.New("error parsing start time to seconds")
	}

	endSeconds, err := timeStringToSeconds(end)
	if err != nil {
		return errors.New("error parsing end time to seconds")
	}

	timeSeconds := endSeconds - startSeconds

	err = ffmpeg.Input(inputFilePath, ffmpeg.KwArgs{"ss": startSeconds}).
		Output(outputFilePath, ffmpeg.KwArgs{"s": size, "t": timeSeconds, "r": fps}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		return err
	}

	return nil
}

// extract audio from video
func ExtractAudio(inputFilePath, outputFilePath string) error {
	err := ffmpeg.Input(inputFilePath).
		Output(outputFilePath, ffmpeg.KwArgs{"q:a": 0, "map": "a"}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		return err
	}

	return nil
}

// convert video from one format to another
func ConvertVideo(inputFilePath, outputFilePath string) error {
	videoCodec := "libx264"
	audioCodec := "aac"

	if strings.HasSuffix(outputFilePath, ".webm") {
		videoCodec = "libvpx"
		audioCodec = "libvorbis"
	}

	err := ffmpeg.Input(inputFilePath).
		Output(outputFilePath, ffmpeg.KwArgs{"c:v": videoCodec, "c:a": audioCodec}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		return err
	}
	return nil
}
