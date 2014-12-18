package board

type Board interface {
	// Get the number of cells
	Len() int
	// Get value of cell
	Get(cell int) rune
	// Set value of cell
	Set(cell int, char rune)
	// Get all adjacent cells
	GetAdjacent(cell int) []int
}

// Get slice of all runes on the table
func Cells(b Board) []rune {
	cells := make([]rune, b.Len())
	for i := range cells {
		cells[i] = b.Get(i)
	}
	return cells
}

// Check if the cell is occupied
func IsOccupied(b Board, cell int) bool {
	return b.Get(cell) != 0
}

// Get only occupied adjacent cells
func GetNeighbours(b Board, cell int) (neighbours []int) {
	for _, cell := range b.GetAdjacent(cell) {
		if IsOccupied(b, cell) {
			neighbours = append(neighbours, cell)
		}
	}
	return
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
