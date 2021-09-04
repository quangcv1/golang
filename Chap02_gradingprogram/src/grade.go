//Project with requirements as below:
//	1. Enter a grade number to check
//	2. If number < 0 then prompt to ask user re-enter value until positive number
//	3. If grade > 60 then passed else failed

//Note: we will study:
//	1. String convert, read string
//	2. Block scope
// 	3. Error handling with blank identifier to ignore error and logFatal
//	4. Function with return value

package  main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func enterGrade() (float64, string) {
	//Enter your grade
	fmt.Print("Enter your grade:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error while reading string with erro: ", err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal("Error while parsing to float64 with erro: ", err)
	}
	return grade, "you entered your grade"
}
func checkGrade(grade float64)  {
	if grade > 60 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}
}

func main() {
	negativeNumber := true
	var grade float64
	for negativeNumber {
		//ignore the second returned value
		//grade, _ = enterGrade()

		//take the second value
		grade, text := enterGrade()
		fmt.Println(text)
		if grade >= 0 {
			negativeNumber = false
		} else {
			negativeNumber = true
			fmt.Println("You entered negative Number. Pls re-enter a new on with positive Number: ")
		}

	}

	checkGrade(grade)
}

