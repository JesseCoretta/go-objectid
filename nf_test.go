package objectid

import (
	"fmt"
	"testing"
)

func TestNewNumberForm(t *testing.T) {
	// even #s = valid
	// odd #s  = invalid
	for idx, num := range []any{
		`3849141823758536772162786183725055278`,
		-103,
		`9399368356398566872162777255735125541`,
		`939936835639856687216277725573512554138275978532897358923759872389572389572893758923758923758923759823`, // >uint128
		`0`,
	} {
		nf, err := NewNumberForm(num)
		ok := err == nil
		if !ok && idx%2 == 0 {
			t.Errorf("%s failed: valid number not parsed: %v", t.Name(), err)
			return
		} else if ok && idx%2 != 0 {
			t.Errorf("%s failed: bogus number parsed without error", t.Name())
			return
		}

		_ = nf.len()
		_ = nf.String()
		_, _ = nf.quoRem64(uint64(idx+1)) // codecov
	}

	var nf NumberForm
	_ = nf.gtLt(nil, true)  // codecov
	_ = nf.gtLt(nil, false) // codecov
}

func TestQuorem64(t *testing.T) {
	for _, z := range []int{
		0,
		1,
		225,
		18457,
		9,
	} {
		x, _ := NewNumberForm(z)
		for _, y := range []uint64{
			// don't send zero
			uint64(13),
			uint64(17284),
			^uint64(0),
			uint64(1),
			uint64(14895623),
			uint64(1234),
			uint64(4895623),
		} {
			_, _ = x.quoRem64(y)
		}
	}
}

func TestBogusNewNumberForm(t *testing.T) {
	bogus := `-48675`
	crap, err := NewNumberForm(bogus)
	if err == nil {
		t.Errorf("%s failed: bogus NumberForm '%v' accepted without error",
			t.Name(), bogus)
		return
	}

	var junk NumberForm

	_ = crap.String()
	_ = junk.String()
}

func TestNumberForm_Gt(t *testing.T) {
	nf, _ := NewNumberForm(4658)
	if !nf.Gt(3700) {
		t.Errorf("%s failed: Gt evaluation returned a bogus value", t.Name())
	}

	if nf, _ = NewNumberForm(`437829765`); nf.Gt(500000000) {
		t.Errorf("%s failed: Gt evaluation returned a bogus value", t.Name())
	}

	if nf, _ = NewNumberForm(uint64(829765)); nf.Gt(500000000) {
		t.Errorf("%s failed: Gt evaluation returned a bogus value", t.Name())
	}
}

func TestNumberForm_Ge(t *testing.T) {
	nf, _ := NewNumberForm(4658)
	if !nf.Ge(3700) {
		t.Errorf("%s failed: Ge evaluation returned a bogus value", t.Name())
	}

	if nf, _ = NewNumberForm(`437829765`); nf.Ge(500000000) {
		t.Errorf("%s failed: Ge evaluation returned a bogus value", t.Name())
	}

	if nf, _ = NewNumberForm(uint64(829765)); nf.Ge(500000000) {
		t.Errorf("%s failed: Ge evaluation returned a bogus value", t.Name())
	}
}

func TestNumberForm_Lt(t *testing.T) {
	nf, _ := NewNumberForm(4658)
	if nf.Lt(3700) {
		t.Errorf("%s failed: Gt evaluation returned a bogus value", t.Name())
	}

	if nf, _ = NewNumberForm(`437829765`); !nf.Lt(500000000) {
		t.Errorf("%s failed: Gt evaluation returned a bogus value", t.Name())
	}

	if nf, _ = NewNumberForm(uint64(329856)); !nf.Lt(500000000) {
		t.Errorf("%s failed: Gt evaluation returned a bogus value", t.Name())
	}
}

