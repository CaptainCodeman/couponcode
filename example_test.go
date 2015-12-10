package couponcode_test

import (
	"fmt"
	"math/rand"

	"github.com/captaincodeman/couponcode"
)

func ExampleGenerate() {
	rand.Seed(42) // to force consistent values for example
	code := couponcode.Generate()
	fmt.Println(code)
	// Output: RCMD-CRVF-FK36
}

func ExampleGenerateCustom() {
	rand.Seed(42) // to force consistent values for example
	cc := couponcode.New(4, 6)
	code := cc.Generate()
	fmt.Println(code)
	// Output: RCMCRH-VFK3TD-V4D182-U0VGHE
}
