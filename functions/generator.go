package functions

//NewIntGenerator return a new int value
func NewIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

/*

func main() {
	gen := functions.NewIntGenerator()
	fmt.Println(gen(), gen(), gen(), gen(), gen())
	fmt.Println(gen(), gen(), gen(), gen(), gen())
}

*/
