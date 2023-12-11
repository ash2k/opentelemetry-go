//go:build go1.21

package attribute

import "slices"

func sliceSortStable(kvs []KeyValue) {
	slices.SortStableFunc(kvs, func(a, b KeyValue) int {
		if a.Key == b.Key {
			return 0
		}
		if a.Key < b.Key {
			return -1
		}
		return 1
	})
}
