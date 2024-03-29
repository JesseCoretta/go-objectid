package objectid

import (
	"fmt"
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
	} {
		_, err = NewNameAndNumberForm(v)
		if idx%2 == 0 && err != nil {
			t.Errorf("%s failed: %v", t.Name(), err)
			return
		} else if err == nil && idx%2 != 0 {
			t.Errorf("%s failed: parsed bogus value without error", t.Name())
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
