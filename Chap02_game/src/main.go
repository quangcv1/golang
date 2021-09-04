//1. Generate a random number from 1 to 100, and store it as a target number for players to guess
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func randomNumber(number int) (int, string) {
	//fmt.Println("This is function randomNumber")
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(number), "Done"
}

func enterNumber() (int, string) {
	//fmt.Println("This is function enterNumber")
	fmt.Print("Enter your guessed number: ")
	reader := bufio.NewReader(os.Stdin)
	input,err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	inputNumber, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return inputNumber, "Done"
}

func checkEnteredNumber() (int,string) {
	negativeNumber := true
	inputNumber := -1
	for negativeNumber {
		inputNumber, _ = enterNumber()
		if inputNumber < 0 {
			fmt.Println("Pls re-enter your number: ")
		} else {
			negativeNumber = false
		}
	}
	return inputNumber, "Done"
}

func main()  {

		target, _ := randomNumber(100)
	for i := 10; i > 0; i-- {
		fmt.Println("Your random number is: ", target)
		inputNumber,_ := checkEnteredNumber()
		fmt.Println("Your entered number is: ", inputNumber)
		if inputNumber == target {
			fmt.Println("Matched")
			break
		} else if inputNumber > target {
			fmt.Println("Higher")
			fmt.Println("left: ", i)
		} else {
			fmt.Println("Lower")
			fmt.Println("left", i)
		}
}
}