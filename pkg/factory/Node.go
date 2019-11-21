// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package factory

import (
	"github.com/spatialcurrent/go-dfl/pkg/dfl"
)

type Node struct {
	name      string
	function  string
	attribute string
	value     interface{}
	left      dfl.Node
	right     dfl.Node
	items     []dfl.Item
	arguments []dfl.Node
	condition dfl.Node
	trueNode  dfl.Node
	falseNode dfl.Node
}

func NewNode(name string) *Node {
	return &Node{
		name: name,
	}
}

func (n *Node) Function(function string) *Node {
	n.function = function
	return n
}

func (n *Node) Attribute(attribute string) *Node {
	n.attribute = attribute
	return n
}

func (n *Node) Condition(condition dfl.Node) *Node {
	n.condition = condition
	return n
}

func (n *Node) True(node dfl.Node) *Node {
	n.trueNode = node
	return n
}

func (n *Node) False(node dfl.Node) *Node {
	n.falseNode = node
	return n
}

func (n *Node) Value(value interface{}) *Node {
	n.value = value
	return n
}

func (n *Node) Left(left dfl.Node) *Node {
	n.left = left
	return n
}

func (n *Node) Right(right dfl.Node) *Node {
	n.right = right
	return n
}

func (n *Node) Items(items []dfl.Item) *Node {
	n.items = items
	return n
}

func (n *Node) Arguments(arguments []dfl.Node) *Node {
	n.arguments = arguments
	return n
}

func (n *Node) Node() dfl.Node {
	switch n.name {
	case "add", "+":
		return &dfl.Add{BinaryOperator: &dfl.BinaryOperator{Left: n.left, Right: n.right}}
	case "subtract", "-":
		return &dfl.Subtract{
			NumericBinaryOperator: &dfl.NumericBinaryOperator{
				BinaryOperator: &dfl.BinaryOperator{
					Left:  n.left,
					Right: n.right,
				},
			},
		}
	case ">":
		return &dfl.GreaterThan{
			NumericBinaryOperator: &dfl.NumericBinaryOperator{
				BinaryOperator: &dfl.BinaryOperator{
					Left:  n.left,
					Right: n.right,
				},
			},
		}
	case ">=":
		return &dfl.GreaterThanOrEqual{
			NumericBinaryOperator: &dfl.NumericBinaryOperator{
				BinaryOperator: &dfl.BinaryOperator{
					Left:  n.left,
					Right: n.right,
				},
			},
		}
	case "<":
		return &dfl.LessThan{
			NumericBinaryOperator: &dfl.NumericBinaryOperator{
				BinaryOperator: &dfl.BinaryOperator{
					Left:  n.left,
					Right: n.right,
				},
			},
		}
	case "<=":
		return &dfl.LessThanOrEqual{
			NumericBinaryOperator: &dfl.NumericBinaryOperator{
				BinaryOperator: &dfl.BinaryOperator{
					Left:  n.left,
					Right: n.right,
				},
			},
		}
	case "ternary":
		return &dfl.TernaryOperator{
			Left:  n.condition,
			True:  n.trueNode,
			False: n.falseNode,
		}
	case "attribute":
		return &dfl.Attribute{Name: n.attribute}
	case "dictionary":
		return &dfl.Dictionary{Items: n.items}
	case "item":
		return &dfl.Item{Key: n.left, Value: n.right}
	case "literal":
		return &dfl.Literal{Value: n.value}
	case "null", "nil":
		return &dfl.Null{}
	case "function":
		return &dfl.Function{
			Name: n.function,
			MultiOperator: &dfl.MultiOperator{
				Arguments: n.arguments,
			},
		}
	}
	return nil
}
