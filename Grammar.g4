grammar Grammar;

// TOKENS


// regular eressions
DOUBLE      :   [0-9]+('.'[0-9]+);
INT         :   [0-9]+;
ID          :   [a-zA-Z_][a-zA-Z0-9_]*;
STRING      :   '"' (~["\r\n] | '""')* '"';
WHITESPACE  :   [ \t\n\r]+ -> skip ; 

// Rules

init
   : block EOF;

block
   : (stmt)*
   ;

stmt
    : declstmt
    | printstmt
    | ifstmt
    | forstmt
    ;

declstmt
    : typeVar='let' ID '=' expr
    | typeVar='var' ID '=' expr
    ;

printstmt
    : 'print' '(' expr ')'
    ;

ifstmt
    : 'if' expr '{' block '}'
    ;

forstmt
    : 'for' ID 'in' ID '{' block '}'
    ;

expr
    : '!' right=expr                            # NotExpr
    | '-' right=expr                            # UnaryExpr
    | left=expr op=('*'|'/'|'%') right=expr     # OpExpr
    | left=expr op=('+'|'-') right=expr         # OpExpr
    | left=expr op=('>='|'>') right=expr        # OpExpr
    | left=expr op=('<='|'<') right=expr        # OpExpr
    | left=expr op=('=='|'!=') right=expr       # OpExpr
    | left=expr '&&' right=expr                 # OpExpr
    | left=expr '||' right=expr                 # OpExpr
    | '(' expr ')'                              # ParExpr
    | INT                                       # IntExpr
    | DOUBLE                                    # DoubleExpr
    | ID                                        # IdExpr
    | STRING                                    # StrExpr
    | ('true'|'false')                          # BoolExpr
    | 'nil'                                     # nilExpr
    ;