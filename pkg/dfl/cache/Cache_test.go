// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package cache

import (
	"testing"

	"github.com/spatialcurrent/go-dfl/dfl"
)

func TestCache(t *testing.T) {

	c := New()

	x := "((@a ?: 10) == 10)"
	y := "(@a ?: 4) + 6"

	z := c.MustParseCompile(x)
	if z.Dfl(dfl.DefaultQuotes, false, 0) != x {
		t.Errorf("cache.Get(%q) == %q, want %q", x, z.Dfl(dfl.DefaultQuotes, false, 0), x)
	}

	if c.Has(y) {
		t.Errorf("cache.Has(%q) == %v, want %v", y, true, false)
	}

	if !c.Has(x) {
		t.Errorf("cache.Has(%q) == %v, want %v", x, false, true)
	}

	if c.Has(y) {
		t.Errorf("cache.Has(%q) == %v, want %v", y, true, false)
	}

	if i, ok := c.Get(x); !ok || i.Dfl(dfl.DefaultQuotes, false, 0) != x {
		t.Errorf("cache.Get(%q) == %q, want %q", x, i.Dfl(dfl.DefaultQuotes, false, 0), x)
	}

}
