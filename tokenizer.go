package parser

type Tokenizer interface {
	Tokenize(dql string) ([]Token)
}

type lexer struct {

}

func (p *lexer) Tokenize (dql string) ([]Token) {
	l := lex("DQL", dql);
	return l.tokens;
}

func NewTokenizer() Tokenizer {
	return &lexer{};
}


