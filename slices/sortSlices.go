package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		Рассмотрим функцию Slice(x interface{}, less func(i, j int) bool).
		В описании функции присутствует неизвестный тип данных interface{}.
		Понятие интерфейса будет рассмотренно в следующих модулях курса.
		Следует запомнить, что пустой интерфейс interface{} в Go означает тип данных, под который подходит любой другой тип.
	*/
	nums := []int{1, 2, 3, 4, 9, 7, 8, 9}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	}) // sorted arr [1 2 3 4 7 8 9 9]

	fmt.Println(nums)
	/*
		Разница между Slice and SliceStable
		Стабильная сортировка сохраняет положение одинаковых элементов, тогда как обычная может не сохранять.
	*/
}
