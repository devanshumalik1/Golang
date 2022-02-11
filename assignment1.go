package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var dishName string
	var quanityOfDish int
	var salesPerOrder = []int{}
	count := 0
	fmt.Println("=======WELCOME TO XYZ CAFE========== ")
	fmt.Println("MENU OF XYZ WITH ALPHABETICAL CODE")
	fmt.Print("C: Coffee  (40rs)\n D: Dosa (80 rs)\n T: Tomato Soup (20rs)\n I : Idli  (20rs)\n V : Vada (25rs) \n B: Bhature and Chole (30rs) \n P: Paneer Pakoda (30rs)\n M: Manchurian (90rs)\n H: Hakka Noodle (70rs)\n F: French Fries (30rs)\n J: Jalebi (30rs)\n L: Lemonade (15rs)\n S: Spring Roll (20rs)\n")
	fmt.Println("-----------------------")
	for true {

		fmt.Println("Press END For closing the day.")
		fmt.Println("Please enter the Dish(Alphabetical code): ")
		fmt.Scan(&dishName)
		dishName = strings.ToUpper(dishName)

		if dishName == "END" {
			total_income(salesPerOrder)
		} else {
			fmt.Println("Please enter the Quanity Of Dish: ")
			fmt.Scan(&quanityOfDish)
		}
		bill := bill(quanityOfDish, dishName)
		fmt.Println("-----------------------")
		fmt.Println("Total bill:", bill, "rs")
		fmt.Println("Thanks for visiting XYZ.\nHave a Good Day.")
		fmt.Println("-----------------------")
		salesPerOrder = append(salesPerOrder, bill)
		count++

	}
}

func total_income(sales []int) {
	var total int = 0

	for i := 0; i < len(sales); i++ {
		total = sales[i] + total
	}
	fmt.Println("Total Earning of the day is : ", total)
	os.Exit(0)
}

func bill(quanityOfDish int, dishName string) int {
	m := map[string]int{
		"C": 40,
		"D": 80,
		"T": 20,
		"I": 20,
		"V": 25,
		"B": 30,
		"P": 30,
		"M": 90,
		"H": 70,
		"F": 30,
		"J": 30,
		"L": 15,
		"S": 20,
	}
	orderAmount := quanityOfDish * m[dishName]
	return orderAmount
}
