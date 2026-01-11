package main

import (
	"os"
	"strings"
)

type Shelf struct {
	rows [][]bool
}

func LoadShelf(filename string) *Shelf {

	// The file will contain "." and "@" characters representing empty and filled spaces respectively.
	file, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	rows := make([][]bool, len(lines)+2)

	// Prepend and append a row of false (empty) to the shelf
	rows[0] = make([]bool, len(lines[0])+2)
	rows[len(rows)-1] = make([]bool, len(lines[0])+2)

	for i, line := range lines {
		rows[i+1] = make([]bool, len(line)+2)

		// Prepend and append false (empty) to each row
		rows[i+1][0] = false
		rows[i+1][len(rows[i+1])-1] = false
		for j, char := range line {
			rows[i+1][j+1] = char == '@'
		}
	}

	return &Shelf{rows: rows}
}

func (shelf *Shelf) String() string {
	var sb strings.Builder
	for i := 0; i < len(shelf.rows); i++ {
		for j := 0; j < len(shelf.rows[i]); j++ {
			if shelf.rows[i][j] {
				sb.WriteString("@")
			} else {
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func BoolToIntConverter(val bool) int {
	if val {
		return 1
	}
	return 0
}

func RollingSumGeneric[T Numeric, TData any](data []TData, window int, converter func(TData) T) []T {
	rollingSum := make([]T, len(data))
	fcq := NewFixedCapacityQueue[T](window)

	// For even windows, prefer backward extent
	// For odd windows, center symmetrically
	var forwardExtent int
	if window%2 == 0 {
		// Even window: prefer looking backward
		forwardExtent = window/2 - 1
	} else {
		// Odd window: center symmetrically
		forwardExtent = (window - 1) / 2
	}

	var cumulativeSum T
	iData := 0
	iSum := 0

	// Now we can start the rolling process
	for iData < forwardExtent && iData < len(data) {
		cumulativeSum += fcq.PushBack(converter(data[iData]))
		iData++
	}

	for iData < len(data) {
		cumulativeSum += fcq.PushBack(converter(data[iData]))
		rollingSum[iSum] = cumulativeSum
		iData++
		iSum++
	}

	for iSum < len(data) {
		cumulativeSum -= fcq.PopFront()
		rollingSum[iSum] = cumulativeSum
		iSum++
	}

	return rollingSum
}

func RollingSumBool(data []bool, window int) []int {
	return RollingSumGeneric(data, window, BoolToIntConverter)
}
func RollingSumInt(data []int, window int) []int {
	return RollingSumGeneric(data, window, func(val int) int { return val })
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type RollAccessibility int

const (
	Accessible RollAccessibility = iota
	Inaccessible
	NoRoll
)

type AccessibleShelf struct {
	shelf           *Shelf
	accessibleRolls [][]RollAccessibility
}

func FindAccessibleShelves(shelf *Shelf) *AccessibleShelf {
	// We're going to do this by making a queue for each row.
	// The queue will have a capacity of 3

	rowQueues := make([]*FixedCapacityQueue[int], len(shelf.rows))
	for i := 0; i < len(shelf.rows); i++ {
		rowQueues[i] = NewFixedCapacityQueue[int](3)
	}

	accessibleRolls := make([][]RollAccessibility, len(shelf.rows))
	for i := 0; i < len(shelf.rows); i++ {
		accessibleRolls[i] = make([]RollAccessibility, len(shelf.rows[i]))
	}

	numberOfSurrounds := make([][]int, len(shelf.rows))
	for i := 0; i < len(shelf.rows); i++ {
		numberOfSurrounds[i] = make([]int, len(shelf.rows[i]))
	}

	const queueCapacity = 3

	for i := 0; i < len(shelf.rows); i++ {
		numberOfSurrounds[i] = RollingSumBool(shelf.rows[i], queueCapacity)
	}

	transposedSurrounds := make([][]int, len(numberOfSurrounds[0]))
	for j := 0; j < len(numberOfSurrounds[0]); j++ {
		transposedSurrounds[j] = make([]int, len(numberOfSurrounds))
		for i := 0; i < len(numberOfSurrounds); i++ {
			transposedSurrounds[j][i] = numberOfSurrounds[i][j]
		}
	}

	const accessibleThreshold = 4

	for j := 0; j < len(transposedSurrounds); j++ {
		thisColumnSurroundingRows := RollingSumInt(transposedSurrounds[j], queueCapacity)
		for i := 0; i < len(thisColumnSurroundingRows); i++ {
			if !shelf.rows[i][j] {
				accessibleRolls[i][j] = NoRoll
			} else if thisColumnSurroundingRows[i] <= accessibleThreshold {
				accessibleRolls[i][j] = Accessible
			} else {
				accessibleRolls[i][j] = Inaccessible
			}
		}
	}

	return &AccessibleShelf{
		shelf:           shelf,
		accessibleRolls: accessibleRolls,
	}
}

func (AccessibleShelf *AccessibleShelf) String() string {
	var sb strings.Builder
	for i := 0; i < len(AccessibleShelf.shelf.rows); i++ {
		for j := 0; j < len(AccessibleShelf.shelf.rows[i]); j++ {
			switch AccessibleShelf.accessibleRolls[i][j] {
			case Accessible:
				sb.WriteString("x")
			case Inaccessible:
				sb.WriteString("@")
			case NoRoll:
				sb.WriteString(".")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (AccessibleShelf *AccessibleShelf) CountAccessibleRolls() int {
	count := 0
	for i := 0; i < len(AccessibleShelf.shelf.rows); i++ {
		for j := 0; j < len(AccessibleShelf.shelf.rows[i]); j++ {
			if AccessibleShelf.accessibleRolls[i][j] == Accessible {
				count++
			}
		}
	}
	return count
}

func (AccessibleShelf *AccessibleShelf) ConvertToShelf() *Shelf {

	rows := make([][]bool, len(AccessibleShelf.shelf.rows))

	for i := 0; i < len(AccessibleShelf.shelf.rows); i++ {
		rows[i] = make([]bool, len(AccessibleShelf.shelf.rows[i]))
		for j := 0; j < len(AccessibleShelf.shelf.rows[i]); j++ {
			rows[i][j] = (AccessibleShelf.accessibleRolls[i][j] == Inaccessible)
		}
	}

	return &Shelf{rows: rows}

}
