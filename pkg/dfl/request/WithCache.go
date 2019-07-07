// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package request

import (
	"context"
)

import (
	"github.com/spatialcurrent/go-dfl/pkg/dfl/cache"
)

func WithCache(requestContext context.Context, c *cache.Cache) context.Context {
	return context.WithValue(requestContext, contextKeyCache, c)
}
