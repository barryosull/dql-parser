package expression

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

var g = &grammar{
	rules: []*rule{
		{
			name: "EXRESSION_TEST",
			pos:  position{line: 11, col: 1, offset: 205},
			expr: &seqExpr{
				pos: position{line: 11, col: 18, offset: 222},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 11, col: 18, offset: 222},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 11, col: 29, offset: 233},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "EXPRESSION",
			pos:  position{line: 13, col: 1, offset: 238},
			expr: &choiceExpr{
				pos: position{line: 13, col: 14, offset: 251},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 13, col: 14, offset: 251},
						name: "QUERY",
					},
					&ruleRefExpr{
						pos:  position{line: 13, col: 22, offset: 259},
						name: "ARITHMETIC",
					},
					&ruleRefExpr{
						pos:  position{line: 13, col: 35, offset: 272},
						name: "COMPARISON",
					},
					&ruleRefExpr{
						pos:  position{line: 13, col: 48, offset: 285},
						name: "ASSIGNMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 13, col: 60, offset: 297},
						name: "LOGICAL",
					},
					&ruleRefExpr{
						pos:  position{line: 13, col: 70, offset: 307},
						name: "ATOMIC",
					},
				},
			},
		},
		{
			name: "ATOMIC",
			pos:  position{line: 15, col: 1, offset: 315},
			expr: &choiceExpr{
				pos: position{line: 15, col: 10, offset: 324},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 15, col: 10, offset: 324},
						name: "PARENTHESIS",
					},
					&ruleRefExpr{
						pos:  position{line: 15, col: 24, offset: 338},
						name: "NEW",
					},
					&ruleRefExpr{
						pos:  position{line: 15, col: 30, offset: 344},
						name: "METHODCALL",
					},
					&ruleRefExpr{
						pos:  position{line: 15, col: 43, offset: 357},
						name: "OBJECTACCESS",
					},
					&ruleRefExpr{
						pos:  position{line: 15, col: 58, offset: 372},
						name: "ARRAY",
					},
					&ruleRefExpr{
						pos:  position{line: 15, col: 66, offset: 380},
						name: "LITERAL",
					},
					&ruleRefExpr{
						pos:  position{line: 15, col: 76, offset: 390},
						name: "UNARY",
					},
				},
			},
		},
		{
			name: "LITERAL",
			pos:  position{line: 17, col: 1, offset: 397},
			expr: &choiceExpr{
				pos: position{line: 17, col: 11, offset: 407},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 17, col: 11, offset: 407},
						name: "STRING",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 20, offset: 416},
						name: "FLOAT",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 28, offset: 424},
						name: "BOOLEAN",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 38, offset: 434},
						name: "NULL",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 45, offset: 441},
						name: "INT",
					},
				},
			},
		},
		{
			name: "NEW",
			pos:  position{line: 19, col: 1, offset: 446},
			expr: &seqExpr{
				pos: position{line: 19, col: 7, offset: 452},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 19, col: 7, offset: 452},
						name: "CLASS_REF_QUOTES",
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 24, offset: 469},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 19, col: 26, offset: 471},
						expr: &ruleRefExpr{
							pos:  position{line: 19, col: 26, offset: 471},
							name: "ARGUMENTS",
						},
					},
				},
			},
		},
		{
			name: "BOOLEAN",
			pos:  position{line: 21, col: 1, offset: 483},
			expr: &choiceExpr{
				pos: position{line: 21, col: 12, offset: 494},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 21, col: 12, offset: 494},
						val:        "true",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 21, col: 19, offset: 501},
						val:        "false",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "NULL",
			pos:  position{line: 23, col: 1, offset: 510},
			expr: &litMatcher{
				pos:        position{line: 23, col: 8, offset: 517},
				val:        "null",
				ignoreCase: false,
			},
		},
		{
			name: "ARRAY",
			pos:  position{line: 25, col: 1, offset: 525},
			expr: &seqExpr{
				pos: position{line: 25, col: 9, offset: 533},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 25, col: 9, offset: 533},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 25, col: 11, offset: 535},
						val:        "[",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 25, col: 15, offset: 539},
						expr: &ruleRefExpr{
							pos:  position{line: 25, col: 15, offset: 539},
							name: "ARGUMENTLIST",
						},
					},
					&litMatcher{
						pos:        position{line: 25, col: 29, offset: 553},
						val:        "]",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 33, offset: 557},
						name: "_",
					},
				},
			},
		},
		{
			name: "STRING",
			pos:  position{line: 27, col: 1, offset: 560},
			expr: &seqExpr{
				pos: position{line: 27, col: 10, offset: 569},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 27, col: 10, offset: 569},
						val:        "\"",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 27, col: 15, offset: 574},
						expr: &charClassMatcher{
							pos:        position{line: 27, col: 15, offset: 574},
							val:        "[a-zA-Z0-9]",
							ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&litMatcher{
						pos:        position{line: 27, col: 28, offset: 587},
						val:        "\"",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "INT",
			pos:  position{line: 29, col: 1, offset: 593},
			expr: &oneOrMoreExpr{
				pos: position{line: 29, col: 7, offset: 599},
				expr: &charClassMatcher{
					pos:        position{line: 29, col: 7, offset: 599},
					val:        "[0-9]",
					ranges:     []rune{'0', '9'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "FLOAT",
			pos:  position{line: 31, col: 1, offset: 607},
			expr: &seqExpr{
				pos: position{line: 31, col: 9, offset: 615},
				exprs: []interface{}{
					&oneOrMoreExpr{
						pos: position{line: 31, col: 9, offset: 615},
						expr: &charClassMatcher{
							pos:        position{line: 31, col: 9, offset: 615},
							val:        "[0-9]",
							ranges:     []rune{'0', '9'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&charClassMatcher{
						pos:        position{line: 31, col: 16, offset: 622},
						val:        "[.]",
						chars:      []rune{'.'},
						ignoreCase: false,
						inverted:   false,
					},
					&oneOrMoreExpr{
						pos: position{line: 31, col: 20, offset: 626},
						expr: &charClassMatcher{
							pos:        position{line: 31, col: 20, offset: 626},
							val:        "[0-9]",
							ranges:     []rune{'0', '9'},
							ignoreCase: false,
							inverted:   false,
						},
					},
				},
			},
		},
		{
			name: "PARENTHESIS",
			pos:  position{line: 33, col: 1, offset: 634},
			expr: &seqExpr{
				pos: position{line: 33, col: 15, offset: 648},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 33, col: 15, offset: 648},
						val:        "(",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 19, offset: 652},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 21, offset: 654},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 32, offset: 665},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 33, col: 34, offset: 667},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "UNARY",
			pos:  position{line: 35, col: 1, offset: 672},
			expr: &choiceExpr{
				pos: position{line: 35, col: 9, offset: 680},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 35, col: 9, offset: 680},
						name: "INCREMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 21, offset: 692},
						name: "DECREMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 33, offset: 704},
						name: "NEGATE",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 42, offset: 713},
						name: "NOT",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 48, offset: 719},
						name: "POSITIVE",
					},
				},
			},
		},
		{
			name: "INCREMENT",
			pos:  position{line: 37, col: 1, offset: 729},
			expr: &seqExpr{
				pos: position{line: 37, col: 13, offset: 741},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 37, col: 13, offset: 741},
						name: "OBJECTACCESS",
					},
					&litMatcher{
						pos:        position{line: 37, col: 26, offset: 754},
						val:        "++",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DECREMENT",
			pos:  position{line: 39, col: 1, offset: 760},
			expr: &seqExpr{
				pos: position{line: 39, col: 13, offset: 772},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 39, col: 13, offset: 772},
						name: "OBJECTACCESS",
					},
					&litMatcher{
						pos:        position{line: 39, col: 26, offset: 785},
						val:        "--",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "NEGATE",
			pos:  position{line: 41, col: 1, offset: 791},
			expr: &seqExpr{
				pos: position{line: 41, col: 10, offset: 800},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 41, col: 10, offset: 800},
						val:        "-",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 14, offset: 804},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "NOT",
			pos:  position{line: 43, col: 1, offset: 818},
			expr: &seqExpr{
				pos: position{line: 43, col: 7, offset: 824},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 43, col: 7, offset: 824},
						val:        "!",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 11, offset: 828},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "POSITIVE",
			pos:  position{line: 45, col: 1, offset: 842},
			expr: &seqExpr{
				pos: position{line: 45, col: 12, offset: 853},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 45, col: 12, offset: 853},
						val:        "+",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 45, col: 16, offset: 857},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "ARITHMETIC",
			pos:  position{line: 47, col: 1, offset: 871},
			expr: &seqExpr{
				pos: position{line: 47, col: 14, offset: 884},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 47, col: 14, offset: 884},
						name: "ATOMIC",
					},
					&ruleRefExpr{
						pos:  position{line: 47, col: 21, offset: 891},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 47, col: 23, offset: 893},
						name: "OPERATOR",
					},
					&ruleRefExpr{
						pos:  position{line: 47, col: 32, offset: 902},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 47, col: 34, offset: 904},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "OPERATOR",
			pos:  position{line: 49, col: 1, offset: 916},
			expr: &choiceExpr{
				pos: position{line: 49, col: 12, offset: 927},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 49, col: 12, offset: 927},
						val:        "+",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 49, col: 18, offset: 933},
						val:        "-",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 49, col: 24, offset: 939},
						val:        "/",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 49, col: 30, offset: 945},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 49, col: 36, offset: 951},
						val:        "%",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "ASSIGNMENT",
			pos:  position{line: 51, col: 1, offset: 956},
			expr: &seqExpr{
				pos: position{line: 51, col: 14, offset: 969},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 51, col: 14, offset: 969},
						name: "OBJECTACCESS",
					},
					&ruleRefExpr{
						pos:  position{line: 51, col: 27, offset: 982},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 51, col: 29, offset: 984},
						val:        "=",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 51, col: 33, offset: 988},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 51, col: 35, offset: 990},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "LOGICAL",
			pos:  position{line: 53, col: 1, offset: 1002},
			expr: &seqExpr{
				pos: position{line: 53, col: 11, offset: 1012},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 53, col: 11, offset: 1012},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 53, col: 22, offset: 1023},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 53, col: 25, offset: 1026},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 53, col: 25, offset: 1026},
								val:        "and",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 53, col: 33, offset: 1034},
								val:        "or",
								ignoreCase: false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 53, col: 39, offset: 1040},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 53, col: 41, offset: 1042},
						name: "ATOMIC",
					},
				},
			},
		},
		{
			name: "COMPARISON",
			pos:  position{line: 55, col: 1, offset: 1050},
			expr: &seqExpr{
				pos: position{line: 55, col: 14, offset: 1063},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 55, col: 14, offset: 1063},
						name: "ATOMIC",
					},
					&ruleRefExpr{
						pos:  position{line: 55, col: 21, offset: 1070},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 55, col: 24, offset: 1073},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 55, col: 24, offset: 1073},
								val:        "===",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 55, col: 32, offset: 1081},
								val:        "!==",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 55, col: 40, offset: 1089},
								val:        "==",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 55, col: 47, offset: 1096},
								val:        "!=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 55, col: 54, offset: 1103},
								val:        "<=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 55, col: 61, offset: 1110},
								val:        ">=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 55, col: 68, offset: 1117},
								val:        "<",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 55, col: 74, offset: 1123},
								val:        ">",
								ignoreCase: false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 55, col: 79, offset: 1128},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 55, col: 81, offset: 1130},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "QUERY",
			pos:  position{line: 57, col: 1, offset: 1142},
			expr: &seqExpr{
				pos: position{line: 57, col: 9, offset: 1150},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 57, col: 9, offset: 1150},
						val:        "run",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 57, col: 16, offset: 1157},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 57, col: 18, offset: 1159},
						val:        "query",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 57, col: 27, offset: 1168},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 57, col: 29, offset: 1170},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 57, col: 41, offset: 1182},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 57, col: 43, offset: 1184},
						expr: &ruleRefExpr{
							pos:  position{line: 57, col: 43, offset: 1184},
							name: "ARGUMENTS",
						},
					},
				},
			},
		},
		{
			name: "OBJECTACCESS",
			pos:  position{line: 59, col: 1, offset: 1196},
			expr: &seqExpr{
				pos: position{line: 59, col: 16, offset: 1211},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 59, col: 16, offset: 1211},
						expr: &seqExpr{
							pos: position{line: 59, col: 17, offset: 1212},
							exprs: []interface{}{
								&choiceExpr{
									pos: position{line: 59, col: 18, offset: 1213},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 59, col: 18, offset: 1213},
											name: "METHODCALL",
										},
										&ruleRefExpr{
											pos:  position{line: 59, col: 31, offset: 1226},
											name: "IDENTIFIER",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 59, col: 43, offset: 1238},
									val:        "->",
									ignoreCase: false,
								},
							},
						},
					},
					&choiceExpr{
						pos: position{line: 59, col: 51, offset: 1246},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 59, col: 51, offset: 1246},
								name: "METHODCALL",
							},
							&ruleRefExpr{
								pos:  position{line: 59, col: 64, offset: 1259},
								name: "IDENTIFIER",
							},
						},
					},
				},
			},
		},
		{
			name: "METHODCALL",
			pos:  position{line: 61, col: 1, offset: 1272},
			expr: &seqExpr{
				pos: position{line: 61, col: 14, offset: 1285},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 61, col: 14, offset: 1285},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 61, col: 25, offset: 1296},
						name: "ARGUMENTS",
					},
				},
			},
		},
		{
			name: "ARGUMENTS",
			pos:  position{line: 63, col: 1, offset: 1307},
			expr: &seqExpr{
				pos: position{line: 63, col: 13, offset: 1319},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 63, col: 13, offset: 1319},
						val:        "(",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 63, col: 17, offset: 1323},
						expr: &ruleRefExpr{
							pos:  position{line: 63, col: 17, offset: 1323},
							name: "ARGUMENTLIST",
						},
					},
					&litMatcher{
						pos:        position{line: 63, col: 31, offset: 1337},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "ARGUMENTLIST",
			pos:  position{line: 65, col: 1, offset: 1342},
			expr: &seqExpr{
				pos: position{line: 65, col: 17, offset: 1358},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 65, col: 17, offset: 1358},
						name: "_",
					},
					&zeroOrMoreExpr{
						pos: position{line: 65, col: 19, offset: 1360},
						expr: &seqExpr{
							pos: position{line: 65, col: 20, offset: 1361},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 65, col: 20, offset: 1361},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 65, col: 22, offset: 1363},
									name: "EXPRESSION",
								},
								&ruleRefExpr{
									pos:  position{line: 65, col: 33, offset: 1374},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 65, col: 35, offset: 1376},
									val:        ",",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 65, col: 39, offset: 1380},
									name: "_",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 65, col: 43, offset: 1384},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 65, col: 54, offset: 1395},
						name: "_",
					},
				},
			},
		},
		{
			name: "CLASS_REF_QUOTES",
			pos:  position{line: 72, col: 1, offset: 1563},
			expr: &seqExpr{
				pos: position{line: 72, col: 20, offset: 1582},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 72, col: 20, offset: 1582},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 72, col: 22, offset: 1584},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 72, col: 26, offset: 1588},
						name: "CLASS_REF",
					},
					&litMatcher{
						pos:        position{line: 72, col: 36, offset: 1598},
						val:        "'",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CLASS_REF",
			pos:  position{line: 74, col: 1, offset: 1603},
			expr: &seqExpr{
				pos: position{line: 74, col: 13, offset: 1615},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 74, col: 13, offset: 1615},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 74, col: 15, offset: 1617},
						name: "CLASS_TYPE",
					},
					&litMatcher{
						pos:        position{line: 74, col: 26, offset: 1628},
						val:        "\\",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 74, col: 31, offset: 1633},
						name: "CLASS_NAME",
					},
				},
			},
		},
		{
			name: "CLASS_TYPE",
			pos:  position{line: 76, col: 1, offset: 1645},
			expr: &seqExpr{
				pos: position{line: 76, col: 14, offset: 1658},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 76, col: 14, offset: 1658},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 76, col: 17, offset: 1661},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 76, col: 17, offset: 1661},
								val:        "value",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 76, col: 27, offset: 1671},
								val:        "entity",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 76, col: 38, offset: 1682},
								val:        "command",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 76, col: 50, offset: 1694},
								val:        "event",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 76, col: 60, offset: 1704},
								val:        "projection",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 76, col: 75, offset: 1719},
								val:        "invariant",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 76, col: 89, offset: 1733},
								val:        "query",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "CLASS_NAME",
			pos:  position{line: 78, col: 1, offset: 1743},
			expr: &seqExpr{
				pos: position{line: 78, col: 14, offset: 1756},
				exprs: []interface{}{
					&oneOrMoreExpr{
						pos: position{line: 78, col: 14, offset: 1756},
						expr: &charClassMatcher{
							pos:        position{line: 78, col: 14, offset: 1756},
							val:        "[a-zA-Z]",
							ranges:     []rune{'a', 'z', 'A', 'Z'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 78, col: 24, offset: 1766},
						expr: &charClassMatcher{
							pos:        position{line: 78, col: 24, offset: 1766},
							val:        "[a-zA-Z0-9_-]",
							chars:      []rune{'_', '-'},
							ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
							ignoreCase: false,
							inverted:   false,
						},
					},
				},
			},
		},
		{
			name: "QUOTED_NAME",
			pos:  position{line: 80, col: 1, offset: 1782},
			expr: &seqExpr{
				pos: position{line: 80, col: 15, offset: 1796},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 80, col: 15, offset: 1796},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 80, col: 19, offset: 1800},
						name: "CLASS_NAME",
					},
					&litMatcher{
						pos:        position{line: 80, col: 30, offset: 1811},
						val:        "'",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "TYPE",
			pos:  position{line: 82, col: 1, offset: 1816},
			expr: &choiceExpr{
				pos: position{line: 82, col: 8, offset: 1823},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 82, col: 8, offset: 1823},
						name: "CLASS_REF",
					},
					&ruleRefExpr{
						pos:  position{line: 82, col: 20, offset: 1835},
						name: "VALUE_TYPE",
					},
				},
			},
		},
		{
			name: "VALUE_TYPE",
			pos:  position{line: 84, col: 1, offset: 1847},
			expr: &seqExpr{
				pos: position{line: 84, col: 14, offset: 1860},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 84, col: 14, offset: 1860},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 84, col: 17, offset: 1863},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 84, col: 17, offset: 1863},
								val:        "string",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 84, col: 28, offset: 1874},
								val:        "boolean",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 84, col: 40, offset: 1886},
								val:        "float",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 84, col: 50, offset: 1896},
								val:        "map",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 84, col: 58, offset: 1904},
								val:        "index",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "CLASS_IMPLIED_REF",
			pos:  position{line: 86, col: 1, offset: 1914},
			expr: &seqExpr{
				pos: position{line: 86, col: 21, offset: 1934},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 86, col: 21, offset: 1934},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 86, col: 23, offset: 1936},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 86, col: 27, offset: 1940},
						name: "CLASS_NAME",
					},
					&litMatcher{
						pos:        position{line: 86, col: 38, offset: 1951},
						val:        "'",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "IDENTIFIER",
			pos:  position{line: 88, col: 1, offset: 1956},
			expr: &seqExpr{
				pos: position{line: 88, col: 14, offset: 1969},
				exprs: []interface{}{
					&oneOrMoreExpr{
						pos: position{line: 88, col: 14, offset: 1969},
						expr: &charClassMatcher{
							pos:        position{line: 88, col: 14, offset: 1969},
							val:        "[a-zA-Z]",
							ranges:     []rune{'a', 'z', 'A', 'Z'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 88, col: 24, offset: 1979},
						expr: &charClassMatcher{
							pos:        position{line: 88, col: 24, offset: 1979},
							val:        "[a-zA-Z0-9_]",
							chars:      []rune{'_'},
							ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
							ignoreCase: false,
							inverted:   false,
						},
					},
				},
			},
		},
		{
			name: "_",
			pos:  position{line: 90, col: 1, offset: 1994},
			expr: &zeroOrMoreExpr{
				pos: position{line: 90, col: 5, offset: 1998},
				expr: &choiceExpr{
					pos: position{line: 90, col: 7, offset: 2000},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 90, col: 7, offset: 2000},
							name: "WHITESPACE",
						},
						&ruleRefExpr{
							pos:  position{line: 90, col: 20, offset: 2013},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "SEMI",
			pos:  position{line: 92, col: 1, offset: 2021},
			expr: &seqExpr{
				pos: position{line: 92, col: 8, offset: 2028},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 92, col: 8, offset: 2028},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 92, col: 10, offset: 2030},
						val:        ";",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 92, col: 14, offset: 2034},
						name: "_",
					},
				},
			},
		},
		{
			name: "WHITESPACE",
			pos:  position{line: 94, col: 1, offset: 2037},
			expr: &charClassMatcher{
				pos:        position{line: 94, col: 14, offset: 2050},
				val:        "[ \\t\\r]",
				chars:      []rune{' ', '\t', '\r'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EOL",
			pos:  position{line: 96, col: 1, offset: 2059},
			expr: &litMatcher{
				pos:        position{line: 96, col: 7, offset: 2065},
				val:        "\n",
				ignoreCase: false,
			},
		},
		{
			name: "EOF",
			pos:  position{line: 98, col: 1, offset: 2071},
			expr: &notExpr{
				pos: position{line: 98, col: 7, offset: 2077},
				expr: &anyMatcher{
					line: 98, col: 8, offset: 2078,
				},
			},
		},
	},
}

