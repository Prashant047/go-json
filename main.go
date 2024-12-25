package main

import (
	"bufio"
	"fmt"
	"go-json/lexer"
	"log"
	// "strconv"

	// "go-json/tokens"

	"go-json/parser"
	"os"
)

func main() {

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

	node, err := Select(ast, "e.f.[0]")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(AstToString(node, "    "))

}
