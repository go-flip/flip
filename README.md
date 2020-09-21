# flip
Routines for flipping slices in Go.

## Usage
```go
package main

import "github.com/go-flip/flip"

// IntSlice is equivalent to flip.IntSlice.
type IntSlice []int

// Len returns the length of the slice.
func (p IntSlice) Len() int {
    return len(p)
}

// Swap swaps the elements at the given indices.
func (p IntSlice) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}

func main() {
    p := []int{1, 2, 3, 4}
    flip.Flip(IntSlice(p))
    fmt.Println(p)
}
```

`flip` can flip the contents of a slice by using the `Len()` and `Swap()` methods provided.
If a type already satisfies `sort.Interface`, it can be passed directly to `flip.Flip`.
Some interfaces for primitive slice types are implemented in the module.
See the [documentation](https://godoc.org/github.com/go-flip/flip) for available types.