var (
	// errNoRule is returned when the grammar to parse has no rule.
	errNoRule = errors.New("grammar has no rule")

	// errInvalidEncoding is returned when the source is not properly
	// utf8-encoded.
	errInvalidEncoding = errors.New("invalid encoding")

	// errNoMatch is returned if no match could be found.
	errNoMatch = errors.New("no match found")
)

// Option is a function that can set an option on the parser. It returns
// the previous setting as an Option.
type Option func(*parser) Option

// Debug creates an Option to set the debug flag to b. When set to true,
// debugging information is printed to stdout while parsing.
//
// The default is false.
func Debug(b bool) Option {
	return func(p *parser) Option {
		old := p.debug
		p.debug = b
		return Debug(old)
	}
}

// Memoize creates an Option to set the memoize flag to b. When set to true,
// the parser will cache all results so each expression is evaluated only
// once. This guarantees linear parsing time even for pathological cases,
// at the expense of more memory and slower times for typical cases.
//
// The default is false.
func Memoize(b bool) Option {
	return func(p *parser) Option {
		old := p.memoize
		p.memoize = b
		return Memoize(old)
	}
}

// Recover creates an Option to set the recover flag to b. When set to
// true, this causes the parser to recover from panics and convert it
// to an error. Setting it to false can be useful while debugging to
// access the full stack trace.
//
// The default is true.
func Recover(b bool) Option {
	return func(p *parser) Option {
		old := p.recover
		p.recover = b
		return Recover(old)
	}
}

