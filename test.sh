../../bin/pigeon $1/parser.peg | ../../bin/goimports > $1/parser.go
echo "Created $1 parser";

go test $1/parser_test.go $1/parser.go