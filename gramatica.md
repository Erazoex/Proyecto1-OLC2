# GRAMATICA USADA PARA EL PROYECTO
## tokens utilizados
T_INT        ---->   'Int'
T_FLOAT      ---->   'Float'
T_CHAR       ---->   'Character'
T_STRING     ---->   'String'
T_BOOL       ---->   'Bool'
# palabras reservadas utilizadas
VAR          ---->   'var'
LET          ---->   'let'
IF           ---->   'if'
ELSE         ---->   'else'
SWITCH       ---->   'switch'
CASE         ---->   'case';
DEFAULT      ---->   'default'
WHILE        ---->   'while'
FOR          ---->   'for'
VECTOR       ---->   'vector'
FUNC         ---->   'func'
GUARD        ---->   'guard'
BREAK        ---->   'break'
RETURN       ---->   'return'
CONTINUE     ---->   'continue' 
## expresiones regulares
DOUBLE       ---->   [0-9]+('.'[0-9]+)
INT          ---->   [0-9]+
ID           ---->   [a-zA-Z_][a-zA-Z0-9_]*
CHAR         ---->   '"' .? '"';
STRING       ---->   '"' ('\\' ["\r\n\t] | ~[\n\r"])* '"'
MULTCOMMENT  ---->   '/*' ( ~[*/] | '*' ~'/')* '*/'-> skip
COMMENT      ---->   '//' ~[\r\n]* ('\r'? '\n' | EOF)-> skip
WHITESPACE   ---->   [ \t\n\r]+ -> skip 

## Gramatica 
init
   : block epsilon;

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
    | guardstmt
    | breakstmt
    | continuestmt
    | returnstmt
    | funcstmt
    | callstmt
    ;

breakstmt
    : BREAK
    ;

continuestmt
    : CONTINUE
    ;

returnstmt
    : RETURN (expr)?
    ;

declstmt
    : vtype=(VAR|LET) ID ':' vartype '=' expr   
    | vtype=(VAR|LET) ID '=' expr               
    | vtype=(VAR|LET) ID ':' vartype '?'        
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
    : IF expr '{' block '}'                                                     
    | IF expr '{' trueCondition=block '}' ELSE '{' falseCondition=block '}'     
    | IF expr '{' block '}' ELSE ifstmt                                         
    ;

switchstmt
    : SWITCH expr '{' (switchcase)+ '}'
    ;

switchcase
    : casetype=CASE expr ':' block
    | casetype=DEFAULT ':' block
    ;

printlnstmt
    : 'print' '(' (exprparams)? ')'
    ;

intstmt
    : 'Int' '(' expr ')'
    ;

floatstmt
    : 'Float' '(' expr ')'
    ;

stringstmt
    : 'String' '(' expr ')'
    ;

exprparams
    : expr ',' exprparams       
    | expr                      
    ;

whilestmt
    : WHILE expr '{' block '}'
    ;

forstmt
    : FOR ID 'in' expr '{' block '}'        
    | FOR ID 'in' forrange '{' block '}'       
    ;

forrange
    : beginsWith=expr '...' endsWith=expr
    ;

guardstmt
    : GUARD expr ELSE '{' block '}' 
    ;

array
    : VECTOR '<' vartype '>' ID array_def
    ;

array_def
    : '=' (expr)*
    ;

funcstmt
    : 'func' ID '(' (listaparametros)? ')' ('->' vartype)? '{' block '}'
    ;

listaparametros
    : parametro ',' listaparametros         
    | parametro                             
    ;

parametro
    : (extVarId=ID)? varId=ID ':' vartype
    ;

callstmt
    : ID '(' exprparams ')'
    ;

vartype
    : tipo=T_INT        
    | tipo=T_FLOAT      
    | tipo=T_CHAR       
    | tipo=T_STRING     
    | tipo=T_BOOL       
;

expr
    : '!' right=expr                            
    | '-' right=expr                            
    | left=expr op=('*'|'/'|'%') right=expr     
    | left=expr op=('+'|'-') right=expr         
    | left=expr op=('>='|'>') right=expr        
    | left=expr op=('<='|'<') right=expr        
    | left=expr op=('=='|'!=') right=expr       
    | left=expr op='&&' right=expr              
    | left=expr op='||' right=expr              
    | '(' expr ')'                              
    | INT                                       
    | DOUBLE                                    
    | ID                                        
    | CHAR                                      
    | STRING                                    
    | ('true'|'false')                          
    | 'nil'                                     
    | intstmt                                   
    | floatstmt                                 
    | stringstmt                                
    | callstmt                                  
    ;