
# Create buildable parser for testing

addPackageHeader() {
    paths=(${1//\// })

    folder=${paths[${#paths[@]} - 1]}

    # Give it the correct header
    echo "
{
    package $folder
}
    "|cat - $1/parser.peg > $1/parser_subset.peg
}

addDependencies() {
    if [ -f $2/dependencies.csv ]; then
        dependenciesline=$(head -n 1 $2/dependencies.csv)
        dependencies=$(echo $dependenciesline | tr ";" "\n")

        for dep in $dependencies; do
            cat "$2/$dep/parser.peg" >> $1/parser_subset.peg;
            echo "Adding deps for $2/$dep"
            addDependencies "$1" "$2/$dep"
        done
    fi
}

buildParser() {
    ../../bin/pigeon $1/parser_subset.peg | ../../bin/goimports > $1/parser.go
    # rm $1/parser_subset.peg
    echo "Created $1 test parser";
}

test() {
    go test $1/parser_test.go $1/parser.go
}

addPackageHeader $1
addDependencies $1 $1
buildParser $1
test $1






