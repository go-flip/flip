package flip

import "reflect"

type reflectSlice struct {
	rv   reflect.Value
	swap func(i, j int)
}

func (p reflectSlice) Len() int {
	return p.rv.Len()
}

func (p reflectSlice) Swap(i, j int) {
	p.swap(i, j)
}

// Slice flips the provided slice. The function panics if the provided
// interface is not a slice.
func Slice(slice interface{}) {
	switch p := slice.(type) {
	case []bool:
		Bools(p)
	case []uint8:
		Uint8s(p)
	case []uint16:
		Uint16s(p)
	case []uint32:
		Uint32s(p)
	case []uint64:
		Uint64s(p)
	case []int8:
		Int8s(p)
	case []int16:
		Int16s(p)
	case []int32:
		Int32s(p)
	case []int64:
		Int64s(p)
	case []float32:
		Float32s(p)
	case []float64:
		Float64s(p)
	case []complex64:
		Complex64s(p)
	case []complex128:
		Complex128s(p)
	case []uint:
		Uints(p)
	case []int:
		Ints(p)
	case []uintptr:
		Uintptrs(p)
	case []string:
		Strings(p)
	case Interface:
		Flip(p)
	default:
		rv := reflectValueOf(slice)
		swap := reflectSwapper(slice)
		Flip(reflectSlice{rv, swap})
	}
}
