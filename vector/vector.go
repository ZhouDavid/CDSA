package vector

import (
	"GoSTL/util"
	"errors"
	"reflect"
)

// Vector is a golang implementation of C++ STL vector
type Vector struct {
	elems []interface{}
}

// New empty vector
func New() *Vector {
	return &Vector{
		elems: make([]interface{}, 0, 1024),
	}
}

// NewVector is a constructor of Vector
func NewVector(size int, elem interface{}) *Vector {
	elems := make([]interface{}, size, size<<1)
	for i := 0; i < size; i++ {
		elems[i] = elem
	}
	return &Vector{
		elems: elems,
	}
}

//NewVectorFromSlice is a constructor of vector
func NewVectorFromSlice(slice interface{}) *Vector {
	s := reflect.ValueOf(slice)
	if s.Kind() == reflect.Slice {
		elems := make([]interface{}, s.Len())
		for i := 0; i < s.Len(); i++ {
			elems[i] = s.Index(i).Interface()
		}
		return &Vector{
			elems: elems,
		}

	} else {
		panic("Given a non-slice type")
	}
}

// Size
func (v *Vector) Len() int {
	return len(v.elems)
}

// Cap
func (v *Vector) Cap() int {
	return cap(v.elems)
}

// Empty
func (v *Vector) Empty() bool {
	if v.Len() == 0 {
		return true
	}
	return false
}

// Get
func (v *Vector) Get(index int) (interface{}, error) {
	if index > v.Len() || index < 0 {
		return nil, errors.New("illegal index, out of bounds")
	}
	return v.elems[index], nil
}

// PushBack
func (v *Vector) PushBack(elem interface{}) {
	v.elems = append(v.elems, elem)
}

// PopBack
func (v *Vector) PopBack() (interface{}, error) {
	n := v.Len()
	if n == 0 {
		return nil, errors.New("vector is empty, cannot pop element")
	} else {
		elem := v.elems[n-1]
		v.elems = v.elems[:n-1]
		return elem, nil
	}
}

// Insert
func (v *Vector) Insert(index int, elem interface{}) error {
	if index > v.Len() || index < 0 {
		return errors.New("illegal index, out of bounds")
	}
	if index == v.Len() {
		v.PushBack(elem)
	} else {
		v.elems = append(append(v.elems[:index], elem), v.elems[index+1:]...)
	}
	return nil
}

// Insert Range
func (v *Vector) InsertRange(beginIndex int, e []interface{}) error {
	if beginIndex > v.Len() || beginIndex < 0 {
		return errors.New("illegal index, out of bounds")
	}
	if beginIndex == v.Len() {
		v.elems = append(v.elems[:beginIndex], e...)
	} else {
		n := v.Len()
		v.elems = append(append(v.elems[:beginIndex], e...), v.elems[beginIndex+n:]...)
	}
	return nil
}

// Remove
func (v Vector) Remove(index int) error {
	if index >= v.Len() || index < 0 {
		return errors.New("illegal index, out of bounds")
	}
	v.elems = append(v.elems[:index], v.elems[index+1:]...)
	return nil
}

// Remove Range
func (v *Vector) RemoveRange(beginIndex int, endIndex int) error {
	if (beginIndex >= v.Len() || beginIndex < 0) ||
		(endIndex > v.Len() || endIndex < 0) || (beginIndex >= endIndex) {
		return errors.New("illegal index, out of bounds")
	}
	if endIndex == v.Len() {
		v.elems = append(v.elems[:beginIndex])
	} else {
		v.elems = append(v.elems[:beginIndex], v.elems[endIndex:]...)
	}
	return nil
}

// Swap
func (v *Vector) Swap(i, j int) {
	v.elems[j], v.elems[i] = v.elems[i], v.elems[j]
}

// Less
func (v *Vector) Less(i, j int) bool {
	return util.Less(&v.elems[i], &v.elems[j])
}
