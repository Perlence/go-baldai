package baldai

import (
	"bufio"
	"github.com/fvbock/trie"
	"os"
	"strings"
)

type Dictionary struct {
	letters []rune
	trie    *trie.Trie
}

func NewDictionary(files ...string) *Dictionary {
	var letters []rune
	trie := trie.NewTrie()
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
			trie.Add(text)
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
		if self.trie.HasPrefix(s) {
			return true
		}
	}
	return false
}

// Find words that correspond to the template
func (self *Dictionary) PossibleWords(str string) (words []string) {
	for _, letter := range self.letters {
		s := strings.Replace(str, "*", string(letter), -1)
		if self.trie.Has(s) {
			words = append(words, s)
		}
	}
	return
}
