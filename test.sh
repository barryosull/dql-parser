../../bin/pigeon class/parser.peg | ../../bin/goimports > class/parser.go
echo "Created class parser";

go test class/parser_test.go class/parser.go