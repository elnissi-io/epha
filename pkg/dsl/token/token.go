package token

type TokenType string

const (
    EOF         TokenType = "EOF"
    Identifier  TokenType = "IDENTIFIER"
    Assign      TokenType = "ASSIGN"
    String      TokenType = "STRING"
    Helm        TokenType = "HELM"
    From        TokenType = "FROM"
    CRD         TokenType = "CRD"
    Import      TokenType = "IMPORT"
    As          TokenType = "AS"
    LParen      TokenType = "("
    RParen      TokenType = ")"
    LBrace      TokenType = "{"
    RBrace      TokenType = "}"
    Comma       TokenType = ","
    Illegal     TokenType = "ILLEGAL"
	Colon 		TokenType = ":"
)

type Token struct {
    Type    TokenType
    Literal string
}
