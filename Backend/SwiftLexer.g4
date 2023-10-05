lexer grammar SwiftLexer;

// --------------- Tokens
// types
RINT: 'Int';
RFLOAT: 'Float';
RBOOL: 'Bool';
RSTRING: 'String';
RCHARACTER: 'Character';

// reserved words
RTRUE: 'true';
RFALSE: 'false';
RPRINT: 'print';
RIF: 'if';
RELSE: 'else';
RWHILE: 'while';
RVAR: 'var';
RLET: 'let';
RNIL: 'nil';
RBREAK: 'break';
RCONTINUE:'continue';
RAPPEND : 'append'  ;
RREMOVELAST : 'removeLast';
RRAT        : 'at';
RREMOVEAT   : 'remove';
RISEMPTY :'isEmpty';
RCOUNT : 'count';
RSWITCH :'switch';
RCASE : 'case';
RDEFAULT : 'default';
RFOR : 'for';
RIN : 'in';
RREPEATING: 'repeating';
RSTRUCT: 'struct';
RGUARD : 'guard';
RFUNC : 'func';
RRETURN : 'return';
RINOUT : 'inout';
RMUTATING: 'mutating';
RSELF : 'self' ;

// primitives
NUMBER : [0-9]+ ('.'[0-9]+)?;
STRING: '"' ('\\' ('\\' | '"' | 'r'| 't'| 'n') | ~('\\' | '"'))* '"';

ID: [a-zA-Z_][a-zA-Z0-9_]*;

// symbols
ARROW :         '->';
UNARYPLUS:      '+=';
UNARYMINUS:     '-=';
DIF:            '!=';
IG_IG:          '==';
NOT:            '!';
OR:             '||';
AND:            '&&';
IG:             '=';
MAY_IG:         '>=';
MEN_IG:         '<=';
MAYOR:          '>';
MENOR:          '<';
MUL:            '*';
DIV:            '/';
ADD:            '+';
SUB:            '-';
PARIZQ:         '(';
PARDER:         ')';
LLAVEIZQ:       '{';
LLAVEDER:       '}';
OBRA:           '[';
CBRA:           ']';
PTOCOMA:        ';';
DOSPTOS:        ':';
MODULE:         '%' ;
COMA:           ',' ;
QM:             '?' ;
PTO:            '.';
AMPERSAND :     '&';
UNDERSCORE :    '_';

// skip
WHITESPACE: [ \\\r\n\t]+ -> skip;
COMMENT : '/*' .*? '*/' -> skip;
LINE_COMMENT : '//' ~[\r\n]* -> skip;

//fragment
//ESC_SEQ
//    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
//    ;
