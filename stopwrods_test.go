package stopwords_test

import (
	"iter"
	"strings"
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/stopwords"
)

func TestContains_English(t *testing.T) {
	is := is.New(t)
	sw := stopwords.MustGet("en")

	is.True(sw.Contains("or"))
	is.True(sw.Contains("and"))
	is.True(sw.Contains("a"))
	is.True(sw.Contains("Almost"))
	is.True(sw.Contains("almost"))

	is.True(!sw.Contains("love"))
	is.True(!sw.Contains("django"))
	is.True(!sw.Contains("Django"))
}

func iter2text(ms iter.Seq[stopwords.Match]) string {
	result := make([]string, 0)
	for m := range ms {
		result = append(result, m.Word)
	}
	return strings.Join(result, " ")
}

func TestFind_English(t *testing.T) {
	is := is.New(t)
	sw := stopwords.MustGet("en")

	is.Equal(iter2text(sw.Find("never gonna give you up")), "never give you up")
	is.Equal(iter2text(sw.Find("I want an")), "I want an")
}

func TestExclude_English(t *testing.T) {
	is := is.New(t)
	sw := stopwords.MustGet("en")

	is.Equal(iter2text(sw.Exclude("never gonna give you up")), "gonna")
	is.Equal(iter2text(sw.Exclude("I want an apple")), "apple")
}
