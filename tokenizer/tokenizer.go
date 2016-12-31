package tokenizer

import (
	tok "parser/token"
)

type Tokenizer interface {
	Tokens() ([]tok.Token, *tok.Error)
	Next() (*tok.Token, *tok.Error)
}

type tokeniser struct {
	l *lexer
	index int
}

func (t *tokeniser) Tokens () ([]tok.Token, *tok.Error) {
	return t.l.tokens, t.l.error;
}

func (t *tokeniser) Next() (*tok.Token, *tok.Error) {
	if (t.index >= len(t.l.tokens)) {
		return nil, t.l.error
	}
	token := t.l.tokens[t.index];
	t.index++;
	return &token, t.l.error;
}

func NewTokenizer(dql string) Tokenizer {
	l := lex("DQL", dql);
	return &tokeniser{l, 0};
}


