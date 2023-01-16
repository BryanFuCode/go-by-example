package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	// fmt.Println(secretNumber)
	fmt.Println("Please input your guess")

	var input string
	for {
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			continue
		}
		inputNumber, err := strconv.Atoi(strings.Trim(input, "\r\n"))
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		if inputNumber > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
		} else if inputNumber < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}

}
