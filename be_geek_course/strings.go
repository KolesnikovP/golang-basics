package main

import "fmt"

func main() {
	textHello := "hello"
	textBe := "begeek"
	text := textHello + " " + textBe
	textTemplate := text + " " + fmt.Sprintf("%d", 10)

	// %s - means string, %d - means digit %T - means type
	fmt.Println(fmt.Sprintf("%s %d %s", textTemplate, 12, "!"))

	var input float32
	fmt.Scanf("%f", &input)
	output := input + 1
	fmt.Println(output)

	const zumba string = "zumba"
}
