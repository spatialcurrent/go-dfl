// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// FunctionMap is a map of functions by string that are reference by name in the Function Node.
type FunctionMap map[string]func(Context, []string) (interface{}, error)
