package parser

type Tokenizer interface {
	Tokenize(dql string) ([]Token, *Token)
}

type tokeniser struct {

}

func (t *tokeniser) Tokenize (dql string) ([]Token, *Token) {
	l := lex("DQL", dql);
	return l.tokens, l.error;
}

func NewTokenizer() Tokenizer {
	return &tokeniser{};
}


