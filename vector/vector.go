package vector

import (
	"GoSTL/util"
	"errors"
	"reflect"
)

// Vector is a golang implementation of C++ STL vector
type Vector []interface{}

// New empty vector
func New() Vector {
	return make(Vector, 0, 1024)
}

// NewVector is a constructor of Vector
func NewVector(size int, elem interface{}) Vector {
	elems := make([]interface{}, size, size<<1)
	for i := 0; i < size; i++ {
		elems[i] = elem
	}
	return elems
}

//NewVectorFromSlice is a constructor of vector
func NewVectorFromSlice(slice interface{}) Vector {
	s := reflect.ValueOf(slice)
	if s.Kind() == reflect.Slice {
		elems := make([]interface{}, s.Len())
		for i := 0; i < s.Len(); i++ {
			elems[i] = s.Index(i).Interface()
		}
		return elems

	} else {
		panic("Given a non-slice type")
	}
}

// Size
// func (v *Vector) Len() int {
// 	return len(*v)
// }

func (v Vector) Len() int {
	return len(v)
}

// Cap
func (v *Vector) Cap() int {
	return cap(*v)
}

// Empty
func (v *Vector) Empty() bool {
	if v.Len() == 0 {
		return true
	}
	return false
}

// // Get
// func (v *Vector) Get(index int) (interface{}, error) {
// 	if index > v.Len() || index < 0 {
// 		return nil, errors.New("illegal index, out of bounds")
// 	}
// 	return (*v)[index], nil
// }

// PushBack
func (v *Vector) PushBack(elem interface{}) {
	*v = append(*v, elem)
}

// PopBack
func (v *Vector) PopBack() (interface{}, error) {
	n := v.Len()
	if n == 0 {
		return nil, errors.New("vector is empty, cannot pop element")
	} else {
		elem := (*v)[n-1]
		*v = (*v)[:n-1]
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
		*v = append(append((*v)[:index], elem), (*v)[index+1:]...)
	}
	return nil
}

// Insert Range
func (v *Vector) InsertRange(beginIndex int, e []interface{}) error {
	if beginIndex > v.Len() || beginIndex < 0 {
		return errors.New("illegal index, out of bounds")
	}
	if beginIndex == v.Len() {
		*v = append((*v)[:beginIndex], e...)
	} else {
		n := v.Len()
		*v = append(append((*v)[:beginIndex], e...), (*v)[beginIndex+n:]...)
	}
	return nil
}

// Remove
func (v *Vector) Remove(index int) error {
	if index >= v.Len() || index < 0 {
		return errors.New("illegal index, out of bounds")
	}
	*v = append((*v)[:index], (*v)[index+1:]...)
	return nil
}

// Remove Range
func (v *Vector) RemoveRange(beginIndex int, endIndex int) error {
	if (beginIndex >= v.Len() || beginIndex < 0) ||
		(endIndex > v.Len() || endIndex < 0) || (beginIndex >= endIndex) {
		return errors.New("illegal index, out of bounds")
	}
	if endIndex == v.Len() {
		*v = append((*v)[:beginIndex])
	} else {
		*v = append((*v)[:beginIndex], (*v)[endIndex:]...)
	}
	return nil
}

// Swap
func (v Vector) Swap(i, j int) {
	v[j], v[i] = v[i], v[j]
}

// Less
func (v Vector) Less(i, j int) bool {
	return util.Less(&v[i], &v[j])
}
