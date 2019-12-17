package worldweather

import (
	"fmt"
	"math/rand"
	"time"
)

//GetCurrentDegree Return current degree
func GetCurrentDegree() int {
	randNum := (time.Now().Unix() + 563) % rand.Int63n(11)

	fmt.Printf("WorldWeather, sleep for %d seconds\n", randNum)

	time.Sleep(time.Duration(randNum) * time.Second)

	return 5
}
