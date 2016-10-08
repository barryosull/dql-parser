package astrepo

import "parser/ast"

type Repo interface {
	Store(node *ast.Expression)
	Fetch() *ast.Expression
	FetchHandlers() *ast.Expression
}
