package main

//first "github.com/burov/workshop/weather/api/belarustoday"
//second "github.com/burov/workshop/weather/api/openweather"
//third "github.com/burov/workshop/weather/api/worldweather"
import (
	"fmt"

	first "github.com/burov/workshop/weather/api/belarustoday"
	second "github.com/burov/workshop/weather/api/openweather"
	third "github.com/burov/workshop/weather/api/worldweather"
)

func main() {
	firstDegree := first.GetCurrentDegree()
	secondDegree := second.GetCurrentDegree()
	thirdDegree := third.GetCurrentDegree()

	fmt.Println(firstDegree)
	fmt.Println(secondDegree)
	fmt.Println(thirdDegree)

	//initialize workers using

	//run workers

	//wait until first result comming

	//print result and exit
}
