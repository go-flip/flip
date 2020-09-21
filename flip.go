package flip

// Interface represents a type whose order of elements can be flipped by the
// routines in this package. The methods require that the elements of the
// collection be enumerated by an integer index.
type Interface interface {
	Len() int
	Swap(i, j int)
}

// Flip flips the data.
func Flip(data Interface) {
	for l, r := 0, data.Len()-1; l < r; l, r = l+1, r-1 {
		data.Swap(l, r)
	}
}

// BoolSlice attaches the methods of Interface to []bool.
type BoolSlice []bool

// Len returns the length of the slice.
func (p BoolSlice) Len() int {
	return len(p)
}

// Swap swaps the elements at the given indices.
func (p BoolSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Bools flips a slice of bools.
func Bools(p []bool) {
	Flip(BoolSlice(p))
}

// ByteSlice attaches the methods of Interface to []byte.
type ByteSlice []byte

// Len returns the length of the slice.
func (p ByteSlice) Len() int {
	return len(p)
}

// Swap swaps the elements at the given indices.
func (p ByteSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Bytes flips a slice of bytes.
func Bytes(p []byte) {
	Flip(ByteSlice(p))
}

// RuneSlice attaches the methods of Interface to []rune.
type RuneSlice []rune

// Len returns the length of the slice.
func (p RuneSlice) Len() int {
	return len(p)
}

// Swap swaps the elements at the given indices.
func (p RuneSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Runes flips a slice of runes.
func Runes(p []rune) {
	Flip(RuneSlice(p))
}

// IntSlice attaches the methods of Interface to []int.
type IntSlice []int

// Len returns the length of the slice.
func (p IntSlice) Len() int {
	return len(p)
}

// Swap swaps the elements at the given indices.
func (p IntSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Ints flips a slice of ints.
func Ints(p []int) {
	Flip(IntSlice(p))
}

// Float32Slice attaches the methods of Interface to []float32.
type Float32Slice []float32

// Len returns the length of the slice.
func (p Float32Slice) Len() int {
	return len(p)
}

// Swap swaps the elements at the given indices.
func (p Float32Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Float32s flips a slice of float32s.
func Float32s(p []float32) {
	Flip(Float32Slice(p))
}

// Float64Slice attaches the methods of Interface to []float64.
type Float64Slice []float64

// Len returns the length of the slice.
func (p Float64Slice) Len() int {
	return len(p)
}

// Swap swaps the elements at the given indices.
func (p Float64Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Float64s flips a slice of float64s.
func Float64s(p []float64) {
	Flip(Float64Slice(p))
}

// StringSlice attaches the methods of Interface to []string.
type StringSlice []string

// Len returns the length of the slice.
func (p StringSlice) Len() int {
	return len(p)
}

// Swap swaps the elements at the given indices.
func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Strings flips a slice of strings.
func Strings(p []string) {
	Flip(StringSlice(p))
}

// String flips a string. This function will return a value as strings are
// immutable in Go.
func String(s string) string {
	p := []rune(s)
	Runes(p)
	return string(p)
}
