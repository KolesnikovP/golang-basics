package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// array
	var array [3]string // объявили массив типа стринг. И поскольку мы не задали значений там будет три пустоты
	// array1 := [3]bool{} // the same
	array2 := [3]int{1, 2, 3}
	array[1] = "be"

	for _, i := range array2 {
		fmt.Println(">>>", i)
	}

	fmt.Println(len(array))

	// slice - just dont choose the length
	slice := []int{2, 3, 4, 56}

	slice = append(slice, 123)

	for _, i := range slice {
		fmt.Println("item of slice", i)
	}

	// alternative of slice
	slice1 := make([]int, 10) // тоже самое. просто указали еще и длину, будет заполнен нулями
	fmt.Println(slice1)

	/*
		Домашнее задание:
		 просмотреть с помощью цикла любую директорию на компьютере, в которой есть папки / файлы и создать:
		  - массив, длинной равной количеству элементов в папке, после чего заполнить массив элементами, которые . находятся в папке ( смотрим папку - создаем массив - вписываем в массив названия папок / файлов ). В конце вывести массив и длину массива
		 - пустой срез, который наполнить элементами из просмотренной папки. Наполнить срез в цикле. В конце вывести срез и длину среза
	*/

	files, err := os.ReadDir("..")
	var filesArr []string

	if err != nil {
		log.Fatalln(err)
	}
	for _, file := range files {
		filesArr = append(filesArr, file.Name())
		fmt.Println(file.Name())
	}
	fmt.Println(filesArr, len(filesArr))
}
