package baldai

import (
	"fmt"
	"github.com/Perlence/go-baldai/board"
	"strings"
)

type Path []int

type CellWord struct {
	cell int
	word string
}

type Solver struct {
	board board.Board
	dict  *Dictionary
}

// Check if the path of sequence is turning 180 degrees on some point
func (self *Solver) checkRepetitions(seq Path) bool {
	if len(seq) > 2 {
		// i := 2
		for prev, cell := range seq[2:] {
			if seq[prev] == cell {
				return true
			}
			// i += 1
		}
		return false
	} else if len(seq) == 2 {
		return seq[0] == seq[1]
	}
	return false
}

// Get possible next step for the given sequence
func (self *Solver) getNormalNeighbours(seq Path) []int {
	b, d := self.board, self.dict
	// Get adjacent cells if all cells in given sequence are occupied
	var neighbours []int
	if board.AllOccupied(b, seq) {
		neighbours = b.GetAdjacent(seq[len(seq)-1])
	} else {
		neighbours = board.GetNeighbours(b, seq[len(seq)-1])
	}
	var result []int
	for _, cell := range neighbours {
		if !self.checkRepetitions(append(seq, cell)) &&
			d.StartsWith(board.GetWord(b, append(seq, cell))) {
			result = append(result, cell)
		}
	}
	return result
}

// Get all possible paths in current state
func (self *Solver) walk() chan Path {
	out := make(chan Path)
	go func() {
		defer close(out)
		var result []Path
		for cell := 0; cell < self.board.Len(); cell++ {
			for _, n := range self.getNormalNeighbours(Path{cell}) {
				path := Path{cell, n}
				result = append(result, path)
				out <- path
			}
		}
		for i := 0; i < len(result); i++ {
			word := result[i]
			for _, n := range self.getNormalNeighbours(word) {
				path := append(word, n)
				result = append(result, path)
				out <- path
			}
		}
	}()
	return out
}

// Get only words that miss a char
func (self *Solver) getWords() chan CellWord {
	out := make(chan CellWord)
	go func() {
		defer close(out)
		for seq := range self.walk() {
			for _, cell := range seq {
				if !board.IsOccupied(self.board, cell) {
					out <- CellWord{cell, board.GetWord(self.board, seq)}
				}
			}
		}
	}()
	return out
}

func (self *Solver) win() {
	for cellWord := range self.getWords() {
		possibilities := self.dict.PossibleWords(cellWord.word)
		if len(possibilities) > 0 {
			fmt.Printf("%d\t%s\t%s\n", cellWord.cell, cellWord.word,
				strings.Join(possibilities, ", "))
		}
	}
}
