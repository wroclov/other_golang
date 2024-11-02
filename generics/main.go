package main
import "fmt"


type Number interface {
    int64 | float64
}

// SumInts adds together the values of main

func SumInts(m map[string]int64) int64 {
	var result int64
	for _, v := range m {
		result += v
	}
	return result
}

func SumFloats(m map[string]float64) float64 {
	var result float64
	for _, v := range m {
		result += v
	}
	return result
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var result V
    for _, v := range m {
        result += v
    }
    return result
}

// SumNumbers sums the values of map m. Its supports both integers
// and floats as map values.
func SumNumbers[Krzeselko comparable, V Number](m map[Krzeselko]V) V {
    var result V
    for _, v := range m {
        result += v
    }
    return result
}

func main() {
	// initialize a map for the integer values
	ints := map[string]int64{
		"first": 33,
		"second": 77,
	}

	// initialize a map for the floats values
	floats := map[string]float64{
		"first": 33.7,
		"second": 77.8,
	}
	
	
	fmt.Printf("Non-Generic Sums: %v and %v\nand same with explicitly provided types %v and %v \n",
		SumInts(ints),
		SumFloats(floats),
		SumIntsOrFloats[string, int64](ints),
        SumIntsOrFloats[string, float64](floats))

	fmt.Printf("You can achieve with generics! %v and %v \n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

    fmt.Printf("Generic Sums with Constraint: %v and %v\n",
        SumNumbers(ints),
        SumNumbers(floats))
}