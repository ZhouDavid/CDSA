package util

type Comparable interface {
	IsEqual(b *interface{}) bool
	Less(b *interface{}) bool
}

func IsComparable(obj *interface{}) bool {
	if _, ok := (*obj).(Comparable); ok {
		return true
	}
	return false
}

func Less(a *interface{}, b *interface{}) bool {
	if IsComparable(a) && IsComparable(b) {
		c := (*a).(Comparable)
		return c.Less(b)
	}
	return basicTypeLess(a, b)
}

func basicTypeLess(a *interface{}, b *interface{}) bool {
	switch (*a).(type) {
	case int:
		return (*a).(int) < (*b).(int)
	case int8:
		return (*a).(int8) < (*b).(int8)
	case int16:
		return (*a).(int16) < (*b).(int16)
	case int32:
		return (*a).(int32) < (*b).(int32)
	case int64:
		return (*a).(int64) < (*b).(int64)
	case float32:
		return (*a).(float32) < (*b).(float32)
	case float64:
		return (*a).(float64) < (*b).(float64)
	case string:
		return (*a).(string) < (*b).(string)
	default:
		panic("a or b is not basic type")
	}
}
