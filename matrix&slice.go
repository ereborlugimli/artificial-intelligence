/*
https://play.golang.org/p/cPtsPTWiK3M
*/
package main

import (
	"fmt"
)

func main() {
	x := [][]int{{42, 43, 44}, {45, 46, 47}, {48, 49, 50}}
	y := []int{42, 43, 44, 45, 46, 47,48, 49, 50, 51}
	fmt.Println(x)
	fmt.Println(y)
	
	x = append(x[:0][:0], x[:0][:0]...) 
	y = append(y[10:],y[10:]...)
	
	fmt.Println(x)
	fmt.Println(y)
}
