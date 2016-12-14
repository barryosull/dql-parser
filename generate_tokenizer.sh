
# Create buildable tokenizer for testing

addPackageHeader() {
    paths=(${1//\// })

    folder=${paths[${#paths[@]} - 1]}

    # Give it the correct header
    echo "
{
    package parser

    func emit(typ TokenType, val interface{}) (interface{}, error) {
    	GetInstanceTokenList().Append(NewToken(typ, val.(string)));
    	return nil, nil;
    }

}
    "|cat - $1/tokenizer.peg > $1/tokenizer_subset.peg
}

addDependencies() {
    if [ -f $2/dependencies.csv ]; then
        dependenciesline=$(head -n 1 $2/dependencies.csv)
        dependencies=$(echo $dependenciesline | tr ";" "\n")

        for dep in $dependencies; do
            cat "$2/$dep/tokenizer.peg" >> $1/tokenizer_subset.peg;
            echo "Adding deps for $2/$dep"
            addDependencies "$1" "$2/$dep"
        done
    fi
}

buildTokenizer() {
    ../../bin/pigeon $1/tokenizer_subset.peg | ../../bin/goimports > tokenizer_generated.go
    rm $1/tokenizer_subset.peg
    echo "Created $1 test tokenizer";
}

addPackageHeader peg/namespace
addDependencies peg/namespace peg/namespace
buildTokenizer peg/namespace






