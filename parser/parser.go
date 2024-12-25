package parser

import (
	"fmt"
	"log"
	"strconv"

	"go-json/ast"
	"go-json/lexer"
	"go-json/tokens"
)

type Parser struct {
	currToken tokens.Token
	lexer     *lexer.Lexer
}

func NewParser(lexer *lexer.Lexer) *Parser {
	return &Parser{
		lexer: lexer,
	}
}

func (p *Parser) error(msg string) {
	log.Fatalf("PARSER ERROR: %v", msg)
}

func (p *Parser) expect(tokenType tokens.TokenType) {
	if p.currToken.TokenType != tokenType {
		p.error(fmt.Sprintf("expected token %v received %v", tokenType, p.currToken.TokenType))
	}
}

func (p *Parser) eat(tokenType tokens.TokenType) {
	p.expect(tokenType)
	p.currToken = p.lexer.GetNextToken()
}

func (p *Parser) parseValue() ast.Node {
	switch p.currToken.TokenType {
	case tokens.STRING:
		node := ast.StrNode{Token: p.currToken, Value: p.currToken.Value}
		p.eat(tokens.STRING)
		return node
	case tokens.NUMBER:
		numberVal, err := strconv.ParseFloat(p.currToken.Value, 64)
		if err != nil {
      p.error(fmt.Sprintf("error converting string to number: %v", err))
			return nil
		}
		node := ast.NumberNode{Token: p.currToken, Value: numberVal}
		p.eat(tokens.NUMBER)
		return node
	case tokens.TRUE:
		boolNode := ast.BoolNode{Token: p.currToken, Value: true}
		p.eat(tokens.TRUE)
		return boolNode
	case tokens.FALSE:
		boolNode := ast.BoolNode{Token: p.currToken, Value: false}
		p.eat(tokens.FALSE)
		return boolNode
	case tokens.NULL:
		node := ast.NullNode{Token: p.currToken}
		p.eat(tokens.NULL)
		return node
	case tokens.LEFT_CURLY:
		return p.parseObject()
	case tokens.LEFT_BRACKET:
		return p.parseArray()
	default:
		p.error(fmt.Sprintf("unexpected value token: %v", p.currToken))
		return nil
	}
}

func (p *Parser) parseArray() ast.ArrayNode {
	arrayNode := ast.ArrayNode{
		Token: p.currToken,
		Items: make([]ast.Node, 0),
	}

	p.eat(tokens.LEFT_BRACKET)

	if p.currToken.TokenType == tokens.RIGHT_BRACKET {
		return arrayNode
	}

	arrayNode.Items = append(arrayNode.Items, p.parseValue())
	for p.currToken.TokenType == tokens.COMMA {
		p.eat(tokens.COMMA)
		arrayNode.Items = append(arrayNode.Items, p.parseValue())
	}

	p.eat(tokens.RIGHT_BRACKET)
	return arrayNode
}

func (p *Parser) parseObject() ast.ObjectNode {
	objectNode := ast.ObjectNode{
		Token: p.currToken,
		Items: make([]ast.KeyVal, 0),
	}

	p.eat(tokens.LEFT_CURLY)

	if p.currToken.TokenType == tokens.RIGHT_CURLY {
		return objectNode
	}

	objItem := ast.KeyVal{}

	p.expect(tokens.STRING)
	objItem.Key = p.currToken.Value
	p.eat(tokens.STRING)
	p.eat(tokens.COLON)

	objItem.Val = p.parseValue()
	objectNode.Items = append(objectNode.Items, objItem)

	for p.currToken.TokenType == tokens.COMMA {
		p.eat(tokens.COMMA)

		objItem = ast.KeyVal{}

		p.expect(tokens.STRING)
		objItem.Key = p.currToken.Value

		p.eat(tokens.STRING)
		p.eat(tokens.COLON)

		objItem.Val = p.parseValue()
		objectNode.Items = append(objectNode.Items, objItem)
	}

	p.eat(tokens.RIGHT_CURLY)
	return objectNode
}

func (p *Parser) Parse() ast.Node {
	p.currToken = p.lexer.GetNextToken()
	switch p.currToken.TokenType {
	case tokens.LEFT_CURLY:
		return p.parseObject()
	case tokens.LEFT_BRACKET:
		return p.parseArray()
	default:
		p.error(fmt.Sprintf("unexpected token %v", p.currToken))
		return nil
	}
}
