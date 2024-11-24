package stopwords

import (
	"bufio"
	"embed"
	"iter"
	"strings"
	"sync"
	"unicode"

	"github.com/derekparker/trie/v3"
)

//go:embed words/*.txt
var files embed.FS

var dicts = sync.OnceValue(func() map[string]*Stopwords {
	dicts := make(map[string]*Stopwords)
	dir, err := files.ReadDir("")
	if err != nil {
		panic(err)
	}
	for _, entry := range dir {
		fileName := entry.Name()
		file, err := files.Open(fileName)
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(file)
		dict := trie.New[struct{}]()
		for scanner.Scan() {
			dict.Add(scanner.Text(), struct{}{})
		}
		dicts[fileName[:2]] = (*Stopwords)(dict)
	}
	files = embed.FS{}
	return dicts
})

func init() {
	go dicts()
}

type Stopwords trie.Trie[struct{}]

func Get(lang string) *Stopwords {
	return dicts()[strings.ToLower(lang[:2])]
}

// A word in a text.
type Match struct {
	// The index of the first rune of the word.
	Start int
	// The index of the first rune after the word.
	End int
	// The star of the show, the complete word in its original case.
	Word string
}

// Iterate over all stopwords in the text.
func (s *Stopwords) Find(text string) iter.Seq[Match] {
	return func(yield func(Match) bool) {
		for match := range iterWords(text) {
			if s.Contains(match.Word) {
				keepGoing := yield(match)
				if !keepGoing {
					break
				}
			}
		}
	}
}

// Iterate over all words in the text except the stopwords.
func (s *Stopwords) Exclude(text string) iter.Seq[Match] {
	return func(yield func(Match) bool) {
		for match := range iterWords(text) {
			if !s.Contains(match.Word) {
				keepGoing := yield(match)
				if !keepGoing {
					break
				}
			}
		}
	}
}

// Check if the given word is a stopword.
func (s *Stopwords) Contains(word string) bool {
	panic("todo")
}

// Iterate over all words
func iterWords(text string) iter.Seq[Match] {
	return func(yield func(Match) bool) {
		start := 0
		end := 0
		for i, r := range text {
			if unicode.IsLetter(r) {
				end = i + 1
				continue
			}
			if start < end {
				keepGoing := yield(Match{
					Start: start,
					End:   end,
					Word:  text[start:end],
				})
				if !keepGoing {
					break
				}
			}
			start = i + 1
		}

		if start < end {
			yield(Match{
				Start: start,
				End:   end,
				Word:  text[start:end],
			})
		}
	}
}
