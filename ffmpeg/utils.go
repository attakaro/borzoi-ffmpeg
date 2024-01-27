package borzoiffmpeg

import (
	"time"
)

func timeStringToSeconds(timeString string) (int, error) {
	parsedTime, err := time.Parse("15:04:05", timeString)
	if err != nil {
		return 0, err
	}

	totalSeconds := int(parsedTime.Hour())*3600 + int(parsedTime.Minute())*60 + parsedTime.Second()

	return totalSeconds, nil
}
