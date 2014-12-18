package baldai

import (
	"fmt"
	"github.com/Perlence/go-baldai/board"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolver(t *testing.T) {
	assert := assert.New(t)

	b := board.NewStdBoard("HELLO", 5)
	dict := NewDictionary("./data/english.txt")
	solver := &Solver{b, dict}

	assert.False(solver.checkRepetitions(Path{10, 11, 12}),
		"path [10 11 12] is not repetitive")
	assert.True(solver.checkRepetitions(Path{10, 11, 10}),
		"path [10 11 10] is repetitive")

	solver.win()
}
