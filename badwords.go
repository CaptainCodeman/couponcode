package couponcode

import (
	"strings"
)

var (
	// ROT13 ... because it doesn't need to be enigma
	badwords = makeBadWords("SHPX PHAG JNAX JNAT CVFF PBPX FUVG GJNG GVGF SNEG URYY ZHSS QVPX XABO NEFR FUNT GBFF FYHG GHEQ FYNT PENC CBBC OHGG SRPX OBBO WVFZ WVMM CUNG")
)

func hasBadWord(code string) bool {
	for _, badword := range badwords {
		if strings.Contains(code, badword) {
			return true
		}
	}
	return false
}

func makeBadWords(badwords string) []string {
	words := make([]rune, len(badwords))
	for i, x := range badwords {
		c := mapRune(x)
		words[i] = c
	}
	return strings.Split(string(words), " ")
}

func mapRune(r rune) rune {
	switch {
	case 65 <= r && r <= 77:
		return r + 13
	case 78 <= r && r <= 90:
		return r - 13
	default:
		return r
	}
}
