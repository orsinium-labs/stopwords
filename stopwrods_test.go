package stopwords_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/stopwords"
)

func TestFind_Contains(t *testing.T) {
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
