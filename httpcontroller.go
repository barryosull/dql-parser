package main

import (
	"net/http"
	"io"
	"github.com/davecgh/go-spew/spew"
	"parser/peg"
	"parser/ast"
	"fmt"
	"io/ioutil"
	"strings"
	"encoding/json"
)

func handleDqlStatement(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Handling request");

	if (r.Method!= "POST") {
		respondWithError("Invalid Request Method", w);
		return
	}

	dql, err := ioutil.ReadAll(r.Body)

	if (string(dql) == "") {
		respondWithError("No request body", w);
		return
	}

	astNode, err := peg.Parse("", dql)
	if err != nil {
		astString := strings.Replace(spew.Sdump(err),"\n","<br>",-1)
		respondWithError("Invalid DQL statement: "+astString, w);
		return
	}

	//io.WriteString(w, spew.Sdump(astNode))

	encoded, _ := json.Marshal(astNode)
	//fmt.Println(string(encoded));
	//io.WriteString(w, string(encoded))
	ast := new (ast.Ast);
	json.Unmarshal(encoded, ast);

	io.WriteString(w, spew.Sdump(ast))
}

func respondWithError(message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	io.WriteString(w, message)
}

func HttpListen() {
	http.HandleFunc("/", handleDqlStatement)
	http.ListenAndServe(":8000", nil)
}

func main() {
	HttpListen();
}