package main

import "fmt"

func FindMaxInt(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func FindMaxFloat(nums []float64) float64 {
	if len(nums) == 0 {
		return -1
	}

	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

// using generics to find max number
func FindMaxGeneric[Num int | float64](nums []Num) Num {
	if len(nums) == 0 {
		return -1
	}

	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func main() {
	// Example usage of FindMaxInt
	max := FindMaxInt([]int{1, 32, 5, 8, 10, 11})
	fmt.Println("The maximum int number is:", max)

	// Example usage of FindMaxFloat
	maxFloat := FindMaxFloat([]float64{1.2, 3.4, 5.6, 7.8, 9.0})
	fmt.Println("The maximum float number is:", maxFloat)

}
