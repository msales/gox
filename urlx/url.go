package urlx

import (
	"fmt"
	"net/url"
	"strings"
)

// ParseQuery parses the URL-encoded query string and returns a Values
// listing the values specified for each key, preserving their original
// order. The returned Values contains every valid query parameter found;
// err describes the first decoding error encountered, if any.
//
// Query is expected to be a list of key=value settings separated by
// ampersands. A setting without an equals sign is interpreted as a key
// set to an empty value. Settings containing a non-URL-encoded semicolon
// are considered invalid.
func ParseQuery(query string) (v Values, err error) {
	for query != "" {
		var key string
		key, query, _ = strings.Cut(query, "&")
		if strings.Contains(key, ";") {
			err = fmt.Errorf("invalid semicolon separator in query")
			continue
		}
		if key == "" {
			continue
		}
		key, value, _ := strings.Cut(key, "=")
		key, err1 := url.QueryUnescape(key)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		value, err1 = url.QueryUnescape(value)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}

		v = append(v, Param{Key: key, Val: value})
	}

	return v, err
}
