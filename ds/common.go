package ds

// Comparable 可比较对象，-1: 小于，0: 等于，1: 大于
type Comparable interface {
	Compare(other Comparable) int
}
