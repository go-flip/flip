package flip

import (
	"reflect"
	"testing"
)

type myint int

type myintSlice []myint

func (p myintSlice) Len() int {
	return len(p)
}

func (p myintSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

var flipTests = []struct {
	in, out interface{}
}{
	{[]bool{false, true, true}, []bool{true, true, false}},
	{[]byte{'a', 'b', 'c'}, []byte{'c', 'b', 'a'}},
	{[]rune{'a', 'b', 'c'}, []rune{'c', 'b', 'a'}},
	{[]int{0, 1, 2}, []int{2, 1, 0}},
	{[]float32{3.14, 2.718, 1.618}, []float32{1.618, 2.718, 3.14}},
	{[]float64{3.14, 2.718, 1.618}, []float64{1.618, 2.718, 3.14}},
	{[]string{"Hello", "world", "!"}, []string{"!", "world", "Hello"}},
	{myintSlice{0, 1, 2}, myintSlice{2, 1, 0}},
	{[]myint{0, 1, 2}, []myint{2, 1, 0}},
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
	}

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
