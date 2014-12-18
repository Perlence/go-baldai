package board

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// [ 0] [ 1] [ 2] [ 3] [ 4]
// [ 5] [ 6] [ 7] [ 8] [ 9]
// [10] [11] [12] [13] [14]
// [15] [16] [17] [18] [19]
// [20] [21] [22] [23] [24]

func TestStdBoard(t *testing.T) {
	assert := assert.New(t)

	b := NewStdBoard("hello", 5)
	assert.Equal(b.Len(), 5*5,
		"number of cells must be square of board size")

	for i := range Cells(b)[:10] {
		assert.Equal(b.Get(i), 0,
			"first 10 cells must be unoccupied")
	}
	for i := range Cells(b)[10:15] {
		assert.NotEqual(b.Get(i+10), 0,
			"cells from 11 to 15 must be occupied")
	}
	for i := range Cells(b)[15:] {
		assert.Equal(b.Get(i+15), 0,
			"rest cells must be unoccupied")
	}

	assert.Equal(b.GetAdjacent(0), []int{1, 5},
		"cell 0 must have 2 neighbours")
	assert.Equal(b.GetAdjacent(4), []int{3, 9},
		"cell 4 must have 2 neighbours")
	assert.Equal(b.GetAdjacent(11), []int{10, 12, 6, 16},
		"cell 11 must have 4 neighbours")
	assert.Equal(b.GetAdjacent(20), []int{21, 15},
		"cell 20 must have 2 neighbours")
	assert.Equal(b.GetAdjacent(21), []int{20, 22, 16},
		"cell 21 must have 3 neighbours")
}
