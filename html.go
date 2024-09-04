package striprtf

import (
	"bytes"
	"io"
	"regexp"
	"strings"
)

type stringScanner struct {
	data   []byte
	pos    int
	result [][]byte
}

var (
	reHtmlTag  = regexp.MustCompile(`^\\\*\\htmltag(\d+) ?`)
	reMHtmlTag = regexp.MustCompile(`^\\\*\\mhtmltag(\d+) ?`)
)

// ExtractHtml removes RTF formatting from HTML content provided via an io.Reader,
// returning the plain HTML as an io.Reader or an error if extraction fails.
func ExtractHtml(r io.Reader) (io.Reader, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if !bytes.Contains(data, []byte("{\\*\\htmltag")) {
		return stripRtf(bytes.NewReader(data), false)
	}

	var html strings.Builder
	var ignoreTag string

	scanner := newStringScanner(data)
	for !scanner.eos() {
		switch {
		case scanner.scanRegex(reHtmlTag):
			if ignoreTag == string(scanner.result[1]) {
				scanner.scanUntil("}")
				ignoreTag = ""
			}

		case scanner.scanRegex(reMHtmlTag):
			ignoreTag = string(scanner.result[1])

		default:
			html.WriteByte(scanner.increment())
		}
	}

	return stripRtf(strings.NewReader(html.String()), true)
}

func newStringScanner(data []byte) *stringScanner {
	return &stringScanner{data: data}
}

func (s *stringScanner) eos() bool {
	return s.pos >= len(s.data)
}

func (s *stringScanner) scanRegex(re *regexp.Regexp) bool {
	if match := re.FindSubmatch(s.data[s.pos:]); match != nil {
		s.result = match
		s.pos += len(match[0])

		return true
	}

	return false
}

func (s *stringScanner) scanUntil(str string) bool {
	if idx := bytes.Index(s.data[s.pos:], []byte(str)); idx != -1 {
		s.pos += idx + len(str)
		return true
	}

	return false
}

func (s *stringScanner) increment() byte {
	if s.pos < len(s.data) {
		b := s.data[s.pos]
		s.pos++

		return b
	}

	return 0
}
