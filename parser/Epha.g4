grammar Epha;

program     : statement* EOF;
statement   : importStmt | variableDecl | helmChart | k8sResource ;
importStmt  : 'import' IDENTIFIER ('from' IDENTIFIER ('import' IDENTIFIER (',' IDENTIFIER)* )?)? ;
variableDecl: IDENTIFIER '=' expression;
helmChart   : 'helm' IDENTIFIER '{' helmStmt* '}';
k8sResource : 'k8s' IDENTIFIER '{' k8sStmt* '}';
helmStmt    : IDENTIFIER ':' expression;
k8sStmt     : IDENTIFIER ':' expression;

expression  : STRING | NUMBER | BOOLEAN | IDENTIFIER | arrayLiteral | hashLiteral;
arrayLiteral: '[' expression (',' expression)* ']';
hashLiteral : '{' (expression ':' expression) (',' expression ':' expression)* '}';

IDENTIFIER  : [a-zA-Z_][a-zA-Z_0-9]*;
STRING      : '"' .*? '"';
NUMBER      : [0-9]+;
BOOLEAN     : 'true' | 'false';
WS          : [ \t\r\n]+ -> skip;
