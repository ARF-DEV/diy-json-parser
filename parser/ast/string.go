package ast

import (
	"fmt"

	"github.com/ARF-DEV/diy-json-parser/lexer"
)

type String struct {
	Token lexer.Token
	Value string
}

func (s *String) TokenValue() lexer.Value {
	return s.Token.Value
}

func (s *String) String() string {
	return fmt.Sprintf("\"%v\"", s.Value)
}
