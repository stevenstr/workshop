package openweather

import (
	"fmt"
	"math/rand"
	"time"
)

//GetCurrentDegree Return current degree
func GetCurrentDegree() int {
	randNum := (time.Now().Unix() + 981) % rand.Int63n(10)

	fmt.Printf("OpenWeather, sleep for %d seconds\n", randNum)

	time.Sleep(time.Duration(randNum) * time.Second)

	return 5
}
