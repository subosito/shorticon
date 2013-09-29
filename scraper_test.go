package shorticon

import (
	"testing"
)

var urlFavicons = map[string]string{
	"http://github.com/":   "https://github.com/favicon.ico",
	"http://subosito.com/": "http://subosito.com/favicon.ico",
	"https://twitter.com/": "https://abs.twimg.com/favicons/favicon.ico",
	"http://nytimes.com/":  "http://css.nyt.com/images/icons/nyt.ico",
}

func TestScraper(t *testing.T) {
	for k, v := range urlFavicons {
		s := NewScraper(k)

		ux, err := s.Favicon()
		if err != nil {
			t.Errorf("s.Favicon returned an error, %s", err)
		}

		if ux.String() != v {
			t.Errorf("s.Favicon returned %s, want %s", ux.String(), v)
		}
	}
}
