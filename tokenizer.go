package parser

type Tokenizer interface {
	Tokenize(dql string) ([]*Token)
}

type peg struct {
	tokenList *TokenList
}

type IdGenerator interface {
	Generate() string
}

func (p *peg) Tokenize (dql string) ([]*Token) {
	Parse("", []byte(dql));
	return p.tokenList.All();
}

func (p *peg) ReferenceError(commandID string, reference string) string {
	return "";
}

func NewTokenizer() Tokenizer {
	return &peg{GetInstanceTokenList()};
}


