package ast

type Expression interface {
    Node
    expressionNode()
    String() string // Define the String() method for Expression
}
// Expression node method implementations
func (i *Identifier) expressionNode() {}
func (sl *StringLiteral) expressionNode() {}
func (hl *HashLiteral) expressionNode() {}
func (il *IntegerLiteral) expressionNode() {}
func (bl *BooleanLiteral) expressionNode() {}
func (al *ArrayLiteral) expressionNode() {}