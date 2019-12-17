package belarustoday

import (
	"fmt"
	"math/rand"
	"time"
)

//GetCurrentDegree Return current degree
func GetCurrentDegree() int {
	randNum := (time.Now().Unix() + 329) % rand.Int63n(9)

	fmt.Printf("BelarusToday, sleep for %d seconds\n", randNum)

	time.Sleep(time.Duration(randNum) * time.Second)

	return 5
}
