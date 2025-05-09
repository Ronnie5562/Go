package main

import "fmt"

func main() {
	fmt.Println("Chapter 4: Complex Types")
	fmt.Println()

	fmt.Println("Array")
	Array()
	fmt.Println()

	fmt.Println("Slice")
	Slice()
	fmt.Println()

	fmt.Println("Maps")
	maps()
	fmt.Println()

	fmt.Println("Structs")
	structs()
	fmt.Println()

	fmt.Println("Struct Embeddings and Composition")
	struct_embeddings_and_composition()
	fmt.Println()
}

func Array() {
	var arr [10]int
	fmt.Println(arr)

	arr2 := [10]string{"Hi", "Hello"}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(len(arr3))
}

func Slice() {
	var slice []int
	fmt.Println(slice)

	slice2 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice2)

	slice3 := make([]int, 10)
	fmt.Println(slice3)

	slice4 := make([]int, 10, 20)
	fmt.Printf("Length: %v, Capacity: %v \n", len(slice4), cap(slice4))
	fmt.Println(slice4)

	slice5 := []int{1, 2, 3, 4, 5}
	slice6 := slice5[0:3]
	fmt.Println(slice6)

	fmt.Println()

	l1, l2, l3 := linked()
	fmt.Println("Linked :", l1, l2, l3)

	nl1, nl2 := noLink()
	fmt.Println("No Link :", nl1, nl2)

	cl1, cl2 := capLinked()
	fmt.Println("Cap Link :", cl1, cl2)

	cnl1, cnl2 := capNoLink()
	fmt.Println("Cap No Link :", cnl1, cnl2)

	copyl, copy2, copied := copyNoLink()
	fmt.Print("Copy No Link: ", copyl, copy2)
	fmt.Printf(" (Number of elements copied %v)\n", copied)

	a1, a2 := appendNoLink()
	fmt.Println("Append No Link:", a1, a2)
}

func linked() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}

	s2 := s1

	s3 := s1[:]

	s1[2] = 30

	return s1[2], s2[2], s3[2]
}

func noLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}

	s2 := s1

	s1 = append(s1, 6)
	s1[3] = 99

	return s1[3], s2[3]
}

func capLinked() (int, int) {
	s1 := make([]int, 5, 10)

	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5

	s2 := s1

	s1 = append(s1, 6)

	s1[3] = 99

	return s1[3], s2[3]
}

func capNoLink() (int, int) {
	s1 := make([]int, 5, 10)

	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5

	s2 := s1

	s1 = append(s1, []int{10: 11}...)

	s1[3] = 99

	return s1[3], s2[3]
}

func copyNoLink() (int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}

	s2 := make([]int, len(s1))

	copied := copy(s2, s1)

	s1[3] = 99

	return s1[3], s2[3], copied
}

func appendNoLink() (int, int) {
	s1 := []int{1, 2, 3, 4, 5}

	s2 := append([]int{}, s1...)

	s1[3] = 99

	return s1[3], s2[3]
}

func maps() {
	var m map[string]int
	fmt.Println(m)

	m2 := map[string]int{"one": 1, "two": 2}
	fmt.Println(m2)

	m3 := make(map[string]int)
	fmt.Println(m3)

	m3["one"] = 1
	m3["two"] = 2
	fmt.Println(m3)

	delete(m3, "one")
	fmt.Println(m3)

	v, ok := m3["one"]
	fmt.Println("Value:", v, "Present:", ok)

	v2, ok2 := m3["two"]
	fmt.Println("Value:", v2, "Present:", ok2)
}

func structs() {
	type person struct {
		name string
		age  int
	}

	p := person{name: "John", age: 20}
	fmt.Println(p)

	p2 := person{"Jane", 30}
	fmt.Println(p2)

	p.age = 21
	fmt.Println(p)

	p3 := &p
	p3.age = 22
	fmt.Println(p)
}

type name string

type location struct {
	x int
	y int
}

type size struct {
	width  int
	height int
}

type dot struct {
	name
	location
	size
}

func getDots() []dot {
	dot0 := dot{}
	var dot1 dot

	dot2 := dot{}

	dot2.name = "A"
	dot2.x = 5
	dot2.y = 10
	dot2.width = 100
	dot2.height = 200

	dot3 := dot{
		name:     "B",
		location: location{x: 10, y: 20},
		size:     size{width: 200, height: 400},
	}

	dot4 := dot{}
	dot4.name = "C"
	dot4.location.x = 101
	dot4.location.y = 202
	dot4.size.width = 303
	dot4.size.height = 404

	return []dot{dot0, dot1, dot2, dot3, dot4}
}

func struct_embeddings_and_composition() {
	dots := getDots()

	for i := 0; i < len(dots); i++ {
		fmt.Printf("dot%v: %#v\n", i, dots[i])
	}
}
