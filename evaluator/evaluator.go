package evaluator

import (
	"github.com/AsyncEngineering/nomad/ast"
	"github.com/AsyncEngineering/nomad/object"
)

func Eval(node ast.Node) object.Object {

	switch node := node.(type) {

	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.NullLiteral:
		return &object.Null{}

	case *ast.NoneLiteral:
		return &object.None{}

	case *ast.Boolean:
		return mapBooleans(node.Value)

	default:
		return nil
	}

}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}

func mapBooleans(val bool) *object.Boolean {
	if val {
		return TRUE
	}
	return FALSE
}

var (
	NONE  = &object.None{}
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)
