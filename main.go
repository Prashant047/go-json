package main

import (
	"bufio"
	"fmt"
	"go-json/lexer"
	"regexp"
	"go-json/parser"
	"os"
)

var helpString string = `operations:
  select : go-json select -q "a.b.[0]" -f some.json
         : go-json select -q "a.b.[0]" <read_from_stdin>
  format : go-json format -i "----" -f some.json
         : go-json format -i "----" <read_from_stdin>
`

func parseArgs(args []string) (map[string]string, error) {
	argsMap := make(map[string]string)
	if len(args) == 0 {
		return argsMap, nil
	}

	i := 0
	for i < len(args) {
		key := args[i]
		matched, err := regexp.MatchString(`^-\w$`, key)
		if err != nil {
			return nil, err
		}
		if !matched {
			return nil, fmt.Errorf("unidetified format for args key: %v", key)
		}

		if i+1 >= len(args) {
			return nil, fmt.Errorf("no value provided for %v", key)
		}
		argsMap[key] = args[i+1]
		i += 2
	}

	return argsMap, nil
}

func execFormat(cliArgs []string) (string, error) {
	args, err := parseArgs(cliArgs[2:])
	if err != nil {
		return "", err
	}

	indentString, ok := args["-t"]
	if !ok {
		indentString = "    "
	}

	fileName, ok := args["-f"]
	if !ok {
		fileName = ""
	}

	file := ""
	if len(fileName) != 0 {
		data, err := os.ReadFile(fileName)
		if err != nil {
			return "", fmt.Errorf("error reading file %v : %v\n", fileName, err)
		}
		file = string(data)
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		src := ""
		for scanner.Scan() {
			line := scanner.Text()
			src += line + string('\n')
		}
		file = src
	}

	lex := lexer.NewLexer(file)
	parser := parser.NewParser(lex)
	ast := parser.Parse()

	return AstToString(ast, indentString), nil
}

func execSelect(cliArgs []string) (string, error) {
	args, err := parseArgs(cliArgs[2:])
	if err != nil {
		return "", err
	}

	query, ok := args["-q"]
	if !ok {
		return "", fmt.Errorf("query string not found")
	}

	fileName, ok := args["-f"]
	if !ok {
		fileName = ""
	}

	file := ""
	if len(fileName) != 0 {
		data, err := os.ReadFile(fileName)
		if err != nil {
			return "", fmt.Errorf("error reading file %v : %v", fileName, err)
		}
		file = string(data)
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		src := ""

		for scanner.Scan() {
			line := scanner.Text()
			src += line + string('\n')
		}
		file = src
	}

	lex := lexer.NewLexer(file)
	parser := parser.NewParser(lex)
	ast := parser.Parse()

	node, err := Select(ast, query)
	if err != nil {
		return "", err
	}

	return AstToString(node, "    "), nil
}

func main() {

	cliArgs := os.Args
	if len(cliArgs) < 2 {
		fmt.Println("no operation specified")
		os.Exit(1)
	}

	switch cliArgs[1] {
	case "help":
		fmt.Println(helpString)
	case "-h":
		fmt.Println(helpString)
	case "format":
		output, err := execFormat(cliArgs)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(output)
	case "select":
		output, err := execSelect(cliArgs)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(output)
	default:
		fmt.Println("invalid operation")
		os.Exit(1)
	}

}
