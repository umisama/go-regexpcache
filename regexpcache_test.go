package regexpcache

import (
	"regexp"
	"testing"
)

func TestContainer(t *testing.T) {
	cases := []struct {
		exp          string
		str          string
		expect_err   bool
		expect_match bool
	}{
		{"^[hc]at", "cat", false, true},
		{"^[hc]at", "hat", false, true},
		{"^[hc]at", "hot", false, false},
		{`^^^[ddd]!!\1\1\1\1`, "hot", true, true},
	}

	// not posix
	cont := newContainer(false)
	for _, c := range cases {
		re, err := cont.Get(c.exp)
		if (err != nil) != c.expect_err {
			t.Error("expect error, but got", err.Error())
		}
		if c.expect_err {
			continue
		}
		match := re.MatchString(c.str)
		if match != c.expect_match {
			t.Error("expect ", c.expect_match, ", but got ", match)
		}
	}

	// posix
	cont = newContainer(true)
	for _, c := range cases {
		re, err := cont.Get(c.exp)
		if (err != nil) != c.expect_err {
			t.Error("expect error, but got", err.Error())
		}
		if c.expect_err {
			continue
		}
		match := re.MatchString(c.str)
		if match != c.expect_match {
			t.Error("expect ", c.expect_match, ", but got ", match)
		}
	}
}

func BenchmarkRegexpPackage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		regexp.MustCompile(`^[hc]at`).MatchString("cat")
	}
}

func BenchmarkRegexpCachePackage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MustCompile(`^[hc]at`).MatchString("cat")
	}
}
