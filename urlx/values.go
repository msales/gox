package urlx

import (
	"net/url"
	"strings"
)

// Param is a single URL query parameter.
type Param struct {
	Key string
	Val string
}

// Values is an ordered list of URL query parameters. Unlike net/url.Values,
// it preserves the insertion order of every key/value pair, which is useful
// when downstream systems treat query order as significant.
//
// Values is not safe for concurrent use.
type Values []Param

// Encode encodes the values into "URL encoded" form ("bar=baz&foo=quux"),
// preserving the order in which entries were added.
func (v Values) Encode() string {
	if len(v) == 0 {
		return ""
	}
	var buf strings.Builder
	for i, p := range v {
		if i > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(url.QueryEscape(p.Key))
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(p.Val))
	}
	return buf.String()
}
