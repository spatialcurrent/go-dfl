// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package cache

import (
	"testing"
)

func TestCache(t *testing.T) {

  c := New()

  x := "((@a ?: 10) == 10)"
  y := "(@a ?: 4) + 6"

  z := c.MustParseCompile(x)
  if z.Dfl(DefaultQuotes, false, 0) != x {
    t.Errorf("cache.Get(%q) == %q, want %q", x, z.Dfl(DefaultQuotes, false, 0), x)
  }

  if cache.Has(y) {
    t.Errorf("cache.Has(%q) == %v, want %v", y, true, false)
  }

  if ! cache.Has(x) {
    t.Errorf("cache.Has(%q) == %v, want %v", x, false, true)
  }

  if cache.Has(y) {
    t.Errorf("cache.Has(%q) == %v, want %v", y, true, false)
  }

  if t, ok := cache.Get(x); !ok || t.Dfl(DefaultQuotes, false, 0) != x {
    t.Errorf("cache.Get(%q) == %q, want %q", x, t.Dfl(DefaultQuotes, false, 0), x)
  }

}
