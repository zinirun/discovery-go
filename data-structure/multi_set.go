package datastructure

import "strings"

// MultiSet is type of map set
type MultiSet map[string]int

// Insert increase the set's value
func (m MultiSet) Insert(val string) {
	m[val]++
}

// Erase decrease the set's value
func (m MultiSet) Erase(val string) {
	if m[val] <= 1 {
		delete(m, val)
	} else {
		m[val]--
	}
}

// Count returns set's value
func (m MultiSet) Count(val string) int {
	return m[val]
}

func (m MultiSet) String() string {
	s := "{ "
	for val, count := range m {
		s += strings.Repeat(val+" ", count)
	}
	return s + "}"
}

/*

func main() {
	m := datastructure.MultiSet{}
	m.Insert("hello")
	m.Erase("hello")
	m.Erase("hi")
	fmt.Println(m.String())
}

*/
