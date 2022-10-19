package helpers

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateRandomData() []int {
	var max_num = 100
	var max = &max_num

	water := rand.Intn(*max + 1)
	wind := rand.Intn(*max + 1)

	payload := []int{water, wind}

	return payload
}

func GetStatus() ([]int, string, string) {

	var (
		waterStatus string
		windStatus  string
	)

	data := GenerateRandomData()

	switch {
	case data[0] <= 5:
		waterStatus = "Aman"
	case data[0] >= 6 && data[0] <= 8:
		waterStatus = "Siaga"
	case data[0] > 8:
		waterStatus = "Bahaya"
	default:
		waterStatus = "Water Value not defined"
	}

	switch {
	case data[1] <= 6:
		windStatus = "Aman"
	case data[1] >= 7 && data[1] <= 15:
		windStatus = "Siaga"
	case data[1] > 15:
		windStatus = "Bahaya"
	default:
		windStatus = "Wind Value not defined"
	}

	return data, waterStatus, windStatus

}
