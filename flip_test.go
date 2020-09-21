package flip

import (
	"reflect"
	"testing"
)

func compare(t *testing.T, a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%v != %v", a, b)
	}
}

type myint int

type myintSlice []myint

func (p myintSlice) Len() int {
	return len(p)
}

func (p myintSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func TestBytes(t *testing.T) {
	in, out := ByteSlice{'a', 'b', 'c'}, ByteSlice{'c', 'b', 'a'}
	Bytes(in)
	compare(t, in, out)
}

func TestRunes(t *testing.T) {
	in, out := RuneSlice{'a', 'b', 'c'}, RuneSlice{'c', 'b', 'a'}
	Runes(in)
	compare(t, in, out)
}

var sliceTests = []struct {
	in, out interface{}
}{
	{[]bool{false, true, true}, []bool{true, true, false}},

	{[]uint8{0, 1, 2}, []uint8{2, 1, 0}},
	{[]uint16{0, 1, 2}, []uint16{2, 1, 0}},
	{[]uint32{0, 1, 2}, []uint32{2, 1, 0}},
	{[]uint64{0, 1, 2}, []uint64{2, 1, 0}},

	{[]int8{0, 1, 2}, []int8{2, 1, 0}},
	{[]int16{0, 1, 2}, []int16{2, 1, 0}},
	{[]int32{0, 1, 2}, []int32{2, 1, 0}},
	{[]int64{0, 1, 2}, []int64{2, 1, 0}},

	{[]byte{'a', 'b', 'c'}, []byte{'c', 'b', 'a'}},
	{[]rune{'a', 'b', 'c'}, []rune{'c', 'b', 'a'}},

	{[]uint{0, 1, 2}, []uint{2, 1, 0}},
	{[]int{0, 1, 2}, []int{2, 1, 0}},
	{[]uintptr{0, 1, 2}, []uintptr{2, 1, 0}},

	{[]float32{3.14, 2.718, 1.618}, []float32{1.618, 2.718, 3.14}},
	{[]float64{3.14, 2.718, 1.618}, []float64{1.618, 2.718, 3.14}},

	{[]complex64{3.14, 2.718, 1.618}, []complex64{1.618, 2.718, 3.14}},
	{[]complex128{3.14, 2.718, 1.618}, []complex128{1.618, 2.718, 3.14}},

	{[]string{"Hello", "world", "!"}, []string{"!", "world", "Hello"}},

	{ByteSlice{0, 1, 2}, ByteSlice{2, 1, 0}},
	{RuneSlice{0, 1, 2}, RuneSlice{2, 1, 0}},

	{myintSlice{0, 1, 2}, myintSlice{2, 1, 0}},
	{[]myint{0, 1, 2}, []myint{2, 1, 0}},
}

func TestSlice(t *testing.T) {
	for _, tt := range sliceTests {
		Slice(tt.in)
		compare(t, tt.in, tt.out)
	}
}

func TestString(t *testing.T) {
	out, in := "Hello, world!", "!dlrow ,olleH"
	if String(in) != out {
		t.Errorf("String(%q) = %q, want %q", in, String(in), out)
	}
}

func BenchmarkFlip(b *testing.B) {
	n := 1000
	p := make([]myint, n)

	b.Run("Flip", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Flip(myintSlice(p))
		}
	})

	b.Run("Slice", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			Slice(p)
		}
	})
}
