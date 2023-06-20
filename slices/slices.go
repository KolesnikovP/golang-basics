package main

import "fmt"

// slices and arrays
func main() {
	// arrays have certain length
	arr := [5]int{}  // [5] - amount of elements, {} - init values. If {} is empty then will be zero values
	fmt.Println(arr) //  [0, 0, 0, 0, 0]
	var arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1) // long version of array which return [1,2,3]
	// We can't call an array index which is outside a diapason of array
	// arr[6] - return error

	// Slices - this is arrays without certain length (dynamic array)
	var nums = []int{1, 2, 3, 4, 5, 6}
	nums1 := []int{4, 5, 6}
	fmt.Println(nums, nums1) // looks like array [1,2,3], but we don't have length restrictions
	// : - operator helps us to get slice of arrays
	fmt.Println(nums[:2], "[1,2]")       // from start of array to index, but index didn't include
	fmt.Println(nums[2:])                // start from index to end of array
	fmt.Println(nums[2:4])               // [3,4] - diapason from index to index (second index didn't include)
	nums = append(nums[:2], nums[3:]...) // spread
	fmt.Println(nums, ">>>")
	nums = append(nums, 7)
	fmt.Println(nums) // added a new value to array nums
	/*
			Slices are more difficult structure than arrays and have two additional properties: len and cap (capacity of array)
			cap - means how many elements an array can contain.
			That's why we can init arrays of a certain length with the built-in "make" function

		Example:
	*/
	// len = 5, cap = 5
	nums3 := make([]int, 5, 5) // [0,0,0,0,0]
	// len = 0, cap = 5. It means the array will be empty, but an array of the right length is already allocated in memory
	// this is good practice
	nums4 := make([]int, 0, 5) // []
	fmt.Println(nums3)
	fmt.Println(nums4)
	/*
		Передача слайса как аргумента функции происходит хитро.
		Длина и вместимость передаются по значению, но массив значений передается по ссылке.
		Вследствие этого получается неявное поведение: добавленные элементы не сохранятся в исходный слайс,
		но изменение существующих останется:
	*/

	// modify slice
	nums5 := []int{1, 2, 3, 4, 5}
	modifySlice(nums5)
	fmt.Println(nums5) // [1,2,10,4,5]

	var nums6 = modifySlice(nums5)
	fmt.Println(nums6) // [1,2,10,11,5,6]
}

func modifySlice(nums []int) []int {
	nums[2] = 10 // the element will be in origin slice
	// YOU MUST REMEMBER that append return new slice and don't change the origin slice
	// FOR example if we change nums array after append func, then our origin array will be without changes
	nums = append(nums, 6) // the element will not be added in origin slice
	nums[3] = 11           // The origin array will be without any changes. Be careful!!!
	return nums
}
