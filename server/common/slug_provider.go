package common

import (
	"io"
	"strings"

	"github.com/gosimple/slug"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type SlugProvider struct{}

func (sl *SlugProvider) GenerateSlug(input string) string {
	// Normalize the string
	t := transform.NewReader(strings.NewReader(input), norm.NFD)
	normalized, _ := io.ReadAll(t)

	// Convert to slug
	return slug.Make(string(normalized))
}
