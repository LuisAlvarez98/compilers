// Code generated by gocc; DO NOT EDIT.

package parser



type (
	ProdTab      [numProductions]ProdTabEntry
	ProdTabEntry struct {
		String     string
		Id         string
		NTType     int
		Index      int
		NumSymbols int
		ReduceFunc func([]Attrib) (Attrib, error)
	}
	Attrib interface {
	}
)

var productionsTable = ProdTab{
	ProdTabEntry{
		String: `S' : Programa	<<  >>`,
		Id:         "S'",
		NTType:     0,
		Index:      0,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Programa : program id semcolon Vars Bloque	<<  >>`,
		Id:         "Programa",
		NTType:     1,
		Index:      1,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Programa : program id semcolon Bloque	<<  >>`,
		Id:         "Programa",
		NTType:     1,
		Index:      2,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Vars : var Recvars	<<  >>`,
		Id:         "Vars",
		NTType:     2,
		Index:      3,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Recvars : Recids colon Tipo semcolon Recvars	<<  >>`,
		Id:         "Recvars",
		NTType:     3,
		Index:      4,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Recvars : Recids colon Tipo semcolon	<<  >>`,
		Id:         "Recvars",
		NTType:     3,
		Index:      5,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Recids : id	<<  >>`,
		Id:         "Recids",
		NTType:     4,
		Index:      6,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Recids : id comma Recids	<<  >>`,
		Id:         "Recids",
		NTType:     4,
		Index:      7,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Tipo : tint	<<  >>`,
		Id:         "Tipo",
		NTType:     5,
		Index:      8,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Tipo : tfloat	<<  >>`,
		Id:         "Tipo",
		NTType:     5,
		Index:      9,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Tipo : tstring	<<  >>`,
		Id:         "Tipo",
		NTType:     5,
		Index:      10,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Bloque : leftkey Estatutos rightkey	<<  >>`,
		Id:         "Bloque",
		NTType:     6,
		Index:      11,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Bloque : leftkey rightkey	<<  >>`,
		Id:         "Bloque",
		NTType:     6,
		Index:      12,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Estatutos : Estatuto Estatutos	<<  >>`,
		Id:         "Estatutos",
		NTType:     7,
		Index:      13,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Estatutos : Estatuto	<<  >>`,
		Id:         "Estatutos",
		NTType:     7,
		Index:      14,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Estatuto : Asignacion	<<  >>`,
		Id:         "Estatuto",
		NTType:     8,
		Index:      15,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Estatuto : Escritura	<<  >>`,
		Id:         "Estatuto",
		NTType:     8,
		Index:      16,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Estatuto : Condicion	<<  >>`,
		Id:         "Estatuto",
		NTType:     8,
		Index:      17,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Asignacion : id equals Expresion semcolon	<<  >>`,
		Id:         "Asignacion",
		NTType:     9,
		Index:      18,
		NumSymbols: 4,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expresion : Exp less Exp	<<  >>`,
		Id:         "Expresion",
		NTType:     10,
		Index:      19,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expresion : Exp greater Exp	<<  >>`,
		Id:         "Expresion",
		NTType:     10,
		Index:      20,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expresion : Exp diff Exp	<<  >>`,
		Id:         "Expresion",
		NTType:     10,
		Index:      21,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Expresion : Exp	<<  >>`,
		Id:         "Expresion",
		NTType:     10,
		Index:      22,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Condicion : if leftpar Expresion rightpar then Bloque semcolon	<<  >>`,
		Id:         "Condicion",
		NTType:     11,
		Index:      23,
		NumSymbols: 7,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Condicion : if leftpar Expresion rightpar then Bloque else Bloque semcolon	<<  >>`,
		Id:         "Condicion",
		NTType:     11,
		Index:      24,
		NumSymbols: 9,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Exp : Termino Recexp	<<  >>`,
		Id:         "Exp",
		NTType:     12,
		Index:      25,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Exp : Termino	<<  >>`,
		Id:         "Exp",
		NTType:     12,
		Index:      26,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Recexp : ex Exp	<<  >>`,
		Id:         "Recexp",
		NTType:     13,
		Index:      27,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Termino : Factor Recterm	<<  >>`,
		Id:         "Termino",
		NTType:     14,
		Index:      28,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Termino : Factor	<<  >>`,
		Id:         "Termino",
		NTType:     14,
		Index:      29,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Recterm : term Termino	<<  >>`,
		Id:         "Recterm",
		NTType:     15,
		Index:      30,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : leftpar Expresion rightpar	<<  >>`,
		Id:         "Factor",
		NTType:     16,
		Index:      31,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : ex Cte	<<  >>`,
		Id:         "Factor",
		NTType:     16,
		Index:      32,
		NumSymbols: 2,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Factor : Cte	<<  >>`,
		Id:         "Factor",
		NTType:     16,
		Index:      33,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Escritura : print leftpar Recesc rightpar semcolon	<<  >>`,
		Id:         "Escritura",
		NTType:     17,
		Index:      34,
		NumSymbols: 5,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Recesc : Expresion comma Recesc	<<  >>`,
		Id:         "Recesc",
		NTType:     18,
		Index:      35,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Recesc : string comma Recesc	<<  >>`,
		Id:         "Recesc",
		NTType:     18,
		Index:      36,
		NumSymbols: 3,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Recesc : Expresion	<<  >>`,
		Id:         "Recesc",
		NTType:     18,
		Index:      37,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Recesc : string	<<  >>`,
		Id:         "Recesc",
		NTType:     18,
		Index:      38,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Cte : id	<<  >>`,
		Id:         "Cte",
		NTType:     19,
		Index:      39,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Cte : int	<<  >>`,
		Id:         "Cte",
		NTType:     19,
		Index:      40,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Cte : string	<<  >>`,
		Id:         "Cte",
		NTType:     19,
		Index:      41,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
	ProdTabEntry{
		String: `Cte : float	<<  >>`,
		Id:         "Cte",
		NTType:     19,
		Index:      42,
		NumSymbols: 1,
		ReduceFunc: func(X []Attrib) (Attrib, error) {
			return X[0], nil
		},
	},
}