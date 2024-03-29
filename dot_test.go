package objectid

import (
	"fmt"
	"testing"
)

func ExampleDotNotation_Index() {
	dot, err := NewDotNotation(`1.3.6.1.4.1.56521.999.5`)
	if err != nil {
		fmt.Println(err)
		return
	}

	arc, _ := dot.Index(1)
	fmt.Printf("%s", arc)
	// Output: 3
}

func ExampleDotNotation_IsZero() {
	var dot DotNotation
	fmt.Printf("Is Zero: %t", dot.IsZero())
	// Output: Is Zero: true
}

func ExampleDotNotation_Valid() {
	var dot DotNotation
	fmt.Printf("Is Valid: %t", dot.Valid())
	// Output: Is Valid: false
}

func ExampleDotNotation_Len() {
	dot, err := NewDotNotation(`1.3.6.1.4.1.56521.999.5`)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Length: %d", dot.Len())
	// Output: Length: 9
}

func ExampleDotNotation_String() {
	dot, err := NewDotNotation(`1.3.6.1.4.1.56521.999.5`)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s", dot)
	// Output: 1.3.6.1.4.1.56521.999.5
}

func TestDotNotation_badInit(t *testing.T) {
	var d DotNotation
	want := false
	got := d.Valid()
	if want != got {
		t.Errorf("%s failed: wanted validity of %t, got %t",
			t.Name(), want, got)
		return
	}
}

func TestDotNotation_Ancestry(t *testing.T) {
	dot, err := NewDotNotation(`1.3.6.1.4.1.56521.999.5`)
	if err != nil {
		t.Errorf("%s failed: %s",
			t.Name(), err.Error())
		return
	}
	anc := dot.Ancestry()

	want := 9
	got := len(anc)

	if want != got {
		t.Errorf("%s failed: wanted length of %d, got %d",
			t.Name(), want, got)
		return
	}
}

func TestDotNotation_NewSubordinate(t *testing.T) {
	dot, _ := NewDotNotation(`1.3.6.1.4.1.56521.999.5`)
	leaf := dot.NewSubordinate(`10001`)

	want := `1.3.6.1.4.1.56521.999.5.10001`
	got := leaf.String()

	if want != got {
		t.Errorf("%s failed: wanted %s, got %s",
			t.Name(), want, got)
		return
	}

	if !dot.Valid() {
		t.Errorf("%s failed %T validity checks", t.Name(), dot)
		return
	}
}

func TestDotNotation_IsZero(t *testing.T) {
	var dot DotNotation
	if !dot.IsZero() {
		t.Errorf("%s failed: bogus IsZero return",
			t.Name())
		return
	}
}

func TestDotNotation_AncestorOf(t *testing.T) {
	dot, _ := NewDotNotation(`1.3.6`)
	chstr := `1.3.6.1.4`
	child, err := NewDotNotation(chstr)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	for _, d := range []any{
		chstr,
		child,
		*child,
	} {
		if !dot.AncestorOf(d) {
			t.Errorf("%s failed: ancestry check returned bogus result",
				t.Name())
			return
		}
	}
}

func TestDotNotation_codecov(t *testing.T) {
	_, err := NewDotNotation(``)
	if err == nil {
		t.Errorf("%s failed: zero length OID parsed without error", t.Name())
		return
	}

	var X DotNotation
	_, _ = X.IntSlice()
}

func TestDotNotation_Index(t *testing.T) {
	dot, err := NewDotNotation(`1.3.6.1.4.1.56521.999.5`)
	if err != nil {
		fmt.Println(err)
		return
	}

	nf, _ := dot.Index(1)
	if nf.String() != `3` {
		t.Errorf("%s failed: unable to call index 1 from %T", t.Name(), nf)
		return
	}

	nf, _ = dot.Index(-1)
	if nf.String() != `5` {
		t.Errorf("%s failed: unable to call index -1 from %T", t.Name(), nf)
		return
	}

	nf, _ = dot.Index(100)
	if nf.String() != `5` {
		t.Errorf("%s failed: unable to call index 100 from %T", t.Name(), nf)
		return
	}
}
