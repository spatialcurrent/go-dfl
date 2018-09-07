// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"github.com/pkg/errors"
)

// AttachLeft attaches the left Node as the left child node to the parent root Node.
func AttachLeft(root Node, left Node) error {
	switch root.(type) {
	case *Declare:
		root.(*Declare).Left = left
	case *Pipe:
		root.(*Pipe).Left = left
	case *And:
		root.(*And).Left = left
	case *Or:
		root.(*Or).Left = left
	case *Xor:
		root.(*Xor).Left = left
	case *Coalesce:
		root.(*Coalesce).Left = left
	case *In:
		root.(*In).Left = left
	case *Like:
		root.(*Like).Left = left
	case *ILike:
		root.(*ILike).Left = left
	case *LessThan:
		root.(*LessThan).Left = left
	case *LessThanOrEqual:
		root.(*LessThanOrEqual).Left = left
	case *GreaterThan:
		root.(*GreaterThan).Left = left
	case *GreaterThanOrEqual:
		root.(*GreaterThanOrEqual).Left = left
	case *Equal:
		root.(*Equal).Left = left
	case *NotEqual:
		root.(*NotEqual).Left = left
	case *Add:
		root.(*Add).Left = left
	case *Subtract:
		root.(*Subtract).Left = left
	case *Divide:
		root.(*Divide).Left = left
	case *Before:
		root.(*Before).Left = left
	case *After:
		root.(*After).Left = left
	default:
		return errors.New("Could not attach left as root is not a binary operator")
	}
	return nil
}
