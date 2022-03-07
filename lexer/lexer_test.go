package lexer

import (
	"one/token"
	"testing"
)

type TestCase struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextToken(t *testing.T) {
	lparen := TestCase{token.LPAREN, "("}
	rparen := TestCase{token.RPAREN, ")"}
	lbrace := TestCase{token.LBRACE, "{"}
	rbrace := TestCase{token.RBRACE, "}"}
	lbracket := TestCase{token.LBRACKET, "["}
	rbracket := TestCase{token.RBRACKET, "]"}
	comma := TestCase{token.COMMA, ","}
	semicolon := TestCase{token.SEMICOLON, ";"}
	eof := TestCase{token.EOF, ""}
	plus := TestCase{token.PLUS, "+"}
	minus := TestCase{token.MINUS, "-"}
	bang := TestCase{token.BANG, "!"}
	slash := TestCase{token.SLASH, "/"}
	asterisk := TestCase{token.ASTERISK, "*"}
	lt := TestCase{token.LT, "<"}
	gt := TestCase{token.GT, ">"}
	eq := TestCase{token.EQ, "=="}
	noteq := TestCase{token.NOT_EQ, "!="}
	let := TestCase{token.LET, "let"}
	assign := TestCase{token.ASSIGN, "="}
	function := TestCase{token.FUNCTION, "fn"}
	true_ := TestCase{token.TRUE, "true"}
	false_ := TestCase{token.FALSE, "false"}
	if_ := TestCase{token.IF, "if"}
	else_ := TestCase{token.ELSE, "else"}
	return_ := TestCase{token.RETURN, "return"}
	colon := TestCase{token.COLON, ":"}

	input := `
let five = 5;
let ten = 10;
let add = fn(x, y) {
  x + y;
};
let result = add(five, ten);
!-/*5;
5 < 10 > 5;
if (5 < 10) {
  return true;
} else {
  return false;
}

10 == 10;
10 != 9;
"foobar"
"foo bar"
[1, 2];
{"foo": "bar"}
`
	tests := []TestCase{
		let,
		{token.IDENT, "five"},
		assign,
		{token.INT, "5"},
		semicolon,
		let,
		{token.IDENT, "ten"},
		assign,
		{token.INT, "10"},
		semicolon,
		let,
		{token.IDENT, "add"},
		assign,
		function,
		lparen,
		{token.IDENT, "x"},
		comma,
		{token.IDENT, "y"},
		rparen,
		lbrace,
		{token.IDENT, "x"},
		plus,
		{token.IDENT, "y"},
		semicolon,
		rbrace,
		semicolon,
		let,
		{token.IDENT, "result"},
		assign,
		{token.IDENT, "add"},
		lparen,
		{token.IDENT, "five"},
		comma,
		{token.IDENT, "ten"},
		rparen,
		semicolon,
		bang,
		minus,
		slash,
		asterisk,
		{token.INT, "5"},
		semicolon,
		{token.INT, "5"},
		lt,
		{token.INT, "10"},
		gt,
		{token.INT, "5"},
		semicolon,
		if_,
		lparen,
		{token.INT, "5"},
		lt,
		{token.INT, "10"},
		rparen,
		lbrace,
		return_,
		true_,
		semicolon,
		rbrace,
		else_,
		lbrace,
		return_,
		false_,
		semicolon,
		rbrace,
		{token.INT, "10"},
		eq,
		{token.INT, "10"},
		semicolon,
		{token.INT, "10"},
		noteq,
		{token.INT, "9"},
		semicolon,
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		lbracket,
		{token.INT, "1"},
		comma,
		{token.INT, "2"},
		rbracket,
		semicolon,
		lbrace,
		{token.STRING, "foo"},
		colon,
		{token.STRING, "bar"},
		rbrace,
		eof,
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
