package striprtf

import (
	"bufio"
	"encoding/hex"
	"io"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/text/encoding"
)

type stackEntry struct {
	numberOfCharactersToSkip int
	ignorable                bool
}

var rtfRegex = regexp.MustCompile(`(?i)\\([a-z]{1,32})(-?\d{1,10})?[ ]?|\\'([0-9a-f]{2})|\\([^a-z])|([{}])|[\r\n]+|(.)`)

// ExtractText reads RTF content from an io.Reader and returns the plain text
// without RTF formatting as an io.Reader, or an error if extraction fails.
func ExtractText(r io.Reader) (io.Reader, error) {
	return stripRtf(r, false)
}

func stripRtf(r io.Reader, html bool) (io.Reader, error) {
	var decoder *encoding.Decoder
	var stack []stackEntry
	var ignorable bool
	ucskip := 1
	curskip := 0

	var res strings.Builder
	reader := bufio.NewReader(r)

	var hexToStr []byte
	for {
		chunk, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return nil, err
		}

		chunklc := strings.ToLower(chunk)
		hyperlink := ""
		inHyperlink := false

		matches := rtfRegex.FindAllStringSubmatch(chunk, -1)
		for _, match := range matches {
			word := match[1]
			arg := match[2]
			chex := match[3]
			character := match[4]
			brace := match[5]
			tchar := match[6]

			if !html {
				if strings.Contains(chunklc, "hyperlink") {
					hyperlink += match[0]

					if strings.Contains(strings.ToLower(hyperlink), "hyperlink \"") {
						hyperlink = match[0]
						inHyperlink = true
					} else if match[0] == "\"" && inHyperlink {
						res.WriteString(strings.Trim(strings.TrimSpace(hyperlink), "\""))
						hyperlink = ""
						inHyperlink = false
					}
				}
			}

			if chex == "" {
				res.WriteString(decodeChars(hexToStr, decoder))
				hexToStr = []byte{}
			}

			switch {
			case tchar != "":
				if curskip > 0 {
					curskip--
				} else if !ignorable {
					if decoder == nil {
						res.WriteString(tchar)
					} else {
						if tcharDec, err := decoder.String(tchar); err == nil {
							res.WriteString(tcharDec)
						}
					}
				}
			case brace != "":
				curskip = 0
				if brace == "{" {
					stack = append(stack, newStackEntry(ucskip, ignorable))
				} else if brace == "}" {
					if l := len(stack); l > 0 {
						entry := stack[l-1]
						stack = stack[:l-1]
						ucskip = entry.numberOfCharactersToSkip
						ignorable = entry.ignorable
					}
				}
			case character != "":
				curskip = 0
				if character == "~" {
					if !ignorable {
						res.WriteString("\xA0")
					}
				} else if strings.Contains("{}\\", character) {
					if !ignorable {
						res.WriteString(character)
					}
				} else if character == "*" {
					ignorable = true
				}
			case word != "":
				curskip = 0
				if _, ok := destinations[word]; html && ok {
					ignorable = true
				} else if destinations[word] {
					ignorable = true
				} else if ignorable {
				} else if specialCharacters[word] != "" {
					res.WriteString(specialCharacters[word])
				} else if word == "ansicpg" {
					if charMap, ok := charmaps[arg]; ok {
						decoder = charMap.NewDecoder()
					} else if encMap, ok := encodings[arg]; ok {
						decoder = encMap.NewDecoder()
					}
				} else if word == "uc" {
					i, _ := strconv.Atoi(arg)
					ucskip = i
				} else if word == "u" {
					c, _ := strconv.Atoi(arg)
					if c < 0 {
						c += 0x10000
					}
					res.WriteRune(rune(c))
					curskip = ucskip
				}
			case chex != "":
				if curskip > 0 {
					curskip--
				} else if !ignorable {
					if c, err := hex.DecodeString(chex); err == nil {
						hexToStr = append(hexToStr, c...)
					}
				}
			}
		}

		if err == io.EOF {
			break
		}
	}

	return strings.NewReader(res.String()), nil
}

func newStackEntry(numberOfCharactersToSkip int, ignorable bool) stackEntry {
	return stackEntry{
		numberOfCharactersToSkip: numberOfCharactersToSkip,
		ignorable:                ignorable,
	}
}

func decodeChars(chars []byte, decoder *encoding.Decoder) string {
	if len(chars) == 0 {
		return ""
	}

	if decoder != nil {
		if dc, err := decoder.Bytes(chars); err == nil {
			return string(dc)
		}
	}

	var res strings.Builder
	for _, h := range chars {
		res.WriteRune(rune(h))
	}

	return res.String()
}
