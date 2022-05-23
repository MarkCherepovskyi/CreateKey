package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	values = map[string]string{
		"0000": "0",
		"0001": "1",
		"0010": "2",
		"0011": "3",
		"0100": "4",
		"0101": "5",
		"0110": "6",
		"0111": "7",
		"1000": "8",
		"1001": "9",
		"1010": "A",
		"1011": "B",
		"1100": "C",
		"1101": "D",
		"1110": "E",
		"1111": "F",
	}

	valuesInt = map[int]string{
		0:  "0",
		1:  "1",
		2:  "2",
		3:  "3",
		4:  "4",
		5:  "5",
		6:  "6",
		7:  "7",
		8:  "8",
		9:  "9",
		10: "A",
		11: "B",
		12: "C",
		13: "D",
		14: "E",
		15: "F",
	}
	sizeRange [10]int = [10]int{8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}
)

func main() {
	str := Create(sizeRange[9])
	Convert(str)
	res := CreateRandomKey(sizeRange[9])
	fmt.Println(res)
	key := Convert(res)
	fmt.Println(key)
	fmt.Println()
	resSer := Search(sizeRange[9], key)
	fmt.Println(resSer)
}

func CreateRandomKey(size int) string {
	var array []int = make([]int, size)
	str := ""
	for index, _ := range array {
		array[index] = rand.Intn(2)
		str += fmt.Sprint(array[index])
	}

	return str
}

func Create(size int) string {

	var array []int = make([]int, size)
	var str string

	for index, _ := range array {
		array[index] = 1
		str += fmt.Sprint(array[index])
	}

	return str
}

func Convert(val string) string {
	var buffer, finalStr string

	for i, dataStr := range val {
		buffer += fmt.Sprintf("%c", dataStr)
		if (i+1)%4 == 0 {
			//fmt.Println(buffer)
			finalStr = fmt.Sprint(finalStr, values[buffer])
			buffer = ""
		}

	}
	return finalStr
}

func f(index int, strByte string) string {

	for i := 0; i < 16; i++ {
		var strBuffer string = fmt.Sprint(valuesInt[i])

		if strBuffer == strByte {

			return valuesInt[i]
		}
	}
	return "*"
}

func Search(size int, str string) string {
	var buffer []string = make([]string, size)

	var finalStr string
	start := time.Now()
	for index := 0; index < size/4; index++ {
		c := fmt.Sprintf("%c", str[index])
		buffer[index] = f(index, c)

	}

	for _, dataStr := range buffer {
		finalStr += fmt.Sprint(dataStr)
	}
	fmt.Println("time :", time.Since(start))

	return finalStr
}
