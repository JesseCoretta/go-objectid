package objectid

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"unicode"
)

var (
	printf     func(string, ...any) (int, error)  = fmt.Printf
	sprintf    func(string, ...any) string        = fmt.Sprintf
	atoi       func(string) (int, error)          = strconv.Atoi
	contains   func(string, string) bool          = strings.Contains
	eq         func(string, string) bool          = strings.EqualFold
	fields     func(string) []string              = strings.Fields
	hasPrefix  func(string, string) bool          = strings.HasPrefix
	hasSuffix  func(string, string) bool          = strings.HasSuffix
	indexRune  func(string, rune) int             = strings.IndexRune
	join       func([]string, string) string      = strings.Join
	split      func(string, string) []string      = strings.Split
	splitAfter func(string, string) []string      = strings.SplitAfter
	splitN     func(string, string, int) []string = strings.SplitN
	trimS      func(string) string                = strings.TrimSpace
	trimL      func(string, string) string        = strings.TrimLeft
	trimR      func(string, string) string        = strings.TrimRight
	isDigit    func(rune) bool                    = unicode.IsDigit
	isLetter   func(rune) bool                    = unicode.IsLetter
	isLower    func(rune) bool                    = unicode.IsLower
	isUpper    func(rune) bool                    = unicode.IsUpper
)

func errorf(msg any, x ...any) error {
	switch tv := msg.(type) {
	case string:
		return errors.New(sprintf(tv, x...))
	case error:
		return errors.New(sprintf(tv.Error(), x...))
	}

	return nil
}

func strInSlice(str string, slice []string) bool {
	if len(str) == 0 || len(slice) == 0 {
		return false
	}

	for _, val := range slice {
		if eq(val, str) {
			return true
		}
	}

	return false
}

/*
is 'val' an unsigned number?
*/
func isNumber(val string) bool {
	if len(val) == 0 {
		return false
	}

	for i := 0; i < len(val); i++ {
		if !isDigit(rune(val[i])) {
			return false
		}
	}
	return true
}

/*
compare slice members of two (2) []int instances.
*/
func intSliceEqual(s1, s2 []int) (equal bool) {
	if len(s1)|len(s2) == 0 || len(s1) != len(s2) {
		return
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return
		}
	}

	equal = true
	return
}

/*
compare slice members of two (2) []string instances.
*/
func strSliceEqual(s1, s2 []string) (equal bool) {
	if len(s1)|len(s2) == 0 || len(s1) != len(s2) {
		return
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return
		}
	}

	equal = true
	return
}

/*
condenseWHSP returns input string b with all contiguous
WHSP characters condensed into single space characters.
*/
func condenseWHSP(b string) (a string) {
	// remove leading and trailing
	// WHSP characters ...
	b = trimS(b)

	var last bool
	for i := 0; i < len(b); i++ {
		c := rune(b[i])
		switch c {
		case ' ':
			if !last {
				last = true
				a += string(c)
			}
		default:
			if last {
				last = false
			}
			a += string(c)
		}
	}

	return
}

/*
sscan is a closure function used by the NumberForm stringer
for interoperability with fmt.Sscan.
*/
func sscan(s fmt.ScanState, ch rune, u *NumberForm) error {
	i := new(big.Int)
	if err := i.Scan(s, ch); err != nil {
		return err
	} else if i.Sign() < 0 {
		return errorf("value cannot be negative")
	} else if i.BitLen() > 128 {
		return errorf("value overflows uint128")
	}
	u.lo = i.Uint64()
	u.hi = i.Rsh(i, 64).Uint64()
	return nil
}
