
# Create buildable tokenizer for testing

addPackageHeader() {
    paths=(${1//\// })

    folder=${paths[${#paths[@]} - 1]}

    # Give it the correct header
    echo "
{
    package $folder
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
    ../../bin/pigeon $1/tokenizer_subset.peg | ../../bin/goimports > $1/tokenizer_generated.go
    // rm $1/tokenizer_subset.peg
    echo "Created $1 test tokenizer";
}

test() {
    go test $1/tokenizer_test.go $1/tokenizer_generated.go
}

addPackageHeader peg/$1
addDependencies peg/$1 peg/$1
buildTokenizer peg/$1
test peg/$1






