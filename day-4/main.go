package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Advent of Code Day 4!")
	shelf := LoadShelf("input.txt")

	fmt.Println(shelf)

	accessibleShelf := FindAccessibleShelves(shelf)

	// Further processing can be done with accessibleShelf
	fmt.Println(accessibleShelf)

	// Count the accessible rolls
	totalAccessibleRolls := accessibleShelf.CountAccessibleRolls()
	// Now, loop until the accessible rolls is zero

	for accessibleShelf.CountAccessibleRolls() > 0 {
		shelf = accessibleShelf.ConvertToShelf()
		accessibleShelf = FindAccessibleShelves(shelf)
		totalAccessibleRolls += accessibleShelf.CountAccessibleRolls()
		fmt.Println(accessibleShelf)
	}
	fmt.Printf("Total Accessible Rolls: %d\n", totalAccessibleRolls)
}
