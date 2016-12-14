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

func emit(typ TokenType, val interface{}) (interface{}, error) {
	GetInstanceTokenList().Append(NewToken(typ, val.(string)))
	return nil, nil
}

var g = &grammar{
	rules: []*rule{
		{
			name: "FILE",
			pos:  position{line: 17, col: 1, offset: 363},
			expr: &seqExpr{
				pos: position{line: 17, col: 8, offset: 370},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 17, col: 8, offset: 370},
						name: "STATEMENTS",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 19, offset: 381},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 21, offset: 383},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "STATEMENTS",
			pos:  position{line: 19, col: 1, offset: 388},
			expr: &labeledExpr{
				pos:   position{line: 19, col: 14, offset: 401},
				label: "statements",
				expr: &zeroOrMoreExpr{
					pos: position{line: 19, col: 25, offset: 412},
					expr: &choiceExpr{
						pos: position{line: 19, col: 26, offset: 413},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 19, col: 26, offset: 413},
								name: "BLOCK_STATEMENT",
							},
							&ruleRefExpr{
								pos:  position{line: 19, col: 43, offset: 430},
								name: "CREATE_OBJECT",
							},
							&ruleRefExpr{
								pos:  position{line: 19, col: 59, offset: 446},
								name: "CREATE_CLASS",
							},
						},
					},
				},
			},
		},
		{
			name: "CREATE_OBJECT",
			pos:  position{line: 21, col: 1, offset: 462},
			expr: &actionExpr{
				pos: position{line: 21, col: 17, offset: 478},
				run: (*parser).callonCREATE_OBJECT1,
				expr: &seqExpr{
					pos: position{line: 21, col: 17, offset: 478},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 21, col: 17, offset: 478},
							name: "_",
						},
						&zeroOrOneExpr{
							pos: position{line: 21, col: 19, offset: 480},
							expr: &ruleRefExpr{
								pos:  position{line: 21, col: 19, offset: 480},
								name: "CREATE_NAMESPACE_OBJECT",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 21, col: 44, offset: 505},
							name: "_",
						},
						&zeroOrOneExpr{
							pos: position{line: 21, col: 46, offset: 507},
							expr: &ruleRefExpr{
								pos:  position{line: 21, col: 46, offset: 507},
								name: "NAMESPACE",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 21, col: 57, offset: 518},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 21, col: 59, offset: 520},
							val:        ";",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "BLOCK_STATEMENT",
			pos:  position{line: 25, col: 1, offset: 563},
			expr: &seqExpr{
				pos: position{line: 25, col: 19, offset: 581},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 25, col: 19, offset: 581},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 21, offset: 583},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 31, offset: 593},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 25, col: 33, offset: 595},
						val:        ":",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 37, offset: 599},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 25, col: 39, offset: 601},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 43, offset: 605},
						name: "STATEMENTS",
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 54, offset: 616},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 25, col: 56, offset: 618},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CREATE_NAMESPACE_OBJECT",
			pos:  position{line: 27, col: 1, offset: 623},
			expr: &actionExpr{
				pos: position{line: 27, col: 27, offset: 649},
				run: (*parser).callonCREATE_NAMESPACE_OBJECT1,
				expr: &seqExpr{
					pos: position{line: 27, col: 27, offset: 649},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 27, col: 27, offset: 649},
							val:        "create",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 27, col: 37, offset: 659},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 27, col: 39, offset: 661},
							label: "typ",
							expr: &ruleRefExpr{
								pos:  position{line: 27, col: 43, offset: 665},
								name: "NAMESPACE_OBJECT_TYPE",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 27, col: 65, offset: 687},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 27, col: 67, offset: 689},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 27, col: 72, offset: 694},
								name: "QUOTED_NAME",
							},
						},
					},
				},
			},
		},
		{
			name: "NAMESPACE_OBJECT_TYPE",
			pos:  position{line: 33, col: 1, offset: 806},
			expr: &actionExpr{
				pos: position{line: 33, col: 25, offset: 830},
				run: (*parser).callonNAMESPACE_OBJECT_TYPE1,
				expr: &labeledExpr{
					pos:   position{line: 33, col: 25, offset: 830},
					label: "typ",
					expr: &choiceExpr{
						pos: position{line: 33, col: 30, offset: 835},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 33, col: 30, offset: 835},
								val:        "database",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 33, col: 43, offset: 848},
								val:        "domain",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 33, col: 54, offset: 859},
								val:        "context",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 33, col: 66, offset: 871},
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
			pos:  position{line: 37, col: 1, offset: 921},
			expr: &zeroOrMoreExpr{
				pos: position{line: 37, col: 13, offset: 933},
				expr: &choiceExpr{
					pos: position{line: 37, col: 14, offset: 934},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 37, col: 14, offset: 934},
							name: "USING_DATABASE",
						},
						&ruleRefExpr{
							pos:  position{line: 37, col: 31, offset: 951},
							name: "FOR_DOMAIN",
						},
						&ruleRefExpr{
							pos:  position{line: 37, col: 44, offset: 964},
							name: "IN_CONTEXT",
						},
						&ruleRefExpr{
							pos:  position{line: 37, col: 57, offset: 977},
							name: "WITHIN_AGGREGATE",
						},
					},
				},
			},
		},
		{
			name: "USING_DATABASE",
			pos:  position{line: 39, col: 1, offset: 997},
			expr: &actionExpr{
				pos: position{line: 39, col: 18, offset: 1014},
				run: (*parser).callonUSING_DATABASE1,
				expr: &seqExpr{
					pos: position{line: 39, col: 18, offset: 1014},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 39, col: 18, offset: 1014},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 39, col: 20, offset: 1016},
							val:        "using",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 39, col: 29, offset: 1025},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 39, col: 31, offset: 1027},
							val:        "database",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 39, col: 43, offset: 1039},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 39, col: 45, offset: 1041},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 39, col: 50, offset: 1046},
								name: "QUOTED_NAME",
							},
						},
					},
				},
			},
		},
		{
			name: "FOR_DOMAIN",
			pos:  position{line: 43, col: 1, offset: 1101},
			expr: &actionExpr{
				pos: position{line: 43, col: 14, offset: 1114},
				run: (*parser).callonFOR_DOMAIN1,
				expr: &seqExpr{
					pos: position{line: 43, col: 14, offset: 1114},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 43, col: 14, offset: 1114},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 43, col: 16, offset: 1116},
							val:        "for",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 43, col: 23, offset: 1123},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 43, col: 25, offset: 1125},
							val:        "domain",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 43, col: 35, offset: 1135},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 43, col: 37, offset: 1137},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 43, col: 42, offset: 1142},
								name: "QUOTED_NAME",
							},
						},
					},
				},
			},
		},
		{
			name: "IN_CONTEXT",
			pos:  position{line: 47, col: 1, offset: 1193},
			expr: &actionExpr{
				pos: position{line: 47, col: 14, offset: 1206},
				run: (*parser).callonIN_CONTEXT1,
				expr: &seqExpr{
					pos: position{line: 47, col: 14, offset: 1206},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 47, col: 14, offset: 1206},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 47, col: 16, offset: 1208},
							val:        "in",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 47, col: 22, offset: 1214},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 47, col: 24, offset: 1216},
							val:        "context",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 47, col: 35, offset: 1227},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 47, col: 37, offset: 1229},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 47, col: 42, offset: 1234},
								name: "QUOTED_NAME",
							},
						},
					},
				},
			},
		},
		{
			name: "WITHIN_AGGREGATE",
			pos:  position{line: 51, col: 1, offset: 1285},
			expr: &actionExpr{
				pos: position{line: 51, col: 20, offset: 1304},
				run: (*parser).callonWITHIN_AGGREGATE1,
				expr: &seqExpr{
					pos: position{line: 51, col: 20, offset: 1304},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 51, col: 20, offset: 1304},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 51, col: 22, offset: 1306},
							val:        "within",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 51, col: 32, offset: 1316},
							name: "_",
						},
						&litMatcher{
							pos:        position{line: 51, col: 34, offset: 1318},
							val:        "aggregate",
							ignoreCase: true,
						},
						&ruleRefExpr{
							pos:  position{line: 51, col: 47, offset: 1331},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 51, col: 49, offset: 1333},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 51, col: 54, offset: 1338},
								name: "QUOTED_NAME",
							},
						},
					},
				},
			},
		},
		{
			name: "CREATE_CLASS",
			pos:  position{line: 55, col: 1, offset: 1395},
			expr: &seqExpr{
				pos: position{line: 55, col: 16, offset: 1410},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 55, col: 16, offset: 1410},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 55, col: 18, offset: 1412},
						name: "CLASS_OPEN",
					},
					&choiceExpr{
						pos: position{line: 55, col: 30, offset: 1424},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 55, col: 30, offset: 1424},
								name: "CREATE_VALUE",
							},
							&ruleRefExpr{
								pos:  position{line: 55, col: 45, offset: 1439},
								name: "CREATE_COMMAND",
							},
							&ruleRefExpr{
								pos:  position{line: 55, col: 62, offset: 1456},
								name: "CREATE_PROJECTION",
							},
							&ruleRefExpr{
								pos:  position{line: 55, col: 82, offset: 1476},
								name: "CREATE_INVARIANT",
							},
							&ruleRefExpr{
								pos:  position{line: 55, col: 101, offset: 1495},
								name: "CREATE_QUERY",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 55, col: 115, offset: 1509},
						name: "CLASS_CLOSE",
					},
				},
			},
		},
		{
			name: "CREATE_VALUE",
			pos:  position{line: 57, col: 1, offset: 1522},
			expr: &actionExpr{
				pos: position{line: 57, col: 16, offset: 1537},
				run: (*parser).callonCREATE_VALUE1,
				expr: &seqExpr{
					pos: position{line: 57, col: 16, offset: 1537},
					exprs: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 57, col: 16, offset: 1537},
							name: "_",
						},
						&choiceExpr{
							pos: position{line: 57, col: 19, offset: 1540},
							alternatives: []interface{}{
								&litMatcher{
									pos:        position{line: 57, col: 19, offset: 1540},
									val:        "value",
									ignoreCase: true,
								},
								&litMatcher{
									pos:        position{line: 57, col: 30, offset: 1551},
									val:        "entity",
									ignoreCase: true,
								},
							},
						},
						&ruleRefExpr{
							pos:  position{line: 57, col: 42, offset: 1563},
							name: "_",
						},
						&labeledExpr{
							pos:   position{line: 57, col: 44, offset: 1565},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 57, col: 49, offset: 1570},
								name: "QUOTED_NAME",
							},
						},
						&ruleRefExpr{
							pos:  position{line: 57, col: 61, offset: 1582},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 57, col: 63, offset: 1584},
							name: "NAMESPACE",
						},
						&ruleRefExpr{
							pos:  position{line: 57, col: 73, offset: 1594},
							name: "_",
						},
						&ruleRefExpr{
							pos:  position{line: 57, col: 75, offset: 1596},
							name: "VALUE_BODY",
						},
						&ruleRefExpr{
							pos:  position{line: 57, col: 86, offset: 1607},
							name: "_",
						},
					},
				},
			},
		},
		{
			name: "CREATE_COMMAND",
			pos:  position{line: 62, col: 1, offset: 1675},
			expr: &seqExpr{
				pos: position{line: 62, col: 18, offset: 1692},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 62, col: 18, offset: 1692},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 62, col: 20, offset: 1694},
						val:        "command",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 32, offset: 1706},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 34, offset: 1708},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 46, offset: 1720},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 48, offset: 1722},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 58, offset: 1732},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 60, offset: 1734},
						name: "COMMAND_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 73, offset: 1747},
						name: "_",
					},
				},
			},
		},
		{
			name: "CREATE_PROJECTION",
			pos:  position{line: 64, col: 1, offset: 1750},
			expr: &seqExpr{
				pos: position{line: 64, col: 21, offset: 1770},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 64, col: 21, offset: 1770},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 64, col: 24, offset: 1773},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 64, col: 24, offset: 1773},
								val:        "aggregate",
								ignoreCase: true,
							},
							&litMatcher{
								pos:        position{line: 64, col: 39, offset: 1788},
								val:        "domain",
								ignoreCase: true,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 50, offset: 1799},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 64, col: 52, offset: 1801},
						val:        "projection",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 67, offset: 1816},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 69, offset: 1818},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 81, offset: 1830},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 83, offset: 1832},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 93, offset: 1842},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 95, offset: 1844},
						name: "PROJECTION_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 111, offset: 1860},
						name: "_",
					},
				},
			},
		},
		{
			name: "CREATE_INVARIANT",
			pos:  position{line: 66, col: 1, offset: 1863},
			expr: &seqExpr{
				pos: position{line: 66, col: 20, offset: 1882},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 66, col: 20, offset: 1882},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 66, col: 23, offset: 1885},
						val:        "invariant",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 37, offset: 1899},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 39, offset: 1901},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 51, offset: 1913},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 66, col: 53, offset: 1915},
						val:        "on",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 59, offset: 1921},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 61, offset: 1923},
						name: "CLASS_REF_QUOTES",
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 78, offset: 1940},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 80, offset: 1942},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 90, offset: 1952},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 92, offset: 1954},
						name: "INVARIANT_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 107, offset: 1969},
						name: "_",
					},
				},
			},
		},
		{
			name: "CREATE_QUERY",
			pos:  position{line: 68, col: 1, offset: 1972},
			expr: &seqExpr{
				pos: position{line: 68, col: 16, offset: 1987},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 68, col: 16, offset: 1987},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 68, col: 19, offset: 1990},
						val:        "query",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 29, offset: 2000},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 31, offset: 2002},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 43, offset: 2014},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 68, col: 45, offset: 2016},
						val:        "on",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 51, offset: 2022},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 53, offset: 2024},
						name: "CLASS_REF_QUOTES",
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 70, offset: 2041},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 72, offset: 2043},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 82, offset: 2053},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 84, offset: 2055},
						name: "QUERY_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 68, col: 95, offset: 2066},
						name: "_",
					},
				},
			},
		},
		{
			name: "CLASS_OPEN",
			pos:  position{line: 70, col: 1, offset: 2069},
			expr: &actionExpr{
				pos: position{line: 70, col: 14, offset: 2082},
				run: (*parser).callonCLASS_OPEN1,
				expr: &litMatcher{
					pos:        position{line: 70, col: 14, offset: 2082},
					val:        "<|",
					ignoreCase: false,
				},
			},
		},
		{
			name: "CLASS_CLOSE",
			pos:  position{line: 74, col: 1, offset: 2136},
			expr: &actionExpr{
				pos: position{line: 74, col: 15, offset: 2150},
				run: (*parser).callonCLASS_CLOSE1,
				expr: &litMatcher{
					pos:        position{line: 74, col: 15, offset: 2150},
					val:        "|>",
					ignoreCase: false,
				},
			},
		},
		{
			name: "CLASS_COMPONENT_TEST",
			pos:  position{line: 83, col: 1, offset: 2380},
			expr: &seqExpr{
				pos: position{line: 83, col: 24, offset: 2403},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 83, col: 24, offset: 2403},
						expr: &choiceExpr{
							pos: position{line: 83, col: 25, offset: 2404},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 83, col: 25, offset: 2404},
									name: "WHEN",
								},
								&ruleRefExpr{
									pos:  position{line: 83, col: 32, offset: 2411},
									name: "COMMAND_HANDLER",
								},
								&ruleRefExpr{
									pos:  position{line: 83, col: 50, offset: 2429},
									name: "PROPERTIES",
								},
								&ruleRefExpr{
									pos:  position{line: 83, col: 63, offset: 2442},
									name: "CHECK",
								},
								&ruleRefExpr{
									pos:  position{line: 83, col: 71, offset: 2450},
									name: "FUNCTION",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 83, col: 82, offset: 2461},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "QUERY_BODY",
			pos:  position{line: 85, col: 1, offset: 2466},
			expr: &zeroOrMoreExpr{
				pos: position{line: 85, col: 14, offset: 2479},
				expr: &choiceExpr{
					pos: position{line: 85, col: 15, offset: 2480},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 85, col: 15, offset: 2480},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 85, col: 28, offset: 2493},
							name: "QUERY_HANDLER",
						},
					},
				},
			},
		},
		{
			name: "INVARIANT_BODY",
			pos:  position{line: 87, col: 1, offset: 2510},
			expr: &zeroOrMoreExpr{
				pos: position{line: 87, col: 18, offset: 2527},
				expr: &choiceExpr{
					pos: position{line: 87, col: 19, offset: 2528},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 87, col: 19, offset: 2528},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 87, col: 32, offset: 2541},
							name: "CHECK",
						},
					},
				},
			},
		},
		{
			name: "PROJECTION_BODY",
			pos:  position{line: 89, col: 1, offset: 2550},
			expr: &zeroOrMoreExpr{
				pos: position{line: 89, col: 19, offset: 2568},
				expr: &choiceExpr{
					pos: position{line: 89, col: 20, offset: 2569},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 89, col: 20, offset: 2569},
							name: "WHEN",
						},
						&ruleRefExpr{
							pos:  position{line: 89, col: 27, offset: 2576},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 89, col: 40, offset: 2589},
							name: "CHECK",
						},
						&ruleRefExpr{
							pos:  position{line: 89, col: 48, offset: 2597},
							name: "FUNCTION",
						},
					},
				},
			},
		},
		{
			name: "WHEN",
			pos:  position{line: 91, col: 1, offset: 2609},
			expr: &seqExpr{
				pos: position{line: 91, col: 8, offset: 2616},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 91, col: 8, offset: 2616},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 91, col: 10, offset: 2618},
						val:        "when",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 91, col: 18, offset: 2626},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 91, col: 20, offset: 2628},
						val:        "event",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 91, col: 29, offset: 2637},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 91, col: 31, offset: 2639},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 91, col: 43, offset: 2651},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 91, col: 45, offset: 2653},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 91, col: 49, offset: 2657},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 91, col: 51, offset: 2659},
						expr: &ruleRefExpr{
							pos:  position{line: 91, col: 51, offset: 2659},
							name: "STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 91, col: 68, offset: 2676},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 91, col: 70, offset: 2678},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 91, col: 74, offset: 2682},
						name: "_",
					},
				},
			},
		},
		{
			name: "COMMAND_BODY",
			pos:  position{line: 93, col: 1, offset: 2685},
			expr: &zeroOrMoreExpr{
				pos: position{line: 93, col: 16, offset: 2700},
				expr: &choiceExpr{
					pos: position{line: 93, col: 17, offset: 2701},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 93, col: 17, offset: 2701},
							name: "COMMAND_HANDLER",
						},
						&ruleRefExpr{
							pos:  position{line: 93, col: 35, offset: 2719},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 93, col: 48, offset: 2732},
							name: "CHECK",
						},
						&ruleRefExpr{
							pos:  position{line: 93, col: 56, offset: 2740},
							name: "FUNCTION",
						},
					},
				},
			},
		},
		{
			name: "COMMAND_HANDLER",
			pos:  position{line: 95, col: 1, offset: 2752},
			expr: &seqExpr{
				pos: position{line: 95, col: 19, offset: 2770},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 95, col: 19, offset: 2770},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 95, col: 21, offset: 2772},
						val:        "handler",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 95, col: 32, offset: 2783},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 95, col: 34, offset: 2785},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 95, col: 38, offset: 2789},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 95, col: 40, offset: 2791},
						expr: &ruleRefExpr{
							pos:  position{line: 95, col: 40, offset: 2791},
							name: "COMMAND_STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 95, col: 65, offset: 2816},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 95, col: 67, offset: 2818},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 95, col: 71, offset: 2822},
						name: "_",
					},
				},
			},
		},
		{
			name: "QUERY_HANDLER",
			pos:  position{line: 97, col: 1, offset: 2825},
			expr: &seqExpr{
				pos: position{line: 97, col: 17, offset: 2841},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 97, col: 17, offset: 2841},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 97, col: 19, offset: 2843},
						val:        "handler",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 97, col: 30, offset: 2854},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 97, col: 32, offset: 2856},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 97, col: 36, offset: 2860},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 97, col: 38, offset: 2862},
						expr: &ruleRefExpr{
							pos:  position{line: 97, col: 38, offset: 2862},
							name: "STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 97, col: 55, offset: 2879},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 97, col: 57, offset: 2881},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 97, col: 61, offset: 2885},
						name: "_",
					},
				},
			},
		},
		{
			name: "COMMAND_STATEMENT_BLOCK",
			pos:  position{line: 99, col: 1, offset: 2888},
			expr: &seqExpr{
				pos: position{line: 99, col: 27, offset: 2914},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 99, col: 27, offset: 2914},
						name: "_",
					},
					&oneOrMoreExpr{
						pos: position{line: 99, col: 29, offset: 2916},
						expr: &ruleRefExpr{
							pos:  position{line: 99, col: 30, offset: 2917},
							name: "COMMAND_STATEMENT",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 99, col: 50, offset: 2937},
						name: "_",
					},
				},
			},
		},
		{
			name: "COMMAND_STATEMENT",
			pos:  position{line: 101, col: 1, offset: 2940},
			expr: &choiceExpr{
				pos: position{line: 101, col: 21, offset: 2960},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 101, col: 21, offset: 2960},
						name: "STATEMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 101, col: 33, offset: 2972},
						name: "ASSERT",
					},
					&ruleRefExpr{
						pos:  position{line: 101, col: 42, offset: 2981},
						name: "APPLY",
					},
				},
			},
		},
		{
			name: "ASSERT",
			pos:  position{line: 103, col: 1, offset: 2988},
			expr: &seqExpr{
				pos: position{line: 103, col: 10, offset: 2997},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 103, col: 10, offset: 2997},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 103, col: 12, offset: 2999},
						val:        "assert",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 103, col: 22, offset: 3009},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 103, col: 24, offset: 3011},
						val:        "invariant",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 103, col: 37, offset: 3024},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 103, col: 39, offset: 3026},
						expr: &litMatcher{
							pos:        position{line: 103, col: 40, offset: 3027},
							val:        "not",
							ignoreCase: true,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 103, col: 49, offset: 3036},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 103, col: 51, offset: 3038},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 103, col: 63, offset: 3050},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 103, col: 65, offset: 3052},
						expr: &ruleRefExpr{
							pos:  position{line: 103, col: 65, offset: 3052},
							name: "ARGUMENTS",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 103, col: 76, offset: 3063},
						name: "SEMI",
					},
				},
			},
		},
		{
			name: "APPLY",
			pos:  position{line: 105, col: 1, offset: 3069},
			expr: &seqExpr{
				pos: position{line: 105, col: 9, offset: 3077},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 105, col: 9, offset: 3077},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 105, col: 11, offset: 3079},
						val:        "apply",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 105, col: 20, offset: 3088},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 105, col: 22, offset: 3090},
						val:        "event",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 105, col: 31, offset: 3099},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 105, col: 33, offset: 3101},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 105, col: 45, offset: 3113},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 105, col: 47, offset: 3115},
						expr: &ruleRefExpr{
							pos:  position{line: 105, col: 47, offset: 3115},
							name: "ARGUMENTS",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 105, col: 58, offset: 3126},
						name: "SEMI",
					},
				},
			},
		},
		{
			name: "VALUE_BODY",
			pos:  position{line: 107, col: 1, offset: 3132},
			expr: &zeroOrMoreExpr{
				pos: position{line: 107, col: 14, offset: 3145},
				expr: &ruleRefExpr{
					pos:  position{line: 107, col: 15, offset: 3146},
					name: "VALUE_COMPONENTS",
				},
			},
		},
		{
			name: "VALUE_COMPONENTS",
			pos:  position{line: 109, col: 1, offset: 3166},
			expr: &choiceExpr{
				pos: position{line: 109, col: 20, offset: 3185},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 109, col: 20, offset: 3185},
						name: "PROPERTIES",
					},
					&ruleRefExpr{
						pos:  position{line: 109, col: 33, offset: 3198},
						name: "CHECK",
					},
					&ruleRefExpr{
						pos:  position{line: 109, col: 41, offset: 3206},
						name: "FUNCTION",
					},
				},
			},
		},
		{
			name: "PROPERTIES",
			pos:  position{line: 111, col: 1, offset: 3216},
			expr: &seqExpr{
				pos: position{line: 111, col: 14, offset: 3229},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 111, col: 14, offset: 3229},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 111, col: 16, offset: 3231},
						val:        "properties",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 111, col: 30, offset: 3245},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 111, col: 32, offset: 3247},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 111, col: 36, offset: 3251},
						name: "PROPERTY_LIST",
					},
					&litMatcher{
						pos:        position{line: 111, col: 50, offset: 3265},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 111, col: 54, offset: 3269},
						name: "_",
					},
				},
			},
		},
		{
			name: "PROPERTY_LIST",
			pos:  position{line: 113, col: 1, offset: 3272},
			expr: &zeroOrMoreExpr{
				pos: position{line: 113, col: 17, offset: 3288},
				expr: &ruleRefExpr{
					pos:  position{line: 113, col: 18, offset: 3289},
					name: "PROPERTY",
				},
			},
		},
		{
			name: "PROPERTY",
			pos:  position{line: 115, col: 1, offset: 3301},
			expr: &seqExpr{
				pos: position{line: 115, col: 12, offset: 3312},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 115, col: 12, offset: 3312},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 115, col: 14, offset: 3314},
						name: "TYPE",
					},
					&ruleRefExpr{
						pos:  position{line: 115, col: 19, offset: 3319},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 115, col: 21, offset: 3321},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 115, col: 32, offset: 3332},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 115, col: 35, offset: 3335},
						expr: &seqExpr{
							pos: position{line: 115, col: 36, offset: 3336},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 115, col: 36, offset: 3336},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 115, col: 40, offset: 3340},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 115, col: 42, offset: 3342},
									name: "EXPRESSION",
								},
								&ruleRefExpr{
									pos:  position{line: 115, col: 53, offset: 3353},
									name: "_",
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 115, col: 57, offset: 3357},
						val:        ";",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 115, col: 61, offset: 3361},
						name: "_",
					},
				},
			},
		},
		{
			name: "CHECK",
			pos:  position{line: 117, col: 1, offset: 3364},
			expr: &seqExpr{
				pos: position{line: 117, col: 9, offset: 3372},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 117, col: 9, offset: 3372},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 117, col: 11, offset: 3374},
						val:        "check",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 117, col: 20, offset: 3383},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 117, col: 22, offset: 3385},
						val:        "(",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 117, col: 26, offset: 3389},
						expr: &ruleRefExpr{
							pos:  position{line: 117, col: 26, offset: 3389},
							name: "STATEMENT_BLOCK",
						},
					},
					&litMatcher{
						pos:        position{line: 117, col: 43, offset: 3406},
						val:        ")",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 117, col: 47, offset: 3410},
						name: "_",
					},
				},
			},
		},
		{
			name: "FUNCTION",
			pos:  position{line: 119, col: 1, offset: 3413},
			expr: &seqExpr{
				pos: position{line: 119, col: 12, offset: 3424},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 119, col: 12, offset: 3424},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 119, col: 14, offset: 3426},
						val:        "function",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 119, col: 26, offset: 3438},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 119, col: 28, offset: 3440},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 119, col: 39, offset: 3451},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 119, col: 41, offset: 3453},
						name: "PARAMETERS",
					},
					&ruleRefExpr{
						pos:  position{line: 119, col: 53, offset: 3465},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 119, col: 55, offset: 3467},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 119, col: 59, offset: 3471},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 119, col: 61, offset: 3473},
						expr: &ruleRefExpr{
							pos:  position{line: 119, col: 61, offset: 3473},
							name: "STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 119, col: 78, offset: 3490},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 119, col: 80, offset: 3492},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 119, col: 84, offset: 3496},
						name: "_",
					},
				},
			},
		},
		{
			name: "PARAMETERS",
			pos:  position{line: 121, col: 1, offset: 3499},
			expr: &seqExpr{
				pos: position{line: 121, col: 14, offset: 3512},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 121, col: 14, offset: 3512},
						val:        "(",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 121, col: 18, offset: 3516},
						name: "PARAMETER_LIST",
					},
					&litMatcher{
						pos:        position{line: 121, col: 33, offset: 3531},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "PARAMETER_LIST",
			pos:  position{line: 123, col: 1, offset: 3536},
			expr: &seqExpr{
				pos: position{line: 123, col: 18, offset: 3553},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 123, col: 18, offset: 3553},
						name: "_",
					},
					&zeroOrMoreExpr{
						pos: position{line: 123, col: 20, offset: 3555},
						expr: &seqExpr{
							pos: position{line: 123, col: 21, offset: 3556},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 123, col: 21, offset: 3556},
									name: "PARAMETER",
								},
								&litMatcher{
									pos:        position{line: 123, col: 31, offset: 3566},
									val:        ",",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 123, col: 35, offset: 3570},
									name: "_",
								},
							},
						},
					},
					&zeroOrOneExpr{
						pos: position{line: 123, col: 40, offset: 3575},
						expr: &ruleRefExpr{
							pos:  position{line: 123, col: 40, offset: 3575},
							name: "PARAMETER",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 123, col: 51, offset: 3586},
						name: "_",
					},
				},
			},
		},
		{
			name: "PARAMETER",
			pos:  position{line: 125, col: 1, offset: 3589},
			expr: &seqExpr{
				pos: position{line: 125, col: 13, offset: 3601},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 125, col: 13, offset: 3601},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 125, col: 15, offset: 3603},
						name: "CLASS_REF",
					},
					&ruleRefExpr{
						pos:  position{line: 125, col: 25, offset: 3613},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 125, col: 27, offset: 3615},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 125, col: 38, offset: 3626},
						name: "_",
					},
				},
			},
		},
		{
			name: "STATEMENT_BLOCK",
			pos:  position{line: 130, col: 1, offset: 3798},
			expr: &seqExpr{
				pos: position{line: 130, col: 19, offset: 3816},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 130, col: 19, offset: 3816},
						name: "_",
					},
					&oneOrMoreExpr{
						pos: position{line: 130, col: 21, offset: 3818},
						expr: &ruleRefExpr{
							pos:  position{line: 130, col: 22, offset: 3819},
							name: "STATEMENT",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 34, offset: 3831},
						name: "_",
					},
				},
			},
		},
		{
			name: "STATEMENT",
			pos:  position{line: 132, col: 1, offset: 3834},
			expr: &choiceExpr{
				pos: position{line: 132, col: 13, offset: 3846},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 132, col: 13, offset: 3846},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 132, col: 13, offset: 3846},
								name: "RETURN",
							},
							&ruleRefExpr{
								pos:  position{line: 132, col: 20, offset: 3853},
								name: "SEMI",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 132, col: 27, offset: 3860},
						name: "IF",
					},
					&ruleRefExpr{
						pos:  position{line: 132, col: 32, offset: 3865},
						name: "FOREACH",
					},
					&seqExpr{
						pos: position{line: 132, col: 42, offset: 3875},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 132, col: 42, offset: 3875},
								name: "EXPRESSION",
							},
							&ruleRefExpr{
								pos:  position{line: 132, col: 53, offset: 3886},
								name: "SEMI",
							},
						},
					},
				},
			},
		},
		{
			name: "IF",
			pos:  position{line: 134, col: 1, offset: 3892},
			expr: &seqExpr{
				pos: position{line: 134, col: 6, offset: 3897},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 134, col: 6, offset: 3897},
						val:        "if",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 134, col: 11, offset: 3902},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 134, col: 13, offset: 3904},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 134, col: 24, offset: 3915},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 134, col: 26, offset: 3917},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 134, col: 30, offset: 3921},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 134, col: 32, offset: 3923},
						expr: &ruleRefExpr{
							pos:  position{line: 134, col: 32, offset: 3923},
							name: "STATEMENT_BLOCK",
						},
					},
					&litMatcher{
						pos:        position{line: 134, col: 49, offset: 3940},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 134, col: 53, offset: 3944},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 134, col: 55, offset: 3946},
						expr: &seqExpr{
							pos: position{line: 134, col: 56, offset: 3947},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 134, col: 56, offset: 3947},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 134, col: 63, offset: 3954},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 134, col: 65, offset: 3956},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 134, col: 69, offset: 3960},
									name: "_",
								},
								&zeroOrOneExpr{
									pos: position{line: 134, col: 71, offset: 3962},
									expr: &ruleRefExpr{
										pos:  position{line: 134, col: 71, offset: 3962},
										name: "STATEMENT_BLOCK",
									},
								},
								&litMatcher{
									pos:        position{line: 134, col: 88, offset: 3979},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 134, col: 92, offset: 3983},
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
			pos:  position{line: 136, col: 1, offset: 3988},
			expr: &seqExpr{
				pos: position{line: 136, col: 11, offset: 3998},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 136, col: 11, offset: 3998},
						val:        "foreach",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 136, col: 21, offset: 4008},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 136, col: 23, offset: 4010},
						val:        "(",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 136, col: 27, offset: 4014},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 136, col: 29, offset: 4016},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 136, col: 40, offset: 4027},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 136, col: 42, offset: 4029},
						val:        "as",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 136, col: 47, offset: 4034},
						expr: &seqExpr{
							pos: position{line: 136, col: 48, offset: 4035},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 136, col: 48, offset: 4035},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 50, offset: 4037},
									name: "IDENTIFIER",
								},
								&ruleRefExpr{
									pos:  position{line: 136, col: 61, offset: 4048},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 136, col: 63, offset: 4050},
									val:        "=>",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 136, col: 70, offset: 4057},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 136, col: 72, offset: 4059},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 136, col: 83, offset: 4070},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 136, col: 85, offset: 4072},
						val:        ")",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 136, col: 89, offset: 4076},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 136, col: 91, offset: 4078},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 136, col: 95, offset: 4082},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 136, col: 97, offset: 4084},
						expr: &ruleRefExpr{
							pos:  position{line: 136, col: 97, offset: 4084},
							name: "STATEMENT_BLOCK",
						},
					},
					&litMatcher{
						pos:        position{line: 136, col: 114, offset: 4101},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "RETURN",
			pos:  position{line: 138, col: 1, offset: 4106},
			expr: &seqExpr{
				pos: position{line: 138, col: 10, offset: 4115},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 138, col: 10, offset: 4115},
						val:        "return",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 138, col: 19, offset: 4124},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 138, col: 21, offset: 4126},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "EXRESSION_TEST",
			pos:  position{line: 144, col: 1, offset: 4309},
			expr: &seqExpr{
				pos: position{line: 144, col: 18, offset: 4326},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 144, col: 18, offset: 4326},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 144, col: 29, offset: 4337},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "EXPRESSION",
			pos:  position{line: 146, col: 1, offset: 4342},
			expr: &choiceExpr{
				pos: position{line: 146, col: 14, offset: 4355},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 146, col: 14, offset: 4355},
						name: "QUERY",
					},
					&ruleRefExpr{
						pos:  position{line: 146, col: 22, offset: 4363},
						name: "ARITHMETIC",
					},
					&ruleRefExpr{
						pos:  position{line: 146, col: 35, offset: 4376},
						name: "COMPARISON",
					},
					&ruleRefExpr{
						pos:  position{line: 146, col: 48, offset: 4389},
						name: "ASSIGNMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 146, col: 60, offset: 4401},
						name: "LOGICAL",
					},
					&ruleRefExpr{
						pos:  position{line: 146, col: 70, offset: 4411},
						name: "ATOMIC",
					},
				},
			},
		},
		{
			name: "ATOMIC",
			pos:  position{line: 148, col: 1, offset: 4419},
			expr: &choiceExpr{
				pos: position{line: 148, col: 10, offset: 4428},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 148, col: 10, offset: 4428},
						name: "PARENTHESIS",
					},
					&ruleRefExpr{
						pos:  position{line: 148, col: 24, offset: 4442},
						name: "NEW",
					},
					&ruleRefExpr{
						pos:  position{line: 148, col: 30, offset: 4448},
						name: "METHODCALL",
					},
					&ruleRefExpr{
						pos:  position{line: 148, col: 43, offset: 4461},
						name: "OBJECTACCESS",
					},
					&ruleRefExpr{
						pos:  position{line: 148, col: 58, offset: 4476},
						name: "ARRAY",
					},
					&ruleRefExpr{
						pos:  position{line: 148, col: 66, offset: 4484},
						name: "LITERAL",
					},
					&ruleRefExpr{
						pos:  position{line: 148, col: 76, offset: 4494},
						name: "UNARY",
					},
				},
			},
		},
		{
			name: "LITERAL",
			pos:  position{line: 150, col: 1, offset: 4501},
			expr: &choiceExpr{
				pos: position{line: 150, col: 11, offset: 4511},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 150, col: 11, offset: 4511},
						name: "STRING",
					},
					&ruleRefExpr{
						pos:  position{line: 150, col: 20, offset: 4520},
						name: "FLOAT",
					},
					&ruleRefExpr{
						pos:  position{line: 150, col: 28, offset: 4528},
						name: "BOOLEAN",
					},
					&ruleRefExpr{
						pos:  position{line: 150, col: 38, offset: 4538},
						name: "NULL",
					},
					&ruleRefExpr{
						pos:  position{line: 150, col: 45, offset: 4545},
						name: "INT",
					},
				},
			},
		},
		{
			name: "NEW",
			pos:  position{line: 152, col: 1, offset: 4550},
			expr: &seqExpr{
				pos: position{line: 152, col: 7, offset: 4556},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 152, col: 7, offset: 4556},
						name: "CLASS_REF_QUOTES",
					},
					&ruleRefExpr{
						pos:  position{line: 152, col: 24, offset: 4573},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 152, col: 26, offset: 4575},
						expr: &ruleRefExpr{
							pos:  position{line: 152, col: 26, offset: 4575},
							name: "ARGUMENTS",
						},
					},
				},
			},
		},
		{
			name: "BOOLEAN",
			pos:  position{line: 154, col: 1, offset: 4587},
			expr: &choiceExpr{
				pos: position{line: 154, col: 12, offset: 4598},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 154, col: 12, offset: 4598},
						val:        "true",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 154, col: 19, offset: 4605},
						val:        "false",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "NULL",
			pos:  position{line: 156, col: 1, offset: 4614},
			expr: &litMatcher{
				pos:        position{line: 156, col: 8, offset: 4621},
				val:        "null",
				ignoreCase: false,
			},
		},
		{
			name: "ARRAY",
			pos:  position{line: 158, col: 1, offset: 4629},
			expr: &seqExpr{
				pos: position{line: 158, col: 9, offset: 4637},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 158, col: 9, offset: 4637},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 158, col: 11, offset: 4639},
						val:        "[",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 158, col: 15, offset: 4643},
						expr: &ruleRefExpr{
							pos:  position{line: 158, col: 15, offset: 4643},
							name: "ARGUMENTLIST",
						},
					},
					&litMatcher{
						pos:        position{line: 158, col: 29, offset: 4657},
						val:        "]",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 158, col: 33, offset: 4661},
						name: "_",
					},
				},
			},
		},
		{
			name: "STRING",
			pos:  position{line: 160, col: 1, offset: 4664},
			expr: &seqExpr{
				pos: position{line: 160, col: 10, offset: 4673},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 160, col: 10, offset: 4673},
						val:        "\"",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 160, col: 15, offset: 4678},
						expr: &charClassMatcher{
							pos:        position{line: 160, col: 15, offset: 4678},
							val:        "[a-zA-Z0-9]",
							ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&litMatcher{
						pos:        position{line: 160, col: 28, offset: 4691},
						val:        "\"",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "INT",
			pos:  position{line: 162, col: 1, offset: 4697},
			expr: &oneOrMoreExpr{
				pos: position{line: 162, col: 7, offset: 4703},
				expr: &charClassMatcher{
					pos:        position{line: 162, col: 7, offset: 4703},
					val:        "[0-9]",
					ranges:     []rune{'0', '9'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "FLOAT",
			pos:  position{line: 164, col: 1, offset: 4711},
			expr: &seqExpr{
				pos: position{line: 164, col: 9, offset: 4719},
				exprs: []interface{}{
					&oneOrMoreExpr{
						pos: position{line: 164, col: 9, offset: 4719},
						expr: &charClassMatcher{
							pos:        position{line: 164, col: 9, offset: 4719},
							val:        "[0-9]",
							ranges:     []rune{'0', '9'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&charClassMatcher{
						pos:        position{line: 164, col: 16, offset: 4726},
						val:        "[.]",
						chars:      []rune{'.'},
						ignoreCase: false,
						inverted:   false,
					},
					&oneOrMoreExpr{
						pos: position{line: 164, col: 20, offset: 4730},
						expr: &charClassMatcher{
							pos:        position{line: 164, col: 20, offset: 4730},
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
			pos:  position{line: 166, col: 1, offset: 4738},
			expr: &seqExpr{
				pos: position{line: 166, col: 15, offset: 4752},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 166, col: 15, offset: 4752},
						val:        "(",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 166, col: 19, offset: 4756},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 166, col: 21, offset: 4758},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 166, col: 32, offset: 4769},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 166, col: 34, offset: 4771},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "UNARY",
			pos:  position{line: 168, col: 1, offset: 4776},
			expr: &choiceExpr{
				pos: position{line: 168, col: 9, offset: 4784},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 168, col: 9, offset: 4784},
						name: "INCREMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 168, col: 21, offset: 4796},
						name: "DECREMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 168, col: 33, offset: 4808},
						name: "NEGATE",
					},
					&ruleRefExpr{
						pos:  position{line: 168, col: 42, offset: 4817},
						name: "NOT",
					},
					&ruleRefExpr{
						pos:  position{line: 168, col: 48, offset: 4823},
						name: "POSITIVE",
					},
				},
			},
		},
		{
			name: "INCREMENT",
			pos:  position{line: 170, col: 1, offset: 4833},
			expr: &seqExpr{
				pos: position{line: 170, col: 13, offset: 4845},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 170, col: 13, offset: 4845},
						name: "OBJECTACCESS",
					},
					&litMatcher{
						pos:        position{line: 170, col: 26, offset: 4858},
						val:        "++",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DECREMENT",
			pos:  position{line: 172, col: 1, offset: 4864},
			expr: &seqExpr{
				pos: position{line: 172, col: 13, offset: 4876},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 172, col: 13, offset: 4876},
						name: "OBJECTACCESS",
					},
					&litMatcher{
						pos:        position{line: 172, col: 26, offset: 4889},
						val:        "--",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "NEGATE",
			pos:  position{line: 174, col: 1, offset: 4895},
			expr: &seqExpr{
				pos: position{line: 174, col: 10, offset: 4904},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 174, col: 10, offset: 4904},
						val:        "-",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 174, col: 14, offset: 4908},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "NOT",
			pos:  position{line: 176, col: 1, offset: 4922},
			expr: &seqExpr{
				pos: position{line: 176, col: 7, offset: 4928},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 176, col: 7, offset: 4928},
						val:        "!",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 176, col: 11, offset: 4932},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "POSITIVE",
			pos:  position{line: 178, col: 1, offset: 4946},
			expr: &seqExpr{
				pos: position{line: 178, col: 12, offset: 4957},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 178, col: 12, offset: 4957},
						val:        "+",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 178, col: 16, offset: 4961},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "ARITHMETIC",
			pos:  position{line: 180, col: 1, offset: 4975},
			expr: &seqExpr{
				pos: position{line: 180, col: 14, offset: 4988},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 180, col: 14, offset: 4988},
						name: "ATOMIC",
					},
					&ruleRefExpr{
						pos:  position{line: 180, col: 21, offset: 4995},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 180, col: 23, offset: 4997},
						name: "OPERATOR",
					},
					&ruleRefExpr{
						pos:  position{line: 180, col: 32, offset: 5006},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 180, col: 34, offset: 5008},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "OPERATOR",
			pos:  position{line: 182, col: 1, offset: 5020},
			expr: &choiceExpr{
				pos: position{line: 182, col: 12, offset: 5031},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 182, col: 12, offset: 5031},
						val:        "+",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 182, col: 18, offset: 5037},
						val:        "-",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 182, col: 24, offset: 5043},
						val:        "/",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 182, col: 30, offset: 5049},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 182, col: 36, offset: 5055},
						val:        "%",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "ASSIGNMENT",
			pos:  position{line: 184, col: 1, offset: 5060},
			expr: &seqExpr{
				pos: position{line: 184, col: 14, offset: 5073},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 184, col: 14, offset: 5073},
						name: "OBJECTACCESS",
					},
					&ruleRefExpr{
						pos:  position{line: 184, col: 27, offset: 5086},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 184, col: 29, offset: 5088},
						val:        "=",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 184, col: 33, offset: 5092},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 184, col: 35, offset: 5094},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "LOGICAL",
			pos:  position{line: 186, col: 1, offset: 5106},
			expr: &seqExpr{
				pos: position{line: 186, col: 11, offset: 5116},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 186, col: 11, offset: 5116},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 186, col: 22, offset: 5127},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 186, col: 25, offset: 5130},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 186, col: 25, offset: 5130},
								val:        "and",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 186, col: 33, offset: 5138},
								val:        "or",
								ignoreCase: false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 186, col: 39, offset: 5144},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 186, col: 41, offset: 5146},
						name: "ATOMIC",
					},
				},
			},
		},
		{
			name: "COMPARISON",
			pos:  position{line: 188, col: 1, offset: 5154},
			expr: &seqExpr{
				pos: position{line: 188, col: 14, offset: 5167},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 188, col: 14, offset: 5167},
						name: "ATOMIC",
					},
					&ruleRefExpr{
						pos:  position{line: 188, col: 21, offset: 5174},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 188, col: 24, offset: 5177},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 188, col: 24, offset: 5177},
								val:        "===",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 188, col: 32, offset: 5185},
								val:        "!==",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 188, col: 40, offset: 5193},
								val:        "==",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 188, col: 47, offset: 5200},
								val:        "!=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 188, col: 54, offset: 5207},
								val:        "<=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 188, col: 61, offset: 5214},
								val:        ">=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 188, col: 68, offset: 5221},
								val:        "<",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 188, col: 74, offset: 5227},
								val:        ">",
								ignoreCase: false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 188, col: 79, offset: 5232},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 188, col: 81, offset: 5234},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "QUERY",
			pos:  position{line: 190, col: 1, offset: 5246},
			expr: &seqExpr{
				pos: position{line: 190, col: 9, offset: 5254},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 190, col: 9, offset: 5254},
						val:        "run",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 190, col: 16, offset: 5261},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 190, col: 18, offset: 5263},
						val:        "query",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 190, col: 27, offset: 5272},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 190, col: 29, offset: 5274},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 190, col: 41, offset: 5286},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 190, col: 43, offset: 5288},
						expr: &ruleRefExpr{
							pos:  position{line: 190, col: 43, offset: 5288},
							name: "ARGUMENTS",
						},
					},
				},
			},
		},
		{
			name: "OBJECTACCESS",
			pos:  position{line: 192, col: 1, offset: 5300},
			expr: &seqExpr{
				pos: position{line: 192, col: 16, offset: 5315},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 192, col: 16, offset: 5315},
						expr: &seqExpr{
							pos: position{line: 192, col: 17, offset: 5316},
							exprs: []interface{}{
								&choiceExpr{
									pos: position{line: 192, col: 18, offset: 5317},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 192, col: 18, offset: 5317},
											name: "METHODCALL",
										},
										&ruleRefExpr{
											pos:  position{line: 192, col: 31, offset: 5330},
											name: "IDENTIFIER",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 192, col: 43, offset: 5342},
									val:        "->",
									ignoreCase: false,
								},
							},
						},
					},
					&choiceExpr{
						pos: position{line: 192, col: 51, offset: 5350},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 192, col: 51, offset: 5350},
								name: "METHODCALL",
							},
							&ruleRefExpr{
								pos:  position{line: 192, col: 64, offset: 5363},
								name: "IDENTIFIER",
							},
						},
					},
				},
			},
		},
		{
			name: "METHODCALL",
			pos:  position{line: 194, col: 1, offset: 5376},
			expr: &seqExpr{
				pos: position{line: 194, col: 14, offset: 5389},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 194, col: 14, offset: 5389},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 194, col: 25, offset: 5400},
						name: "ARGUMENTS",
					},
				},
			},
		},
		{
			name: "ARGUMENTS",
			pos:  position{line: 196, col: 1, offset: 5411},
			expr: &seqExpr{
				pos: position{line: 196, col: 13, offset: 5423},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 196, col: 13, offset: 5423},
						val:        "(",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 196, col: 17, offset: 5427},
						expr: &ruleRefExpr{
							pos:  position{line: 196, col: 17, offset: 5427},
							name: "ARGUMENTLIST",
						},
					},
					&litMatcher{
						pos:        position{line: 196, col: 31, offset: 5441},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "ARGUMENTLIST",
			pos:  position{line: 198, col: 1, offset: 5446},
			expr: &seqExpr{
				pos: position{line: 198, col: 17, offset: 5462},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 198, col: 17, offset: 5462},
						name: "_",
					},
					&zeroOrMoreExpr{
						pos: position{line: 198, col: 19, offset: 5464},
						expr: &seqExpr{
							pos: position{line: 198, col: 20, offset: 5465},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 198, col: 20, offset: 5465},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 198, col: 22, offset: 5467},
									name: "EXPRESSION",
								},
								&ruleRefExpr{
									pos:  position{line: 198, col: 33, offset: 5478},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 198, col: 35, offset: 5480},
									val:        ",",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 198, col: 39, offset: 5484},
									name: "_",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 198, col: 43, offset: 5488},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 198, col: 54, offset: 5499},
						name: "_",
					},
				},
			},
		},
		{
			name: "CLASS_REF_QUOTES",
			pos:  position{line: 205, col: 1, offset: 5667},
			expr: &seqExpr{
				pos: position{line: 205, col: 20, offset: 5686},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 205, col: 20, offset: 5686},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 205, col: 22, offset: 5688},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 205, col: 26, offset: 5692},
						name: "CLASS_REF",
					},
					&litMatcher{
						pos:        position{line: 205, col: 36, offset: 5702},
						val:        "'",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CLASS_REF",
			pos:  position{line: 207, col: 1, offset: 5707},
			expr: &seqExpr{
				pos: position{line: 207, col: 13, offset: 5719},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 207, col: 13, offset: 5719},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 207, col: 15, offset: 5721},
						name: "CLASS_TYPE",
					},
					&litMatcher{
						pos:        position{line: 207, col: 26, offset: 5732},
						val:        "\\",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 207, col: 31, offset: 5737},
						name: "CLASS_NAME",
					},
				},
			},
		},
		{
			name: "CLASS_TYPE",
			pos:  position{line: 209, col: 1, offset: 5749},
			expr: &seqExpr{
				pos: position{line: 209, col: 14, offset: 5762},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 209, col: 14, offset: 5762},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 209, col: 17, offset: 5765},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 209, col: 17, offset: 5765},
								val:        "value",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 209, col: 27, offset: 5775},
								val:        "entity",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 209, col: 38, offset: 5786},
								val:        "command",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 209, col: 50, offset: 5798},
								val:        "event",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 209, col: 60, offset: 5808},
								val:        "projection",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 209, col: 75, offset: 5823},
								val:        "invariant",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 209, col: 89, offset: 5837},
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
			pos:  position{line: 211, col: 1, offset: 5847},
			expr: &actionExpr{
				pos: position{line: 211, col: 14, offset: 5860},
				run: (*parser).callonCLASS_NAME1,
				expr: &seqExpr{
					pos: position{line: 211, col: 14, offset: 5860},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 211, col: 14, offset: 5860},
							expr: &charClassMatcher{
								pos:        position{line: 211, col: 14, offset: 5860},
								val:        "[a-zA-Z]",
								ranges:     []rune{'a', 'z', 'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 211, col: 24, offset: 5870},
							expr: &charClassMatcher{
								pos:        position{line: 211, col: 24, offset: 5870},
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
		},
		{
			name: "QUOTED_NAME",
			pos:  position{line: 215, col: 1, offset: 5922},
			expr: &actionExpr{
				pos: position{line: 215, col: 15, offset: 5936},
				run: (*parser).callonQUOTED_NAME1,
				expr: &seqExpr{
					pos: position{line: 215, col: 15, offset: 5936},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 215, col: 15, offset: 5936},
							val:        "'",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 215, col: 19, offset: 5940},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 215, col: 24, offset: 5945},
								name: "CLASS_NAME",
							},
						},
						&litMatcher{
							pos:        position{line: 215, col: 35, offset: 5956},
							val:        "'",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TYPE",
			pos:  position{line: 219, col: 1, offset: 5987},
			expr: &choiceExpr{
				pos: position{line: 219, col: 8, offset: 5994},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 219, col: 8, offset: 5994},
						name: "CLASS_REF",
					},
					&ruleRefExpr{
						pos:  position{line: 219, col: 20, offset: 6006},
						name: "VALUE_TYPE",
					},
				},
			},
		},
		{
			name: "VALUE_TYPE",
			pos:  position{line: 221, col: 1, offset: 6018},
			expr: &seqExpr{
				pos: position{line: 221, col: 14, offset: 6031},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 221, col: 14, offset: 6031},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 221, col: 17, offset: 6034},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 221, col: 17, offset: 6034},
								val:        "string",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 221, col: 28, offset: 6045},
								val:        "boolean",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 221, col: 40, offset: 6057},
								val:        "float",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 221, col: 50, offset: 6067},
								val:        "map",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 221, col: 58, offset: 6075},
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
			pos:  position{line: 223, col: 1, offset: 6085},
			expr: &seqExpr{
				pos: position{line: 223, col: 21, offset: 6105},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 223, col: 21, offset: 6105},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 223, col: 23, offset: 6107},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 223, col: 27, offset: 6111},
						name: "CLASS_NAME",
					},
					&litMatcher{
						pos:        position{line: 223, col: 38, offset: 6122},
						val:        "'",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "IDENTIFIER",
			pos:  position{line: 225, col: 1, offset: 6127},
			expr: &seqExpr{
				pos: position{line: 225, col: 14, offset: 6140},
				exprs: []interface{}{
					&oneOrMoreExpr{
						pos: position{line: 225, col: 14, offset: 6140},
						expr: &charClassMatcher{
							pos:        position{line: 225, col: 14, offset: 6140},
							val:        "[a-zA-Z]",
							ranges:     []rune{'a', 'z', 'A', 'Z'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 225, col: 24, offset: 6150},
						expr: &charClassMatcher{
							pos:        position{line: 225, col: 24, offset: 6150},
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
			pos:  position{line: 227, col: 1, offset: 6165},
			expr: &zeroOrMoreExpr{
				pos: position{line: 227, col: 5, offset: 6169},
				expr: &choiceExpr{
					pos: position{line: 227, col: 7, offset: 6171},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 227, col: 7, offset: 6171},
							name: "WHITESPACE",
						},
						&ruleRefExpr{
							pos:  position{line: 227, col: 20, offset: 6184},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "SEMI",
			pos:  position{line: 229, col: 1, offset: 6192},
			expr: &seqExpr{
				pos: position{line: 229, col: 8, offset: 6199},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 229, col: 8, offset: 6199},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 229, col: 10, offset: 6201},
						val:        ";",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 229, col: 14, offset: 6205},
						name: "_",
					},
				},
			},
		},
		{
			name: "WHITESPACE",
			pos:  position{line: 231, col: 1, offset: 6208},
			expr: &charClassMatcher{
				pos:        position{line: 231, col: 14, offset: 6221},
				val:        "[ \\t\\r]",
				chars:      []rune{' ', '\t', '\r'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EOL",
			pos:  position{line: 233, col: 1, offset: 6230},
			expr: &litMatcher{
				pos:        position{line: 233, col: 7, offset: 6236},
				val:        "\n",
				ignoreCase: false,
			},
		},
		{
			name: "EOF",
			pos:  position{line: 235, col: 1, offset: 6242},
			expr: &notExpr{
				pos: position{line: 235, col: 7, offset: 6248},
				expr: &anyMatcher{
					line: 235, col: 8, offset: 6249,
				},
			},
		},
	},
}

func (c *current) onCREATE_OBJECT1() (interface{}, error) {
	return emit(Apostrophe, ";")
}

func (p *parser) callonCREATE_OBJECT1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCREATE_OBJECT1()
}

