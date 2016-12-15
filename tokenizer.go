package parser

type Tokenizer interface {
	Tokenize(dql string) ([]Token)
}

type tokeniser struct {

}

func (t *tokeniser) Tokenize (dql string) ([]Token) {
	l := lex("DQL", dql);
	return l.tokens;
}

func NewTokenizer() Tokenizer {
	return &tokeniser{};
}


