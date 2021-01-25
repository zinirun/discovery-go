package datastructure

// AddOne Add 1 to nums array
func AddOne(nums []int) {
	for i := range nums {
		nums[i]++
	}
}

/*

func main() {
	n := []int{1, 2, 3, 4}
	datastructure.AddOne(n)
	fmt.Println(n)
}

*/
