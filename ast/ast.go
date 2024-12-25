package ast

import (
	"go-json/tokens"
)

type Node interface {
	TokenLiteral() string
}

type Str interface {
	Node
	stringNode()
}

type Number interface {
	Node
	numberNode()
}

type Bool interface {
  Node
  boolNode()
}

type Null interface {
  Node
  nullNode()
}

type Object interface {
	Node
	objectNode()
}

type Array interface {
	Node
	arrayNode()
}

type ObjectNode struct {
	Token tokens.Token
	Items map[string]Node
}

func (on *ObjectNode) objectNode() {}
func (on ObjectNode) TokenLiteral() string {
	return on.Token.Value
}

type ArrayNode struct {
	Token tokens.Token
	Items []Node
}
func(an *ArrayNode) arrayNode(){}
func(an ArrayNode) TokenLiteral() string {
  return an.Token.Value
}

type StrNode struct {
  Token tokens.Token
  Value string
}
func (sn *StrNode) stringNode(){}
func (sn StrNode) TokenLiteral() string {
  return sn.Value
}

type NumberNode struct {
  Token tokens.Token
  Value float64
}
func (nn *NumberNode) numberNode(){}
func (nn NumberNode) TokenLiteral() string {
  return nn.Token.Value
}

type BoolNode struct {
  Token tokens.Token
  Value bool
}
func (nn *BoolNode) boolNode(){}
func (nn BoolNode) TokenLiteral() string {
  return nn.Token.Value
}

type NullNode struct {
  Token tokens.Token
}
func (bn *NullNode) nullNode(){}
func (bn NullNode) TokenLiteral() string {
  return bn.Token.Value
}
