package main

import (
	"bufio"
	"fmt"
	"go-json/lexer"
	// "strconv"

	// "go-json/tokens"

	"go-json/parser"
	"os"
)

func main() {

  // TODO: IMPLEMENT FIELD SELECTION
  // a.b.c[0]
  // a.b.c[0].d

	scanner := bufio.NewScanner(os.Stdin)
	src := ""

	for scanner.Scan() {
		line := scanner.Text()
		src += line + string('\n')
	}

	lex := lexer.NewLexer(src)
	parser := parser.NewParser(lex)
	ast := parser.Parse()

	fmt.Println(AstToString(ast, "    "))

	// number, err := strconv.ParseFloat("2e2000", 64)
	// if err != nil {
	//   fmt.Println(err)
	// }
	// fmt.Println(number)

	// s, _ := json.MarshalIndent(ast, "", "  ")
	// fmt.Println(string(s))

	// tok := lex.GetNextToken()
	// for tok.TokenType != tokens.EOF {
	//   fmt.Println(tok)
	//   tok = lex.GetNextToken()
	// }

}
