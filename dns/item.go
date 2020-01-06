package dns

import "github.com/Workiva/go-datastructures/queue"

// Item 实现了queue.Item接口
type Item struct {
	ID        string
	Timestamp int64
}

// Compare returns a bool that can be used to determine
// ordering in the priority queue.  Assuming the queue
// is in ascending order, this should return > logic.
// Return 1 to indicate this object is greater than the
// the other logic, 0 to indicate equality, and -1 to indicate
// less than other.
func (i Item) Compare(other queue.Item) int {
	myOther := other.(Item)
	if i.Timestamp < myOther.Timestamp {
		return -1
	} else if i.Timestamp == myOther.Timestamp {
		return 0
	}
	return 1
}
