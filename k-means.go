package main

import (
	"fmt"
	"math"
)

var centroid1_x, centroid2_x, step int

func main() {

	var xc1, yc1, xc2, yc2 float64
	row := [][]float64{}
	row2 := []float64{}
	row3 := []float64{}
	euclidean1 := []float64{}
	euclidean2 := []float64{}

	fmt.Println("--- Select two centroid ---\n")
	object := [][]float64{
		{1, 1}, //centroid1
		{2, 1},
		{4, 3},
		{5, 4}, //centroid2
	}
	fmt.Println("A\t", "B\t", "C\t", "D\t")
	fmt.Println(object, "\n")

	fmt.Print("Enter the Centroid index : ")
	fmt.Scanf("%d", &centroid1_x)
	fmt.Print("Enter the Centroid index : ")
	fmt.Scanf("%d", &centroid2_x)

	row2 = append(row2, object[centroid1_x][0], object[centroid1_x][1])
	row3 = append(row3, object[centroid2_x][0], object[centroid2_x][1])
	row = append(row, row2, row3)

	fmt.Println("\n--- Your centroids ---\n")
	fmt.Println(row, "\n")
	step = 0
	for step < 5 {

		if step == 0 {
			var a, b float64
			for i := 0; i < len(object); i++ {

				a, b = calculator(object, row, i)
				euclidean1 = append(euclidean1, a)
				euclidean2 = append(euclidean2, b)

			}

		} else {
			fmt.Println("Euclidean1 = ", euclidean1)
			fmt.Println("Euclidean2 = ", euclidean2)
			xc1, yc1, xc2, yc2 = newCentroid(object, euclidean1, euclidean2)
			euclidean1 = append(euclidean1[len(object):], euclidean1[len(object):]...) //delete
			euclidean2 = append(euclidean2[len(object):], euclidean2[len(object):]...) //delete
			row2 = append(row2[2:], row2[2:]...)                                       //delete
			row3 = append(row3[2:], row3[2:]...)                                       //delete
			row = append(row[:0][:0], row[:0][:0]...)                                  //delete
			row2 = append(row2, xc1, yc1)
			row3 = append(row3, xc2, yc2)
			row = append(row, row2, row3)
			var a, b float64
			for i := 0; i < len(object); i++ {

				a, b = calculator(object, row, i)
				euclidean1 = append(euclidean1, a)
				euclidean2 = append(euclidean2, b)

			}
		}
		step++
	}
}

func calculator(object [][]float64, row [][]float64, x int) (float64, float64) {

	c1 := math.Sqrt(math.Pow((object[x][0]-row[0][0]), 2) + math.Pow((object[x][1]-row[0][1]), 2))
	c2 := math.Sqrt(math.Pow((object[x][0]-row[1][0]), 2) + math.Pow((object[x][1]-row[1][1]), 2))

	return c1, c2
}

func newCentroid(object [][]float64, euclidean1 []float64, euclidean2 []float64) (float64, float64, float64, float64) {
	var xc1, yc1, xc2, yc2, count1, count2 float64

	for i := 0; i < len(euclidean1); i++ {
		if euclidean1[i] < euclidean2[i] {
			xc1 += object[i][0]
			yc1 += object[i][1]
			count1++
		} else {
			xc2 += object[i][0]
			yc2 += object[i][1]
			count2++
		}
	}
	xc1 /= count1
	yc1 /= count1
	xc2 /= count2
	yc2 /= count2
	return xc1, yc1, xc2, yc2
}
