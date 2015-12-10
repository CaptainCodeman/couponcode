# CouponCode for Go

An implementation of Perl's [Algorithm::CouponCode][couponcode] for Golang.

# Synopsis #

This package helps generate and validate coupon codes which could be used for ecommerce coupons. It avoids creating
coupons containing certain 'bad words' which may be offensive.

The package provides a default configuration for codes of 3 parts of 4 characters separated by a '-'. The default
configuration can be accessed using the `Default` package field or, for convenience, two top level functions:

`Generate() string` to generate a coupon code
`Validate(code string) (string, error)` to normalize and validate a code

```
package couponcode_test

import (
	"fmt"

	"github.com/captaincodeman/couponcode"
)

func main() {
	code := couponcode.Generate()
	fmt.Println(code)
	// Output: RCMD-CRVF-FK36
}
```

To use a custom part and part length, create a new generator using the constructor and then use the functions
provided by that instance:

```
cc := couponcode.New(4, 6)

code := cc.Generate() // note function of cc, not couponcode package

validated, err := cc.Validate(code)
```

Now, when someone types their code in, you can check that it is valid. This means that letters like `O`
are converted to `0` prior to checking. The validate command returns a normalized code (useful for any
consistent database lookups of the coupon code) and an error to indicate if the code was valid.

```
// same code, just lowercased
code, err := couponcode.Validate('55g2-dhm0-50nn');
// '55G2-DHM0-50NN'

// various letters instead of numbers
code, err := couponcode.Validate('SSGZ-DHMO-SONN');
// '55G2-DHM0-50NN'

// wrong last character
code, err := couponcode.Validate('55G2-DHM0-50NK');
// err != nil

// not enough chars in the 2nd part
code, err := couponcode.Validate('55G2-DHM-50NN');
// err != nil
```

The first thing we do to each code is uppercase it. Then we convert the following letters to numbers:

* O -> 0
* I -> 1
* Z -> 2
* S -> 5

This means [oizs], [OIZS] and [0125] are considered the same code.

# Example #

Let's say you want a user to verify they got something, whether that is an email, letter, fax or carrier pigeon. To
prove they received it, they have to type the code you sent them into a certain page on your website. You create a code
which they have to type in:

```
code := couponcode.Generate();
// 55G2-DHM0-50NN
```

Time passes, letters get wet, carrier pigeons go on adventures and faxes are just as bad as they ever were. Now the
user has to type their code into your website. The problem is, they can hardly read what the code was. Luckily we're
somewhat forgiving since Z's and 2's are considered the same, O's and 0's, I's and 1's and S's and 5's are also mapped
to each other. But even more than that, the 4th character of each group is a checkdigit which can determine if the
other three in that group are correct. The user types this:

```
[s5g2-dhmo-50nn]
```

Because our codes are case insensitive and have good conversions for similar chars, the code is accepted as correct.

Also, since we have a checkdigit, we can use a client-side plugin to highlight to the user any mistake in their code
before they submit it. Please see the original project ([Algorithm::CouponCode][couponcode]) for more details of client
side validation.

# Installation

The easiest way to get it is via `go get`:

``` bash
$ go get -u github.com/captaincodeman/couponcode
```

# Build

The default build uses the random number generator from the `math/rand` package which generates consistent random numbers
for a given seed which can be changed. This is helpful for examples and when testing and, assuming you use a random seed
it may be good enough for live use if you store codes in a database and mark them as consumed.

There is a risk though that someone could work out the seed value and use it to generate their own valid codes so there is
an implementation that uses the `crypto/rand` package for more secure random number generation. This can be used by adding
the conditional compilation build tag `coupon-crypto` to your app build command, e.g.

    go build -tags coupon-crypto

# Tests

To run the tests, use go test:

```
$ go test
```

# Author

* Written by [Simon Green](http://captaincodeman.com)

# Inspired By

[Grant McLean](grant)'s [Algorithm::CouponCode][couponcode] - with thanks. :)

# License

MIT.