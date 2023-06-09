package objectid

/*
nf.go provides NumberForm methods and types.

NOTE: uint128-related code written by Luke Champine, per https://github.com/lukechampine/uint128.
It has been incorporated into this package (unexported), and simplified to serve in the capacity
of OID numberForm storage.
*/

import (
	"fmt"
	"math/bits"
)

var nilNF NumberForm

/*
NumberForm is an unsigned 128-bit number. This type is based on
github.com/lukechampine/uint128. It has been incorporated
into this package to produce unsigned 128-bit OID numberForm
support (i.e.: UUID-based OIDs).
*/
type NumberForm struct {
	lo, hi uint64
	parsed bool
}

// isZero returns true if a == 0.
func (a NumberForm) IsZero() bool {
	if &a == nil {
		return true
	}

	// NOTE: we do not compare against Zero, because that
	// is a global variable that could be modified.
	return a.lo == uint64(0) && a.hi == uint64(0)
}

/*
Equal returns a boolean value indicative of whether the receiver is equal to
the value provided. Valid input types are string, uint64, int and NumberForm.

Any input that represents a negative number guarantees a false return.
*/
func (a NumberForm) Equal(n any) bool {
	switch tv := n.(type) {
	case NumberForm:
		return a == tv
	case string:
		if nf, err := NewNumberForm(tv); err == nil {
			return a == nf
		}
	case uint64:
		return a.lo == tv && a.hi == 0
	case int:
		if 0 <= tv {
			return a.lo == uint64(tv) && a.hi == 0
		}
	}

	return false
}

/*
Gt returns a boolean value indicative of whether the receiver is greater than
the value provided. Valid input types are string, uint64, int and NumberForm.

Any input that represents a negative number guarantees a false return.
*/
func (a NumberForm) Gt(n any) bool {
	switch tv := n.(type) {
	case NumberForm:
		return a.hi > tv.hi || (a.hi == tv.hi && a.lo > tv.lo)
	case string:
		if nf, err := NewNumberForm(tv); err == nil {
			return a.hi > nf.hi || (a.hi == nf.hi && a.lo > nf.lo)
		}
	case uint64:
		return a.lo > tv && a.hi == uint64(0)
	case int:
		if 0 <= tv {
			return a.lo > uint64(tv) && a.hi == uint64(0)
		}
	}
	return false
}

/*
Lt returns a boolean value indicative of whether the receiver is less than
the value provided. Valid input types are string, uint64, int and NumberForm.

Any input that represents a negative number guarantees a false return.
*/
func (a NumberForm) Lt(n any) bool {
	switch tv := n.(type) {
	case NumberForm:
		return a.hi < tv.hi || (a.hi == tv.hi && a.lo < tv.lo)
	case string:
		if nf, err := NewNumberForm(tv); err == nil {
			return a.hi < nf.hi || (a.hi == nf.hi && a.lo < nf.lo)
		} else {
			printf("%s", err.Error())
		}
	case uint64:
		return a.lo < tv && a.hi == uint64(0)
	case int:
		if 0 <= tv {
			return a.lo < uint64(tv) && a.hi == uint64(0)
		}
	}
	return false
}

/*
Valid returns a boolean valud indicative of proper instantiation.
*/
func (a NumberForm) Valid() bool {
	return a.parsed
}

/*
leadingZeros returns the number of leading zero bits in u;
the result is 128 for a == 0.
*/
func (a NumberForm) leadingZeros() int {
	if a.hi > 0 {
		return bits.LeadingZeros64(a.hi)
	}
	return 64 + bits.LeadingZeros64(a.lo)
}

/*
Len returns the minimum number of bits required to represent u;
the result is 0 for a == 0.
*/
func (a NumberForm) len() int {
	return 128 - a.leadingZeros()
}

// String returns the base-10 representation of a as a string.
func (a NumberForm) String() string {
	if a.IsZero() {
		return "0"
	} else if !a.parsed {
		return "0"
	}

	buf := []byte("0000000000000000000000000000000000000000") // log10(2^128) < 40
	for i := len(buf); ; i -= 19 {
		q, r := a.quoRem64(1e19) // largest power of 10 that fits in a uint64
		var n int
		for ; r != 0; r /= 10 {
			n++
			buf[i-n] += byte(r % 10)
		}
		if q.IsZero() {
			return string(buf[i-n:])
		}
		a = q
	}
}

/*
Scan implements fmt.Scanner, and is only present to allow conversion
of an NumberForm into a string value per fmt.Sscan.  Users need not
execute this method directly.
*/
func (a *NumberForm) Scan(s fmt.ScanState, ch rune) error {
	return sscan(s, ch, a)
}

// quoRem64 returns q = u/v and r = u%v.
func (a NumberForm) quoRem64(v uint64) (q NumberForm, r uint64) {
	if a.hi < v {
		q.lo, r = bits.Div64(a.hi, a.lo, v)
	} else {
		q.hi, r = bits.Div64(0, a.hi, v)
		q.lo, r = bits.Div64(r, a.lo, v)
	}
	return
}

// NewNumberForm returns the NumberForm value.
func newNumberForm(lo, hi uint64) NumberForm {
	return NumberForm{lo: lo, hi: hi, parsed: true}
}

/*
NewNumberForm converts v into an instance of NumberForm, which
is returned alongside an error.

Acceptable input types are string, int and uint64. No decimal value,
whether string or int, can ever be negative.
*/
func NewNumberForm(v any) (a NumberForm, err error) {
	switch tv := v.(type) {
	case string:
		_, err = fmt.Sscan(tv, &a)
		a.parsed = err == nil
	case int:
		if tv < 0 {
			err = errorf("A NumberForm cannot be negative")
			break
		}
		a = newNumberForm(uint64(tv), 0)
	case uint64:
		a = newNumberForm(tv, 0)
	}

	return
}
