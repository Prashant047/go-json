package main

import (
	"fmt"
	"go-json/ast"
	"regexp"
	"strconv"
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
			return "\"" + n.Value + "\""
		case ast.BoolNode:
			n := node.(ast.BoolNode)
			return n.Token.Value
		case ast.NullNode:
			n := node.(ast.NullNode)
			return n.Token.Value
		case ast.ObjectNode:
			n := node.(ast.ObjectNode)
			mapLength := len(n.Items)
			i := 0

			res += "{\n"
			for key, value := range n.Items {
				res += strings.Repeat(indentChar, tab+1) + "\"" + key + "\"" + ": "
				res += toString(value, tab+1)
				if i != mapLength-1 {
					res += ","
				}
				res += "\n"
				i += 1
			}
			res += strings.Repeat(indentChar, tab) + "}"
		case ast.ArrayNode:
			n := node.(ast.ArrayNode)
			res += "[\n"
			for i, item := range n.Items {
				res += strings.Repeat(indentChar, tab+1) + toString(item, tab+1)
				if i != len(n.Items)-1 {
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
		for _, value := range n.Items {
			TraverseAst(value)
		}
	case ast.ArrayNode:
		n := node.(ast.ArrayNode)
		for _, item := range n.Items {
			TraverseAst(item)
		}
	}
}

func Select(rootNode ast.Node, query string) (ast.Node, error) {
	pattern := `^((\w+|\[\d+\])\.)*(\w+|\[\d+\])$`
	matched, err := regexp.MatchString(pattern, query)
	if err != nil {
		return nil, err
	}

	if !matched {
		return nil, fmt.Errorf("invalid query string: %v", query)
	}

	keys := strings.Split(query, ".")
	currNode := rootNode

	for _, q := range keys {
		if q[0] != '[' {
			node, ok := currNode.(ast.ObjectNode)
			if !ok {
        return nil, fmt.Errorf("invalid query: %v ; %v doesn't exist", query, q)
			}
			val, ok := node.Items[q]
			if !ok {
        return nil, fmt.Errorf("invalid query: %v ; key %v doesn't exist", query, q)
			}
			currNode = val
		} else {
			node, ok := currNode.(ast.ArrayNode)
			if !ok {
				return nil, fmt.Errorf("%v is not an array", q)
			}
			index, err := strconv.Atoi(q[1 : len(q)-1])
			if err != nil {
				return nil, fmt.Errorf("error converting integer to index: %v", err)
			}
			if index < 0 || index >= len(node.Items) {
				return nil, fmt.Errorf("index %v out of range of %v", index, len(node.Items))
			}
			val := node.Items[index]
			currNode = val
		}
	}
	return currNode, nil
}
