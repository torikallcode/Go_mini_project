package main

import (
	"fmt"
)

func average(a []int) int {
	var sum int
	for _, num := range a {
		sum += num
	}
	return sum / len(a)
}

func main() {
	var number int
	var value int

	fmt.Print("berapa banyak nilai yang ingin anda masukkan: ")
	fmt.Scanln(&number)

	if number <= 0 {
		fmt.Println("Jumlah nilai harus lebih dari nol.")
		return
	}

	values := make([]int, number)

	for i := 0; i < number; i++ {
		fmt.Printf("Masukkan angka ke-%d: ", i+1)
		fmt.Scanln(&value)
		values[i] = value
	}

	result := average(values)
	fmt.Println("Average:", result)
}
