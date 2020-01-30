package main

import (
	"fmt"
	"strings"
)

func main() {
	var userInput string
	_, err := fmt.Scanf("%s", &userInput)
	if err != nil {
		fmt.Print(err)
	}
	smallerCaseUserInput := strings.ToLower(userInput)
	if strings.HasPrefix(smallerCaseUserInput, "i") && strings.HasSuffix(smallerCaseUserInput, "n") && strings.Contains(smallerCaseUserInput, "a") {
		fmt.Println("Found!\n")
	} else {
		fmt.Println("Not Found!\n")
	}
}