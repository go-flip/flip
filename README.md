# flip
[![buddy pipeline](https://app.buddy.works/kotone/flip/pipelines/pipeline/280627/badge.svg?token=78de2d527ff01f524ebf5cbeb752835d299dfa834590af212642b02673ecba52 "buddy pipeline")](https://app.buddy.works/kotone/flip/pipelines/pipeline/280627)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-flip/flip)](https://goreportcard.com/report/github.com/go-flip/flip)
[![GoDoc](http://godoc.org/github.com/go-flip/flip?status.svg)](http://godoc.org/github.com/go-flip/flip)

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