package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	banknotes := map[int]int{5: 200, 10: 100, 100: 50, 200: 25, 500: 25, 1000: 10}

	for {
		fmt.Println("Choose option:")
		fmt.Println("1 - cash withdrawal")
		fmt.Println("2 - exit")

		readOption := bufio.NewReader(os.Stdin)
		option, _ := readOption.ReadString('\n')
		switch option {
		case "1":
			reader := bufio.NewReader(os.Stdin)
			requestedAmount, _ := reader.ReadString('\n')
			fmt.Println("Requested Amount is", requestedAmount)
			fmt.Println("Available baknotes: ", banknotes)
		case "2":
			return
		default:
			fmt.Println("No such option exists")
		}

	}
}
