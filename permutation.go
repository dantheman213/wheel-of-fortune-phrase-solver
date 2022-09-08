package main

import (
	"fmt"
	"strings"
)

// Iterator of a slice index. `len` equals to the length of the slice
type IdxIter struct {
	idx uint
	len uint
}

// Returns true if the slice is empty.
func (i IdxIter) Empty() bool {
	return i.len == 0
}

// Returns true is the iteration is over.
func (i IdxIter) Done() bool {
	return i.idx >= i.len
}

// Builds the next iteration value. If called for the last index,
// the next value's `Done` returns `true`.
func (i *IdxIter) Next() {
	i.idx++
}

// Resets the iterator
func (i *IdxIter) Reset() {
	i.idx = 0
}

// The index value
func (i IdxIter) Idx() uint {
	return i.idx
}

// Utility function for debugging
func (i IdxIter) String() string {
	return fmt.Sprintf("%d/%d", i.idx, i.len)
}

// Creates new iterator
func NewIdxIter(max uint) IdxIter {
	return IdxIter{idx: 0, len: max}
}

// Index iterator for a slice of slices
type IdxVectorIter []IdxIter

// Returns true is the iteration is over.
func (ii IdxVectorIter) Done() bool {
	last := len(ii) - 1
	return ii[last].Done()
}

// Builds the next iteration value. If called for the last index vector,
// the next value's `Done` returns `true`.
func (ii IdxVectorIter) Next() {
	if len(ii) == 0 {
		return
	}
	last := len(ii) - 1
	for pos := range ii[:last] {
		ii[pos].Next()
		if ii[pos].Done() {
			ii[pos].Reset()
		} else {
			return
		}
	}
	ii[last].Next()
}

// Utility function for debugging
func (ii IdxVectorIter) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for _, i := range ii {
		sb.WriteString(i.String())
		sb.WriteString(",")
	}
	sb.WriteString("]")
	return sb.String()
}

// Creates new vector iterator
func NewIdxVector(max ...uint) IdxVectorIter {
	res := make(IdxVectorIter, len(max))
	for pos := range res {
		res[pos].len = max[pos]
	}
	return res
}

// Creates new vector iterator for a slice of slices
func NewIdxVectorFromSlices[T any](slices [][]T) IdxVectorIter {
	max := make([]uint, len(slices))
	for pos := range slices {
		max[pos] = uint(len(slices[pos]))
	}
	return NewIdxVector(max...)
}

// Returns a slice of values that correspond to the given state of the iterator
// Unsafe, doesn't check correctnes of the input data
func Get[T any](slices [][]T, ii IdxVectorIter) []T {
	res := make([]T, len(ii))
	GetTo(slices, res, ii)
	return res
}

// Copies the values that correspond to the given state of the iterator, to the given slice
// Unsafe, doesn't check correctnes of the input data
func GetTo[T any](slices [][]T, dst []T, ii IdxVectorIter) {
	for pos := range ii {
		if !ii[pos].Empty() {
			dst[pos] = slices[pos][ii[pos].Idx()]
		}
	}
}
