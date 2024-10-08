package hashtag

import (
	"log"
	"regexp"
	"strings"
)

const (
	reHashTag = `#\w+`
)

type Parser struct {
	re *regexp.Regexp

	caseSensitive bool
}

func NewParser(caseSensitive bool) *Parser {
	parser := &Parser{
		caseSensitive: caseSensitive,
	}

	parser.re = regexp.MustCompile(reHashTag)

	return parser
}

// Parse Parsing text for hashtags. If no hashtags were found, nil will be returned.
// Returns a map where hashtag is the key and value is the number uses of the hashtag in the text.
func (receiver *Parser) Parse(text string) map[string]int {
	if len(text) == 0 {
		return nil
	}

	text = strings.ReplaceAll(text, ` `, "\n")
	log.Printf("t: `%s`", text)

	log.Printf("found: %++v\n", receiver.re.FindAllStringSubmatch(text, -1))
	found := receiver.re.FindAllString(text, -1)
	log.Printf("found: %++v\n", found)

	l := len(found)

	if l == 0 {
		return nil
	}

	result := make(map[string]int, l)

	for _, s := range found {
		if s[len(s)-1:] == "#" {
			continue
		}

		key := strings.TrimSpace(s)[1:]
		if !receiver.caseSensitive {
			key = strings.ToLower(key)
		}

		log.Printf("key: `%s`", key)

		result[key]++
	}

	if len(result) == 0 {
		return nil
	}

	return result
}
