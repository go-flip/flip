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
	rv := reflectValueOf(slice)
	swap := reflectSwapper(slice)
	Flip(reflectSlice{rv, swap})
}
