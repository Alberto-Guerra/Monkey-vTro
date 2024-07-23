package lexer

import (
	"testing"

	"github.com/Alberto-Guerra/Monkey-vTro/token"
)

func TestNextToken(t *testing.T) {

	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	l := New(input)

	for _, test := range tests {
		token := l.NextToken()
		if token.Type != test.expectedType {
			t.Errorf("Expected type %s, got %s", test.expectedType, token.Type)
		}
		if token.Literal != test.expectedLiteral {
			t.Errorf("Expected literal %s, got %s", test.expectedLiteral, token.Literal)
		}
	}
}
