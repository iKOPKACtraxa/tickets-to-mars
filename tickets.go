package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())                                   //задает в параметрах рандомизатора использование текущего времени
	distance := 62100000                                               //расстояние в Км
	maxSpeed := 30                                                     //максимальная скорость Км/с
	minSpeed := 16                                                     //минимальная скорость Км/с
	secondsInDays := 24 * 60 * 60                                      //секунд в дне
	var minDays float64 = float64(distance / maxSpeed / secondsInDays) //минимальное количество дней для перелета
	var maxDays float64 = float64(distance / minSpeed / secondsInDays) //максимальное...
	minCost := 36                                                      //минимальная цена билета
	var addCostToMax float64 = 50 - 36                                 //доплата до максимальной цены билета

	//шапка таблицы
	fmt.Printf("%-17v %4v %10v %5v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("=======================================")

	//запуск цикла для вывода строк таблицы
	for i := 0; i < 10; i++ {

		//Spaceline поле
		spaceline := "Space Adventures"
		switch rand.Intn(3) {
		case 0:
			spaceline = "Virgin Galactic"
		case 1:
			spaceline = "SpaceX"
		}

		//Days поле
		speed := rand.Intn(maxSpeed-minSpeed+1) + minSpeed //для дней нужно сначала рассчитать скорость в диапазоне MinMax
		var days float64 = float64(distance / speed / secondsInDays)

		//Trip type
		roundtrip := 2 //по умолчанию у нас полет туда-обратно
		roundtripName := "Round-trip"
		if rand.Intn(2) == 1 { //переопределение на one-way
			roundtrip = 1
			roundtripName = "One-way"
		}

		//Price поле, рассчитывается сложной формулой так как самые быстрые полеты стоят дополнительную плату,
		//а самые медленные полеты будут без дополнительной платы
		price := roundtrip * (minCost + int(addCostToMax*(1-(days-minDays)/(maxDays-minDays))))

		//line, вывод строки
		fmt.Printf("%-17v %4v %10v %5v\n", spaceline, days, roundtripName, price)
	}
}
