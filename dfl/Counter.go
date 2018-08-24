// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"sort"
)

// Counter is used for creating a frequency histogram of values
type Counter map[string]int

func NewCounter() Counter {
	return Counter(map[string]int{})
}

func (c Counter) Len() int {
	return len(c)
}

func (c Counter) Has(key string) bool {
	_, ok := c[key]
	return ok
}

func (c Counter) Increment(key string) {
	if count, ok := c[key]; ok {
		c[key] = count + 1
	} else {
		c[key] = 1
	}
}

func (c Counter) Top(n int, min int) []string {

	if n == 0 {
		return make([]string, 0)
	}

	items := make([]struct {
		Value     string
		Frequency int
	}, 0)
	for value, frequency := range c {
		if frequency >= min {
			items = append(items, struct {
				Value     string
				Frequency int
			}{Value: value, Frequency: frequency})
		}
	}

	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Frequency > items[j].Frequency
	})

	if n > 0 && n < len(items) {
		values := make([]string, 0, n)
		for _, item := range items {
			values = append(values, item.Value)
		}
		return values[:n]
	}

	values := make([]string, 0, len(items))
	for _, item := range items {
		values = append(values, item.Value)
	}
	return values
}

func (c Counter) Bottom(n int, max int) []string {

	if n == 0 {
		return make([]string, 0)
	}

	items := make([]struct {
		Value     string
		Frequency int
	}, 0)
	for value, frequency := range c {
		if frequency >= max {
			items = append(items, struct {
				Value     string
				Frequency int
			}{Value: value, Frequency: frequency})
		}
	}

	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Frequency < items[j].Frequency
	})

	if n > 0 && n < len(items) {
		values := make([]string, 0, n)
		for _, item := range items {
			values = append(values, item.Value)
		}
		return values[:n]
	}

	values := make([]string, 0, len(items))
	for _, item := range items {
		values = append(values, item.Value)
	}
	return values
}
