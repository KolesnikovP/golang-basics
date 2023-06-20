package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input = readLine("Enter yes or no: ")

	switch input {
	case "yes", "да":
		fmt.Println("YOu've agreed")
		// fallthrough - позволяет провалиться к следующему кейсу независимо от ситуации
	case "no", "нет":
		fmt.Println("You've disagreed")
	default:
		fmt.Println("I don't understand")
	}

	switch {
	case input == "yes" || input == "да":
		fmt.Println("YOu've agreed")
	case input == "no" || input == "нет":
		fmt.Println("You've disagreed")
	default:
		fmt.Println("I don't understand")
	}
}

func readLine(greeting string) string {
	fmt.Print(greeting)
	reader := bufio.NewReader(os.Stdin)
	text, _, _ := reader.ReadLine()
	return string(text)
}
