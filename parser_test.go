package hashtag_test

import (
	"fmt"
	"log"
	"testing"

	go_hashtag "github.com/itbasis/go-hashtag"
	"github.com/stretchr/testify/assert"
)

func TestParser_Parse(t *testing.T) {
	tests := []struct {
		caseSensitive bool
		text          string
		want          map[string]int
	}{
		{},
		{text: "qw"},
		{text: "qw qw"},
		{text: "$qw"},
		{text: "#1qw2", want: map[string]int{"1qw2": 1}},
		{text: "#12_", want: map[string]int{"12_": 1}},
		{text: "#1_2", want: map[string]int{"1_2": 1}},
		{text: "#12", want: map[string]int{"12": 1}},
		{text: "#qw", want: map[string]int{"qw": 1}},
		{text: "#qw #", want: map[string]int{"qw": 1}},
		{text: " #qw", want: map[string]int{"qw": 1}},
		{text: "#q1w", want: map[string]int{"q1w": 1}},
		{text: "#q12w", want: map[string]int{"q12w": 1}},
		{text: "#qw1", want: map[string]int{"qw1": 1}},
		{text: "#qw12", want: map[string]int{"qw12": 1}},
		{text: "#qw$", want: map[string]int{"qw": 1}},
		{text: "#qw qw", want: map[string]int{"qw": 1}},
		{text: "#qw #qw", want: map[string]int{"qw": 2}},
		{text: "#qw #Qw", want: map[string]int{"qw": 2}},
		{text: "#qw #Qw", want: map[string]int{"qw": 2}},
		{caseSensitive: true, text: "#qw #Qw", want: map[string]int{"qw": 1, "Qw": 1}},
		{text: "#qw sd #qw", want: map[string]int{"qw": 2}},
	}
	for i, test := range tests {
		t.Run(
			fmt.Sprintf("#%d", i), func(t *testing.T) {
				log.Println("test text:", test.text)
				assert.EqualValues(t, test.want, go_hashtag.NewParser(test.caseSensitive).Parse(test.text))
			},
		)
	}
}
