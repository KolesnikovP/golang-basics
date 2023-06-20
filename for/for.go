package main

import "fmt"

func main() {
	nums := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		nums = append(nums, i)
		fmt.Printf("for %d ", i)
	}

	// for as while
	i := 0
	for i < 10 {
		fmt.Printf("iteration %d \n", i)
		i++
	}

	// if we didn't set condition it means true
	k := 0
	for {
		if k == 10 {
			break
		}
		fmt.Printf("iteration without cond %d \n", k)
		k++
	}

	// boolean value
	j := 0
	for ok := true; ok; ok = j <= 5 {
		fmt.Println(j)
		j++
	}

	// range
	names := []string{"a", "b", "c", "d"}
	for i, el := range names {
		fmt.Printf("index %d value %s \n", i, el)
	}

	for i := range names {
		fmt.Printf("i is %d .It's just an index \n", i)
	}

}
