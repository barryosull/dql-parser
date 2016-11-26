
# Create buildable parser for testing
paths=(${1//\// })

folder=${paths[${#paths[@]} - 1]}

echo "
{
    package $folder
}
"|cat - $1/parser.peg > $1/parser_subset.peg

# Add the dependencies
if [ -f $1/dependencies.csv ]; then
    dependenciesline=$(head -n 1 $1/dependencies.csv)
    dependencies=$(echo $dependenciesline | tr ";" "\n")

    for dep in $dependencies; do
        cat "$1/$dep/parser.peg" >> $1/parser_subset.peg;
    done

else
    echo "No dependencies found"
fi

# Build the parser
../../bin/pigeon $1/parser_subset.peg | ../../bin/goimports > $1/parser_subset.go
//rm $1/parser_subset.peg

echo "Created $1 test parser";

# Test the parser
go test $1/parser_test.go $1/parser_subset.go

# Destroy the test file, its not needed
rm $1/parser_subset.go