// ParseFile parses the file identified by filename.
func ParseFile(filename string, opts ...Option) (interface{}, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseReader(filename, f, opts...)
}

// ParseReader parses the data from r using filename as information in the
// error messages.
func ParseReader(filename string, r io.Reader, opts ...Option) (interface{}, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return Parse(filename, b, opts...)
}

// Parse parses the data from b using filename as information in the
// error messages.
func Parse(filename string, b []byte, opts ...Option) (interface{}, error) {
	return newParser(filename, b, opts...).parse(g)
}

// position records a position in the text.
type position struct {
	line, col, offset int
}

func (p position) String() string {
	return fmt.Sprintf("%d:%d [%d]", p.line, p.col, p.offset)
}

// savepoint stores all state required to go back to this point in the
// parser.
type savepoint struct {
	position
	rn rune
	w  int
}

type current struct {
	pos  position // start position of the match
	text []byte   // raw text of the match
}

// the AST types...

type grammar struct {
	pos   position
	rules []*rule
}

type rule struct {
	pos         position
	name        string
	displayName string
	expr        interface{}
}

type choiceExpr struct {
	pos          position
	alternatives []interface{}
}

type actionExpr struct {
	pos  position
	expr interface{}
	run  func(*parser) (interface{}, error)
}

type seqExpr struct {
	pos   position
	exprs []interface{}
}

