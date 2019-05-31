package vector

import "errors"

// Vector is a golang implementation of C++ STL vector
type Vector []interface{}

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

// PushBack
func (arr *Vector) PushBack(elem interface{}) {
	*arr = append(*arr, elem)
}

// Pop
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
	if index >= arr.Size() || index < 0 {
		return errors.New("illegal index, out of bounds")
	}
	*arr = append(append((*arr)[:index], elem), (*arr)[index+1:]...)
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
