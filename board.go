package board

import (
	"fmt"
	"strings"
)

type Board interface {
	// Iterable containing all cells in linear order
	Cells() []rune
	// Get value of cell
	Get(cell int) rune
	// Set value of cell
	Set(cell int, char rune)
	// Get all adjacent cells
	GetAdjacent(cell int) []int
}

func IsOccupied(b Board, cell int) bool {
	return b.Get(cell) != 0
}

// Get only occupied adjacent cells
func GetNeighbours(b Board, cell int) []int {
	var result []int
	for _, cell := range b.GetAdjacent(cell) {
		if IsOccupied(b, cell) {
			result = append(result, cell)
		}
	}
	return result
}

// Get word from sequence of cells
func GetWord(b Board, seq []int) string {
	word := ""
	for _, cell := range seq {
		if b.Get(cell) > 0 {
			word += string(b.Get(cell))
		} else {
			word += "*"
		}
	}
	return word
}

type StdBoard struct {
	size  int
	cells []rune
}

func NewStdBoard(word string, size int) *StdBoard {
	squareSize := size * size
	cells := make([]rune, squareSize)
	mid := size * (size / 2)
	if word > "" {
		runeWord := []rune(word)
		for i := range cells[mid : mid+size] {
			cells[mid+i] = runeWord[i]
		}
	}
	return &StdBoard{size, cells}
}

func (self *StdBoard) Cells() []rune {
	return self.cells
}

func (self *StdBoard) Get(cell int) rune {
	return self.cells[cell]
}

func (self *StdBoard) Set(cell int, char rune) {
	self.cells[cell] = char
}

func (self *StdBoard) GetAdjacent(cell int) []int {
	horizontal := []int{-1, 1}
	vertical := []int{-self.size, self.size}
	var result []int
	for _, dc := range horizontal {
		if (cell+dc)%self.size == cell%self.size+dc {
			result = append(result, cell+dc)
		}
	}
	for _, dc := range vertical {
		result = append(result, cell+dc)
	}

	var filtered []int
	for _, c := range result {
		if 0 <= c && c < self.size*self.size {
			filtered = append(filtered, c)
		}
	}
	return filtered
}

func (self *StdBoard) String() string {
	var rows []string
	result := ""
	for cell, char := range self.cells {
		if char != 0 {
			result += fmt.Sprintf("[%c] ", char)
		} else {
			result += "[ ] "
		}
		if (cell+1)%self.size == 0 {
			rows = append(rows, result)
			result = ""
		}
	}
	return strings.Join(rows, "\n")
}
