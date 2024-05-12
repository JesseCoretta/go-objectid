package objectid

/*
nf.go provides NumberForm methods and types.
*/

import "math/big"

var nilNF NumberForm

/*
NumberForm is an unbounded, unsigned number.
*/
type NumberForm big.Int

/*
IsZero returns a Boolean value indicative of whether the
receiver instance is nil, or unset.
*/
func (a *NumberForm) IsZero() (is bool) {
	return len(a.cast().Bytes()) == 0
}

func (a NumberForm) cast() *big.Int {
	x := big.Int(a)
	return &x
}

/*
Equal returns a boolean value indicative of whether the receiver is equal to
the value provided.

Valid input types are string, uint64, int, uint, *[math/big.Int] and [NumberForm].

Any input that represents a negative or unspecified number guarantees a false return.
*/
func (a NumberForm) Equal(n any) (is bool) {
	switch tv := n.(type) {
	case *big.Int:
		is = a.cast().Cmp(tv) == 0
	case NumberForm:
		is = a.cast().Cmp(tv.cast()) == 0
	case string:
		nf, ok := big.NewInt(0).SetString(tv, 10)
		if !ok {
			is = ok
			break
		}
		is = a.cast().Cmp(nf) == 0
	case uint64:
		is = a.cast().Uint64() == tv
	case uint:
		is = a.cast().Uint64() == uint64(tv)
	case int:
		if 0 <= tv {
			is = a.cast().Uint64() == uint64(tv)
		}
	}

	return
}

/*
Gt returns a boolean value indicative of whether the receiver is greater than
the value provided.

Valid input types are string, uint64, int, uint, *[math/big.Int] and [NumberForm].

Any input that represents a negative or unspecified number guarantees a false return.
*/
func (a NumberForm) Gt(n any) (is bool) {
	switch tv := n.(type) {
	case *big.Int:
		is = a.cast().Cmp(tv) == 1
	case NumberForm:
		is = a.cast().Cmp(tv.cast()) == 1
	case string:
		nf, ok := big.NewInt(0).SetString(tv, 10)
		if !ok {
			is = ok
			break
		}
		is = a.cast().Cmp(nf) == 1
	case uint64:
		is = a.cast().Uint64() > tv
	case uint:
		is = a.cast().Uint64() > uint64(tv)
	case int:
		if 0 <= tv {
			is = a.cast().Uint64() > uint64(tv)
		}
	}
	return
}

/*
Ge returns a boolean value indicative of whether the receiver is greater than
or equal to the value provided.

This method is merely a convenient wrapper to an ORed call of the [NumberForm.Gt]
and [NumberForm.Equal] methods.

Valid input types are string, uint64, int, uint, *[math/big.Int] and [NumberForm].

Any input that represents a negative or unspecified number guarantees a false return.
*/
func (a NumberForm) Ge(n any) (is bool) {
	return a.Gt(n) || a.Equal(n)
}

/*
Lt returns a boolean value indicative of whether the receiver is less than
the value provided.

Valid input types are string, uint64, int, uint, *[math/big.Int] and [NumberForm].

Any input that represents a negative or unspecified number guarantees a false return.
*/
func (a NumberForm) Lt(n any) (is bool) {
	switch tv := n.(type) {
	case *big.Int:
		is = a.cast().Cmp(tv) == -1
	case NumberForm:
		is = a.cast().Cmp(tv.cast()) == -1
	case string:
		nf, ok := big.NewInt(0).SetString(tv, 10)
		if !ok {
			is = ok
			break
		}
		is = a.cast().Cmp(nf) == -1
	case uint64:
		is = a.cast().Uint64() < tv
	case uint:
		is = a.cast().Uint64() < uint64(tv)
	case int:
		if 0 <= tv {
			is = a.cast().Uint64() < uint64(tv)
		}
	}
	return
}

/*
Le returns a boolean value indicative of whether the receiver is less than or
equal to the value provided.

This method is merely a convenient wrapper to an ORed call of the [NumberForm.Lt]
and [NumberForm.Equal] methods.

Valid input types are string, uint64, int, uint, *[math/big.Int] and [NumberForm].

Any input that represents a negative or unspecified number guarantees a false return.
*/
func (a NumberForm) Le(n any) (is bool) {
	return a.Lt(n) || a.Equal(n)
}

/*
Valid returns a Boolean value indicative of proper initialization.
*/
func (a NumberForm) Valid() bool {
	return &a != nil
}

/*
String returns the base-10 string representation of the receiver
instance.
*/
func (a NumberForm) String() string {
	return a.cast().String()
}

func newStringNF(tv string) (nf *big.Int, err error) {
	if len(tv) == 0 {
		err = errorf("Zero length NumberForm %T", tv)
		return
	} else if tv[0] == '-' {
		err = errorf("A NumberForm cannot be negative")
		return
	}

	var ok bool
	if nf, ok = big.NewInt(0).SetString(tv, 10); !ok {
		err = errorf("Failed to read '%s' into NumberForm", tv)
	}

	return
}

/*
NewNumberForm converts v into an instance of [NumberForm], which is
returned alongside an error.

Valid input types are string, uint64, int, uint, and *[math/big.Int].

Any input that represents a negative or unspecified number guarantees an error.
*/
func NewNumberForm(v any) (a NumberForm, err error) {
	switch tv := v.(type) {
	case *big.Int:
		a = NumberForm(*tv)
	case string:
		var _a *big.Int
		if _a, err = newStringNF(tv); err == nil {
			a = NumberForm(*_a)
		}
	case int:
		if tv < 0 {
			err = errorf("A NumberForm cannot be negative")
			break
		}

		_a := big.NewInt(int64(tv))
		a = NumberForm(*_a)
	case uint64:
		_a := big.NewInt(0).SetUint64(tv)
		a = NumberForm(*_a)
	case uint:
		_a := big.NewInt(0).SetUint64(uint64(tv))
		a = NumberForm(*_a)
	default:
		err = errorf("Unsupported %T type '%T'", a, tv)
	}

	return
}
