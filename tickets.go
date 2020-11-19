package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())                                   //use random number as a function of time
	distance := 62100000                                               //kilometers
	maxSpeed := 30                                                     //max speed km/s
	minSpeed := 16                                                     //min speed km/s
	secondsInDays := 24 * 60 * 60                                      //
	var minDays float64 = float64(distance / maxSpeed / secondsInDays) //min amount of days for one way flight
	var maxDays float64 = float64(distance / minSpeed / secondsInDays) //max amount of days for one way flight
	minCost := 36                                                      //min costs for ticket
	var addCostToMax float64 = 50 - 36                                 //surcharge to get the maximum ticket price

	//top of the table
	fmt.Printf("%-17v %4v %10v %5v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("=======================================")

	for i := 0; i < 10; i++ {

		//Spaceline field
		spaceline := "Space Adventures"
		switch rand.Intn(3) {
		case 0:
			spaceline = "Virgin Galactic"
		case 1:
			spaceline = "SpaceX"
		}

		//Days field
		speed := rand.Intn(maxSpeed-minSpeed+1) + minSpeed //for Days we need to get the speed between min max speed
		var days float64 = float64(distance / speed / secondsInDays)

		//Trip type
		roundtrip := 2 //by default we have two-way trip
		roundtripName := "Round-trip"
		if rand.Intn(2) == 1 { //change for one-way trip
			roundtrip = 1
			roundtripName = "One-way"
		}

		//Price field, calculated by complex formula, as faster trips costs extra payment,
		//and slowest will be for minimum amount
		price := roundtrip * (minCost + int(addCostToMax*(1-(days-minDays)/(maxDays-minDays))))

		//print one line
		fmt.Printf("%-17v %4v %10v %5v\n", spaceline, days, roundtripName, price)
	}
}