type labeledExpr struct {
	pos   position
	label string
	expr  interface{}
}

type expr struct {
	pos  position
	expr interface{}
}

type andExpr expr
type notExpr expr
type zeroOrOneExpr expr
type zeroOrMoreExpr expr
type oneOrMoreExpr expr

type ruleRefExpr struct {
	pos  position
	name string
}

type andCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type notCodeExpr struct {
	pos position
	run func(*parser) (bool, error)
}

type litMatcher struct {
	pos        position
	val        string
	ignoreCase bool
}

type charClassMatcher struct {
	pos        position
	val        string
	chars      []rune
	ranges     []rune
	classes    []*unicode.RangeTable
	ignoreCase bool
	inverted   bool
}

type anyMatcher position

// errList cumulates the errors found by the parser.
type errList []error

func (e *errList) add(err error) {
	*e = append(*e, err)
}

func (e errList) err() error {
	if len(e) == 0 {
		return nil
	}
	e.dedupe()
	return e
}

func (e *errList) dedupe() {
	var cleaned []error
	set := make(map[string]bool)
	for _, err := range *e {
		if msg := err.Error(); !set[msg] {
			set[msg] = true
			cleaned = append(cleaned, err)
		}
	}
	*e = cleaned
}

func (e errList) Error() string {
	switch len(e) {
	case 0:
		return ""
	case 1:
		return e[0].Error()
	default:
		var buf bytes.Buffer

		for i, err := range e {
			if i > 0 {
				buf.WriteRune('\n')
			}
			buf.WriteString(err.Error())
		}
		return buf.String()
	}
}

