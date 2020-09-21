package flip

import (
	"reflect"
	"testing"
)

var flipTests = []struct {
	in, out interface{}
}{
	{BoolSlice{false, true}, BoolSlice{true, false}},
	{ByteSlice{'a', 'b', 'c'}, ByteSlice{'c', 'b', 'a'}},
	{RuneSlice{'a', 'b', 'c'}, RuneSlice{'c', 'b', 'a'}},
	{IntSlice{0, 1, 2}, IntSlice{2, 1, 0}},
	{Float32Slice{3.14, 2.718, 1.618}, Float32Slice{1.618, 2.718, 3.14}},
	{Float64Slice{3.14, 2.718, 1.618}, Float64Slice{1.618, 2.718, 3.14}},
	{StringSlice{"Hello", "world", "!"}, StringSlice{"!", "world", "Hello"}},
}

func compare(t *testing.T, a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%v != %v", a, b)
	}
}

func TestFlip(t *testing.T) {
	for _, tt := range flipTests {
		Slice(tt.in)
		compare(t, tt.in, tt.out)
		Slice(tt.in)

		switch in := tt.in.(type) {
		case BoolSlice:
			if out, ok := tt.out.(BoolSlice); ok {
				Bools(in)
				compare(t, in, out)
			}
		case ByteSlice:
			if out, ok := tt.out.(ByteSlice); ok {
				Bytes(in)
				compare(t, in, out)
			}
		case RuneSlice:
			if out, ok := tt.out.(RuneSlice); ok {
				Runes(in)
				compare(t, in, out)
			}
		case IntSlice:
			if out, ok := tt.out.(IntSlice); ok {
				Ints(in)
				compare(t, in, out)
			}
		case Float32Slice:
			if out, ok := tt.out.(Float32Slice); ok {
				Float32s(in)
				compare(t, in, out)
			}
		case Float64Slice:
			if out, ok := tt.out.(Float64Slice); ok {
				Float64s(in)
				compare(t, in, out)
			}
		case StringSlice:
			if out, ok := tt.out.(StringSlice); ok {
				Strings(in)
				compare(t, in, out)
			}
		}
	}

	out, in := "Hello, world!", "!dlrow ,olleH"
	if String(in) != out {
		t.Errorf("String(%q) = %q, want %q", in, String(in), out)
	}
}
