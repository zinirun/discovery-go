package functions

import "errors"

var errNotTrue = errors.New("not true")

// TestExceptionIsTrue return true if t is true or error
func TestExceptionIsTrue(t bool) error {
	if t == true {
		return nil
	}
	return errNotTrue
}

/*

func main() {
	err := functions.TestExceptionIsTrue(false)
	if err != nil {
		log.Fatal(err)
	}
}

*/
