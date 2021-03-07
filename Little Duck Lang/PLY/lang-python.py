#------------------------------------------------------
# lang-python.py
#
# A simple compiler
#------------------------------------------------------

tokens = (
    'ID','IF','THEN','ELSE',
    'LESS','GREATER','DIFF','EX','TERM',
    'LEFTPAR','RIGHTPAR','LEFTKEY',
    'RIGHTKEY','EQUALS','SEMCOLON','PROGRAM',
    'COLON','COMMA','VAR','TINT','TFLOAT','TSTRING','INT','FLOAT','STRING','PRINT'
)

reserved = {
    'int':'TINT',
    'float':'TFLOAT',
    'string':'TSTRING',
    'if':'IF',
    'then':'THEN',
    'else':'ELSE',
    'var':'VAR',
    'print':'PRINT',
    'program':'PROGRAM',
}

t_LESS = r'\<'
t_DIFF =r'\<\>'
t_GREATER = r'\>'
t_EX = r'[\+|\-]'
t_TERM = r'[\*|\/]'
t_LEFTPAR = r'\('
t_RIGHTPAR = r'\)'
t_LEFTKEY = r'\{'
t_RIGHTKEY = r'\}'
t_EQUALS  = r'\='
t_SEMCOLON = r';'
t_COLON = r':'
t_COMMA = r','

def t_FLOAT(t):
    r'\d+\.\d+'
    t.value = float(t.value)
    print(t)
    return t 

def t_INT(t):
    r'\d+'
    t.value = int(t.value)
    return t


def t_STRING(t):
    r'\"[a-zA-Z_][a-zA-Z0-9_]*\"'
    return t

# Ignored characters
t_ignore = " \t"

def t_newline(t):
    r'\n+'
    t.lexer.lineno += t.value.count("\n")


def t_ID(t):
    r'[a-zA-Z_][a-zA-Z0-9_]*'
    if t.value in reserved:
        t.type = reserved[t.value]
    else:
        t.type = 'ID'
    return t    

def t_error(t):
    print(t)
    print(f"Illegal character {t.value[0]!r}")
    t.lexer.skip(1)

import ply.lex as lex
lex.lex()
 
def p_programa(p): #.
    ''' programa : PROGRAM ID SEMCOLON vars bloque
                 | PROGRAM ID SEMCOLON bloque '''
def p_vars(p): #.
    ''' vars : VAR recvars '''

def p_recvars(p): #.
    ''' recvars : recids COLON tipo SEMCOLON recvars
                | recids COLON tipo SEMCOLON '''
def p_recids(p):  #.
    ''' recids : ID 
               | ID COMMA recids '''

def p_tipo(p): #.
    ''' tipo : TINT
             | TFLOAT 
             | TSTRING
             '''

def p_bloque(p): #.
    ''' bloque : LEFTKEY estatutos RIGHTKEY
                | LEFTKEY RIGHTKEY ''' 

def p_estatutos(p):#.
    ''' estatutos : estatuto estatutos
                  | estatuto ''' 

def p_estatuto(p): #.
    ''' estatuto : asignacion 
                 | escritura
                 | condicion
                  '''

def p_asignacion(p): #.
    ''' asignacion : ID EQUALS expresion SEMCOLON '''

def p_expresion(p): #.
    ''' expresion : exp LESS exp
                  | exp GREATER exp
                  | exp DIFF exp
                  | exp '''

def p_condicion(p): #.
    ''' condicion : IF LEFTPAR expresion RIGHTPAR THEN bloque SEMCOLON
                  | IF LEFTPAR expresion RIGHTPAR THEN bloque ELSE bloque SEMCOLON
     '''

def p_exp(p):#.
    ''' exp : termino recexp
            | termino '''

def p_recexp(p):#.
    ''' recexp : EX exp '''

def p_termino(p):#.
    ''' termino : factor recterm 
                | factor '''

def p_recterm(p):#.
    ''' recterm : TERM termino '''

def p_factor(p): #.
    ''' factor : LEFTPAR expresion RIGHTPAR 
               | EX cte
               | cte
    '''

def p_escritura(p):
    ''' escritura : PRINT LEFTPAR recesc RIGHTPAR SEMCOLON '''
    
def p_recesc(p):
    ''' recesc : expresion COMMA recesc 
               | STRING COMMA recesc 
               | expresion 
               | STRING
    '''

def p_cte(p): #.
    ''' cte : ID 
            | INT
            | STRING
            | FLOAT
    '''

def p_error(p):
    print(f"Syntax error at {p.value!r}")
    exit()

import ply.yacc as yacc
yacc.yacc()


f = open('code.txt')
s = f.read()
f.close()

yacc.parse(s)

print('Code is okay.')