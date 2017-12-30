package main

import (
	"bufio"
	"fmt"
	//"io/ioutil"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var name, class, text string
	var age, income, k, t, total, r1, r2, r3, near, counterA, counterB int
	var a []int
	m := make(map[int]string)
	m2 := make(map[int]string)
	counterA = 0
	counterB = 0
	t = 0

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
	//fmt.Println(m2)
	for _, s := range m2 {
		if t < 3 {
			if s == "A" {
				counterA = counterA + 1
			} else if s == "B" {
				counterB = counterB + 1
			}
		}
		t++
	}

	if counterA > counterB {
		class = "A"
		fmt.Println("A sınıfına ait")
	} else {
		class = "B"
		fmt.Println("B sınıfına ait")
	}
	text = name + "," + strconv.Itoa(age) + "," + strconv.Itoa(income) + "," + strconv.Itoa(k) + "," + class + "\n"
	x, err := os.OpenFile("knnData.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer x.Close()

	if _, err = x.WriteString(text); err != nil {
		panic(err)
	}

}

