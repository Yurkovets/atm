package main

import (
	"fmt"
)

func main() {

	banknotes := map[int]int{5: 200, 10: 100, 100: 50, 200: 25, 500: 25, 1000: 10}

	for {
		fmt.Println("Choose option:")
		fmt.Println("1 - cash withdrawal")
		fmt.Println("2 - exit")

		var option int
		fmt.Scan(&option)
		switch option {
		case 1:
			fmt.Println("Available baknotes: ", banknotes)
		case 2:
			return
		default:
			fmt.Println("No such option exists")
			fmt.Println(option)
		}
	}
}
