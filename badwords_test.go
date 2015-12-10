package couponcode

import (
	"testing"
)

func TestBadWordsNegative(t *testing.T) {
	if hasBadWord("LOVE") {
		t.Error("LOVE shouldn't be a bad word")
	}
}

func TestBadWordsPositive(t *testing.T) {
	if !hasBadWord("BOOBIES") {
		t.Error("BOOB should be a bad word")
	}
}
