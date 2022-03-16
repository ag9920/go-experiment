package go118

/*
	Please be sure to read the official articles first
	- Why Generics? https://go.dev/blog/why-generics
	- Generics Tutorial: Tutorial: Getting started with generics
*/

// non generics implementation

// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// add a single generic function that can receive a map containing either integer or float values

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
// type parameters: K and V (inside the square brackets)
// m of type map[K]V, one argument that uses the type parameters
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// Declare the Number interface type to use as a type constraint.
// Declare a union of int64 and float64 inside the interface.
// Essentially, you’re moving the union from the function declaration into a new type constraint.
// That way, when you want to constrain a type parameter to either int64 or float64, you can use this Number type constraint instead of writing out int64 | float64.
type Number interface {
	int64 | float64
}

// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
