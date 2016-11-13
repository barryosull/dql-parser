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

//import "fmt"

var g = &grammar{
	rules: []*rule{
		{
			name: "DOC",
			pos:  position{line: 7, col: 1, offset: 48},
			expr: &actionExpr{
				pos: position{line: 7, col: 7, offset: 54},
				run: (*parser).callonDOC1,
				expr: &seqExpr{
					pos: position{line: 7, col: 7, offset: 54},
					exprs: []interface{}{
						&labeledExpr{
							pos:   position{line: 7, col: 7, offset: 54},
							label: "stats",
							expr: &ruleRefExpr{
								pos:  position{line: 7, col: 13, offset: 60},
								name: "STATEMENTBLOCK",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 7, col: 28, offset: 75},
							name: "EOF",
						},
					},
				},
			},
		},
		{
			name: "STATEMENTBLOCK",
			pos:  position{line: 11, col: 1, offset: 107},
			expr: &zeroOrMoreExpr{
				pos: position{line: 11, col: 18, offset: 124},
				expr: &ruleRefExpr{
					pos:  position{line: 11, col: 18, offset: 124},
					name: "STATEMENT",
				},
			},
		},
		{
			name: "STATEMENT",
			pos:  position{line: 13, col: 1, offset: 136},
			expr: &choiceExpr{
				pos: position{line: 13, col: 13, offset: 148},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 13, col: 13, offset: 148},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 13, col: 13, offset: 148},
								name: "RETURN",
							},
							&ruleRefExpr{
								pos:  position{line: 13, col: 20, offset: 155},
								name: "SEMI",
							},
						},
					},
					&seqExpr{
						pos: position{line: 13, col: 27, offset: 162},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 13, col: 27, offset: 162},
								name: "EXPRESSION",
							},
							&ruleRefExpr{
								pos:  position{line: 13, col: 38, offset: 173},
								name: "SEMI",
							},
						},
					},
				},
			},
		},
		{
			name: "RETURN",
			pos:  position{line: 15, col: 1, offset: 179},
			expr: &seqExpr{
				pos: position{line: 15, col: 10, offset: 188},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 15, col: 10, offset: 188},
						val:        "return",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 15, col: 19, offset: 197},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 15, col: 21, offset: 199},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "EXPRESSION",
			pos:  position{line: 17, col: 1, offset: 211},
			expr: &choiceExpr{
				pos: position{line: 17, col: 14, offset: 224},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 17, col: 14, offset: 224},
						name: "ARITHMETIC",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 27, offset: 237},
						name: "ASSIGNMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 39, offset: 249},
						name: "LOGICAL",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 49, offset: 259},
						name: "COMPARISON",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 62, offset: 272},
						name: "ATOMIC",
					},
				},
			},
		},
		{
			name: "ATOMIC",
			pos:  position{line: 19, col: 1, offset: 280},
			expr: &choiceExpr{
				pos: position{line: 19, col: 10, offset: 289},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 19, col: 10, offset: 289},
						name: "PARENTHESIS",
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 24, offset: 303},
						name: "NEW",
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 30, offset: 309},
						name: "METHODCALL",
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 43, offset: 322},
						name: "OBJECTACCESS",
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 58, offset: 337},
						name: "LITERAL",
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 68, offset: 347},
						name: "UNARY",
					},
				},
			},
		},
		{
			name: "LITERAL",
			pos:  position{line: 21, col: 1, offset: 354},
			expr: &choiceExpr{
				pos: position{line: 21, col: 11, offset: 364},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 21, col: 11, offset: 364},
						name: "STRING",
					},
					&ruleRefExpr{
						pos:  position{line: 21, col: 20, offset: 373},
						name: "FLOAT",
					},
					&ruleRefExpr{
						pos:  position{line: 21, col: 28, offset: 381},
						name: "BOOLEAN",
					},
					&ruleRefExpr{
						pos:  position{line: 21, col: 38, offset: 391},
						name: "NULL",
					},
					&actionExpr{
						pos: position{line: 21, col: 45, offset: 398},
						run: (*parser).callonLITERAL6,
						expr: &ruleRefExpr{
							pos:  position{line: 21, col: 45, offset: 398},
							name: "INT",
						},
					},
				},
			},
		},
		{
			name: "NEW",
			pos:  position{line: 25, col: 1, offset: 435},
			expr: &seqExpr{
				pos: position{line: 25, col: 7, offset: 441},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 25, col: 7, offset: 441},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 11, offset: 445},
						name: "IDENTIFIER",
					},
					&litMatcher{
						pos:        position{line: 25, col: 22, offset: 456},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 26, offset: 460},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 28, offset: 462},
						name: "ARGUMENTS",
					},
				},
			},
		},
		{
			name: "BOOLEAN",
			pos:  position{line: 27, col: 1, offset: 473},
			expr: &choiceExpr{
				pos: position{line: 27, col: 12, offset: 484},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 27, col: 12, offset: 484},
						val:        "true",
						ignoreCase: false,
					},
					&actionExpr{
						pos: position{line: 27, col: 19, offset: 491},
						run: (*parser).callonBOOLEAN3,
						expr: &litMatcher{
							pos:        position{line: 27, col: 19, offset: 491},
							val:        "false",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "NULL",
			pos:  position{line: 31, col: 1, offset: 531},
			expr: &actionExpr{
				pos: position{line: 31, col: 8, offset: 538},
				run: (*parser).callonNULL1,
				expr: &litMatcher{
					pos:        position{line: 31, col: 8, offset: 538},
					val:        "null",
					ignoreCase: false,
				},
			},
		},
		{
			name: "STRING",
			pos:  position{line: 35, col: 1, offset: 573},
			expr: &actionExpr{
				pos: position{line: 35, col: 10, offset: 582},
				run: (*parser).callonSTRING1,
				expr: &seqExpr{
					pos: position{line: 35, col: 10, offset: 582},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 35, col: 10, offset: 582},
							val:        "\"",
							ignoreCase: false,
						},
						&zeroOrMoreExpr{
							pos: position{line: 35, col: 15, offset: 587},
							expr: &charClassMatcher{
								pos:        position{line: 35, col: 15, offset: 587},
								val:        "[a-zA-Z0-9]",
								ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&litMatcher{
							pos:        position{line: 35, col: 28, offset: 600},
							val:        "\"",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "INT",
			pos:  position{line: 39, col: 1, offset: 635},
			expr: &actionExpr{
				pos: position{line: 39, col: 7, offset: 641},
				run: (*parser).callonINT1,
				expr: &oneOrMoreExpr{
					pos: position{line: 39, col: 7, offset: 641},
					expr: &charClassMatcher{
						pos:        position{line: 39, col: 7, offset: 641},
						val:        "[0-9]",
						ranges:     []rune{'0', '9'},
						ignoreCase: false,
						inverted:   false,
					},
				},
			},
		},
		{
			name: "FLOAT",
			pos:  position{line: 43, col: 1, offset: 675},
			expr: &actionExpr{
				pos: position{line: 43, col: 9, offset: 683},
				run: (*parser).callonFLOAT1,
				expr: &seqExpr{
					pos: position{line: 43, col: 9, offset: 683},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 43, col: 9, offset: 683},
							expr: &charClassMatcher{
								pos:        position{line: 43, col: 9, offset: 683},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&charClassMatcher{
							pos:        position{line: 43, col: 16, offset: 690},
							val:        "[.]",
							chars:      []rune{'.'},
							ignoreCase: false,
							inverted:   false,
						},
						&oneOrMoreExpr{
							pos: position{line: 43, col: 20, offset: 694},
							expr: &charClassMatcher{
								pos:        position{line: 43, col: 20, offset: 694},
								val:        "[0-9]",
								ranges:     []rune{'0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name: "IDENTIFIER",
			pos:  position{line: 47, col: 1, offset: 730},
			expr: &actionExpr{
				pos: position{line: 47, col: 14, offset: 743},
				run: (*parser).callonIDENTIFIER1,
				expr: &seqExpr{
					pos: position{line: 47, col: 14, offset: 743},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 47, col: 14, offset: 743},
							expr: &charClassMatcher{
								pos:        position{line: 47, col: 14, offset: 743},
								val:        "[a-zA-Z]",
								ranges:     []rune{'a', 'z', 'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 47, col: 24, offset: 753},
							expr: &charClassMatcher{
								pos:        position{line: 47, col: 24, offset: 753},
								val:        "[a-zA-Z0-9\\\\]",
								chars:      []rune{'\\'},
								ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
								ignoreCase: false,
								inverted:   false,
							},
						},
					},
				},
			},
		},
		{
			name: "PARENTHESIS",
			pos:  position{line: 51, col: 1, offset: 802},
			expr: &actionExpr{
				pos: position{line: 51, col: 15, offset: 816},
				run: (*parser).callonPARENTHESIS1,
				expr: &seqExpr{
					pos: position{line: 51, col: 15, offset: 816},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 51, col: 15, offset: 816},
							val:        "(",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 51, col: 19, offset: 820},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 51, col: 21, offset: 822},
							name: "EXPRESSION",
						},
						&ruleRefExpr{
							pos:  position{line: 51, col: 32, offset: 833},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 51, col: 34, offset: 835},
							val:        ")",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "UNARY",
			pos:  position{line: 55, col: 1, offset: 874},
			expr: &choiceExpr{
				pos: position{line: 55, col: 9, offset: 882},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 55, col: 9, offset: 882},
						name: "INCREMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 55, col: 21, offset: 894},
						name: "DECREMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 55, col: 33, offset: 906},
						name: "NEGATE",
					},
					&ruleRefExpr{
						pos:  position{line: 55, col: 42, offset: 915},
						name: "NOT",
					},
					&ruleRefExpr{
						pos:  position{line: 55, col: 48, offset: 921},
						name: "POSITIVE",
					},
				},
			},
		},
		{
			name: "INCREMENT",
			pos:  position{line: 57, col: 1, offset: 931},
			expr: &actionExpr{
				pos: position{line: 57, col: 13, offset: 943},
				run: (*parser).callonINCREMENT1,
				expr: &seqExpr{
					pos: position{line: 57, col: 13, offset: 943},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 57, col: 13, offset: 943},
							name: "OBJECTACCESS",
						},
						&litMatcher{
							pos:        position{line: 57, col: 26, offset: 956},
							val:        "++",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "DECREMENT",
			pos:  position{line: 61, col: 1, offset: 995},
			expr: &actionExpr{
				pos: position{line: 61, col: 13, offset: 1007},
				run: (*parser).callonDECREMENT1,
				expr: &seqExpr{
					pos: position{line: 61, col: 13, offset: 1007},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 61, col: 13, offset: 1007},
							name: "OBJECTACCESS",
						},
						&litMatcher{
							pos:        position{line: 61, col: 26, offset: 1020},
							val:        "--",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "NEGATE",
			pos:  position{line: 65, col: 1, offset: 1059},
			expr: &actionExpr{
				pos: position{line: 65, col: 10, offset: 1068},
				run: (*parser).callonNEGATE1,
				expr: &seqExpr{
					pos: position{line: 65, col: 10, offset: 1068},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 65, col: 10, offset: 1068},
							val:        "-",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 65, col: 14, offset: 1072},
							name: "OBJECTACCESS",
						},
					},
				},
			},
		},
		{
			name: "NOT",
			pos:  position{line: 69, col: 1, offset: 1116},
			expr: &actionExpr{
				pos: position{line: 69, col: 7, offset: 1122},
				run: (*parser).callonNOT1,
				expr: &seqExpr{
					pos: position{line: 69, col: 7, offset: 1122},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 69, col: 7, offset: 1122},
							val:        "!",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 69, col: 11, offset: 1126},
							name: "OBJECTACCESS",
						},
					},
				},
			},
		},
		{
			name: "POSITIVE",
			pos:  position{line: 73, col: 1, offset: 1167},
			expr: &actionExpr{
				pos: position{line: 73, col: 12, offset: 1178},
				run: (*parser).callonPOSITIVE1,
				expr: &seqExpr{
					pos: position{line: 73, col: 12, offset: 1178},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 73, col: 12, offset: 1178},
							val:        "+",
							ignoreCase: false,
						},
						&ruleRefExpr{
							pos:  position{line: 73, col: 16, offset: 1182},
							name: "OBJECTACCESS",
						},
					},
				},
			},
		},
		{
			name: "ARITHMETIC",
			pos:  position{line: 77, col: 1, offset: 1228},
			expr: &actionExpr{
				pos: position{line: 77, col: 14, offset: 1241},
				run: (*parser).callonARITHMETIC1,
				expr: &seqExpr{
					pos: position{line: 77, col: 14, offset: 1241},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 77, col: 14, offset: 1241},
							name: "ATOMIC",
						},
						&ruleRefExpr{
							pos:  position{line: 77, col: 21, offset: 1248},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 77, col: 23, offset: 1250},
							name: "OPERATOR",
						},
						&ruleRefExpr{
							pos:  position{line: 77, col: 32, offset: 1259},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 77, col: 34, offset: 1261},
							name: "EXPRESSION",
						},
					},
				},
			},
		},
		{
			name: "OPERATOR",
			pos:  position{line: 81, col: 1, offset: 1307},
			expr: &choiceExpr{
				pos: position{line: 81, col: 12, offset: 1318},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 81, col: 12, offset: 1318},
						val:        "+",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 81, col: 18, offset: 1324},
						val:        "-",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 81, col: 24, offset: 1330},
						val:        "/",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 81, col: 30, offset: 1336},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 81, col: 36, offset: 1342},
						val:        "%",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "ASSIGNMENT",
			pos:  position{line: 83, col: 1, offset: 1347},
			expr: &seqExpr{
				pos: position{line: 83, col: 14, offset: 1360},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 83, col: 14, offset: 1360},
						name: "OBJECTACCESS",
					},
					&ruleRefExpr{
						pos:  position{line: 83, col: 27, offset: 1373},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 83, col: 29, offset: 1375},
						val:        "=",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 83, col: 33, offset: 1379},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 83, col: 35, offset: 1381},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "LOGICAL",
			pos:  position{line: 85, col: 1, offset: 1393},
			expr: &seqExpr{
				pos: position{line: 85, col: 11, offset: 1403},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 85, col: 11, offset: 1403},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 85, col: 22, offset: 1414},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 85, col: 25, offset: 1417},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 85, col: 25, offset: 1417},
								val:        "and",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 85, col: 33, offset: 1425},
								val:        "or",
								ignoreCase: false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 85, col: 39, offset: 1431},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 85, col: 41, offset: 1433},
						name: "ATOMIC",
					},
				},
			},
		},
		{
			name: "COMPARISON",
			pos:  position{line: 87, col: 1, offset: 1441},
			expr: &seqExpr{
				pos: position{line: 87, col: 14, offset: 1454},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 87, col: 14, offset: 1454},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 87, col: 25, offset: 1465},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 87, col: 28, offset: 1468},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 87, col: 28, offset: 1468},
								val:        "===",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 87, col: 36, offset: 1476},
								val:        "!==",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 87, col: 44, offset: 1484},
								val:        "==",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 87, col: 51, offset: 1491},
								val:        "!=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 87, col: 58, offset: 1498},
								val:        "<=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 87, col: 65, offset: 1505},
								val:        ">=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 87, col: 72, offset: 1512},
								val:        "<",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 87, col: 78, offset: 1518},
								val:        ">",
								ignoreCase: false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 87, col: 83, offset: 1523},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 87, col: 85, offset: 1525},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "METHODCALL",
			pos:  position{line: 89, col: 1, offset: 1537},
			expr: &seqExpr{
				pos: position{line: 89, col: 14, offset: 1550},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 89, col: 14, offset: 1550},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 89, col: 25, offset: 1561},
						name: "ARGUMENTS",
					},
				},
			},
		},
		{
			name: "OBJECTACCESS",
			pos:  position{line: 91, col: 1, offset: 1572},
			expr: &seqExpr{
				pos: position{line: 91, col: 16, offset: 1587},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 91, col: 16, offset: 1587},
						expr: &seqExpr{
							pos: position{line: 91, col: 17, offset: 1588},
							exprs: []interface{}{
								&choiceExpr{
									pos: position{line: 91, col: 18, offset: 1589},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 91, col: 18, offset: 1589},
											name: "METHODCALL",
										},
										&ruleRefExpr{
											pos:  position{line: 91, col: 31, offset: 1602},
											name: "IDENTIFIER",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 91, col: 43, offset: 1614},
									val:        "->",
									ignoreCase: false,
								},
							},
						},
					},
					&choiceExpr{
						pos: position{line: 91, col: 51, offset: 1622},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 91, col: 51, offset: 1622},
								name: "METHODCALL",
							},
							&ruleRefExpr{
								pos:  position{line: 91, col: 64, offset: 1635},
								name: "IDENTIFIER",
							},
						},
					},
				},
			},
		},
		{
			name: "ARGUMENTS",
			pos:  position{line: 93, col: 1, offset: 1648},
			expr: &seqExpr{
				pos: position{line: 93, col: 13, offset: 1660},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 93, col: 13, offset: 1660},
						val:        "(",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 93, col: 17, offset: 1664},
						expr: &ruleRefExpr{
							pos:  position{line: 93, col: 17, offset: 1664},
							name: "ARGUMENTLIST",
						},
					},
					&litMatcher{
						pos:        position{line: 93, col: 31, offset: 1678},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "ARGUMENTLIST",
			pos:  position{line: 95, col: 1, offset: 1683},
			expr: &seqExpr{
				pos: position{line: 95, col: 16, offset: 1698},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 95, col: 16, offset: 1698},
						expr: &seqExpr{
							pos: position{line: 95, col: 17, offset: 1699},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 95, col: 17, offset: 1699},
									name: "EXPRESSION",
								},
								&ruleRefExpr{
									pos:  position{line: 95, col: 28, offset: 1710},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 95, col: 30, offset: 1712},
									val:        ",",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 95, col: 34, offset: 1716},
									name: "_",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 95, col: 38, offset: 1720},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "_",
			pos:  position{line: 97, col: 1, offset: 1732},
			expr: &zeroOrMoreExpr{
				pos: position{line: 97, col: 5, offset: 1738},
				expr: &choiceExpr{
					pos: position{line: 97, col: 7, offset: 1740},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 97, col: 7, offset: 1740},
							name: "Whitespace",
						},
						&ruleRefExpr{
							pos:  position{line: 97, col: 20, offset: 1753},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "SEMI",
			pos:  position{line: 99, col: 1, offset: 1761},
			expr: &seqExpr{
				pos: position{line: 99, col: 8, offset: 1768},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 99, col: 8, offset: 1768},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 99, col: 10, offset: 1770},
						val:        ";",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 99, col: 14, offset: 1774},
						name: "_",
					},
				},
			},
		},
		{
			name: "Whitespace",
			pos:  position{line: 101, col: 1, offset: 1777},
			expr: &charClassMatcher{
				pos:        position{line: 101, col: 14, offset: 1792},
				val:        "[ \\t\\r]",
				chars:      []rune{' ', '\t', '\r'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EOL",
			pos:  position{line: 102, col: 1, offset: 1800},
			expr: &litMatcher{
				pos:        position{line: 102, col: 7, offset: 1808},
				val:        "\n",
				ignoreCase: false,
			},
		},
		{
			name: "EOF",
			pos:  position{line: 103, col: 1, offset: 1813},
			expr: &notExpr{
				pos: position{line: 103, col: 7, offset: 1821},
				expr: &anyMatcher{
					line: 103, col: 8, offset: 1822,
				},
			},
		},
	},
}

