package hashtag_test

import (
	"github.com/itbasis/go-hashtag"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe(
	"Parsing", func() {
		ginkgo.DescribeTable(
			"no hashtags", func(text string, sensitivity bool) {
				gomega.Ω(hashtag.NewParser(sensitivity).Parse(text)).To(gomega.BeNil())
			},
			ginkgo.Entry(nil, "", false),
			ginkgo.Entry(nil, "qw", false),
			ginkgo.Entry(nil, "qw", true),
			ginkgo.Entry(nil, "Qw", false),
			ginkgo.Entry(nil, "Qw", true),
			ginkgo.Entry(nil, "qw qw", false),
			ginkgo.Entry(nil, "$qw", false),
		)

		ginkgo.DescribeTable(
			"Success parsing", func(text string, sensitivity bool, want map[string]int) {
				parser := hashtag.NewParser(sensitivity)
				gomega.Ω(parser.Parse(text)).To(gomega.Equal(want))
			},
			ginkgo.Entry(nil, "#1qw2", false, map[string]int{"1qw2": 1}),
			ginkgo.Entry(nil, "#12_", false, map[string]int{"12_": 1}),
			ginkgo.Entry(nil, "#1_2", false, map[string]int{"1_2": 1}),
			ginkgo.Entry(nil, "#12", false, map[string]int{"12": 1}),
			ginkgo.Entry(nil, "#qw", false, map[string]int{"qw": 1}),
			ginkgo.Entry(nil, "#qw #", false, map[string]int{"qw": 1}),
			ginkgo.Entry(nil, " #qw", false, map[string]int{"qw": 1}),
			ginkgo.Entry(nil, "#q1w", false, map[string]int{"q1w": 1}),
			ginkgo.Entry(nil, "#q12w", false, map[string]int{"q12w": 1}),
			ginkgo.Entry(nil, "#qw1", false, map[string]int{"qw1": 1}),
			ginkgo.Entry(nil, "#qw12", false, map[string]int{"qw12": 1}),
			ginkgo.Entry(nil, "#qw$", false, map[string]int{"qw": 1}),
			ginkgo.Entry(nil, "#qw qw", false, map[string]int{"qw": 1}),
			ginkgo.Entry(nil, "#qw #qw", false, map[string]int{"qw": 2}),
			ginkgo.Entry(nil, "#qw #Qw", false, map[string]int{"qw": 2}),
			ginkgo.Entry(nil, "#qw #Qw", false, map[string]int{"qw": 2}),
			ginkgo.Entry(nil, "#qw #Qw", true, map[string]int{"qw": 1, "Qw": 1}),
			ginkgo.Entry(nil, "#qw sd #qw", false, map[string]int{"qw": 2}),
		)
	},
)
