package borzoiffmpeg

import (
	"fmt"
	"testing"
)

func TestResize(t *testing.T) {
	err := Resize("test.mp4", "test_resized.mp4", "1280", "720")
	if err != nil {
		fmt.Println(err.Error())
		t.Error("resize failed")
	}
}

func TestTrim(t *testing.T) {
	err := Trim("test.mp4", "test_trimmed.mp4", "00:00:05", "00:00:10")
	if err != nil {
		fmt.Println(err.Error())
		t.Error("trim failed")
	}
}

func TestVideoToGif(t *testing.T) {
	err := VideoToGif("test.mp4", "test.gif", "480x360", "00:00:05", "00:00:10", "10")
	if err != nil {
		fmt.Println(err.Error())
		t.Error("video to gif failed")
	}
}

func TestExtractAudio(t *testing.T) {
	err := ExtractAudio("test.mp4", "test.mp3")
	if err != nil {
		fmt.Println(err.Error())
		t.Error("extract audio failed")
	}
}

func TestConvertVideo(t *testing.T) {
	err := ConvertVideo("test.mp4", "test.3gp")
	if err != nil {
		fmt.Println(err.Error())
		t.Error("convert video failed")
	}
}
