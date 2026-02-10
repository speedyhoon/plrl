package plrl_test

import (
	"fmt"
	"github.com/speedyhoon/plrl"
	"testing"
)

func TestManyEmpty(t *testing.T) {
	runAll(t, "", "", "", "s")
}

func TestManyBoth(t *testing.T) {
	runAll(t, " is", "s are", " is", "s are")
}

func TestManyEmptyMultiple(t *testing.T) {
	runAll(t, "items", "", "items", "s")
}

func TestManyEmptySingle(t *testing.T) {
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
		want, ok := list[i]
		if !ok {
			t.Errorf("test isn't setup correctly, index %d not defined", i)
		}
		if got := plrl.Many(i, single, multiple); got != want {
			t.Errorf("index %d, got `%s`, want `%s`", i, want, got)
		}
	}
}

func TestAny(t *testing.T) {
	tests := []struct {
		length   int
		single   string
		multiple string
		want     string
	}{
		{length: -3, single: "radius", multiple: "radii", want: "radii"},
		{length: -2, single: "radius", multiple: "radii", want: "radii"},
		{length: -1, single: "radius", multiple: "radii", want: "radius"},
		{length: 0, single: "radius", multiple: "radii", want: "radii"},
		{length: 1, single: "radius", multiple: "radii", want: "radius"},
		{length: 2, single: "radius", multiple: "radii", want: "radii"},
		{length: 3, single: "radius", multiple: "radii", want: "radii"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test[%d]", i), func(t *testing.T) {
			if got := plrl.Any(tt.length, tt.single, tt.multiple); got != tt.want {
				t.Errorf("Any(%d, %q, %q) = %q, want %q", tt.length, tt.single, tt.multiple, got, tt.want)
			}
		})
	}
}

func TestS(t *testing.T) {
	tests := []struct {
		length int
		want   string
	}{
		{length: -3, want: "s"},
		{length: -2, want: "s"},
		{length: -1, want: ""},
		{length: 0, want: "s"},
		{length: 1, want: ""},
		{length: 2, want: "s"},
		{length: 3, want: "s"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test[%d]", i), func(t *testing.T) {
			if got := plrl.S(tt.length); got != tt.want {
				t.Errorf("S(%d) = %s, want %s", tt.length, got, tt.want)
			}
		})
	}
}
