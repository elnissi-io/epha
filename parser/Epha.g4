grammar Epha;

program : statement+ ;

statement
    : importStatement
    | resourceDefinition
    ;

importStatement
    : 'import' IDENTIFIER ('as' IDENTIFIER)?
    ;

resourceDefinition
    : IDENTIFIER IDENTIFIER '{' resourceBody '}'
    ;

resourceBody
    : resourceProperty*
    ;

resourceProperty
    : propertyKey '=' value
    | propertyKey '=' valueList
    | propertyKey '{' resourcePropertyBody '}'
    ;

resourcePropertyBody
    : resourceProperty*
    ;

propertyKey
    : IDENTIFIER ('.' IDENTIFIER)*
    ;

value
    : STRING
    | NUMBER
    ;

valueList
    : '[' value (',' value)* ']'
    ;

IDENTIFIER
    : [a-zA-Z_] [a-zA-Z_0-9\-]*
    ;

STRING
    : '"' (~["\\] | '\\' .)* '"'
    ;

NUMBER
    : [0-9]+ ('.' [0-9]+)?
    ;

// Comment handling
LINE_COMMENT
    : '#' ~[\r\n]* -> skip
    ;

WHITESPACE
    : [ \t\r\n]+ -> skip
    ;
