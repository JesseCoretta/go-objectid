package objectid

import (
	"testing"
)

func TestDotNotation_Ancestry(t *testing.T) {
	dot, err := NewDotNotation(`1.3.6.1.4.1.56521.999.5`)
	if err != nil {
		t.Errorf("%s failed: %s",
			t.Name(), err.Error())
	}
	anc := dot.Ancestry()

	want := 9
	got := len(anc)

	if want != got {
		t.Errorf("%s failed: wanted length of %d, got %d",
			t.Name(), want, got)
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
	}
}

func TestDotNotation_IsZero(t *testing.T) {
	var dot DotNotation
	if !dot.IsZero() {
		t.Errorf("%s failed: bogus IsZero return",
			t.Name())
	}
}

func TestDotNotation_AncestorOf(t *testing.T) {
	dot, _ := NewDotNotation(`1.3.6`)
	child, err := NewDotNotation(`1.3.6.1.4`)
	if err != nil {
		t.Errorf(err.Error())
	}
	if !dot.AncestorOf(child) {
		t.Errorf("%s failed: ancestry check returned bogus result",
			t.Name())
	}
}
