package parser

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

func emit(typ TokenType, val interface{}) {
	//if val == nil {
	//	val = ;
	//}
	//GetInstanceTokenList().Append(NewToken(typ, val.(string)));
}

var g = &grammar{
	rules: []*rule{
		{
			name: "FILE",
			pos:  position{line: 19, col: 1, offset: 370},
			expr: &seqExpr{
				pos: position{line: 19, col: 8, offset: 377},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 19, col: 8, offset: 377},
						name: "STATEMENTS",
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 19, offset: 388},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 21, offset: 390},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "STATEMENTS",
			pos:  position{line: 21, col: 1, offset: 395},
			expr: &labeledExpr{
				pos:   position{line: 21, col: 14, offset: 408},
				label: "statements",
				expr: &zeroOrMoreExpr{
					pos: position{line: 21, col: 25, offset: 419},
					expr: &choiceExpr{
						pos: position{line: 21, col: 26, offset: 420},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 21, col: 26, offset: 420},
								name: "BLOCK_STATEMENT",
							},
							&ruleRefExpr{
								pos:  position{line: 21, col: 43, offset: 437},
								name: "CREATE_OBJECT",
							},
							&ruleRefExpr{
								pos:  position{line: 21, col: 59, offset: 453},
								name: "CREATE_CLASS",
							},
						},
					},
				},
			},
		},
		{
			name: "CREATE_OBJECT",
			pos:  position{line: 23, col: 1, offset: 469},
			expr: &seqExpr{
				pos: position{line: 23, col: 17, offset: 485},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 23, col: 17, offset: 485},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 23, col: 19, offset: 487},
						expr: &ruleRefExpr{
							pos:  position{line: 23, col: 19, offset: 487},
							name: "CREATE_NAMESPACE_OBJECT",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 44, offset: 512},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 23, col: 46, offset: 514},
						expr: &ruleRefExpr{
							pos:  position{line: 23, col: 46, offset: 514},
							name: "NAMESPACE",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 57, offset: 525},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 23, col: 59, offset: 527},
						val:        ";",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "BLOCK_STATEMENT",
			pos:  position{line: 25, col: 1, offset: 532},
			expr: &seqExpr{
				pos: position{line: 25, col: 19, offset: 550},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 25, col: 19, offset: 550},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 21, offset: 552},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 31, offset: 562},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 25, col: 33, offset: 564},
						val:        ":",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 37, offset: 568},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 25, col: 39, offset: 570},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 43, offset: 574},
						name: "STATEMENTS",
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 54, offset: 585},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 25, col: 56, offset: 587},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CREATE_NAMESPACE_OBJECT",
			pos:  position{line: 27, col: 1, offset: 592},
			expr: &actionExpr{
				pos: position{line: 27, col: 27, offset: 618},
				run: (*parser).callonCREATE_NAMESPACE_OBJECT1,
				expr: &seqExpr{
					pos: position{line: 27, col: 27, offset: 618},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 27, col: 27, offset: 618},
							val:        "create",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 27, col: 37, offset: 628},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 27, col: 39, offset: 630},
							label: "typ",
							expr: &ruleRefExpr{
								pos:  position{line: 27, col: 43, offset: 634},
								name: "NAMESPACE_OBJECT_TYPE",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 27, col: 65, offset: 656},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 27, col: 67, offset: 658},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 27, col: 72, offset: 663},
								name: "QUOTED_NAME",
							},
						},
					},
				},
			},
		},
		{
			name: "NAMESPACE_OBJECT_TYPE",
			pos:  position{line: 34, col: 1, offset: 816},
			expr: &actionExpr{
				pos: position{line: 34, col: 25, offset: 840},
				run: (*parser).callonNAMESPACE_OBJECT_TYPE1,
				expr: &labeledExpr{
					pos:   position{line: 34, col: 25, offset: 840},
					label: "typ",
					expr: &choiceExpr{
						pos: position{line: 34, col: 30, offset: 845},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 34, col: 30, offset: 845},
								val:        "database",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 34, col: 43, offset: 858},
								val:        "domain",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 34, col: 54, offset: 869},
								val:        "context",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 34, col: 66, offset: 881},
								val:        "aggregate",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "NAMESPACE",
			pos:  position{line: 38, col: 1, offset: 931},
			expr: &zeroOrMoreExpr{
				pos: position{line: 38, col: 13, offset: 943},
				expr: &choiceExpr{
					pos: position{line: 38, col: 14, offset: 944},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 38, col: 14, offset: 944},
							name: "USING_DATABASE",
						},
						&ruleRefExpr{
							pos:  position{line: 38, col: 31, offset: 961},
							name: "FOR_DOMAIN",
						},
						&ruleRefExpr{
							pos:  position{line: 38, col: 44, offset: 974},
							name: "IN_CONTEXT",
						},
						&ruleRefExpr{
							pos:  position{line: 38, col: 57, offset: 987},
							name: "WITHIN_AGGREGATE",
						},
					},
				},
			},
		},
		{
			name: "USING_DATABASE",
			pos:  position{line: 40, col: 1, offset: 1007},
			expr: &actionExpr{
				pos: position{line: 40, col: 18, offset: 1024},
				run: (*parser).callonUSING_DATABASE1,
				expr: &seqExpr{
					pos: position{line: 40, col: 18, offset: 1024},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 40, col: 18, offset: 1024},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 40, col: 20, offset: 1026},
							val:        "using",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 40, col: 29, offset: 1035},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 40, col: 31, offset: 1037},
							val:        "database",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 40, col: 43, offset: 1049},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 40, col: 45, offset: 1051},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 40, col: 50, offset: 1056},
								name: "QUOTED_NAME",
							},
						},
					},
				},
			},
		},
		{
			name: "FOR_DOMAIN",
			pos:  position{line: 45, col: 1, offset: 1134},
			expr: &actionExpr{
				pos: position{line: 45, col: 14, offset: 1147},
				run: (*parser).callonFOR_DOMAIN1,
				expr: &seqExpr{
					pos: position{line: 45, col: 14, offset: 1147},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 45, col: 14, offset: 1147},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 45, col: 16, offset: 1149},
							val:        "for",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 45, col: 23, offset: 1156},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 45, col: 25, offset: 1158},
							val:        "domain",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 45, col: 35, offset: 1168},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 45, col: 37, offset: 1170},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 45, col: 42, offset: 1175},
								name: "QUOTED_NAME",
							},
						},
					},
				},
			},
		},
		{
			name: "IN_CONTEXT",
			pos:  position{line: 50, col: 1, offset: 1249},
			expr: &actionExpr{
				pos: position{line: 50, col: 14, offset: 1262},
				run: (*parser).callonIN_CONTEXT1,
				expr: &seqExpr{
					pos: position{line: 50, col: 14, offset: 1262},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 50, col: 14, offset: 1262},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 50, col: 16, offset: 1264},
							val:        "in",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 50, col: 22, offset: 1270},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 50, col: 24, offset: 1272},
							val:        "context",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 50, col: 35, offset: 1283},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 50, col: 37, offset: 1285},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 50, col: 42, offset: 1290},
								name: "QUOTED_NAME",
							},
						},
					},
				},
			},
		},
		{
			name: "WITHIN_AGGREGATE",
			pos:  position{line: 55, col: 1, offset: 1364},
			expr: &actionExpr{
				pos: position{line: 55, col: 20, offset: 1383},
				run: (*parser).callonWITHIN_AGGREGATE1,
				expr: &seqExpr{
					pos: position{line: 55, col: 20, offset: 1383},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 55, col: 20, offset: 1383},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 55, col: 22, offset: 1385},
							val:        "within",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 55, col: 32, offset: 1395},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 55, col: 34, offset: 1397},
							val:        "aggregate",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 55, col: 47, offset: 1410},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 55, col: 49, offset: 1412},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 55, col: 54, offset: 1417},
								name: "QUOTED_NAME",
							},
						},
					},
				},
			},
		},
		{
			name: "CREATE_CLASS",
			pos:  position{line: 60, col: 1, offset: 1497},
			expr: &choiceExpr{
				pos: position{line: 60, col: 16, offset: 1512},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 60, col: 16, offset: 1512},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 60, col: 16, offset: 1512},
								name: "_",
							},
							&ruleRefExpr{
								pos:  position{line: 60, col: 19, offset: 1515},
								name: "CREATE_VALUE",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 60, col: 34, offset: 1530},
						name: "CREATE_COMMAND",
					},
					&ruleRefExpr{
						pos:  position{line: 60, col: 51, offset: 1547},
						name: "CREATE_PROJECTION",
					},
					&ruleRefExpr{
						pos:  position{line: 60, col: 71, offset: 1567},
						name: "CREATE_INVARIANT",
					},
					&ruleRefExpr{
						pos:  position{line: 60, col: 90, offset: 1586},
						name: "CREATE_QUERY",
					},
				},
			},
		},
		{
			name: "CREATE_VALUE",
			pos:  position{line: 62, col: 1, offset: 1600},
			expr: &seqExpr{
				pos: position{line: 62, col: 16, offset: 1615},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 62, col: 16, offset: 1615},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 62, col: 18, offset: 1617},
						val:        "<|",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 23, offset: 1622},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 62, col: 26, offset: 1625},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 62, col: 26, offset: 1625},
								val:        "value",
								ignoreCase: true,
							},
							&litMatcher{
								pos:        position{line: 62, col: 37, offset: 1636},
								val:        "entity",
								ignoreCase: true,
							},
							&litMatcher{
								pos:        position{line: 62, col: 49, offset: 1648},
								val:        "event",
								ignoreCase: true,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 60, offset: 1659},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 62, offset: 1661},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 74, offset: 1673},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 76, offset: 1675},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 86, offset: 1685},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 88, offset: 1687},
						name: "VALUE_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 99, offset: 1698},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 62, col: 101, offset: 1700},
						val:        "|>",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CREATE_COMMAND",
			pos:  position{line: 64, col: 1, offset: 1706},
			expr: &seqExpr{
				pos: position{line: 64, col: 18, offset: 1723},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 64, col: 18, offset: 1723},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 64, col: 20, offset: 1725},
						val:        "<|",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 25, offset: 1730},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 64, col: 27, offset: 1732},
						val:        "command",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 39, offset: 1744},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 41, offset: 1746},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 53, offset: 1758},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 55, offset: 1760},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 65, offset: 1770},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 67, offset: 1772},
						name: "COMMAND_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 80, offset: 1785},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 64, col: 82, offset: 1787},
						val:        "|>",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CREATE_PROJECTION",
			pos:  position{line: 66, col: 1, offset: 1793},
			expr: &seqExpr{
				pos: position{line: 66, col: 21, offset: 1813},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 66, col: 21, offset: 1813},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 66, col: 23, offset: 1815},
						val:        "<|",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 28, offset: 1820},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 66, col: 31, offset: 1823},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 66, col: 31, offset: 1823},
								val:        "aggregate",
								ignoreCase: true,
							},
							&litMatcher{
								pos:        position{line: 66, col: 46, offset: 1838},
								val:        "domain",
								ignoreCase: true,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 57, offset: 1849},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 66, col: 59, offset: 1851},
						val:        "projection",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 74, offset: 1866},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 76, offset: 1868},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 88, offset: 1880},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 90, offset: 1882},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 100, offset: 1892},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 102, offset: 1894},
						name: "PROJECTION_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 118, offset: 1910},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 66, col: 120, offset: 1912},
						val:        "|>",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CREATE_INVARIANT",
			pos:  position{line: 68, col: 1, offset: 1918},
			expr: &seqExpr{
				pos: position{line: 68, col: 20, offset: 1937},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 68, col: 20, offset: 1937},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 68, col: 22, offset: 1939},
						val:        "<|",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 27, offset: 1944},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 68, col: 29, offset: 1946},
						val:        "invariant",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 43, offset: 1960},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 45, offset: 1962},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 57, offset: 1974},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 68, col: 59, offset: 1976},
						val:        "on",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 65, offset: 1982},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 67, offset: 1984},
						name: "CLASS_REF_QUOTES",
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 84, offset: 2001},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 86, offset: 2003},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 96, offset: 2013},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 98, offset: 2015},
						name: "INVARIANT_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 113, offset: 2030},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 68, col: 115, offset: 2032},
						val:        "|>",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CREATE_QUERY",
			pos:  position{line: 70, col: 1, offset: 2038},
			expr: &seqExpr{
				pos: position{line: 70, col: 16, offset: 2053},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 70, col: 16, offset: 2053},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 70, col: 18, offset: 2055},
						val:        "<|",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 70, col: 23, offset: 2060},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 70, col: 25, offset: 2062},
						val:        "query",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 70, col: 35, offset: 2072},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 70, col: 37, offset: 2074},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 70, col: 49, offset: 2086},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 70, col: 51, offset: 2088},
						val:        "on",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 70, col: 57, offset: 2094},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 70, col: 59, offset: 2096},
						name: "CLASS_REF_QUOTES",
					},
					&ruleRefExpr{
						pos:  position{line: 70, col: 76, offset: 2113},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 70, col: 78, offset: 2115},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 70, col: 88, offset: 2125},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 70, col: 90, offset: 2127},
						name: "QUERY_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 70, col: 101, offset: 2138},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 70, col: 103, offset: 2140},
						val:        "|>",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CLASS_COMPONENT_TEST",
			pos:  position{line: 76, col: 1, offset: 2322},
			expr: &seqExpr{
				pos: position{line: 76, col: 24, offset: 2345},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 76, col: 24, offset: 2345},
						expr: &choiceExpr{
							pos: position{line: 76, col: 25, offset: 2346},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 76, col: 25, offset: 2346},
									name: "WHEN",
								},
								&ruleRefExpr{
									pos:  position{line: 76, col: 32, offset: 2353},
									name: "COMMAND_HANDLER",
								},
								&ruleRefExpr{
									pos:  position{line: 76, col: 50, offset: 2371},
									name: "PROPERTIES",
								},
								&ruleRefExpr{
									pos:  position{line: 76, col: 63, offset: 2384},
									name: "CHECK",
								},
								&ruleRefExpr{
									pos:  position{line: 76, col: 71, offset: 2392},
									name: "FUNCTION",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 76, col: 82, offset: 2403},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "QUERY_BODY",
			pos:  position{line: 78, col: 1, offset: 2408},
			expr: &zeroOrMoreExpr{
				pos: position{line: 78, col: 14, offset: 2421},
				expr: &choiceExpr{
					pos: position{line: 78, col: 15, offset: 2422},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 78, col: 15, offset: 2422},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 78, col: 28, offset: 2435},
							name: "QUERY_HANDLER",
						},
					},
				},
			},
		},
		{
			name: "INVARIANT_BODY",
			pos:  position{line: 80, col: 1, offset: 2452},
			expr: &zeroOrMoreExpr{
				pos: position{line: 80, col: 18, offset: 2469},
				expr: &choiceExpr{
					pos: position{line: 80, col: 19, offset: 2470},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 80, col: 19, offset: 2470},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 80, col: 32, offset: 2483},
							name: "CHECK",
						},
					},
				},
			},
		},
		{
			name: "PROJECTION_BODY",
			pos:  position{line: 82, col: 1, offset: 2492},
			expr: &zeroOrMoreExpr{
				pos: position{line: 82, col: 19, offset: 2510},
				expr: &choiceExpr{
					pos: position{line: 82, col: 20, offset: 2511},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 82, col: 20, offset: 2511},
							name: "WHEN",
						},
						&ruleRefExpr{
							pos:  position{line: 82, col: 27, offset: 2518},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 82, col: 40, offset: 2531},
							name: "CHECK",
						},
						&ruleRefExpr{
							pos:  position{line: 82, col: 48, offset: 2539},
							name: "FUNCTION",
						},
					},
				},
			},
		},
		{
			name: "WHEN",
			pos:  position{line: 84, col: 1, offset: 2551},
			expr: &seqExpr{
				pos: position{line: 84, col: 8, offset: 2558},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 84, col: 8, offset: 2558},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 84, col: 10, offset: 2560},
						val:        "when",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 84, col: 18, offset: 2568},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 84, col: 20, offset: 2570},
						val:        "event",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 84, col: 29, offset: 2579},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 84, col: 31, offset: 2581},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 84, col: 43, offset: 2593},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 84, col: 45, offset: 2595},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 84, col: 49, offset: 2599},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 84, col: 51, offset: 2601},
						expr: &ruleRefExpr{
							pos:  position{line: 84, col: 51, offset: 2601},
							name: "STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 84, col: 68, offset: 2618},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 84, col: 70, offset: 2620},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 84, col: 74, offset: 2624},
						name: "_",
					},
				},
			},
		},
		{
			name: "COMMAND_BODY",
			pos:  position{line: 86, col: 1, offset: 2627},
			expr: &zeroOrMoreExpr{
				pos: position{line: 86, col: 16, offset: 2642},
				expr: &choiceExpr{
					pos: position{line: 86, col: 17, offset: 2643},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 86, col: 17, offset: 2643},
							name: "COMMAND_HANDLER",
						},
						&ruleRefExpr{
							pos:  position{line: 86, col: 35, offset: 2661},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 86, col: 48, offset: 2674},
							name: "CHECK",
						},
						&ruleRefExpr{
							pos:  position{line: 86, col: 56, offset: 2682},
							name: "FUNCTION",
						},
					},
				},
			},
		},
		{
			name: "COMMAND_HANDLER",
			pos:  position{line: 88, col: 1, offset: 2694},
			expr: &seqExpr{
				pos: position{line: 88, col: 19, offset: 2712},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 88, col: 19, offset: 2712},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 88, col: 21, offset: 2714},
						val:        "handler",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 88, col: 32, offset: 2725},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 88, col: 34, offset: 2727},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 88, col: 38, offset: 2731},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 88, col: 40, offset: 2733},
						expr: &ruleRefExpr{
							pos:  position{line: 88, col: 40, offset: 2733},
							name: "COMMAND_STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 88, col: 65, offset: 2758},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 88, col: 67, offset: 2760},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 88, col: 71, offset: 2764},
						name: "_",
					},
				},
			},
		},
		{
			name: "QUERY_HANDLER",
			pos:  position{line: 90, col: 1, offset: 2767},
			expr: &seqExpr{
				pos: position{line: 90, col: 17, offset: 2783},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 90, col: 17, offset: 2783},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 90, col: 19, offset: 2785},
						val:        "handler",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 90, col: 30, offset: 2796},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 90, col: 32, offset: 2798},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 90, col: 36, offset: 2802},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 90, col: 38, offset: 2804},
						expr: &ruleRefExpr{
							pos:  position{line: 90, col: 38, offset: 2804},
							name: "STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 90, col: 55, offset: 2821},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 90, col: 57, offset: 2823},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 90, col: 61, offset: 2827},
						name: "_",
					},
				},
			},
		},
		{
			name: "COMMAND_STATEMENT_BLOCK",
			pos:  position{line: 92, col: 1, offset: 2830},
			expr: &seqExpr{
				pos: position{line: 92, col: 27, offset: 2856},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 92, col: 27, offset: 2856},
						name: "_",
					},
					&oneOrMoreExpr{
						pos: position{line: 92, col: 29, offset: 2858},
						expr: &ruleRefExpr{
							pos:  position{line: 92, col: 30, offset: 2859},
							name: "COMMAND_STATEMENT",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 92, col: 50, offset: 2879},
						name: "_",
					},
				},
			},
		},
		{
			name: "COMMAND_STATEMENT",
			pos:  position{line: 94, col: 1, offset: 2882},
			expr: &choiceExpr{
				pos: position{line: 94, col: 21, offset: 2902},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 94, col: 21, offset: 2902},
						name: "STATEMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 94, col: 33, offset: 2914},
						name: "ASSERT",
					},
					&ruleRefExpr{
						pos:  position{line: 94, col: 42, offset: 2923},
						name: "APPLY",
					},
				},
			},
		},
		{
			name: "ASSERT",
			pos:  position{line: 96, col: 1, offset: 2930},
			expr: &seqExpr{
				pos: position{line: 96, col: 10, offset: 2939},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 96, col: 10, offset: 2939},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 96, col: 12, offset: 2941},
						val:        "assert",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 22, offset: 2951},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 96, col: 24, offset: 2953},
						val:        "invariant",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 37, offset: 2966},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 96, col: 39, offset: 2968},
						expr: &litMatcher{
							pos:        position{line: 96, col: 40, offset: 2969},
							val:        "not",
							ignoreCase: true,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 49, offset: 2978},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 51, offset: 2980},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 63, offset: 2992},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 96, col: 65, offset: 2994},
						expr: &ruleRefExpr{
							pos:  position{line: 96, col: 65, offset: 2994},
							name: "ARGUMENTS",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 76, offset: 3005},
						name: "SEMI",
					},
				},
			},
		},
		{
			name: "APPLY",
			pos:  position{line: 98, col: 1, offset: 3011},
			expr: &seqExpr{
				pos: position{line: 98, col: 9, offset: 3019},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 98, col: 9, offset: 3019},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 98, col: 11, offset: 3021},
						val:        "apply",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 98, col: 20, offset: 3030},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 98, col: 22, offset: 3032},
						val:        "event",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 98, col: 31, offset: 3041},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 98, col: 33, offset: 3043},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 98, col: 45, offset: 3055},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 98, col: 47, offset: 3057},
						expr: &ruleRefExpr{
							pos:  position{line: 98, col: 47, offset: 3057},
							name: "ARGUMENTS",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 98, col: 58, offset: 3068},
						name: "SEMI",
					},
				},
			},
		},
		{
			name: "VALUE_BODY",
			pos:  position{line: 100, col: 1, offset: 3074},
			expr: &zeroOrMoreExpr{
				pos: position{line: 100, col: 14, offset: 3087},
				expr: &ruleRefExpr{
					pos:  position{line: 100, col: 15, offset: 3088},
					name: "VALUE_COMPONENTS",
				},
			},
		},
		{
			name: "VALUE_COMPONENTS",
			pos:  position{line: 102, col: 1, offset: 3108},
			expr: &choiceExpr{
				pos: position{line: 102, col: 20, offset: 3127},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 102, col: 20, offset: 3127},
						name: "PROPERTIES",
					},
					&ruleRefExpr{
						pos:  position{line: 102, col: 33, offset: 3140},
						name: "CHECK",
					},
					&ruleRefExpr{
						pos:  position{line: 102, col: 41, offset: 3148},
						name: "FUNCTION",
					},
				},
			},
		},
		{
			name: "PROPERTIES",
			pos:  position{line: 104, col: 1, offset: 3158},
			expr: &seqExpr{
				pos: position{line: 104, col: 14, offset: 3171},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 104, col: 14, offset: 3171},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 104, col: 16, offset: 3173},
						val:        "properties",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 104, col: 30, offset: 3187},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 104, col: 32, offset: 3189},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 104, col: 36, offset: 3193},
						name: "PROPERTY_LIST",
					},
					&litMatcher{
						pos:        position{line: 104, col: 50, offset: 3207},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 104, col: 54, offset: 3211},
						name: "_",
					},
				},
			},
		},
		{
			name: "PROPERTY_LIST",
			pos:  position{line: 106, col: 1, offset: 3214},
			expr: &zeroOrMoreExpr{
				pos: position{line: 106, col: 17, offset: 3230},
				expr: &ruleRefExpr{
					pos:  position{line: 106, col: 18, offset: 3231},
					name: "PROPERTY",
				},
			},
		},
		{
			name: "PROPERTY",
			pos:  position{line: 108, col: 1, offset: 3243},
			expr: &seqExpr{
				pos: position{line: 108, col: 12, offset: 3254},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 108, col: 12, offset: 3254},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 108, col: 14, offset: 3256},
						name: "TYPE",
					},
					&ruleRefExpr{
						pos:  position{line: 108, col: 19, offset: 3261},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 108, col: 21, offset: 3263},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 108, col: 32, offset: 3274},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 108, col: 35, offset: 3277},
						expr: &seqExpr{
							pos: position{line: 108, col: 36, offset: 3278},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 108, col: 36, offset: 3278},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 108, col: 40, offset: 3282},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 108, col: 42, offset: 3284},
									name: "EXPRESSION",
								},
								&ruleRefExpr{
									pos:  position{line: 108, col: 53, offset: 3295},
									name: "_",
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 108, col: 57, offset: 3299},
						val:        ";",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 108, col: 61, offset: 3303},
						name: "_",
					},
				},
			},
		},
		{
			name: "CHECK",
			pos:  position{line: 110, col: 1, offset: 3306},
			expr: &seqExpr{
				pos: position{line: 110, col: 9, offset: 3314},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 110, col: 9, offset: 3314},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 110, col: 11, offset: 3316},
						val:        "check",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 110, col: 20, offset: 3325},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 110, col: 22, offset: 3327},
						val:        "(",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 110, col: 26, offset: 3331},
						expr: &ruleRefExpr{
							pos:  position{line: 110, col: 26, offset: 3331},
							name: "STATEMENT_BLOCK",
						},
					},
					&litMatcher{
						pos:        position{line: 110, col: 43, offset: 3348},
						val:        ")",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 110, col: 47, offset: 3352},
						name: "_",
					},
				},
			},
		},
		{
			name: "FUNCTION",
			pos:  position{line: 112, col: 1, offset: 3355},
			expr: &seqExpr{
				pos: position{line: 112, col: 12, offset: 3366},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 112, col: 12, offset: 3366},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 112, col: 14, offset: 3368},
						val:        "function",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 26, offset: 3380},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 28, offset: 3382},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 39, offset: 3393},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 41, offset: 3395},
						name: "PARAMETERS",
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 53, offset: 3407},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 112, col: 55, offset: 3409},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 59, offset: 3413},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 112, col: 61, offset: 3415},
						expr: &ruleRefExpr{
							pos:  position{line: 112, col: 61, offset: 3415},
							name: "STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 78, offset: 3432},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 112, col: 80, offset: 3434},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 84, offset: 3438},
						name: "_",
					},
				},
			},
		},
		{
			name: "PARAMETERS",
			pos:  position{line: 114, col: 1, offset: 3441},
			expr: &seqExpr{
				pos: position{line: 114, col: 14, offset: 3454},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 114, col: 14, offset: 3454},
						val:        "(",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 114, col: 18, offset: 3458},
						name: "PARAMETER_LIST",
					},
					&litMatcher{
						pos:        position{line: 114, col: 33, offset: 3473},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "PARAMETER_LIST",
			pos:  position{line: 116, col: 1, offset: 3478},
			expr: &seqExpr{
				pos: position{line: 116, col: 18, offset: 3495},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 116, col: 18, offset: 3495},
						name: "_",
					},
					&zeroOrMoreExpr{
						pos: position{line: 116, col: 20, offset: 3497},
						expr: &seqExpr{
							pos: position{line: 116, col: 21, offset: 3498},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 116, col: 21, offset: 3498},
									name: "PARAMETER",
								},
								&litMatcher{
									pos:        position{line: 116, col: 31, offset: 3508},
									val:        ",",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 116, col: 35, offset: 3512},
									name: "_",
								},
							},
						},
					},
					&zeroOrOneExpr{
						pos: position{line: 116, col: 40, offset: 3517},
						expr: &ruleRefExpr{
							pos:  position{line: 116, col: 40, offset: 3517},
							name: "PARAMETER",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 116, col: 51, offset: 3528},
						name: "_",
					},
				},
			},
		},
		{
			name: "PARAMETER",
			pos:  position{line: 118, col: 1, offset: 3531},
			expr: &seqExpr{
				pos: position{line: 118, col: 13, offset: 3543},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 118, col: 13, offset: 3543},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 118, col: 15, offset: 3545},
						name: "CLASS_REF",
					},
					&ruleRefExpr{
						pos:  position{line: 118, col: 25, offset: 3555},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 118, col: 27, offset: 3557},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 118, col: 38, offset: 3568},
						name: "_",
					},
				},
			},
		},
		{
			name: "STATEMENT_BLOCK",
			pos:  position{line: 123, col: 1, offset: 3740},
			expr: &seqExpr{
				pos: position{line: 123, col: 19, offset: 3758},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 123, col: 19, offset: 3758},
						name: "_",
					},
					&oneOrMoreExpr{
						pos: position{line: 123, col: 21, offset: 3760},
						expr: &ruleRefExpr{
							pos:  position{line: 123, col: 22, offset: 3761},
							name: "STATEMENT",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 123, col: 34, offset: 3773},
						name: "_",
					},
				},
			},
		},
		{
			name: "STATEMENT",
			pos:  position{line: 125, col: 1, offset: 3776},
			expr: &choiceExpr{
				pos: position{line: 125, col: 13, offset: 3788},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 125, col: 13, offset: 3788},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 125, col: 13, offset: 3788},
								name: "RETURN",
							},
							&ruleRefExpr{
								pos:  position{line: 125, col: 20, offset: 3795},
								name: "SEMI",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 125, col: 27, offset: 3802},
						name: "IF",
					},
					&ruleRefExpr{
						pos:  position{line: 125, col: 32, offset: 3807},
						name: "FOREACH",
					},
					&seqExpr{
						pos: position{line: 125, col: 42, offset: 3817},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 125, col: 42, offset: 3817},
								name: "EXPRESSION",
							},
							&ruleRefExpr{
								pos:  position{line: 125, col: 53, offset: 3828},
								name: "SEMI",
							},
						},
					},
				},
			},
		},
		{
			name: "IF",
			pos:  position{line: 127, col: 1, offset: 3834},
			expr: &seqExpr{
				pos: position{line: 127, col: 6, offset: 3839},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 127, col: 6, offset: 3839},
						val:        "if",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 127, col: 11, offset: 3844},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 127, col: 13, offset: 3846},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 127, col: 24, offset: 3857},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 127, col: 26, offset: 3859},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 127, col: 30, offset: 3863},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 127, col: 32, offset: 3865},
						expr: &ruleRefExpr{
							pos:  position{line: 127, col: 32, offset: 3865},
							name: "STATEMENT_BLOCK",
						},
					},
					&litMatcher{
						pos:        position{line: 127, col: 49, offset: 3882},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 127, col: 53, offset: 3886},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 127, col: 55, offset: 3888},
						expr: &seqExpr{
							pos: position{line: 127, col: 56, offset: 3889},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 127, col: 56, offset: 3889},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 127, col: 63, offset: 3896},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 127, col: 65, offset: 3898},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 127, col: 69, offset: 3902},
									name: "_",
								},
								&zeroOrOneExpr{
									pos: position{line: 127, col: 71, offset: 3904},
									expr: &ruleRefExpr{
										pos:  position{line: 127, col: 71, offset: 3904},
										name: "STATEMENT_BLOCK",
									},
								},
								&litMatcher{
									pos:        position{line: 127, col: 88, offset: 3921},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 127, col: 92, offset: 3925},
									name: "_",
								},
							},
						},
					},
				},
			},
		},
		{
			name: "FOREACH",
			pos:  position{line: 129, col: 1, offset: 3930},
			expr: &seqExpr{
				pos: position{line: 129, col: 11, offset: 3940},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 129, col: 11, offset: 3940},
						val:        "foreach",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 129, col: 21, offset: 3950},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 129, col: 23, offset: 3952},
						val:        "(",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 129, col: 27, offset: 3956},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 129, col: 29, offset: 3958},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 129, col: 40, offset: 3969},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 129, col: 42, offset: 3971},
						val:        "as",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 129, col: 47, offset: 3976},
						expr: &seqExpr{
							pos: position{line: 129, col: 48, offset: 3977},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 129, col: 48, offset: 3977},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 129, col: 50, offset: 3979},
									name: "IDENTIFIER",
								},
								&ruleRefExpr{
									pos:  position{line: 129, col: 61, offset: 3990},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 129, col: 63, offset: 3992},
									val:        "=>",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 129, col: 70, offset: 3999},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 129, col: 72, offset: 4001},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 129, col: 83, offset: 4012},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 129, col: 85, offset: 4014},
						val:        ")",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 129, col: 89, offset: 4018},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 129, col: 91, offset: 4020},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 129, col: 95, offset: 4024},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 129, col: 97, offset: 4026},
						expr: &ruleRefExpr{
							pos:  position{line: 129, col: 97, offset: 4026},
							name: "STATEMENT_BLOCK",
						},
					},
					&litMatcher{
						pos:        position{line: 129, col: 114, offset: 4043},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "RETURN",
			pos:  position{line: 131, col: 1, offset: 4048},
			expr: &seqExpr{
				pos: position{line: 131, col: 10, offset: 4057},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 131, col: 10, offset: 4057},
						val:        "return",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 131, col: 19, offset: 4066},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 131, col: 21, offset: 4068},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "EXRESSION_TEST",
			pos:  position{line: 137, col: 1, offset: 4251},
			expr: &seqExpr{
				pos: position{line: 137, col: 18, offset: 4268},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 137, col: 18, offset: 4268},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 137, col: 29, offset: 4279},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "EXPRESSION",
			pos:  position{line: 139, col: 1, offset: 4284},
			expr: &choiceExpr{
				pos: position{line: 139, col: 14, offset: 4297},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 139, col: 14, offset: 4297},
						name: "QUERY",
					},
					&ruleRefExpr{
						pos:  position{line: 139, col: 22, offset: 4305},
						name: "ARITHMETIC",
					},
					&ruleRefExpr{
						pos:  position{line: 139, col: 35, offset: 4318},
						name: "COMPARISON",
					},
					&ruleRefExpr{
						pos:  position{line: 139, col: 48, offset: 4331},
						name: "ASSIGNMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 139, col: 60, offset: 4343},
						name: "LOGICAL",
					},
					&ruleRefExpr{
						pos:  position{line: 139, col: 70, offset: 4353},
						name: "ATOMIC",
					},
				},
			},
		},
		{
			name: "ATOMIC",
			pos:  position{line: 141, col: 1, offset: 4361},
			expr: &choiceExpr{
				pos: position{line: 141, col: 10, offset: 4370},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 141, col: 10, offset: 4370},
						name: "PARENTHESIS",
					},
					&ruleRefExpr{
						pos:  position{line: 141, col: 24, offset: 4384},
						name: "NEW",
					},
					&ruleRefExpr{
						pos:  position{line: 141, col: 30, offset: 4390},
						name: "METHODCALL",
					},
					&ruleRefExpr{
						pos:  position{line: 141, col: 43, offset: 4403},
						name: "OBJECTACCESS",
					},
					&ruleRefExpr{
						pos:  position{line: 141, col: 58, offset: 4418},
						name: "ARRAY",
					},
					&ruleRefExpr{
						pos:  position{line: 141, col: 66, offset: 4426},
						name: "LITERAL",
					},
					&ruleRefExpr{
						pos:  position{line: 141, col: 76, offset: 4436},
						name: "UNARY",
					},
				},
			},
		},
		{
			name: "LITERAL",
			pos:  position{line: 143, col: 1, offset: 4443},
			expr: &choiceExpr{
				pos: position{line: 143, col: 11, offset: 4453},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 143, col: 11, offset: 4453},
						name: "STRING",
					},
					&ruleRefExpr{
						pos:  position{line: 143, col: 20, offset: 4462},
						name: "FLOAT",
					},
					&ruleRefExpr{
						pos:  position{line: 143, col: 28, offset: 4470},
						name: "BOOLEAN",
					},
					&ruleRefExpr{
						pos:  position{line: 143, col: 38, offset: 4480},
						name: "NULL",
					},
					&ruleRefExpr{
						pos:  position{line: 143, col: 45, offset: 4487},
						name: "INT",
					},
				},
			},
		},
		{
			name: "NEW",
			pos:  position{line: 145, col: 1, offset: 4492},
			expr: &seqExpr{
				pos: position{line: 145, col: 7, offset: 4498},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 145, col: 7, offset: 4498},
						name: "CLASS_REF_QUOTES",
					},
					&ruleRefExpr{
						pos:  position{line: 145, col: 24, offset: 4515},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 145, col: 26, offset: 4517},
						expr: &ruleRefExpr{
							pos:  position{line: 145, col: 26, offset: 4517},
							name: "ARGUMENTS",
						},
					},
				},
			},
		},
		{
			name: "BOOLEAN",
			pos:  position{line: 147, col: 1, offset: 4529},
			expr: &choiceExpr{
				pos: position{line: 147, col: 12, offset: 4540},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 147, col: 12, offset: 4540},
						val:        "true",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 147, col: 19, offset: 4547},
						val:        "false",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "NULL",
			pos:  position{line: 149, col: 1, offset: 4556},
			expr: &litMatcher{
				pos:        position{line: 149, col: 8, offset: 4563},
				val:        "null",
				ignoreCase: false,
			},
		},
		{
			name: "ARRAY",
			pos:  position{line: 151, col: 1, offset: 4571},
			expr: &seqExpr{
				pos: position{line: 151, col: 9, offset: 4579},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 151, col: 9, offset: 4579},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 151, col: 11, offset: 4581},
						val:        "[",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 151, col: 15, offset: 4585},
						expr: &ruleRefExpr{
							pos:  position{line: 151, col: 15, offset: 4585},
							name: "ARGUMENTLIST",
						},
					},
					&litMatcher{
						pos:        position{line: 151, col: 29, offset: 4599},
						val:        "]",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 151, col: 33, offset: 4603},
						name: "_",
					},
				},
			},
		},
		{
			name: "STRING",
			pos:  position{line: 153, col: 1, offset: 4606},
			expr: &seqExpr{
				pos: position{line: 153, col: 10, offset: 4615},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 153, col: 10, offset: 4615},
						val:        "\"",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 153, col: 15, offset: 4620},
						expr: &charClassMatcher{
							pos:        position{line: 153, col: 15, offset: 4620},
							val:        "[a-zA-Z0-9]",
							ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&litMatcher{
						pos:        position{line: 153, col: 28, offset: 4633},
						val:        "\"",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "INT",
			pos:  position{line: 155, col: 1, offset: 4639},
			expr: &oneOrMoreExpr{
				pos: position{line: 155, col: 7, offset: 4645},
				expr: &charClassMatcher{
					pos:        position{line: 155, col: 7, offset: 4645},
					val:        "[0-9]",
					ranges:     []rune{'0', '9'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "FLOAT",
			pos:  position{line: 157, col: 1, offset: 4653},
			expr: &seqExpr{
				pos: position{line: 157, col: 9, offset: 4661},
				exprs: []interface{}{
					&oneOrMoreExpr{
						pos: position{line: 157, col: 9, offset: 4661},
						expr: &charClassMatcher{
							pos:        position{line: 157, col: 9, offset: 4661},
							val:        "[0-9]",
							ranges:     []rune{'0', '9'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&charClassMatcher{
						pos:        position{line: 157, col: 16, offset: 4668},
						val:        "[.]",
						chars:      []rune{'.'},
						ignoreCase: false,
						inverted:   false,
					},
					&oneOrMoreExpr{
						pos: position{line: 157, col: 20, offset: 4672},
						expr: &charClassMatcher{
							pos:        position{line: 157, col: 20, offset: 4672},
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
			pos:  position{line: 159, col: 1, offset: 4680},
			expr: &seqExpr{
				pos: position{line: 159, col: 15, offset: 4694},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 159, col: 15, offset: 4694},
						val:        "(",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 159, col: 19, offset: 4698},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 159, col: 21, offset: 4700},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 159, col: 32, offset: 4711},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 159, col: 34, offset: 4713},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "UNARY",
			pos:  position{line: 161, col: 1, offset: 4718},
			expr: &choiceExpr{
				pos: position{line: 161, col: 9, offset: 4726},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 161, col: 9, offset: 4726},
						name: "INCREMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 161, col: 21, offset: 4738},
						name: "DECREMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 161, col: 33, offset: 4750},
						name: "NEGATE",
					},
					&ruleRefExpr{
						pos:  position{line: 161, col: 42, offset: 4759},
						name: "NOT",
					},
					&ruleRefExpr{
						pos:  position{line: 161, col: 48, offset: 4765},
						name: "POSITIVE",
					},
				},
			},
		},
		{
			name: "INCREMENT",
			pos:  position{line: 163, col: 1, offset: 4775},
			expr: &seqExpr{
				pos: position{line: 163, col: 13, offset: 4787},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 163, col: 13, offset: 4787},
						name: "OBJECTACCESS",
					},
					&litMatcher{
						pos:        position{line: 163, col: 26, offset: 4800},
						val:        "++",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DECREMENT",
			pos:  position{line: 165, col: 1, offset: 4806},
			expr: &seqExpr{
				pos: position{line: 165, col: 13, offset: 4818},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 165, col: 13, offset: 4818},
						name: "OBJECTACCESS",
					},
					&litMatcher{
						pos:        position{line: 165, col: 26, offset: 4831},
						val:        "--",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "NEGATE",
			pos:  position{line: 167, col: 1, offset: 4837},
			expr: &seqExpr{
				pos: position{line: 167, col: 10, offset: 4846},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 167, col: 10, offset: 4846},
						val:        "-",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 167, col: 14, offset: 4850},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "NOT",
			pos:  position{line: 169, col: 1, offset: 4864},
			expr: &seqExpr{
				pos: position{line: 169, col: 7, offset: 4870},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 169, col: 7, offset: 4870},
						val:        "!",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 169, col: 11, offset: 4874},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "POSITIVE",
			pos:  position{line: 171, col: 1, offset: 4888},
			expr: &seqExpr{
				pos: position{line: 171, col: 12, offset: 4899},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 171, col: 12, offset: 4899},
						val:        "+",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 171, col: 16, offset: 4903},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "ARITHMETIC",
			pos:  position{line: 173, col: 1, offset: 4917},
			expr: &seqExpr{
				pos: position{line: 173, col: 14, offset: 4930},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 173, col: 14, offset: 4930},
						name: "ATOMIC",
					},
					&ruleRefExpr{
						pos:  position{line: 173, col: 21, offset: 4937},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 173, col: 23, offset: 4939},
						name: "OPERATOR",
					},
					&ruleRefExpr{
						pos:  position{line: 173, col: 32, offset: 4948},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 173, col: 34, offset: 4950},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "OPERATOR",
			pos:  position{line: 175, col: 1, offset: 4962},
			expr: &choiceExpr{
				pos: position{line: 175, col: 12, offset: 4973},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 175, col: 12, offset: 4973},
						val:        "+",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 175, col: 18, offset: 4979},
						val:        "-",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 175, col: 24, offset: 4985},
						val:        "/",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 175, col: 30, offset: 4991},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 175, col: 36, offset: 4997},
						val:        "%",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "ASSIGNMENT",
			pos:  position{line: 177, col: 1, offset: 5002},
			expr: &seqExpr{
				pos: position{line: 177, col: 14, offset: 5015},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 177, col: 14, offset: 5015},
						name: "OBJECTACCESS",
					},
					&ruleRefExpr{
						pos:  position{line: 177, col: 27, offset: 5028},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 177, col: 29, offset: 5030},
						val:        "=",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 177, col: 33, offset: 5034},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 177, col: 35, offset: 5036},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "LOGICAL",
			pos:  position{line: 179, col: 1, offset: 5048},
			expr: &seqExpr{
				pos: position{line: 179, col: 11, offset: 5058},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 179, col: 11, offset: 5058},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 179, col: 22, offset: 5069},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 179, col: 25, offset: 5072},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 179, col: 25, offset: 5072},
								val:        "and",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 179, col: 33, offset: 5080},
								val:        "or",
								ignoreCase: false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 179, col: 39, offset: 5086},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 179, col: 41, offset: 5088},
						name: "ATOMIC",
					},
				},
			},
		},
		{
			name: "COMPARISON",
			pos:  position{line: 181, col: 1, offset: 5096},
			expr: &seqExpr{
				pos: position{line: 181, col: 14, offset: 5109},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 181, col: 14, offset: 5109},
						name: "ATOMIC",
					},
					&ruleRefExpr{
						pos:  position{line: 181, col: 21, offset: 5116},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 181, col: 24, offset: 5119},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 181, col: 24, offset: 5119},
								val:        "===",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 181, col: 32, offset: 5127},
								val:        "!==",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 181, col: 40, offset: 5135},
								val:        "==",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 181, col: 47, offset: 5142},
								val:        "!=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 181, col: 54, offset: 5149},
								val:        "<=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 181, col: 61, offset: 5156},
								val:        ">=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 181, col: 68, offset: 5163},
								val:        "<",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 181, col: 74, offset: 5169},
								val:        ">",
								ignoreCase: false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 181, col: 79, offset: 5174},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 181, col: 81, offset: 5176},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "QUERY",
			pos:  position{line: 183, col: 1, offset: 5188},
			expr: &seqExpr{
				pos: position{line: 183, col: 9, offset: 5196},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 183, col: 9, offset: 5196},
						val:        "run",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 183, col: 16, offset: 5203},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 183, col: 18, offset: 5205},
						val:        "query",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 183, col: 27, offset: 5214},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 183, col: 29, offset: 5216},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 183, col: 41, offset: 5228},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 183, col: 43, offset: 5230},
						expr: &ruleRefExpr{
							pos:  position{line: 183, col: 43, offset: 5230},
							name: "ARGUMENTS",
						},
					},
				},
			},
		},
		{
			name: "OBJECTACCESS",
			pos:  position{line: 185, col: 1, offset: 5242},
			expr: &seqExpr{
				pos: position{line: 185, col: 16, offset: 5257},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 185, col: 16, offset: 5257},
						expr: &seqExpr{
							pos: position{line: 185, col: 17, offset: 5258},
							exprs: []interface{}{
								&choiceExpr{
									pos: position{line: 185, col: 18, offset: 5259},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 185, col: 18, offset: 5259},
											name: "METHODCALL",
										},
										&ruleRefExpr{
											pos:  position{line: 185, col: 31, offset: 5272},
											name: "IDENTIFIER",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 185, col: 43, offset: 5284},
									val:        "->",
									ignoreCase: false,
								},
							},
						},
					},
					&choiceExpr{
						pos: position{line: 185, col: 51, offset: 5292},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 185, col: 51, offset: 5292},
								name: "METHODCALL",
							},
							&ruleRefExpr{
								pos:  position{line: 185, col: 64, offset: 5305},
								name: "IDENTIFIER",
							},
						},
					},
				},
			},
		},
		{
			name: "METHODCALL",
			pos:  position{line: 187, col: 1, offset: 5318},
			expr: &seqExpr{
				pos: position{line: 187, col: 14, offset: 5331},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 187, col: 14, offset: 5331},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 187, col: 25, offset: 5342},
						name: "ARGUMENTS",
					},
				},
			},
		},
		{
			name: "ARGUMENTS",
			pos:  position{line: 189, col: 1, offset: 5353},
			expr: &seqExpr{
				pos: position{line: 189, col: 13, offset: 5365},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 189, col: 13, offset: 5365},
						val:        "(",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 189, col: 17, offset: 5369},
						expr: &ruleRefExpr{
							pos:  position{line: 189, col: 17, offset: 5369},
							name: "ARGUMENTLIST",
						},
					},
					&litMatcher{
						pos:        position{line: 189, col: 31, offset: 5383},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "ARGUMENTLIST",
			pos:  position{line: 191, col: 1, offset: 5388},
			expr: &seqExpr{
				pos: position{line: 191, col: 17, offset: 5404},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 191, col: 17, offset: 5404},
						name: "_",
					},
					&zeroOrMoreExpr{
						pos: position{line: 191, col: 19, offset: 5406},
						expr: &seqExpr{
							pos: position{line: 191, col: 20, offset: 5407},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 191, col: 20, offset: 5407},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 22, offset: 5409},
									name: "EXPRESSION",
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 33, offset: 5420},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 191, col: 35, offset: 5422},
									val:        ",",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 191, col: 39, offset: 5426},
									name: "_",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 191, col: 43, offset: 5430},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 191, col: 54, offset: 5441},
						name: "_",
					},
				},
			},
		},
		{
			name: "CLASS_REF_QUOTES",
			pos:  position{line: 198, col: 1, offset: 5609},
			expr: &seqExpr{
				pos: position{line: 198, col: 20, offset: 5628},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 198, col: 20, offset: 5628},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 198, col: 22, offset: 5630},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 198, col: 26, offset: 5634},
						name: "CLASS_REF",
					},
					&litMatcher{
						pos:        position{line: 198, col: 36, offset: 5644},
						val:        "'",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CLASS_REF",
			pos:  position{line: 200, col: 1, offset: 5649},
			expr: &seqExpr{
				pos: position{line: 200, col: 13, offset: 5661},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 200, col: 13, offset: 5661},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 200, col: 15, offset: 5663},
						name: "CLASS_TYPE",
					},
					&litMatcher{
						pos:        position{line: 200, col: 26, offset: 5674},
						val:        "\\",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 200, col: 31, offset: 5679},
						name: "CLASS_NAME",
					},
				},
			},
		},
		{
			name: "CLASS_TYPE",
			pos:  position{line: 202, col: 1, offset: 5691},
			expr: &seqExpr{
				pos: position{line: 202, col: 14, offset: 5704},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 202, col: 14, offset: 5704},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 202, col: 17, offset: 5707},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 202, col: 17, offset: 5707},
								val:        "value",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 202, col: 27, offset: 5717},
								val:        "entity",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 202, col: 38, offset: 5728},
								val:        "command",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 202, col: 50, offset: 5740},
								val:        "event",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 202, col: 60, offset: 5750},
								val:        "projection",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 202, col: 75, offset: 5765},
								val:        "invariant",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 202, col: 89, offset: 5779},
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
			pos:  position{line: 204, col: 1, offset: 5789},
			expr: &seqExpr{
				pos: position{line: 204, col: 14, offset: 5802},
				exprs: []interface{}{
					&oneOrMoreExpr{
						pos: position{line: 204, col: 14, offset: 5802},
						expr: &charClassMatcher{
							pos:        position{line: 204, col: 14, offset: 5802},
							val:        "[a-zA-Z]",
							ranges:     []rune{'a', 'z', 'A', 'Z'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 204, col: 24, offset: 5812},
						expr: &charClassMatcher{
							pos:        position{line: 204, col: 24, offset: 5812},
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
			pos:  position{line: 206, col: 1, offset: 5828},
			expr: &seqExpr{
				pos: position{line: 206, col: 15, offset: 5842},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 206, col: 15, offset: 5842},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 206, col: 19, offset: 5846},
						name: "CLASS_NAME",
					},
					&litMatcher{
						pos:        position{line: 206, col: 30, offset: 5857},
						val:        "'",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "TYPE",
			pos:  position{line: 208, col: 1, offset: 5862},
			expr: &choiceExpr{
				pos: position{line: 208, col: 8, offset: 5869},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 208, col: 8, offset: 5869},
						name: "CLASS_REF",
					},
					&ruleRefExpr{
						pos:  position{line: 208, col: 20, offset: 5881},
						name: "VALUE_TYPE",
					},
				},
			},
		},
		{
			name: "VALUE_TYPE",
			pos:  position{line: 210, col: 1, offset: 5893},
			expr: &seqExpr{
				pos: position{line: 210, col: 14, offset: 5906},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 210, col: 14, offset: 5906},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 210, col: 17, offset: 5909},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 210, col: 17, offset: 5909},
								val:        "string",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 210, col: 28, offset: 5920},
								val:        "boolean",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 210, col: 40, offset: 5932},
								val:        "float",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 210, col: 50, offset: 5942},
								val:        "map",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 210, col: 58, offset: 5950},
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
			pos:  position{line: 212, col: 1, offset: 5960},
			expr: &seqExpr{
				pos: position{line: 212, col: 21, offset: 5980},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 212, col: 21, offset: 5980},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 212, col: 23, offset: 5982},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 212, col: 27, offset: 5986},
						name: "CLASS_NAME",
					},
					&litMatcher{
						pos:        position{line: 212, col: 38, offset: 5997},
						val:        "'",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "IDENTIFIER",
			pos:  position{line: 214, col: 1, offset: 6002},
			expr: &seqExpr{
				pos: position{line: 214, col: 14, offset: 6015},
				exprs: []interface{}{
					&oneOrMoreExpr{
						pos: position{line: 214, col: 14, offset: 6015},
						expr: &charClassMatcher{
							pos:        position{line: 214, col: 14, offset: 6015},
							val:        "[a-zA-Z]",
							ranges:     []rune{'a', 'z', 'A', 'Z'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 214, col: 24, offset: 6025},
						expr: &charClassMatcher{
							pos:        position{line: 214, col: 24, offset: 6025},
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
			pos:  position{line: 216, col: 1, offset: 6040},
			expr: &zeroOrMoreExpr{
				pos: position{line: 216, col: 5, offset: 6044},
				expr: &choiceExpr{
					pos: position{line: 216, col: 7, offset: 6046},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 216, col: 7, offset: 6046},
							name: "WHITESPACE",
						},
						&ruleRefExpr{
							pos:  position{line: 216, col: 20, offset: 6059},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "SEMI",
			pos:  position{line: 218, col: 1, offset: 6067},
			expr: &seqExpr{
				pos: position{line: 218, col: 8, offset: 6074},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 218, col: 8, offset: 6074},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 218, col: 10, offset: 6076},
						val:        ";",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 218, col: 14, offset: 6080},
						name: "_",
					},
				},
			},
		},
		{
			name: "WHITESPACE",
			pos:  position{line: 220, col: 1, offset: 6083},
			expr: &charClassMatcher{
				pos:        position{line: 220, col: 14, offset: 6096},
				val:        "[ \\t\\r]",
				chars:      []rune{' ', '\t', '\r'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EOL",
			pos:  position{line: 222, col: 1, offset: 6105},
			expr: &litMatcher{
				pos:        position{line: 222, col: 7, offset: 6111},
				val:        "\n",
				ignoreCase: false,
			},
		},
		{
			name: "EOF",
			pos:  position{line: 224, col: 1, offset: 6117},
			expr: &notExpr{
				pos: position{line: 224, col: 7, offset: 6123},
				expr: &anyMatcher{
					line: 224, col: 8, offset: 6124,
				},
			},
		},
	},
}

func (c *current) onCREATE_NAMESPACE_OBJECT1(typ, name interface{}) (interface{}, error) {
	//emit(parser.Create, "create");
	//emit(parser.NamespaceObject, typ);
	//emit(parser.QuotedName, name);
	return nil, nil
}

func (p *parser) callonCREATE_NAMESPACE_OBJECT1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCREATE_NAMESPACE_OBJECT1(stack["typ"], stack["name"])
}

func (c *current) onNAMESPACE_OBJECT_TYPE1(typ interface{}) (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonNAMESPACE_OBJECT_TYPE1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onNAMESPACE_OBJECT_TYPE1(stack["typ"])
}

func (c *current) onUSING_DATABASE1(name interface{}) (interface{}, error) {
	//emit(parser.UsingDatabase, name);
	return nil, nil
}

func (p *parser) callonUSING_DATABASE1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUSING_DATABASE1(stack["name"])
}

func (c *current) onFOR_DOMAIN1(name interface{}) (interface{}, error) {
	//emit(parser.ForDomain, name);
	return nil, nil
}

func (p *parser) callonFOR_DOMAIN1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFOR_DOMAIN1(stack["name"])
}

func (c *current) onIN_CONTEXT1(name interface{}) (interface{}, error) {
	//emit(parser.InContext, name);
	return nil, nil
}

func (p *parser) callonIN_CONTEXT1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIN_CONTEXT1(stack["name"])
}

func (c *current) onWITHIN_AGGREGATE1(name interface{}) (interface{}, error) {
	//emit(parser.WithinAggregate, name);
	return nil, nil
}

func (p *parser) callonWITHIN_AGGREGATE1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onWITHIN_AGGREGATE1(stack["name"])
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
