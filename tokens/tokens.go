package tokens

import "fmt"

type TokenType string

const (
	LEFT_CURLY    TokenType = "LEFT_CURLY"
	RIGHT_CURLY             = "RIGHT_CURLY"
	RIGHT_BRACKET           = "RIGHT_BRACKET"
	LEFT_BRACKET            = "LEFT_BRACKET"
	COMMA                   = "COMMA"
	COLON                   = "COLON"
	STRING                  = "STRING"
	NUMBER                  = "NUMBER"
	TRUE                    = "TRUE"
	FALSE                   = "FALSE"
	NULL                    = "NULL"
	EOF                     = "EOF"
)

type Token struct {
	TokenType TokenType
	Value     string
}

func (t Token) String() string {
	return fmt.Sprintf("TOKEN<| %v, %v |>", t.TokenType, t.Value)
  // return fmt.Sprintf("")
}
