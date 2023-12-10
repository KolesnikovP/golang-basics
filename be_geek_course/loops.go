package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 0; i <= 5; i++ {
		fmt.Println(i, "<<")
	}

	text := []string{"be", "geek", "hello", "!"}
	fmt.Println(text[1])

	for index, i := range text {
		fmt.Println(i, index)
	}

	// alternative to while
	index := 0
	for {
		fmt.Println("geek")
		index = index + 1
		fmt.Println(">>>", text[index])
		if text[index] == "!" {
			fmt.Println("break")
			break
		}
	}

	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
