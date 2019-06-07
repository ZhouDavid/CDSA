package vector

import "errors"

// Vector is a golang implementation of C++ STL vector
type Vector []interface{}

// New empty vector
func New() *Vector {
	elems := make(Vector, 0, 1024)
	return &elems
}

// NewVector is a constructor of Vector
func NewVector(size int, elem interface{}) *Vector {
	elems := make(Vector, size, size<<1)
	for i := 0; i < size; i++ {
		elems[i] = elem
	}
	return &elems
}

//NewVectorFromSlice is a constructor of vector
func NewVectorFromSlice(elems []interface{}) *Vector {
	arr := Vector(elems)
	return &arr
}

// Size
func (arr *Vector) Size() int {
	return len(*arr)
}

// Cap
func (arr *Vector) Cap() int {
	return cap(*arr)
}

// Empty
func (arr *Vector) Empty() bool {
	if arr.Size() == 0 {
		return true
	}
	return false
}

// Get
func (arr *Vector) Get(index int) (interface{}, error) {
	if index > arr.Size() || index < 0 {
		return nil, errors.New("illegal index, out of bounds")
	}
	return (*arr)[index], nil
}

// PushBack
func (arr *Vector) PushBack(elem interface{}) {
	*arr = append(*arr, elem)
}

// PopBack
func (arr *Vector) PopBack() (interface{}, error) {
	size := len(*arr)
	if size == 0 {
		return nil, errors.New("vector is empty, cannot pop element")
	} else {
		elem := (*arr)[size-1]
		*arr = (*arr)[:size-1]
		return elem, nil
	}
}

// Insert
func (arr *Vector) Insert(index int, elem interface{}) error {
	if index > arr.Size() || index < 0 {
		return errors.New("illegal index, out of bounds")
	}
	if index == arr.Size() {
		arr.PushBack(elem)
	} else {
		*arr = append(append((*arr)[:index], elem), (*arr)[index+1:]...)
	}
	return nil
}

// Insert Range
func (arr *Vector) InsertRange(beginIndex int, v []interface{}) error {
	if beginIndex > arr.Size() || beginIndex < 0 {
		return errors.New("illegal index, out of bounds")
	}
	if beginIndex == arr.Size() {
		*arr = append((*arr)[:beginIndex], v...)
	} else {
		n := len(v)
		*arr = append(append((*arr)[:beginIndex], v...), (*arr)[beginIndex+n:]...)
	}
	return nil
}

// Remove
func (arr *Vector) Remove(index int) error {
	if index >= arr.Size() || index < 0 {
		return errors.New("illegal index, out of bounds")
	}
	*arr = append((*arr)[:index], (*arr)[index+1:]...)
	return nil
}

// Remove Range
func (arr *Vector) RemoveRange(beginIndex int, endIndex int) error {
	if (beginIndex >= arr.Size() || beginIndex < 0) ||
		(endIndex >= arr.Size() || endIndex < 0) || (beginIndex >= endIndex) {
		return errors.New("illegal index, out of bounds")
	}
	*arr = append((*arr)[:beginIndex], (*arr)[endIndex:]...)
	return nil
}
