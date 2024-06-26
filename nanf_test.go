package objectid

import (
	"fmt"
	"math/big"
	"testing"
)

func ExampleNewNameAndNumberForm() {
	nanf, err := NewNameAndNumberForm(`enterprise(1)`)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", nanf)
	// Output: enterprise(1)
}

func ExampleNameAndNumberForm_String() {
	nanf, err := NewNameAndNumberForm(`enterprise(1)`)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", nanf)
	// Output: enterprise(1)
}

func ExampleNameAndNumberForm_IsZero() {
	nanf, err := NewNameAndNumberForm(`enterprise(1)`)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Zero: %t", nanf.IsZero())
	// Output: Zero: false
}

func ExampleNameAndNumberForm_Identifier() {
	nanf, err := NewNameAndNumberForm(`enterprise(1)`)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", nanf.Identifier())
	// Output: enterprise
}

func ExampleNameAndNumberForm_NumberForm() {
	nanf, err := NewNameAndNumberForm(`enterprise(1)`)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", nanf.NumberForm())
	// Output: 1
}

func ExampleNameAndNumberForm_Equal() {
	var nanf1, nanf2 *NameAndNumberForm
	var err error

	if nanf1, err = NewNameAndNumberForm(`enterprise(1)`); err != nil {
		fmt.Println(err)
		return
	}

	// bogus
	if nanf2, err = NewNameAndNumberForm(`enterprise(10)`); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Equal: %t", nanf1.Equal(nanf2))
	// Output: Equal: false
}

func TestNewNameAndNumberForm(t *testing.T) {
	var err error
	var nf NumberForm
	if nf, err = NewNumberForm(177); err != nil {
		t.Errorf("%s failed: %s", t.Name(), err.Error())
		return
	}

	for idx, v := range []any{
		`enterprise(1)`,
		`enterprise1)`,
		77,
		``,
		nf,
		`blarg`,
		big.NewInt(42),
		big.NewInt(-42),
		uint64(3),
		-4,
		uint(0),
		new(big.Int),
		`thing(432897659847395789374568903476893476937468934769843)`,
	} {
		_, err = NewNameAndNumberForm(v)
		if idx%2 == 0 && err != nil {
			t.Errorf("%s[%d] failed: %v", t.Name(), idx, err)
			return
		} else if err == nil && idx%2 != 0 {
			t.Errorf("%s[%d] failed: parsed bogus value without error", t.Name(), idx)
			return
		}
	}

	_, _ = NewNameAndNumberForm(nil)
}

func TestBogusNameAndNumberForm(t *testing.T) {
	if _, err := NewNameAndNumberForm("Enterprise(1)"); err == nil {
		t.Errorf("%s failed: parsed bogus string value without error", t.Name())
	}
}

func TestParseNaNFStr(t *testing.T) {
	for idx, slice := range []string{
		`cn(3)`,
		`identifier(-3)`,
	} {
		_, err := parseNaNFstr(slice)
		if err != nil {
			if idx%2 == 0 {
				t.Errorf("%s failed: unexpected error: %v", t.Name(), err)
			}
			continue
		} else {
			if idx%2 != 0 {
				t.Errorf("%s failed: expected error, got nothing", t.Name())
				continue
			}
		}
	}
}