func (c *current) onDOC1(stats interface{}) (interface{}, error) {
	return stats, nil
}

func (p *parser) callonDOC1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDOC1(stack["stats"])
}

func (c *current) onLITERAL6() (interface{}, error) {
	return "LITERAL", nil
}

func (p *parser) callonLITERAL6() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onLITERAL6()
}

func (c *current) onBOOLEAN3() (interface{}, error) {
	return "BOOLEAN", nil
}

func (p *parser) callonBOOLEAN3() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onBOOLEAN3()
}

func (c *current) onNULL1() (interface{}, error) {
	return "NULL", nil
}

func (p *parser) callonNULL1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNULL1()
}

func (c *current) onSTRING1() (interface{}, error) {
	return "STRING", nil
}

func (p *parser) callonSTRING1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onSTRING1()
}

func (c *current) onINT1() (interface{}, error) {
	return "INT", nil
}

func (p *parser) callonINT1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onINT1()
}

func (c *current) onFLOAT1() (interface{}, error) {
	return "FLOAT", nil
}

func (p *parser) callonFLOAT1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFLOAT1()
}

func (c *current) onIDENTIFIER1() (interface{}, error) {
	return "IDENTIFIER", nil
}

func (p *parser) callonIDENTIFIER1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIDENTIFIER1()
}

func (c *current) onPARENTHESIS1() (interface{}, error) {
	return "PARENTHESIS", nil
}

func (p *parser) callonPARENTHESIS1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPARENTHESIS1()
}

func (c *current) onINCREMENT1() (interface{}, error) {
	return "INCREMENT", nil
}

func (p *parser) callonINCREMENT1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onINCREMENT1()
}

func (c *current) onDECREMENT1() (interface{}, error) {
	return "DECREMENT", nil
}

func (p *parser) callonDECREMENT1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onDECREMENT1()
}

func (c *current) onNEGATE1() (interface{}, error) {
	return "NEGATE", nil
}

func (p *parser) callonNEGATE1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNEGATE1()
}

func (c *current) onNOT1() (interface{}, error) {
	return "NOT", nil
}

func (p *parser) callonNOT1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNOT1()
}

func (c *current) onPOSITIVE1() (interface{}, error) {
	return "POSITIVE", nil
}

func (p *parser) callonPOSITIVE1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onPOSITIVE1()
}

func (c *current) onARITHMETIC1() (interface{}, error) {
	return "ARITHMETIC", nil
}

func (p *parser) callonARITHMETIC1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onARITHMETIC1()
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
