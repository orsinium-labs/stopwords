package stopwords

import (
	"iter"
	"unicode"
)

// A word in a text.
type Match struct {
	// The index of the first rune of the word.
	Start int
	// The index of the first rune after the word.
	End int
	// The star of the show, the complete word in its original case.
	Word string
}

// Iterate over all stopwords in the text
func Find(text string) iter.Seq[Match] {
	return func(yield func(Match) bool) {
		for match := range iterWords(text) {
			if IsStopword(match.Word) {
				keepGoing := yield(match)
				if !keepGoing {
					break
				}
			}
		}
	}
}

// Iterate over all words in the text except the stopwords.
func Exclude(text string) iter.Seq[Match] {
	return func(yield func(Match) bool) {
		for match := range iterWords(text) {
			if !IsStopword(match.Word) {
				keepGoing := yield(match)
				if !keepGoing {
					break
				}
			}
		}
	}
}

// Check if the given word is a stopword.
func IsStopword(word string) bool {
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
