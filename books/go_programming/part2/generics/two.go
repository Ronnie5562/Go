package main

func AcceptAnythingAndReturnIt[T any](val T) T {
	return val
}

func main() {
	_ = AcceptAnythingAndReturnIt(42)
	_ = AcceptAnythingAndReturnIt("Hello, World!")
	_ = AcceptAnythingAndReturnIt(3.14)
	_ = AcceptAnythingAndReturnIt(true)
	_ = AcceptAnythingAndReturnIt([]int{1, 2, 3})

	max := FindMaxGeneric([]int{1, 32, 5, 8, 10, 11})
	println("The maximum int number is:", max)
}

// Type constraints in Go generics refer to interfaces that define sets of types.

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func FindMaxGeneric[Num Number](nums []Num) Num {
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
