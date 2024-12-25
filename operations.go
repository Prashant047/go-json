package main

import (
	"fmt"
	"go-json/ast"
  "strings"
)

func AstToString(node ast.Node, indentChar string) string {

  var toString func(node ast.Node, tab int) string
  toString = func(node ast.Node, tab int) string {
    res := ""

    switch node.(type) {
    case ast.NumberNode:
      n := node.(ast.NumberNode)
      return n.Token.Value
    case ast.StrNode:
      n := node.(ast.StrNode)
      return "\""+n.Value+"\""
    case ast.BoolNode:
      n := node.(ast.BoolNode)
      return n.Token.Value
    case ast.NullNode:
      n := node.(ast.NullNode)
      return n.Token.Value
    case ast.ObjectNode:
      n := node.(ast.ObjectNode)
      res += "{\n"
      for i, item := range n.Items {
        res += strings.Repeat(indentChar, tab+1) + "\"" + item.Key + "\"" + ": "
        res += toString(item.Val, tab+1)
        if i != len(n.Items)-1{
          res += ","
        }
        res += "\n"
      }
      res += strings.Repeat(indentChar, tab) + "}"
    case ast.ArrayNode:
      n := node.(ast.ArrayNode)
      res += "[\n"
      for i, item := range n.Items {
        res += strings.Repeat(indentChar, tab+1) + toString(item, tab+1)
        if i != len(n.Items)-1{
          res += ","
        }
        res += "\n"
      }
      res += strings.Repeat(indentChar, tab) + "]"
    }
    return res
  }

  return toString(node, 0)
}

func TraverseAst(node ast.Node) {
	switch node.(type) {
	case ast.NumberNode:
		n := node.(ast.NumberNode)
		fmt.Println(n.Value)
	case ast.StrNode:
		n := node.(ast.StrNode)
		fmt.Println(n.Value)
	case ast.BoolNode:
		n := node.(ast.BoolNode)
		fmt.Println(n.Value)
	case ast.NullNode:
		n := node.(ast.NullNode)
		fmt.Println(n.Token.Value)
	case ast.ObjectNode:
		n := node.(ast.ObjectNode)
		for _, item := range n.Items {
			TraverseAst(item.Val)
		}
	case ast.ArrayNode:
		n := node.(ast.ArrayNode)
		for _, item := range n.Items {
			TraverseAst(item)
		}
	}
}

func Select(rootNode ast.Node, query string) ast.Node {
  // TODO: Implement
  return nil
}

