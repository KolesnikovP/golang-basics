package main

import "fmt"

// если переменная не проинициализирована каким либо значением то ей будет присвоен zero-value свойственный для кождого типа
// Нулевое значение - это ключевая особенность Go, чтобы выделенную память нельзя было прочитать
func main() {
	var a int
	var sum int = 10
	var sum0 = 10
	sum1 := 10

	fmt.Println(sum, sum1, a, sum0)
}

// simple types
// bool
// string
// int8, uint8, int16, uint16, int32, uint42, int64, uint64, byte, rune (для UTF-8 код поинтов)
// float32, float64
