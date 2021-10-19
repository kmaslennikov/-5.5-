package main

import "fmt"

func main() {
	fmt.Println("Введите три процентные ставки:")
	var firstRate, secondRate, thirdRate int
	fmt.Scan(&firstRate)
	fmt.Scan(&secondRate)
	fmt.Scan(&thirdRate)

	if firstRate == secondRate && firstRate == thirdRate {
		fmt.Println("Ставки равны")
	} else {
		fmt.Println("Ставки:")
		if firstRate <= secondRate {
			if firstRate <= thirdRate {
				fmt.Println(secondRate)
				fmt.Println(thirdRate)
			} else {
				fmt.Println(firstRate)
				fmt.Println(secondRate)
			}
		} else {
			if secondRate <= thirdRate {
				fmt.Println(firstRate)
				fmt.Println(thirdRate)
			} else {
				fmt.Println(firstRate)
				fmt.Println(secondRate)
			}
		}
	}
}
