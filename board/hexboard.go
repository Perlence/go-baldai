package board

import (
	"fmt"
	"strings"
)

type HexBoard struct {
	size  int
	cells [][]rune
}

func NewHexBoard(word string, size int) *HexBoard {
	cells := make([][]rune, size)
	for i := range cells {
		if i == size/2 && len(word) > 0 {
			cells[i] = []rune(word)
		} else if i%2 == 0 {
			cells[i] = make([]rune, size)
		} else {
			cells[i] = make([]rune, size-1)
		}
	}
	return &HexBoard{size, cells}
}

// Get absolute number of cell in linear order from given row and col
func (self *HexBoard) absolute(r, c int) int {
	result := 0
	for i, row := range self.cells {
		if r == i {
			break
		}
		result += len(row)
	}
	return result + c
}

// Get row and col from absolute cell number
func (self *HexBoard) coords(cell int) (r, c int) {
	firstCell := 0
	for r, row := range self.cells {
		if firstCell <= cell && cell < firstCell+len(row) {
			return r, cell - firstCell
		}
		firstCell += len(row)
	}
	return -1, -1
}

func (self *HexBoard) Len() (length int) {
	for _, row := range self.cells {
		length += len(row)
	}
	return
}

func (self *HexBoard) Get(cell int) rune {
	r, c := self.coords(cell)
	return self.cells[r][c]
}

func (self *HexBoard) Set(cell int, char rune) {
	r, c := self.coords(cell)
	self.cells[r][c] = char
}

func (self *HexBoard) GetAdjacent(cell int) []int {
	r, c := self.coords(cell)
	diagonal := [][]int{
		{-1, -1}, {-1, 0},
		{1, -1}, {1, 0},
	}
	horizontal := [][]int{
		{0, -1}, {0, 1},
	}
	var result []int
	for _, coord := range diagonal {
		dr, dc := coord[0], coord[1]
		dr, dc = dr, dc+r%2
		if 0 <= r+dr && r+dr < self.size &&
			0 <= c+dc && c+dc < self.size-(r+dr)%2 {
			result = append(result, self.absolute(r+dr, c+dc))
		}
	}
	for _, coord := range horizontal {
		dr, dc := coord[0], coord[1]
		if 0 <= c+dc && c+dc < self.size-r%2 {
			result = append(result, self.absolute(r+dr, c+dc))
		}
	}
	return result
}

func (self *HexBoard) String() string {
	var rows []string
	for _, row := range self.cells {
		result := ""
		if len(row) < self.size {
			result += "  "
		}
		for _, char := range row {
			if char > 0 {
				result += fmt.Sprintf("[%c] ", char)
			} else {
				result += "[ ] "
			}
		}
		rows = append(rows, result)
	}
	return strings.Join(rows, "\n")
}
