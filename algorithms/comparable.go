package algorithms

type Comparable interface {
	IsEqual(b interface{}) bool
	Less(b interface{}) bool
}
