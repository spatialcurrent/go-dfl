// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"sort"
	"strings"
)

// StringSet is a logical set of string values using a map[string]struct{} backend.
// The use of a string -> empty struct backend provides a higher write performance versus a slice backend.
type StringSet map[string]struct{}

func (set StringSet) Dfl(quotes []string, pretty bool, tabs int) string {
	values := make([]string, 0, len(set))
	for v := range set {
		values = append(values, TryFormatLiteral(v, quotes, pretty, tabs))
	}
	return "{" + strings.Join(values, ", ") + "}"
}

func (set StringSet) Contains(x string) bool {
	_, ok := set[x]
	return ok
}

func (set StringSet) Len() int {
	return len(set)
}

// Add is a variadic function to add values to a set
//	- https://gobyexample.com/variadic-functions
func (set StringSet) Add(values ...string) {
	for _, v := range values {
		set[v] = struct{}{}
	}
}

func (set StringSet) Union(values interface{}) StringSet {
	union := NewStringSet()
	for x := range set {
		union.Add(x)
	}
	switch values := values.(type) {
	case []string:
		for _, v := range values {
			union.Add(v)
		}
	case StringSet:
		for v := range values {
			union.Add(v)
		}
	case map[string]struct{}:
		for v := range values {
			union.Add(v)
		}
	}
	return union
}

func (set StringSet) Intersection(values interface{}) StringSet {
	intersection := NewStringSet()
	switch values := values.(type) {
	case []string:
		for _, v := range values {
			if set.Contains(v) {
				intersection.Add(v)
			}
		}
	case StringSet:
		for v := range values {
			if set.Contains(v) {
				intersection.Add(v)
			}
		}
	case map[string]struct{}:
		for v := range values {
			if set.Contains(v) {
				intersection.Add(v)
			}
		}
	}

	return intersection
}

func (set StringSet) Intersects(values interface{}) bool {

	switch values := values.(type) {
	case []string:
		for _, v := range values {
			if set.Contains(v) {
				return true
			}
		}
	case StringSet:
		for v := range values {
			if set.Contains(v) {
				return true
			}
		}
	case map[string]struct{}:
		for v := range values {
			if set.Contains(v) {
				return true
			}
		}
	}
	return false
}

// Slice returns a slice representation of this set.
// If parameter sorted is true, then sorts the values using natural sort order.
func (set StringSet) Slice(sorted bool) sort.StringSlice {
	slice := sort.StringSlice(make([]string, len(set)))
	for x := range set {
		slice = append(slice, x)
	}
	if sorted {
		slice.Sort()
	}
	return slice
}

// NewStringSet returns a new StringSet.
func NewStringSet() StringSet {
	return make(map[string]struct{})
}
