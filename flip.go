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

// Uint8Slice attaches the methods of Interface to []uint8.
type Uint8Slice []uint8

// Len returns the length of the slice.
func (p Uint8Slice) Len() int {
	return len(p)
}

// Swap swaps the elements at the given indices.
func (p Uint8Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Uint8s flips a slice of uint8s.
func Uint8s(p []uint8) {
	Flip(Uint8Slice(p))
}

// Uint16Slice attaches the methods of Interface to []uint16.
type Uint16Slice []uint16

// Len returns the length of the slice.
func (p Uint16Slice) Len() int {
	return len(p)
}

// Swap swaps the elements at the given indices.
func (p Uint16Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Uint16s flips a slice of uint16s.
func Uint16s(p []uint16) {
	Flip(Uint16Slice(p))
}

// Uint32Slice attaches the methods of Interface to []uint32.
type Uint32Slice []uint32

// Len returns the length of the slice.
func (p Uint32Slice) Len() int {
	return len(p)
}

// Swap swaps the elements at the given indices.
func (p Uint32Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Uint32s flips a slice of uint32s.
func Uint32s(p []uint32) {
	Flip(Uint32Slice(p))
}

// Uint64Slice attaches the methods of Interface to []uint64.
type Uint64Slice []uint64

// Len returns the length of the slice.
func (p Uint64Slice) Len() int {
	return len(p)
}

// Swap swaps the elements at the given indices.
func (p Uint64Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Uint64s flips a slice of uint64s.
func Uint64s(p []uint64) {
	Flip(Uint64Slice(p))
}

// Int8Slice attaches the methods of Interface to []int8.
type Int8Slice []int8

// Len returns the length of the slice.
func (p Int8Slice) Len() int {
	return len(p)
}

// Swap swaps the elements at the given indices.
func (p Int8Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Int8s flips a slice of int8s.
func Int8s(p []int8) {
	Flip(Int8Slice(p))
}

// Int16Slice attaches the methods of Interface to []int16.
type Int16Slice []int16

// Len returns the length of the slice.
func (p Int16Slice) Len() int {
	return len(p)
}

// Swap swaps the elements at the given indices.
func (p Int16Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Int16s flips a slice of int16s.
func Int16s(p []int16) {
	Flip(Int16Slice(p))
}

// Int32Slice attaches the methods of Interface to []int32.
type Int32Slice []int32

// Len returns the length of the slice.
func (p Int32Slice) Len() int {
	return len(p)
}

// Swap swaps the elements at the given indices.
func (p Int32Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Int32s flips a slice of int32s.
func Int32s(p []int32) {
	Flip(Int32Slice(p))
}

// Int64Slice attaches the methods of Interface to []int64.
type Int64Slice []int64

// Len returns the length of the slice.
func (p Int64Slice) Len() int {
	return len(p)
}

// Swap swaps the elements at the given indices.
func (p Int64Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Int64s flips a slice of int64s.
func Int64s(p []int64) {
	Flip(Int64Slice(p))
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

// Complex64Slice attaches the methods of Interface to []complex64.
type Complex64Slice []complex64

// Len returns the length of the slice.
func (p Complex64Slice) Len() int {
	return len(p)
}

// Swap swaps the elements at the given indices.
func (p Complex64Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Complex64s flips a slice of complex64s.
func Complex64s(p []complex64) {
	Flip(Complex64Slice(p))
}

// Complex128Slice attaches the methods of Interface to []complex128.
type Complex128Slice []complex128

// Len returns the length of the slice.
func (p Complex128Slice) Len() int {
	return len(p)
}

// Swap swaps the elements at the given indices.
func (p Complex128Slice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Complex128s flips a slice of complex128s.
func Complex128s(p []complex128) {
	Flip(Complex128Slice(p))
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

// UintSlice attaches the methods of Interface to []uint.
type UintSlice []uint

// Len returns the length of the slice.
func (p UintSlice) Len() int {
	return len(p)
}

// Swap swaps the elements at the given indices.
func (p UintSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Uints flips a slice of uints.
func Uints(p []uint) {
	Flip(UintSlice(p))
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

// UintptrSlice attaches the methods of Interface to []uintptr.
type UintptrSlice []uintptr

// Len returns the length of the slice.
func (p UintptrSlice) Len() int {
	return len(p)
}

// Swap swaps the elements at the given indices.
func (p UintptrSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Uintptrs flips a slice of uintptrs.
func Uintptrs(p []uintptr) {
	Flip(UintptrSlice(p))
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
