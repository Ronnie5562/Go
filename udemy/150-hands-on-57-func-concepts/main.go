package main

func main() {

}

func sum(li []int) (total int) {
	total = 0
	for _, v := range li {
		total += v
	}
	return
}