// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package request

type contextKey string

func (c contextKey) String() string {
    return "dfl."+string(c)
}

var (
  contextKeyFunctions = contextKey("funcs")
  contextKeyCache = contextKey("cache")
  contextKeyVariables = contextKey("vars")
  contextKeyContext = contextKey("ctx")
  contextKeyExpression = contextKey("exp")
)
