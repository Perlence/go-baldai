package board_test

import (
	"github.com/Perlence/go-baldai"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoard(t *testing.T) {
	assert := assert.New(t)

	b := board.NewStdBoard("hello", 5)
	assert.Equal(len(b.Cells()), 5*5,
		"number of cells must be square of board size")

	assert.Equal(b.Get(0), 0, "empty cell must be 0")
	assert.Equal(b.Get(10), 'h', "cell 10 must be 'h'")

	for i := range b.Cells()[:10] {
		assert.False(board.IsOccupied(b, i),
			"first 10 cells must be unoccupied")
	}
	for i := range b.Cells()[10:15] {
		assert.True(board.IsOccupied(b, i+10),
			"cells from 11 to 15 must be occupied")
	}
	for i := range b.Cells()[15:] {
		assert.False(board.IsOccupied(b, i+15),
			"rest cells must be unoccupied")
	}

	assert.Empty(board.GetNeighbours(b, 0),
		"0 cell must not have occupied neighbours")
	assert.Equal(board.GetNeighbours(b, 11), []int{10, 12},
		"cell 11 must have 10 and 12 as neighbours")
}
