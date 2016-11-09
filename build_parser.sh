../../bin/pigeon peg/parser.peg | ../../bin/goimports > peg/parser.go
echo "Created parser";

../../bin/pigeon expression/parser.peg | ../../bin/goimports > expression/parser.go
echo "Created expression parser";

