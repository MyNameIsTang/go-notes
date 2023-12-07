package greetings

import (
	"fmt"
	"time"
)

func GoodDay() {
	fmt.Println("Good Day")
}

func GoodNight() {
	fmt.Println("Good Night")
}

func IsAM() bool {
	currentTime := time.Now()
	hour := currentTime.Hour()

	return hour < 12
}

func IsAfternoon() bool {
	currentTime := time.Now()
	hour := currentTime.Hour()

	return hour > 12 && hour < 18
}

func IsEvening() bool {
	currentTime := time.Now()
	hour := currentTime.Hour()

	return hour > 18 && hour < 24
}
