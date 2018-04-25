package dfl

import (
	"fmt"
)

import (
	"github.com/pkg/errors"
)

func ParseLiteral(v interface{}, remainder string) (Node, error) {

	if len(remainder) == 0 {
		return &Literal{Value: v}, nil
	}

	left := &Literal{Value: v}
	root, err := Parse(remainder)
	if err != nil {
		return root, err
	}
	switch root.(type) {
	case *And:
		root.(*And).Left = left
	case *Or:
		root.(*Or).Left = left
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
	default:
		return root, errors.New("Invalid expression syntax for " + fmt.Sprint(v) + ".  Root is not a binary operator")
	}
	return root, nil
}