// parserError wraps an error with a prefix indicating the rule in which
// the error occurred. The original error is stored in the Inner field.
type parserError struct {
	Inner  error
	pos    position
	prefix string
}

// Error returns the error message.
func (p *parserError) Error() string {
	return p.prefix + ": " + p.Inner.Error()
}

// newParser creates a parser with the specified input source and options.
func newParser(filename string, b []byte, opts ...Option) *parser {
	p := &parser{
		filename: filename,
		errs:     new(errList),
		data:     b,
		pt:       savepoint{position: position{line: 1}},
		recover:  true,
	}
	p.setOptions(opts)
	return p
}

// setOptions applies the options to the parser.
func (p *parser) setOptions(opts []Option) {
	for _, opt := range opts {
		opt(p)
	}
}

type resultTuple struct {
	v   interface{}
	b   bool
	end savepoint
}

type parser struct {
	filename string
	pt       savepoint
	cur      current

	data []byte
	errs *errList

	recover bool
	debug   bool
	depth   int

	memoize bool
	// memoization table for the packrat algorithm:
	// map[offset in source] map[expression or rule] {value, match}
	memo map[int]map[interface{}]resultTuple

	// rules table, maps the rule identifier to the rule node
	rules map[string]*rule
	// variables stack, map of label to value
	vstack []map[string]interface{}
	// rule stack, allows identification of the current rule in errors
	rstack []*rule

	// stats
	exprCnt int
}

