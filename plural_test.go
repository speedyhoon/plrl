package utl

import (
	"testing"
)

func TestPluralEmpty(t *testing.T) {
	runAll(t, "", "", "", "s")
}

func TestPluralBoth(t *testing.T) {
	runAll(t, " is", "s are", " is", "s are")
}

func TestPluralEmptyMultiple(t *testing.T) {
	runAll(t, "items", "", "items", "s")
}

func TestPluralEmptySingle(t *testing.T) {
	runAll(t, "", "foobar", "", "foobar")
}

func runAll(t *testing.T, single, multiple, assertSingle, assertMultiple string) {
	list := map[int]string{
		-3: assertMultiple,
		-2: assertMultiple,
		-1: assertSingle,
		0:  assertMultiple,
		1:  assertSingle,
		2:  assertMultiple,
		3:  assertMultiple,
	}
	for i := -3; i <= 3; i++ {
		assert, ok := list[i]
		if !ok {
			t.Errorf("test isn't setup correctly")
		}
		if plural := Plural(i, single, multiple); assert != plural {
			t.Errorf("len(%d) `%v` != `%v`", i, assert, plural)
		}
	}
}
