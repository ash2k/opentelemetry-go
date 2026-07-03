// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package internal provides internal functionality for the metric package.
package internal // import "go.opentelemetry.io/otel/sdk/metric/internal"
import (
	"slices"
)

// ReuseSlice ensures slice length is n and slice capacity is at least n elements.
func ReuseSlice[T any](slice []T, n int) []T {
	// Go allocates memory using size classes.
	// When a slice is allocated using 'append', it uses a chunk of memory that is
	// at least as big as the requested slice capacity.
	// While the length may be smaller, the capacity of that slice is the size of the memory class.
	// Hence, it is more efficient to use 'append' to allocate slices when the extra capacity may come in handy.
	// slices.Grow uses 'append' for that exact reason.
	//
	// Go 1.25 had fixed size classes: https://github.com/golang/go/blob/go1.25.11/src/internal/runtime/gc/sizeclasses.go
	// Go 1.26 works differently but the idea is the same: https://github.com/golang/go/blob/go1.26.4/src/runtime/malloc.go#L10-L15
	//
	// See https://go.dev/play/p/zK2t1ai1_vc.
	return slices.Grow(slice[:0], n)[:n]
}
