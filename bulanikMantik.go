package main

import (
	"fmt"
)

func main() {

	var dishes, impurity, x, y, result float32
	var result2 string
	process := map[string]string{
		"00": "P1",
		"01": "P2",
		"10": "P3",
		"11": "P4",
	}
	fmt.Print("Enter the amount of the dishes : ")
	fmt.Scanf("%f", &dishes)
	fmt.Print("Enter the rate of the impurity : ")
	fmt.Scanf("%f", &impurity)
	x = blurring(dishes)
	y = clarification(impurity)
	fmt.Println("Blurring = ", x)
	fmt.Println("Clarification = ", y)
	result, result2 = rules(dishes, impurity, x, y, process)
	calculator(x, y, result, result2)

}

//bulanıklaştırma
func blurring(dishes float32) float32 { //geridonus degiskeni ekle

	var keep float32
	if 0 <= dishes && dishes <= 30 { //az
		keep = (triangle(0.0, 30.0, dishes))
	} else if 30 < dishes && dishes <= 70 { //çok
		keep = (trapezoid(30.0, 70.0, 40.0, 60.0, dishes))
	}
	return keep
}

func triangle(min float32, max float32, x float32) float32 { //üçgen

	var m float32
	m = max + min/2
	var result float32
	if x <= min || x >= max {
		result = 0.0
	} else if x > min && x <= m {
		result = (x - min) / (m - min)
	} else if x > m && x <= max {
		result = (max - x) / (max - m)
	}
	return result
}
func trapezoid(min float32, max float32, midleft float32, midright float32, x float32) float32 {

	var result float32
	if x < min || x > max {
		result = 0.0
	} else if min <= x && x <= midleft {
		result = (x - min) / (midleft - min)
	} else if midleft <= x && x <= midright {
		result = 1.0
	} else if midright <= x && x <= max {
		result = (max - x) / (max - midright)
	}
	return result
}

//durulaştırma
func clarification(impurity float32) float32 {

	var keep float32
	if 0 <= impurity && impurity <= 150 {
		keep = trapezoid(0.0, 150.0, 30.0, 120.0, impurity)
	} else if 150 < impurity && impurity <= 230 {
		keep = trapezoid(150.0, 230.0, 170.0, 210.0, impurity)
	}
	return keep
}

func rules(dishes float32, impurity float32, x float32, y float32, process map[string]string) (float32, string) {

	var result float32
	var result2 string
	if dishes <= 30.0 && impurity <= 150.0 {
		fmt.Println("Az bulasık-Az kirli ", process["00"])
		if x > y {
			result = y
		} else {
			result = x
		}
		result2 = process["00"]
	} else if dishes <= 30.0 && impurity > 150.0 {
		fmt.Println("Az bulasik-Cok kirli", process["01"])
		if x > y {
			result = y
		} else {
			result = x
		}
		result2 = process["01"]
	} else if dishes > 30.0 && impurity <= 150.0 {
		fmt.Println("Cok bulasik-Az kirli", process["10"])
		if x > y {
			result = y
		} else {
			result = x
		}
		result2 = process["10"]
	} else if dishes > 30.0 && impurity > 150.0 {
		fmt.Println("Cok bulasik-Cok kirli", process["11"])
		if x > y {
			result = y
		} else {
			result = x
		}
		result2 = process["11"]
	}
	fmt.Println(result)
	fmt.Println(result2)
	return result, result2
}

//ağırlık ortalamalarına göre hesapla
func calculator(x float32, y float32, result float32, result2 string) {

	var total float32

	if result2 == "P1" { //30
		total = (result*0 + result*5 + result*15 + result*20 + result*25 + result*30) / (result * 6)
	} else if result2 == "P2" { //20-45
		total = (result*20 + result*25 + result*30 + result*35 + result*40 + result*45) / (result * 6)
	} else if result2 == "P3" { //45-60
		total = (result*45 + result*50 + result*55 + result*60) / (result * 4)
	} else if result2 == "P4" { //50-90
		total = (result*50 + result*55 + result*60 + result*65 + result*70 + result*75 + result*80 + result*85 + result*90) / (result * 9)
	}
	fmt.Println("Ağırlıklı ortalama sonucu = ", total)
}
