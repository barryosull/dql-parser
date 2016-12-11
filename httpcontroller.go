package parser

/*
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
	"bytes"
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

	encoded, _ := json.Marshal(astNode)

	decoded := new (ast.Node);
	json.Unmarshal(encoded, decoded);

	io.WriteString(w, string(encoded))

	reEncoded, _ := json.Marshal(astNode);
	if (bytes.Compare(reEncoded, encoded) == 0) {
		io.WriteString(w, "Encoding and decoding ASTs worked.")
	} else {
		io.WriteString(w, "Encoded ASTs are not the same")
	}
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
*/