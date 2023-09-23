package hashtag_test

import (
	"github.com/itbasis/go-hashtag"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe(
	"Parsing", func() {
		DescribeTable(
			"no hashtags", func(text string, sensitivity bool) {
				Ω(hashtag.NewParser(sensitivity).Parse(text)).To(BeNil())
			},
			Entry(nil, "", false),
			Entry(nil, "qw", false),
			Entry(nil, "qw", true),
			Entry(nil, "Qw", false),
			Entry(nil, "Qw", true),
			Entry(nil, "qw qw", false),
			Entry(nil, "$qw", false),
		)

		DescribeTable(
			"Success parsing", func(text string, sensitivity bool, want map[string]int) {
				parser := hashtag.NewParser(sensitivity)
				Ω(parser.Parse(text)).To(Equal(want))
			},
			Entry(nil, "#1qw2", false, map[string]int{"1qw2": 1}),
			Entry(nil, "#12_", false, map[string]int{"12_": 1}),
			Entry(nil, "#1_2", false, map[string]int{"1_2": 1}),
			Entry(nil, "#12", false, map[string]int{"12": 1}),
			Entry(nil, "#qw", false, map[string]int{"qw": 1}),
			Entry(nil, "#qw #", false, map[string]int{"qw": 1}),
			Entry(nil, " #qw", false, map[string]int{"qw": 1}),
			Entry(nil, "#q1w", false, map[string]int{"q1w": 1}),
			Entry(nil, "#q12w", false, map[string]int{"q12w": 1}),
			Entry(nil, "#qw1", false, map[string]int{"qw1": 1}),
			Entry(nil, "#qw12", false, map[string]int{"qw12": 1}),
			Entry(nil, "#qw$", false, map[string]int{"qw": 1}),
			Entry(nil, "#qw qw", false, map[string]int{"qw": 1}),
			Entry(nil, "#qw #qw", false, map[string]int{"qw": 2}),
			Entry(nil, "#qw #Qw", false, map[string]int{"qw": 2}),
			Entry(nil, "#qw #Qw", false, map[string]int{"qw": 2}),
			Entry(nil, "#qw #Qw", true, map[string]int{"qw": 1, "Qw": 1}),
			Entry(nil, "#qw sd #qw", false, map[string]int{"qw": 2}),
		)
	},
)
