package baldai

import (
	// "fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDictionary(t *testing.T) {
	assert := assert.New(t)

	dict := NewDictionary("./data/english.txt")

	assert.True(dict.trie.Has([]byte("ABBOT")),
		"trie must have 'ABBOT'")

	assert.True(dict.StartsWith("ADVE*"),
		"there must be at least 4 words starting with 'ADVE'")
}
