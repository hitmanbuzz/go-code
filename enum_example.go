//go:build ignore

package main

import "fmt"

// [uint8] is more than enough for most enum
type FruitKind uint8

const (
	MANGO FruitKind = iota
	APPLE
)

// 1st method
// this is mutable so the below [Stringer] interface implementation [String] method of switch case is better
var FRUITS = map[FruitKind]string{
	MANGO: "mango",
	APPLE: "apple",
}

// 2nd method
// readonly
func (f FruitKind) String() string {
	switch f {
	case MANGO:
		return "mango"
	case APPLE:
		return "apple"
	default:
		return "" // if all enum fields case are covered then it wouldn't reach the default
	}
}

// part of 2nd method
func Do(e FruitKind) {
	if e == MANGO {
		fmt.Println("It is mango")
	}

	fmt.Printf("Enum String: %s\n", e.String())
	fmt.Printf("Enum Number: %d\n", e) // return the uint8 value
}

func main() {
	fruit := MANGO

	// example with map
	name := FRUITS[fruit]
	fmt.Println("From Map:", name)

	// example with switch case
	Do(fruit)
}
