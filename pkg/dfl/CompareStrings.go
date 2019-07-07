// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"regexp"
	"strings"
)

import (
	"github.com/pkg/errors"
)

// CompareStrings compares parameter a and parameter b.
// The parameters must be of type string.
// Returns true if a like b.
func CompareStrings(lvs string, rvs string) (bool, error) {

	if len(rvs) == 0 {
		return len(lvs) == 0, nil
	}

	pattern, err := regexp.Compile("^" + strings.Replace(rvs, "%", ".*", -1) + "$")
	if err != nil {
		return false, errors.Wrap(err, "Error comparing strings \""+lvs+"\" and \""+rvs+"\"")
	}

	return pattern.MatchString(lvs), nil
}
