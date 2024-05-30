package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	// create map to store exsistance of number
	numbers := make(map[int]struct{}) 


	for _, number := range nums {
		//Check if number exsist in map 
		_, exsists := numbers[number]
		if exsists {
			return true
		} else {
		numbers[number] = struct{}{} // using struct instead of any value for memory usage
		}

	}

	return false
}

func main(){
	nums := []int{1, 1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))

}