// push a variable set on the vstack.
func (p *parser) pushV() {
	if cap(p.vstack) == len(p.vstack) {
		// create new empty slot in the stack
		p.vstack = append(p.vstack, nil)
	} else {
		// slice to 1 more
		p.vstack = p.vstack[:len(p.vstack)+1]
	}

	// get the last args set
	m := p.vstack[len(p.vstack)-1]
	if m != nil && len(m) == 0 {
		// empty map, all good
		return
	}

	m = make(map[string]interface{})
	p.vstack[len(p.vstack)-1] = m
}

// pop a variable set from the vstack.
func (p *parser) popV() {
	// if the map is not empty, clear it
	m := p.vstack[len(p.vstack)-1]
	if len(m) > 0 {
		// GC that map
		p.vstack[len(p.vstack)-1] = nil
	}
	p.vstack = p.vstack[:len(p.vstack)-1]
}

func (p *parser) print(prefix, s string) string {
	if !p.debug {
		return s
	}

	fmt.Printf("%s %d:%d:%d: %s [%#U]\n",
		prefix, p.pt.line, p.pt.col, p.pt.offset, s, p.pt.rn)
	return s
}

func (p *parser) in(s string) string {
	p.depth++
	return p.print(strings.Repeat(" ", p.depth)+">", s)
}

func (p *parser) out(s string) string {
	p.depth--
	return p.print(strings.Repeat(" ", p.depth)+"<", s)
}

func (p *parser) addErr(err error) {
	p.addErrAt(err, p.pt.position)
}

func (p *parser) addErrAt(err error, pos position) {
	var buf bytes.Buffer
	if p.filename != "" {
		buf.WriteString(p.filename)
	}
	if buf.Len() > 0 {
		buf.WriteString(":")
	}
	buf.WriteString(fmt.Sprintf("%d:%d (%d)", pos.line, pos.col, pos.offset))
	if len(p.rstack) > 0 {
		if buf.Len() > 0 {
			buf.WriteString(": ")
		}
		rule := p.rstack[len(p.rstack)-1]
		if rule.displayName != "" {
			buf.WriteString("rule " + rule.displayName)
		} else {
			buf.WriteString("rule " + rule.name)
		}
	}
	pe := &parserError{Inner: err, pos: pos, prefix: buf.String()}
	p.errs.add(pe)
}

// read advances the parser to the next rune.
func (p *parser) read() {
	p.pt.offset += p.pt.w
	rn, n := utf8.DecodeRune(p.data[p.pt.offset:])
	p.pt.rn = rn
	p.pt.w = n
	p.pt.col++
	if rn == '\n' {
		p.pt.line++
		p.pt.col = 0
	}

	if rn == utf8.RuneError {
		if n == 1 {
			p.addErr(errInvalidEncoding)
		}
	}
}

