/*

Bayes Theorem - Golang
Eksik düzenlenmeli  ! 
*/
package main

import (
	"fmt"
	"os"
)

func main() {

	//outlook temp. humidity windy play city
	dataSet := [][]string{
		{"Sunny", "Hot", "High", "False", "No", "Istanbul"},
		{"Sunny", "Hot", "High", "True", "No", "Ankara"},
		{"Overcast", "Hot", "High", "False", "Yes", "Istanbul"},
		{"Rainy", "Mild", "High", "False", "Yes", "Ankara"},
		{"Rainy", "Cool", "Normal", "False", "Yes", "Ankara"},
		{"Rainy", "Cool", "Normal", "True", "No", "Ankara"},
		{"Overcast", "Cool", "Normal", "True", "Yes", "Istanbul"}}

	var choice int
	total := float32(1)
	var valueCity, valueLook, valueWindy, valueHum, valuePlay, valueTemp string

	for {
		fmt.Println("--- Look at this ---\n\n[1]Outlook\n[2]Tempature\n[3]Humidity\n[4]Windy\n[5]Play\n[6]Total\n[7]Exit\n")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			fmt.Println("--- City ---\n\n[Istanbul]\n[Ankara]\n")
			fmt.Scanf("%s", &valueCity)
			fmt.Println("\n\n--- OutLook ---\n[Sunny]\n[Rainy]\n[Overcast]\n")
			fmt.Scanf("%s", &valueLook)
			total *= (test(dataSet, valueCity, valueLook))
			fmt.Println("Percentage :", test(dataSet, valueCity, valueLook))
		case 2:
			fmt.Println("--- City ---\n\n[Istanbul]\n[Ankara]\n")
			fmt.Scanf("%s", &valueCity)
			fmt.Println("\n\n--- Tempature ---\n[Hot]\n[Mild]\n[Cool]\n")
			fmt.Scanf("%s", &valueTemp)
			total *= test(dataSet, valueCity, valueTemp)
			fmt.Println("Percentage :", test(dataSet, valueCity, valueTemp))
		case 3:
			fmt.Println("--- City ---\n\n[Istanbul]\n[Ankara]\n")
			fmt.Scanf("%s", &valueCity)
			fmt.Println("\n\n--- Humidity ---\n[High]\n[Normal]\n")
			fmt.Scanf("%s", &valueHum)
			total *= test(dataSet, valueCity, valueHum)
			fmt.Println("Percentage :", test(dataSet, valueCity, valueHum))
		case 4:
			fmt.Println("--- City ---\n\n[Istanbul]\n[Ankara]\n")
			fmt.Scanf("%s", &valueCity)
			fmt.Println("\n\n--- Windy ---\n\n[False]\n[True]\n")
			fmt.Scanf("%s", &valueWindy)
			total *= test(dataSet, valueCity, valueWindy)
			fmt.Println("Percentage :", test(dataSet, valueCity, valueWindy))
		case 5:
			fmt.Println("--- City ---\n\n[Istanbul]\n[Ankara]\n")
			fmt.Scanf("%s", &valueCity)
			fmt.Println("\n\n--- Play ---\n\n[Yes]\n[No]\n")
			fmt.Scanf("%s", &valuePlay)
			total *= test(dataSet, valueCity, valuePlay)
			fmt.Println("Percentage :", test(dataSet, valueCity, valuePlay))
		case 6:
			fmt.Println("Total : ", total)
		case 7:
			os.Exit(1)

		default:
			fmt.Println("Wrong Choice ! Try again")

		}
	}

}
func test(dataSet [][]string, valueCity string, value string) float32 {

	var city, weather float32
	control := 0

	for _, x := range dataSet {

		for y := range x {

			if x[y] == valueCity {
				city += 1
				control = 1
			} else {
				control = 0
			}
		}
		for y := range x {

			if x[y] == value && control == 1 {
				weather += 1
				control = 0
			}
		}
	}
	if weather/city == 0 { //no multiplication by zero
		return (weather / city) + 1

	} else {

		return weather / city

	}
}
