package regexpcache

import (
	"io"
	"regexp"
	"strconv"
	"sync"
)

var (
	regexpContainer, posixContainer container
)

// Compile parses a regular expression.
// This compatible with regexp.Compile but this uses a cache.
func Compile(str string) (*regexp.Regexp, error) {
	return regexpContainer.Get(str)
}

// CompilePOSIX parses a regular expression with POSIX ERE syntax.
// This compatible with regexp.CompilePOSIX but this uses a cache.
func CompilePOSIX(str string) (*regexp.Regexp, error) {
	return posixContainer.Get(str)
}

// MustCompile parses a regular expression and panic if str can not parsed.
// This compatible with regexp.MustCompile but this uses a cache.
func MustCompile(str string) *regexp.Regexp {
	re, err := regexpContainer.Get(str)
	if err != nil {
		// regexp.MustCompile() like message
		panic(`regexpcache: Compile(` + quote(str) + `): ` + err.Error())
	}
	return re
}

// MustCompilePOSIX parses a regular expression with POSIX ERE syntax and panic if str can not parsed.
// This compatible with regexp.MustCompilePOSIX but this uses a cache.
func MustCompilePOSIX(str string) *regexp.Regexp {
	re, err := posixContainer.Get(str)
	if err != nil {
		// regexp.MustCompilePOSIX() like message
		panic(`regexpcache: CompilePOSIX(` + quote(str) + `): ` + err.Error())
	}
	return re
}

// Match checks whether a textual regular expression matches a byte slice.
// This compatible with regexp.Match but this uses a cache.
func Match(pattern string, b []byte) (matched bool, err error) {
	re, err := Compile(pattern)
	if err != nil {
		return false, err
	}
	return re.Match(b), nil
}

// Match checks whether a textual regular expression matches the RuneReader.
// This compatible with regexp.MatchReader but this uses a cache.
func MatchReader(pattern string, r io.RuneReader) (matched bool, err error) {
	re, err := Compile(pattern)
	if err != nil {
		return false, err
	}
	return re.MatchReader(r), nil
}

// Match checks whether a textual regular expression matches a string.
// This compatible with regexp.MatchString but this uses a cache.
func MatchString(pattern string, s string) (matched bool, err error) {
	re, err := Compile(pattern)
	if err != nil {
		return false, err
	}
	return re.MatchString(s), nil
}

type container struct {
	regexps map[string]*regexp.Regexp
	posix   bool
	mu      *sync.Mutex
}

func newContainer(posix bool) container {
	return container{
		regexps: make(map[string]*regexp.Regexp),
		posix:   posix,
		mu:      &sync.Mutex{},
	}
}

func (s *container) Get(str string) (*regexp.Regexp, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	re, ok := s.regexps[str]
	if ok {
		return re, nil
	}

	var err error
	if s.posix {
		re, err = regexp.CompilePOSIX(str)
	} else {
		re, err = regexp.Compile(str)
	}
	if err != nil {
		return nil, err
	}
	s.regexps[str] = re

	reCopy := new(regexp.Regexp)
	*reCopy = *re
	return reCopy, nil
}

func quote(s string) string {
	if strconv.CanBackquote(s) {
		return "`" + s + "`"
	}
	return strconv.Quote(s)
}

func init() {
	regexpContainer = newContainer(false)
	posixContainer = newContainer(true)
}
