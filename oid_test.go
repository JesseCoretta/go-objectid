package objectid

import (
	"fmt"
	"math/big"
	"testing"
)

func TestNewOID(t *testing.T) {
	for idx, typ := range []any{
		[]string{
			`iso(1)`,
			`identified-organization(3)`,
			`dod(6)`,
			`internet(1)`,
			`private(4)`,
			`enterprise(1)`,
			`56521`,
			`example(999)`,
		},
		nil,
		`{iso(1) identified-organization(3) dod(6) internet(1) private(4) enterprise(1) 56521 example(999)}`,
		float32(1),
	} {
		_, err := NewOID(typ)
		if err != nil && idx%2 == 0 {
			t.Errorf("%s failed: %v", t.Name(), err)
			return
		} else if err == nil && idx%2 != 0 {
			t.Errorf("%s failed: no error where one was expected", t.Name())
			return
		}
	}
}

func ExampleOID_Dot() {
	raw := `{iso(1) identified-organization(3) dod(6) internet(1) private(4) enterprise(1) 56521 example(999)}`
	id, err := NewOID(raw)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", id.Dot())
	// Output: 1.3.6.1.4.1.56521.999
}

/*
This example demonstrates a bogus [DotNotation] output due to the presence of
less than two (2) [NameAndNumberForm] instances within the receiver.

[DotNotation] ALWAYS requires two (2) or more elements to be considered valid.
*/
func ExampleOID_Dot_bogus() {
	id, err := NewOID(`{iso(1)}`)
	if err != nil {
		fmt.Println(err)
		return
	}
	dot := id.Dot()
	fmt.Println(dot)
	// Output:
}

func ExampleOID_Len() {
	raw := `{iso(1) identified-organization(3) dod(6) internet(1) private(4) enterprise(1) 56521 example(999)}`
	id, err := NewOID(raw)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%d", id.Len())
	// Output: 8
}

func ExampleOID_ASN() {
	raw := `{iso(1) identified-organization(3) dod(6) internet(1) private(4) enterprise(1) 56521 example(999)}`
	id, err := NewOID(raw)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", id.ASN())
	// Output: {iso(1) identified-organization(3) dod(6) internet(1) private(4) enterprise(1) 56521 example(999)}
}

func ExampleOID_IsZero() {
	var z OID
	fmt.Printf("Zero: %t", z.IsZero())
	// Output: Zero: true
}

func ExampleOID_Valid() {
	var o OID
	fmt.Printf("Valid: %t", o.Valid())
	// Output: Valid: false
}

func ExampleOID_Leaf() {
	a := `{joint-iso-itu-t(2) uuid(25) ans(987895962269883002155146617097157934)}`
	id, err := NewOID(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Leaf node: %s", id.Leaf())
	// Output: Leaf node: ans(987895962269883002155146617097157934)
}

func ExampleOID_Parent() {
	a := `{joint-iso-itu-t(2) uuid(25) ans(987895962269883002155146617097157934)}`
	id, err := NewOID(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Leaf node parent: %s", id.Parent())
	// Output: Leaf node parent: uuid(25)
}

func ExampleOID_Root() {
	a := `{joint-iso-itu-t(2) uuid(25) ans(987895962269883002155146617097157934)}`
	id, err := NewOID(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Root node: %s", id.Root())
	// Output: Root node: joint-iso-itu-t(2)
}

func TestOID_bogus(t *testing.T) {
	if _, err := NewOID(testASN1Bogus); err == nil {
		t.Errorf("%s successfully parsed bogus value; expected an error", t.Name())
		return
	}

	if _, err := NewOID(`iso(3) identified-organization(3)`); err == nil {
		t.Errorf("%s successfully parsed bogus value; expected an error", t.Name())
		return
	}

	if _, err := NewOID(`itu-t recommendation(-3)`); err == nil {
		t.Errorf("%s successfully parsed bogus value; expected an error", t.Name())
		return
	}

	if _, err := NewOID(`joint-iso-itu-t thing`); err == nil {
		t.Errorf("%s successfully parsed bogus value; expected an error", t.Name())
		return
	}

	if _, err := NewOID([]NameAndNumberForm{
		{identifier: `iso`, primaryIdentifier: NumberForm(*big.NewInt(1)), parsed: true},
		{identifier: `identified-organization`, primaryIdentifier: NumberForm(*big.NewInt(3)), parsed: true},
	}); err != nil {
		t.Errorf("%s error: %v", t.Name(), err)
		return
	}

	if _, err := NewOID([]NameAndNumberForm{}); err == nil {
		t.Errorf("%s error: %v", t.Name(), err)
		return
	}
}
