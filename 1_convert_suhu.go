package main

import (
	"fmt"
)

func cls(c int) int {
	return (c * 9 / 5) + 32
}

func fht(f int) int {
	return (f - 32) * 5 / 9
}

func main() {

	var selected int
	var celcius, fahrenheit int

	for {
		fmt.Println("======================")
		fmt.Println("Aplikasi Convert Suhu")
		fmt.Println("======================")
		fmt.Println("1. Celcius to Fahrenheit")
		fmt.Println("2. Fahrenheit to Celcius")
		fmt.Println("3. Exit")
		fmt.Print("Select a conversion type (1 or 2): ")
		fmt.Scanln(&selected)
		switch selected {
		case 1:
			{
				fmt.Print("Input celcius: ")
				fmt.Scanln(&celcius)
				result := cls(celcius)
				fmt.Printf("%v celcius = %v fahrenheit\n", celcius, result)
			}
		case 2:
			{
				fmt.Print("Input fahrenheit: ")
				fmt.Scanln(&fahrenheit)
				result := fht(fahrenheit)
				fmt.Printf("%v fahrenheit = %v celcius\n", fahrenheit, result)
			}
		case 3:
			{
				fmt.Print("Exit")
			}
		default:
			fmt.Println("Enter a valid number!")
		}
		if selected == 3 {
			break
		}
	}
}
