package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	//s"reflect"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var name string
	var age, income, k, total, r1, r2, r3, near, counterA, counterB int
	var a []int
	m := make(map[int]string)
	m2 := make(map[int]string)

	fmt.Println("Enter the customer's name : ")
	fmt.Scanf("%s", &name)
	fmt.Println("Enter the customer's age : ")
	fmt.Scanf("%d", &age)
	fmt.Println("Enter the customer's income : ")
	fmt.Scanf("%d", &income)
	fmt.Println("Enter the customer's k.k number : ")
	fmt.Scanf("%d", &k)
	fmt.Println("Enter the near : ")
	fmt.Scanf("%d", &near)

	f, _ := os.Open("knnData.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text() //satır satır yazdırır
		parts := strings.Split(line, ",")

		total = 0
		for i := range parts {
			valueStr := parts[i]
			if valueInt, err := strconv.Atoi(valueStr); err == nil {

				if i == 1 {
					r1 = (age - valueInt)
					total += int(math.Pow(float64(r1), 2))
				} else if i == 2 {
					r2 = (income - valueInt)
					total += int(math.Pow(float64(r2), 2))
				} else if i == 3 {
					r3 = (k - valueInt)
					total += int(math.Pow(float64(r3), 2))
				}
			}
		}
		total = int(math.Sqrt(float64(total)))
		m[total] = parts[4]
	}
	for k := range m {
		a = append(a, k)
	}
	//sort.Sort(sort.Reverse(sort.IntSlice(a))) büyükten küçüğe sıralar
	sort.Sort((sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range m[k] {
			m2[k] = string(s)
		}
	}
	counterA = 0
	counterB = 0
	for t := 0; t < near; t++ {
		for _, k := range m2 {
			if m2[k] == "A" {
				counterA++
			} else if m2[k] == "B" {
				counterB++
			}

		}
	}
	fmt.Println(m2)
	fmt.Println(counterA, counterB)
}
