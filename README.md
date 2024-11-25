# stopwords

Go package for detecting and removing stopwords from a text.

## Installation

```bash
go get github.com/orsinium-labs/stopwords
```

## Usage

```go
language := "en"
sw := stopwords.MustGet(language)
for match := sw.Find(input) {
    fmt.Println(match.Word)
}
```