func (c *current) onCREATE_NAMESPACE_OBJECT1(typ, name interface{}) (interface{}, error) {
	emit(Create, "create")
	emit(NamespaceObject, typ)
	return emit(QuotedName, name)
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
	return emit(UsingDatabase, name)
}

func (p *parser) callonUSING_DATABASE1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onUSING_DATABASE1(stack["name"])
}

func (c *current) onFOR_DOMAIN1(name interface{}) (interface{}, error) {
	return emit(ForDomain, name)
}

func (p *parser) callonFOR_DOMAIN1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onFOR_DOMAIN1(stack["name"])
}

func (c *current) onIN_CONTEXT1(name interface{}) (interface{}, error) {
	return emit(InContext, name)
}

func (p *parser) callonIN_CONTEXT1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onIN_CONTEXT1(stack["name"])
}

func (c *current) onWITHIN_AGGREGATE1(name interface{}) (interface{}, error) {
	return emit(WithinAggregate, name)
}

func (p *parser) callonWITHIN_AGGREGATE1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onWITHIN_AGGREGATE1(stack["name"])
}

func (c *current) onCREATE_VALUE1(name interface{}) (interface{}, error) {
	emit(Class, "value")
	return emit(QuotedName, name)
}

func (p *parser) callonCREATE_VALUE1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCREATE_VALUE1(stack["name"])
}

func (c *current) onCLASS_OPEN1() (interface{}, error) {
	return emit(ClassOpen, string(c.text))
}

func (p *parser) callonCLASS_OPEN1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCLASS_OPEN1()
}

func (c *current) onCLASS_CLOSE1() (interface{}, error) {
	return emit(ClassOpen, string(c.text))
}

func (p *parser) callonCLASS_CLOSE1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCLASS_CLOSE1()
}

func (c *current) onCLASS_NAME1() (interface{}, error) {
	return string(c.text), nil
}

func (p *parser) callonCLASS_NAME1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onCLASS_NAME1()
}

func (c *current) onQUOTED_NAME1(name interface{}) (interface{}, error) {
	return name, nil
}

func (p *parser) callonQUOTED_NAME1() (interface{}, error) {
	stack := p.vstack[len(p.vstack)-1]
	_ = stack
	return p.cur.onQUOTED_NAME1(stack["name"])
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
