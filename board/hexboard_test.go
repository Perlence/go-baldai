package board_test

import (
	"github.com/Perlence/go-baldai/board"
	"github.com/stretchr/testify/assert"
	"testing"
)

// [ 0] [ 1] [ 2] [ 3] [ 4]
//   [ 5] [ 6] [ 7] [ 8]
// [ 9] [10] [11] [12] [13]
//   [14] [15] [16] [17]
// [18] [19] [20] [21] [22]

func TestHexBoard(t *testing.T) {
	assert := assert.New(t)

	b := board.NewHexBoard("hello", 5)
	assert.Equal(b.Len(), 23,
		"number of cells must be 23")

	for i := range board.Cells(b)[:9] {
		assert.Equal(b.Get(i), 0,
			"first 10 cells must be unoccupied")
	}
	for i := range board.Cells(b)[9:14] {
		assert.NotEqual(b.Get(i+9), 0,
			"cells from 9 to 14 must be occupied")
	}
	for i := range board.Cells(b)[14:] {
		assert.Equal(b.Get(i+14), 0,
			"rest cells must be unoccupied")
	}

	assert.Equal(b.GetAdjacent(0), []int{5, 1},
		"cell 0 must have 2 neighbours")
	assert.Equal(b.GetAdjacent(1), []int{5, 6, 0, 2},
		"cell 1 must have 4 neighbours")
}