// restore parser position to the savepoint pt.
func (p *parser) restore(pt savepoint) {
	if p.debug {
		defer p.out(p.in("restore"))
	}
	if pt.offset == p.pt.offset {
		return
	}
	p.pt = pt
}

// get the slice of bytes from the savepoint start to the current position.
func (p *parser) sliceFrom(start savepoint) []byte {
	return p.data[start.position.offset:p.pt.position.offset]
}

func (p *parser) getMemoized(node interface{}) (resultTuple, bool) {
	if len(p.memo) == 0 {
		return resultTuple{}, false
	}
	m := p.memo[p.pt.offset]
	if len(m) == 0 {
		return resultTuple{}, false
	}
	res, ok := m[node]
	return res, ok
}

func (p *parser) setMemoized(pt savepoint, node interface{}, tuple resultTuple) {
	if p.memo == nil {
		p.memo = make(map[int]map[interface{}]resultTuple)
	}
	m := p.memo[pt.offset]
	if m == nil {
		m = make(map[interface{}]resultTuple)
		p.memo[pt.offset] = m
	}
	m[node] = tuple
}

func (p *parser) buildRulesTable(g *grammar) {
	p.rules = make(map[string]*rule, len(g.rules))
	for _, r := range g.rules {
		p.rules[r.name] = r
	}
}

func (p *parser) parse(g *grammar) (val interface{}, err error) {
	if len(g.rules) == 0 {
		p.addErr(errNoRule)
		return nil, p.errs.err()
	}

	// TODO : not super critical but this could be generated
	p.buildRulesTable(g)

	if p.recover {
		// panic can be used in action code to stop parsing immediately
		// and return the panic as an error.
		defer func() {
			if e := recover(); e != nil {
				if p.debug {
					defer p.out(p.in("panic handler"))
				}
				val = nil
				switch e := e.(type) {
				case error:
					p.addErr(e)
				default:
					p.addErr(fmt.Errorf("%v", e))
				}
				err = p.errs.err()
			}
		}()
	}

	// start rule is rule [0]
	p.read() // advance to first rune
	val, ok := p.parseRule(g.rules[0])
	if !ok {
		if len(*p.errs) == 0 {
			// make sure this doesn't go out silently
			p.addErr(errNoMatch)
		}
		return nil, p.errs.err()
	}
	return val, p.errs.err()
}

func (p *parser) parseRule(rule *rule) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRule " + rule.name))
	}

	if p.memoize {
		res, ok := p.getMemoized(rule)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
	}

	start := p.pt
	p.rstack = append(p.rstack, rule)
	p.pushV()
	val, ok := p.parseExpr(rule.expr)
	p.popV()
	p.rstack = p.rstack[:len(p.rstack)-1]
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}

	if p.memoize {
		p.setMemoized(start, rule, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseExpr(expr interface{}) (interface{}, bool) {
	var pt savepoint
	var ok bool

	if p.memoize {
		res, ok := p.getMemoized(expr)
		if ok {
			p.restore(res.end)
			return res.v, res.b
		}
		pt = p.pt
	}

	p.exprCnt++
	var val interface{}
	switch expr := expr.(type) {
	case *actionExpr:
		val, ok = p.parseActionExpr(expr)
	case *andCodeExpr:
		val, ok = p.parseAndCodeExpr(expr)
	case *andExpr:
		val, ok = p.parseAndExpr(expr)
	case *anyMatcher:
		val, ok = p.parseAnyMatcher(expr)
	case *charClassMatcher:
		val, ok = p.parseCharClassMatcher(expr)
	case *choiceExpr:
		val, ok = p.parseChoiceExpr(expr)
	case *labeledExpr:
		val, ok = p.parseLabeledExpr(expr)
	case *litMatcher:
		val, ok = p.parseLitMatcher(expr)
	case *notCodeExpr:
		val, ok = p.parseNotCodeExpr(expr)
	case *notExpr:
		val, ok = p.parseNotExpr(expr)
	case *oneOrMoreExpr:
		val, ok = p.parseOneOrMoreExpr(expr)
	case *ruleRefExpr:
		val, ok = p.parseRuleRefExpr(expr)
	case *seqExpr:
		val, ok = p.parseSeqExpr(expr)
	case *zeroOrMoreExpr:
		val, ok = p.parseZeroOrMoreExpr(expr)
	case *zeroOrOneExpr:
		val, ok = p.parseZeroOrOneExpr(expr)
	default:
		panic(fmt.Sprintf("unknown expression type %T", expr))
	}
	if p.memoize {
		p.setMemoized(pt, expr, resultTuple{val, ok, p.pt})
	}
	return val, ok
}

func (p *parser) parseActionExpr(act *actionExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseActionExpr"))
	}

	start := p.pt
	val, ok := p.parseExpr(act.expr)
	if ok {
		p.cur.pos = start.position
		p.cur.text = p.sliceFrom(start)
		actVal, err := act.run(p)
		if err != nil {
			p.addErrAt(err, start.position)
		}
		val = actVal
	}
	if ok && p.debug {
		p.print(strings.Repeat(" ", p.depth)+"MATCH", string(p.sliceFrom(start)))
	}
	return val, ok
}

func (p *parser) parseAndCodeExpr(and *andCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndCodeExpr"))
	}

	ok, err := and.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, ok
}