func TestNumberForm_Le(t *testing.T) {
	nf, _ := NewNumberForm(4658)
	if nf.Le(3700) {
		t.Errorf("%s failed: Le evaluation returned a bogus value", t.Name())
	}

	if nf, _ = NewNumberForm(`437829765`); !nf.Le(500000000) {
		t.Errorf("%s failed: Le evaluation returned a bogus value", t.Name())
	}

	if nf, _ = NewNumberForm(uint64(329856)); !nf.Le(500000000) {
		t.Errorf("%s failed: Le evaluation returned a bogus value", t.Name())
	}
}

func TestNumberForm_Equal(t *testing.T) {
	nf, _ := NewNumberForm(4658)
	if nf.Equal(3700) {
		t.Errorf("%s failed: Gt evaluation returned a bogus value", t.Name())
	}

	if nf, _ = NewNumberForm(`437829765`); !nf.Equal(`437829765`) {
		t.Errorf("%s failed: Gt evaluation returned a bogus value", t.Name())
	}

	if nf, _ = NewNumberForm(uint64(329856)); !nf.Equal(uint64(329856)) {
		t.Errorf("%s failed: Gt evaluation returned a bogus value", t.Name())
	}
}

func ExampleNumberForm_Equal() {
	nf1, _ := NewNumberForm(4658)
	nf2, _ := NewNumberForm(4657)
	fmt.Printf("Instances are equal: %t", nf1.Equal(nf2))
	// Output: Instances are equal: false
}

func ExampleNumberForm_Valid() {
	nf, _ := NewNumberForm(4658)
	fmt.Printf("Valid: %t", nf.Valid())
	// Output: Valid: true
}

func ExampleNumberForm_String() {
	nf, _ := NewNumberForm(4658)
	fmt.Printf("%s", nf)
	// Output: 4658
}

func ExampleNumberForm_IsZero() {
	var nf NumberForm
	fmt.Printf("Zero: %t", nf.IsZero())
	// Output: Zero: true
}

func ExampleNumberForm_Ge() {
	nf, _ := NewNumberForm(4658)
	oth, _ := NewNumberForm(4501)
	fmt.Printf("%s >= %s: %t", nf, oth, nf.Ge(oth))
	// Output: 4658 >= 4501: true
}

func ExampleNumberForm_Gt() {
	nf, _ := NewNumberForm(4658)
	oth, _ := NewNumberForm(4501)
	fmt.Printf("%s > %s: %t", nf, oth, nf.Gt(oth))
	// Output: 4658 > 4501: true
}

func ExampleNumberForm_Gt_byString() {
	nf, _ := NewNumberForm(`4658`)
	oth := `4501`
	fmt.Printf("%s > %s: %t", nf, oth, nf.Gt(oth))
	// Output: 4658 > 4501: true
}

func ExampleNumberForm_Gt_byUint64() {
	nf, _ := NewNumberForm(uint64(4658))
	oth := uint64(4501)
	fmt.Printf("%s > %d: %t", nf, oth, nf.Gt(oth))
	// Output: 4658 > 4501: true
}

func ExampleNumberForm_Lt() {
	nf, _ := NewNumberForm(4658)
	oth, _ := NewNumberForm(4501)
	fmt.Printf("%s < %s: %t", nf, oth, nf.Lt(oth))
	// Output: 4658 < 4501: false
}

func ExampleNumberForm_Le() {
	nf, _ := NewNumberForm(4658)
	oth, _ := NewNumberForm(4501)
	fmt.Printf("%s =< %s: %t", nf, oth, nf.Le(oth))
	// Output: 4658 =< 4501: false
}

func ExampleNumberForm_Lt_byString() {
	nf, _ := NewNumberForm(`4658`)
	oth := `4501`
	fmt.Printf("%s < %s: %t", nf, oth, nf.Lt(oth))
	// Output: 4658 < 4501: false
}

func ExampleNumberForm_Lt_byUint64() {
	nf, _ := NewNumberForm(uint64(4658))
	oth := uint64(4501)
	fmt.Printf("%s < %d: %t", nf, oth, nf.Lt(oth))
	// Output: 4658 < 4501: false
}
