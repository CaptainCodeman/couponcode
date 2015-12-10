package couponcode

import (
	"regexp"
	"strings"
	"testing"
)

func TestDefault(t *testing.T) {
	if Default.parts != 3 && Default.partLen != 4 {
		t.Error("wrong defaults")
	}
}

func TestNew(t *testing.T) {
	g := New(5, 3)
	if g.parts != 5 && g.partLen != 3 {
		t.Error("wonky constructor")
	}
}

func TestCheckDigit(t *testing.T) {
	var runs = []struct {
		code string
		part int
	}{
		{"55G2", 1},
		{"DHM0", 2},
		{"50NN", 3},
		{"U5H9", 1},
		{"HKDH", 2},
		{"8RNX", 3},
		{"1EX7", 4},
		{"WYLKQM", 1},
		{"U35V40", 2},
		{"9N84DA", 3},
	}
	for _, run := range runs {
		check := checkCharacter(run.code[:len(run.code)-1], run.part)
		if !strings.HasSuffix(run.code, check) {
			t.Errorf("check digit failed for %s got %s", run.code, check)
		}
	}
}

func TestValidCodes(t *testing.T) {
	var runs = []struct {
		g    *generator
		code string
	}{
		{Default, "55G2-DHM0-50NN"},
		{New(4, 4), "U5H9-HKDH-8RNX-1EX7"},
		{New(3, 6), "WYLKQM-U35V40-9N84DA"},
		{Default, "55g2-dhm0-50nn"},
		{Default, "SSGZ-DHMO-SONN"},
		{New(7, 12), "QBXA5CV4Q85E-HNYV4U3UD69M-B7XU1BHF3FYE-HXT9LD4Q0DAH-U6WMKC1WNF4N-5PCG5C4JF0GL-5DTUNJ40LRB5"},
		{New(1, 4), "1K7Q"},
		{New(2, 4), "1K7Q-CTFM"},
		{New(3, 4), "1K7Q-CTFM-LMTC"},
		{New(4, 4), "7YQH-1FU7-E1HX-0BG9"},
		{New(5, 4), "YENH-UPJK-PTE0-20U6-QYME"},
		{New(6, 4), "YENH-UPJK-PTE0-20U6-QYME-RBK1"},
	}
	for _, run := range runs {
		validated, err := run.g.Validate(run.code)
		if err != nil {
			t.Errorf("code %s should be valid got %s %s", run.code, validated, err)
		}
	}
}

func TestInvalidCodes(t *testing.T) {
	var runs = []struct {
		g    *generator
		code string
	}{
		{Default, "55G2-DHM0-50NK"}, // wrong check
		{Default, "55G2-DHM-50NN"},  // not enough characters
		{New(3, 4), "1K7Q-CTFM"},    // too short
		{New(1, 4), "1K7C"},
		{New(2, 4), "1K7Q-CTFW"},
		{New(3, 4), "1K7Q-CTFM-LMT1"},
		{New(4, 4), "7YQH-1FU7-E1HX-0BGP"},
		{New(5, 4), "YENH-UPJK-PTE0-20U6-QYMT"},
		{New(6, 4), "YENH-UPJK-PTE0-20U6-QYME-RBK2"},
	}
	for _, run := range runs {
		validated, err := run.g.Validate(run.code)
		if err == nil {
			t.Errorf("code %s should be invalid got %s", run.code, validated)
		}
	}
}

func TestCodesNormalized(t *testing.T) {
	var runs = []struct {
		g    *generator
		code string
		exp  string
	}{
		{Default, "i9oD/V467/8Dsz", "190D-V467-8D52"},   // alternate separator
		{Default, " i9oD V467 8Dsz ", "190D-V467-8D52"}, // whitespace accepted
		{Default, " i9oD_V467_8Dsz ", "190D-V467-8D52"}, // underscores accepted
		{Default, "i9oDV4678Dsz", "190D-V467-8D52"},     // no separator required
	}
	for _, run := range runs {
		validated, err := run.g.Validate(run.code)
		if err != nil || validated != run.exp {
			t.Errorf("code %s should be %s got %s %s", run.code, run.exp, validated, err)
		}
	}
}

func TestPattern(t *testing.T) {
	code := Generate()
	matched, _ := regexp.MatchString(`^[0-9A-Z-]+$`, code)
	if !matched {
		t.Error("should only contain uppercase letters, digits, and dashes")
	}
	matched, _ = regexp.MatchString(`^\w{4}-\w{4}-\w{4}$`, code)
	if !matched {
		t.Error("should look like XXXX-XXXX-XXXX")
	}
	g := New(2, 5)
	code = g.Generate()
	matched, _ = regexp.MatchString(`^\w{5}-\w{5}$`, code)
	if !matched {
		t.Error("should generate an arbitrary number of parts")
	}
}

func TestDefaultSelfContained(t *testing.T) {
	for i := 0; i < 10; i++ {
		code := Generate()
		validated, err := Validate(code)
		if err != nil {
			t.Errorf("generated %s got %s %s", code, validated, err)
		}
	}
}

func TestCustomSelfContained(t *testing.T) {
	g := New(4, 6)
	for i := 0; i < 10; i++ {
		code := g.Generate()
		validated, err := g.Validate(code)
		if err != nil {
			t.Errorf("generated %s got %s %s", code, validated, err)
		}
	}
}
