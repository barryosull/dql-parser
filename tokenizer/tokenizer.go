package tokenizer

type Tokenizer interface {
	Tokens() ([]Token, *Token)
	Next() (*Token, *Token)
}

type tokeniser struct {
	l *lexer
	index int
}

func (t *tokeniser) Tokens () ([]Token, *Token) {
	return t.l.tokens, t.l.error;
}

func (t *tokeniser) Next() (*Token, *Token) {
	if (t.index >= len(t.l.tokens)) {
		return nil, t.l.error
	}
	token := t.l.tokens[t.index];
	t.index++;
	return &token, nil;
}

func NewTokenizer(dql string) Tokenizer {
	l := lex("DQL", dql);
	return &tokeniser{l, 0};
}


