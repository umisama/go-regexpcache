# umisama/go-regexpcache [![Build Status](https://travis-ci.org/umisama/go-regexpcache.svg)](https://travis-ci.org/umisama/go-regexpcache)

## what's this?
This is a cacheable regexp package.  If expression is used two times or more, this returns cached regexp.Regexp object in very less time.
go-regexpcache has compatible interface with official regexp package.

Here is simple usage:

```go
import "github.com/umisama/go-regexpcache"

// First time, returns regexp is parsed.
match := regexpcache.MustCompile("^[hc]at").MatchString("cat")
println(match)

// After that, returns regexp from cache in less time.
match = regexpcache.MustCompile("^[hc]at").MatchString("cat")
println(match)
```

You don't have to pollute global variables with go-regexpcache.

before:

```go
var re = regexp.MustCompile("^[hc]at")

func varidation(str string) bool {
	// regexp.Regexp object is global for avoid compile time
	return re.MatchString(str)
}
```

after:

```go
func varidation(str string) bool {
	// You can write regexp inside function
	return regexpcache.MustCompile("^[hc]at").MatchString(str)
}
```

## intallation
```
go get github.com/umisama/go-regexpcache
``` 

## document
 * [godoc.org](http://godoc.org/github.com/umisama/go-regexpcache)

## license
under the MIT License
