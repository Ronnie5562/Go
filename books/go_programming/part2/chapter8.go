package main

import "fmt"

func main() {
	fmt.Println("Generics")

	max_int := findMaxGeneric([]int{1, 32, 5, 8, 10, 11})
	fmt.Println("Max Integer value: ", max_int)

	max_float := findMaxGeneric(
		[]float64{1.1, 32.2, 5.5, 8.8, 10.0, 11.1},
	)
	fmt.Println("Max Float value: ", max_float)

	PrintNumOrString(10)
	PrintNumOrString(10.5)
}


func findMaxGeneric[Num int | float64](nums []Num) Num {
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


type RonnieNumberOrString interface {
	int | float64 | string
}


func PrintNumOrString[Num RonnieNumberOrString](num Num) Num {
	fmt.Println(num)
	return num
}
