package couponcode // import "github.com/captaincodeman/couponcode"

import (
	"fmt"
	"regexp"
	"strings"
)

type (
	generator struct {
		parts   int
		partLen int
	}
)

const (
	symbols   = "0123456789ABCDEFGHJKLMNPQRTUVWXY"
	length    = len(symbols) - 1
	separator = "-"
)

var (
	Default         = New(3, 4)
	removeInvalidRe = regexp.MustCompile(`[^0-9A-Z]+`)
)

func Generate() string {
	return Default.Generate()
}

func Validate(code string) (string, error) {
	return Default.Validate(code)
}

func New(parts, partLen int) *generator {
	return &generator{
		parts:   parts,
		partLen: partLen,
	}
}

func (g *generator) Generate() string {
	parts := make([]string, g.parts)
	i := 0
	for i < g.parts {
		code := randString(g.partLen - 1)
		check := checkCharacter(code, i+1)
		parts[i] = code + check
		if !hasBadWord(strings.Join(parts, "")) {
			i += 1
		}
	}
	return strings.Join(parts, separator)
}

func (g *generator) Validate(code string) (string, error) {
	// make uppercase
	code = strings.ToUpper(code)

	// remove invalid characters
	code = removeInvalidRe.ReplaceAllLiteralString(code, "")

	// convert special letters to numbers
	code = strings.Replace(code, "O", "0", -1)
	code = strings.Replace(code, "I", "1", -1)
	code = strings.Replace(code, "Z", "2", -1)
	code = strings.Replace(code, "S", "5", -1)

	// split into parts
	parts := []string{}
	tmp := code
	for len(tmp) > 0 {
		max := g.partLen
		if max > len(tmp) {
			max = len(tmp)
		}
		parts = append(parts, tmp[:max])
		tmp = tmp[max:len(tmp)]
	}

	// join with separator (shouldn't we test that)
	code = strings.Join(parts, separator)

	if len(parts) != g.parts {
		return code, fmt.Errorf("wrong number of parts (%d)", len(parts))
	}
	for i, part := range parts {
		if len(part) != g.partLen {
			return code, fmt.Errorf("wrong length of part %d", i)
		}
		check := checkCharacter(part[:len(part)-1], i+1)
		if !strings.HasSuffix(part, check) {
			return code, fmt.Errorf("wrong part %d (%s) check character %s", i+1, part, check)
		}
	}

	return code, nil
}

func checkCharacter(code string, check int) string {
	for _, r := range code {
		k := strings.IndexRune(symbols, r)
		check = check*19 + k
	}
	return string(symbols[check%int(length)])
}