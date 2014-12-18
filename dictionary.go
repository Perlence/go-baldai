package baldai

import (
	"bufio"
	"fmt"
	"github.com/timtadh/data-structures/trie"
	"os"
	"strings"
)

type Dictionary struct {
	letters []rune
	trie    *trie.TST
}

func NewDictionary(files ...string) *Dictionary {
	var letters []rune
	trie := new(trie.TST)
	for _, filename := range files {
		// fp = codecs.open(filename)
		// words += fp.read().splitlines()
		// self.letters = set(''.join(words))
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			text := scanner.Text()
			trie.Put([]byte(text), nil)
			for _, ch := range []rune(text) {
				if !strings.ContainsRune(string(letters), ch) {
					letters = append(letters, ch)
				}
			}
		}

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
	return &Dictionary{letters, trie}
}

// Check if dictionary has words starting with the given prefix
func (self *Dictionary) StartsWith(str string) bool {
	for _, letter := range self.letters {
		s := strings.Replace(str, "*", string(letter), -1)
		iterator := self.trie.PrefixFind([]byte(s))
		if k, v, next := iterator(); k != nil {
			return true
		}
	}
	return false
}

// Find words that correspond to the template
func (self *Dictionary) PossibleWords(str string) (words []string) {
	for _, letter := range self.letters {
		s := strings.Replace(str, "*", string(letter), -1)
		if self.trie.Has([]byte(s)) {
			words = append(words, s)
		}
	}
	return
}
