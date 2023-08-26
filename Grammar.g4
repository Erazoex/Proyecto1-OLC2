grammar Grammar;

// TOKENS
T_INT       :   'Int';
T_FLOAT     :   'Float';
T_CHAR      :   'Character';
T_STRING    :   'String';
T_BOOL      :   'Bool';

// Palabras reservadas
VAR         :   'var';
LET         :   'let';
IF          :   'if';
ELSE        :   'else';
SWITCH      :   'switch';
CASE        :   'case';
DEFAULT     :   'default';
WHILE       :   'while';
FOR         :   'for';

// regular expressions
DOUBLE      :   [0-9]+('.'[0-9]+);
INT         :   [0-9]+;
ID          :   [a-zA-Z_][a-zA-Z0-9_]*;
CHAR        :   '"' .? '"';
STRING      :   '"' ('\\' ["\r\n\t] | ~[\n\r"])* '"';
MULTCOMMENT :   '/*' ( ~[*/] | '*' ~'/')* '*/'-> skip;
COMMENT     :   '//' ~[\r\n]* ('\r'? '\n' | EOF)-> skip;
WHITESPACE  :   [ \t\n\r]+ -> skip ; 

// Rules

init
   : block EOF;

block
   : (stmt)*
   ;

stmt
    : declstmt
    | asignstmt
    | incstmt
    | decstmt
    | ifstmt
    | switchstmt
    | printlnstmt
    | whilestmt
    | forstmt
    ;

declstmt
    : vtype=(VAR|LET) ID ':' vartype '=' expr   #declstmtWithTypeAndExpr
    | vtype=(VAR|LET) ID '=' expr               #declstmtWithExpr
    | vtype=(VAR|LET) ID ':' vartype '?'        #declstmtWithType
    ;

asignstmt
    : ID '=' expr
    ;

incstmt
    : ID '+=' expr
    ;

decstmt
    : ID '-=' expr
    ;

ifstmt
    : IF expr '{' block '}'                                                     # ifSimple
    | IF expr '{' trueCondition=block '}' ELSE '{' falseCondition=block '}'     # ifWithElse
    | IF expr '{' block '}' ELSE ifstmt                                         # ifWithElseIf
    ;

switchstmt
    : SWITCH expr '{' (switchcase)+ '}'
    ;


switchcase
    : casetype=CASE expr ':' block
    | casetype=DEFAULT ':' block
    ;

printlnstmt
    : 'println' '(' expr ')'
    ;

whilestmt
    : WHILE expr '{' block '}'
    ;

forstmt
    : FOR ID 'in' expr '{' block '}'        #forWithExpr
    | FOR ID 'in' forrange '{' block '}'       #forWithRange
    ;

forrange
    : beginsWith=INT '...' endsWith=INT
    ;

vartype
    : tipo=T_INT        # primitiveType
    | tipo=T_FLOAT      # primitiveType
    | tipo=T_CHAR       # primitiveType
    | tipo=T_STRING     # primitiveType
    | tipo=T_BOOL       # primitiveType
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
    | CHAR                                      # CharExpr
    | STRING                                    # StrExpr
    | ('true'|'false')                          # BoolExpr
    | 'nil'                                     # nilExpr
    ;