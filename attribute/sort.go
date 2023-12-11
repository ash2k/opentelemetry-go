//go:build !go1.21

package attribute

import "sort"

func sliceSortStable(kvs []KeyValue) {
	sort.SliceStable(kvs, func(i, j int) bool {
		return kvs[i].Key < kvs[j].Key
	})
}