func (p *parser) parseAndExpr(and *andExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAndExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(and.expr)
	p.popV()
	p.restore(pt)
	return nil, ok
}

func (p *parser) parseAnyMatcher(any *anyMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseAnyMatcher"))
	}

	if p.pt.rn != utf8.RuneError {
		start := p.pt
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseCharClassMatcher(chr *charClassMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseCharClassMatcher"))
	}

	cur := p.pt.rn
	// can't match EOF
	if cur == utf8.RuneError {
		return nil, false
	}
	start := p.pt
	if chr.ignoreCase {
		cur = unicode.ToLower(cur)
	}

	// try to match in the list of available chars
	for _, rn := range chr.chars {
		if rn == cur {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of ranges
	for i := 0; i < len(chr.ranges); i += 2 {
		if cur >= chr.ranges[i] && cur <= chr.ranges[i+1] {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	// try to match in the list of Unicode classes
	for _, cl := range chr.classes {
		if unicode.Is(cl, cur) {
			if chr.inverted {
				return nil, false
			}
			p.read()
			return p.sliceFrom(start), true
		}
	}

	if chr.inverted {
		p.read()
		return p.sliceFrom(start), true
	}
	return nil, false
}

func (p *parser) parseChoiceExpr(ch *choiceExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseChoiceExpr"))
	}

	for _, alt := range ch.alternatives {
		p.pushV()
		val, ok := p.parseExpr(alt)
		p.popV()
		if ok {
			return val, ok
		}
	}
	return nil, false
}

func (p *parser) parseLabeledExpr(lab *labeledExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLabeledExpr"))
	}

	p.pushV()
	val, ok := p.parseExpr(lab.expr)
	p.popV()
	if ok && lab.label != "" {
		m := p.vstack[len(p.vstack)-1]
		m[lab.label] = val
	}
	return val, ok
}

func (p *parser) parseLitMatcher(lit *litMatcher) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseLitMatcher"))
	}

	start := p.pt
	for _, want := range lit.val {
		cur := p.pt.rn
		if lit.ignoreCase {
			cur = unicode.ToLower(cur)
		}
		if cur != want {
			p.restore(start)
			return nil, false
		}
		p.read()
	}
	return p.sliceFrom(start), true
}

func (p *parser) parseNotCodeExpr(not *notCodeExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotCodeExpr"))
	}

	ok, err := not.run(p)
	if err != nil {
		p.addErr(err)
	}
	return nil, !ok
}

func (p *parser) parseNotExpr(not *notExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseNotExpr"))
	}

	pt := p.pt
	p.pushV()
	_, ok := p.parseExpr(not.expr)
	p.popV()
	p.restore(pt)
	return nil, !ok
}

func (p *parser) parseOneOrMoreExpr(expr *oneOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseOneOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			if len(vals) == 0 {
				// did not match once, no match
				return nil, false
			}
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseRuleRefExpr(ref *ruleRefExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseRuleRefExpr " + ref.name))
	}

	if ref.name == "" {
		panic(fmt.Sprintf("%s: invalid rule: missing name", ref.pos))
	}

	rule := p.rules[ref.name]
	if rule == nil {
		p.addErr(fmt.Errorf("undefined rule: %s", ref.name))
		return nil, false
	}
	return p.parseRule(rule)
}

func (p *parser) parseSeqExpr(seq *seqExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseSeqExpr"))
	}

	var vals []interface{}

	pt := p.pt
	for _, expr := range seq.exprs {
		val, ok := p.parseExpr(expr)
		if !ok {
			p.restore(pt)
			return nil, false
		}
		vals = append(vals, val)
	}
	return vals, true
}

func (p *parser) parseZeroOrMoreExpr(expr *zeroOrMoreExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrMoreExpr"))
	}

	var vals []interface{}

	for {
		p.pushV()
		val, ok := p.parseExpr(expr.expr)
		p.popV()
		if !ok {
			return vals, true
		}
		vals = append(vals, val)
	}
}

func (p *parser) parseZeroOrOneExpr(expr *zeroOrOneExpr) (interface{}, bool) {
	if p.debug {
		defer p.out(p.in("parseZeroOrOneExpr"))
	}

	p.pushV()
	val, _ := p.parseExpr(expr.expr)
	p.popV()
	// whether it matched or not, consider it a match
	return val, true
}

func rangeTable(class string) *unicode.RangeTable {
	if rt, ok := unicode.Categories[class]; ok {
		return rt
	}
	if rt, ok := unicode.Properties[class]; ok {
		return rt
	}
	if rt, ok := unicode.Scripts[class]; ok {
		return rt
	}

	// cannot happen
	panic(fmt.Sprintf("invalid Unicode class: %s", class))
}
