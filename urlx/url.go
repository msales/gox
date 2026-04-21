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
func ParseQuery(query string) (Values, error) {
	var v Values
	for query != "" {
		var key string
		key, query, _ = strings.Cut(query, "&")
		if strings.Contains(key, ";") {
			return nil, fmt.Errorf("invalid semicolon separator in query")
		}
		if key == "" {
			continue
		}
		key, value, _ := strings.Cut(key, "=")
		key, err := url.QueryUnescape(key)
		if err != nil {
			return nil, fmt.Errorf("unescaping key %q: %w", key, err)
		}
		value, err = url.QueryUnescape(value)
		if err != nil {
			return nil, fmt.Errorf("unescaping value %q: %w", value, err)
		}

		v = append(v, Param{Key: key, Val: value})
	}

	return v, nil
}
