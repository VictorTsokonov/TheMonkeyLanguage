package ast

type Node interface {
	TokenLiteral() string // returns the literal value of the token it's associated with
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct { // root node of every AST the parser produces
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// continuing to read about parsers so no coding progress today,
// but I will still commit cuz I'm doing the work guys!
// p.s. also don't wanna lose my commit streak! :)
// still reading, STILL READING
// still reading
// damn, still reading...
// still reading, ye
// still reading
// still reading
// still reading
