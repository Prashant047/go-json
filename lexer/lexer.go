package lexer

import (
	"go-json/tokens"
	"log"
  "fmt"
)

type Lexer struct {
	src      string
	currPos  int
	currChar byte
}

func NewLexer(src string) *Lexer {
	return &Lexer{
		src:      src,
		currPos:  0,
		currChar: src[0],
	}
}

func (l *Lexer) advance() {
	l.currPos += 1
	if l.currPos >= len(l.src) {
		l.currChar = 0
	} else {
		l.currChar = l.src[l.currPos]
	}
}

func (l *Lexer) error(msg string) {
  log.Fatalf("LEXER ERROR: %v", msg)
}

func (l *Lexer) parseDigit() {
  if !isDigit(l.currChar) {
    l.error(fmt.Sprintf("expected digit received %v", l.currChar))
    return
  }

  for isDigit(l.currChar) {
    l.advance()
  }
}

func (l *Lexer) parseFraction() {
  if l.currChar != '.' {
    return
  }

  l.advance()
  l.parseDigit()
}

func (l *Lexer) parseNumberExpo() {
  if !(l.currChar == 'E' || l.currChar == 'e') {
    return
  }

  l.advance()
  if l.currChar == '-' || l.currChar == '+' {
    l.advance()
  }
  l.parseDigit()
}

func (l *Lexer) handleNumber() string {
	start := l.currPos
  if l.currChar == '-' {
    l.advance()
  }

  if l.currChar == '0' {
    l.advance()
    l.parseFraction()
    l.parseNumberExpo()
    if isDigit(l.currChar) {
      l.error("numbers can't have trailing zeroes")
    }
    end := l.currPos
    return l.src[start:end]
  }

  l.parseDigit()
  l.parseFraction()
  l.parseNumberExpo()

	end := l.currPos
	return l.src[start:end]
}

func (l *Lexer) handleString() string {
	l.advance()
	start := l.currPos

	for l.currChar != '"' {
		l.advance()
	}
	end := l.currPos
	literal := l.src[start:end]
	l.advance()

	return literal
}

func (l *Lexer) handleAlpha() string {
	start := l.currPos
	for isAlpha(l.currChar) {
		l.advance()
	}
	end := l.currPos
	return l.src[start:end]
}

func (l *Lexer) handleWhiteSpaces() {
	for l.currChar == '\n' || l.currChar == '\t' || l.currChar == ' ' {
		l.advance()
	}
}

func (l *Lexer) GetNextToken() tokens.Token {
	l.handleWhiteSpaces()
	if l.currPos < len(l.src) {
    var tok tokens.Token

		switch l.currChar {
		case '{':
			tok = tokens.Token{TokenType: tokens.LEFT_CURLY, Value: "{"}
			l.advance()
		case '}':
			tok = tokens.Token{TokenType: tokens.RIGHT_CURLY, Value: "}"}
			l.advance()
		case '[':
			tok = tokens.Token{TokenType: tokens.LEFT_BRACKET, Value: "["}
			l.advance()
		case ']':
			tok = tokens.Token{TokenType: tokens.RIGHT_BRACKET, Value: "]"}
			l.advance()
		case ',':
			tok = tokens.Token{TokenType: tokens.COMMA, Value: ","}
			l.advance()
		case ':':
			tok = tokens.Token{TokenType: tokens.COLON, Value: ":"}
			l.advance()
		default:
			if isDigit(l.currChar) || l.currChar == '-' {
				literal := l.handleNumber()
				tok = tokens.Token{TokenType: tokens.NUMBER, Value: literal}
			} else if isAlpha(l.currChar) {
				literal := l.handleAlpha()
				switch literal {
				case "true":
					tok =  tokens.Token{TokenType: tokens.TRUE, Value: "true"}
				case "false":
					tok =  tokens.Token{TokenType: tokens.FALSE, Value: "false"}
				case "null":
					tok =  tokens.Token{TokenType: tokens.NULL, Value: "null"}
				default:
					l.error(fmt.Sprintf("invalid keyword: %v", literal))
				}
			} else if l.currChar == '"' {
				literal := l.handleString()
				tok = tokens.Token{TokenType: tokens.STRING, Value: literal}
			} else {
        l.error(fmt.Sprintf("invalid token: %v", string(l.currChar)))
			}
		}
    return tok
	}

	return tokens.Token{TokenType: tokens.EOF, Value: ""}
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isAlpha(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}
