package namespace

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
			name: "FILE",
			pos:  position{line: 11, col: 1, offset: 202},
			expr: &seqExpr{
				pos: position{line: 11, col: 8, offset: 209},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 11, col: 8, offset: 209},
						name: "STATEMENTS",
					},
					&ruleRefExpr{
						pos:  position{line: 11, col: 19, offset: 220},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 11, col: 21, offset: 222},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "STATEMENTS",
			pos:  position{line: 13, col: 1, offset: 227},
			expr: &labeledExpr{
				pos:   position{line: 13, col: 14, offset: 240},
				label: "statements",
				expr: &zeroOrMoreExpr{
					pos: position{line: 13, col: 25, offset: 251},
					expr: &choiceExpr{
						pos: position{line: 13, col: 26, offset: 252},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 13, col: 26, offset: 252},
								name: "BLOCK_STATEMENT",
							},
							&ruleRefExpr{
								pos:  position{line: 13, col: 43, offset: 269},
								name: "CREATE_OBJECT",
							},
							&ruleRefExpr{
								pos:  position{line: 13, col: 59, offset: 285},
								name: "CREATE_CLASS",
							},
						},
					},
				},
			},
		},
		{
			name: "CREATE_OBJECT",
			pos:  position{line: 15, col: 1, offset: 301},
			expr: &seqExpr{
				pos: position{line: 15, col: 17, offset: 317},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 15, col: 17, offset: 317},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 15, col: 19, offset: 319},
						expr: &ruleRefExpr{
							pos:  position{line: 15, col: 19, offset: 319},
							name: "CREATE_NAMESPACE_OBJECT",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 15, col: 44, offset: 344},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 15, col: 46, offset: 346},
						expr: &ruleRefExpr{
							pos:  position{line: 15, col: 46, offset: 346},
							name: "NAMESPACE",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 15, col: 57, offset: 357},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 15, col: 59, offset: 359},
						val:        ";",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "BLOCK_STATEMENT",
			pos:  position{line: 17, col: 1, offset: 364},
			expr: &seqExpr{
				pos: position{line: 17, col: 19, offset: 382},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 17, col: 19, offset: 382},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 21, offset: 384},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 31, offset: 394},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 17, col: 33, offset: 396},
						val:        ":",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 37, offset: 400},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 17, col: 39, offset: 402},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 43, offset: 406},
						name: "STATEMENTS",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 54, offset: 417},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 17, col: 56, offset: 419},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CREATE_NAMESPACE_OBJECT",
			pos:  position{line: 19, col: 1, offset: 424},
			expr: &seqExpr{
				pos: position{line: 19, col: 27, offset: 450},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 19, col: 27, offset: 450},
						val:        "create",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 37, offset: 460},
						name: "_",
					},
					&labeledExpr{
						pos:   position{line: 19, col: 39, offset: 462},
						label: "typ",
						expr: &ruleRefExpr{
							pos:  position{line: 19, col: 43, offset: 466},
							name: "NAMESPACE_OBJECT_TYPE",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 65, offset: 488},
						name: "_",
					},
					&labeledExpr{
						pos:   position{line: 19, col: 67, offset: 490},
						label: "name",
						expr: &ruleRefExpr{
							pos:  position{line: 19, col: 72, offset: 495},
							name: "QUOTED_NAME",
						},
					},
				},
			},
		},
		{
			name: "NAMESPACE_OBJECT_TYPE",
			pos:  position{line: 21, col: 1, offset: 508},
			expr: &labeledExpr{
				pos:   position{line: 21, col: 25, offset: 532},
				label: "typ",
				expr: &choiceExpr{
					pos: position{line: 21, col: 30, offset: 537},
					alternatives: []interface{}{
						&litMatcher{
							pos:        position{line: 21, col: 30, offset: 537},
							val:        "database",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 21, col: 43, offset: 550},
							val:        "domain",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 21, col: 54, offset: 561},
							val:        "context",
							ignoreCase: false,
						},
						&litMatcher{
							pos:        position{line: 21, col: 66, offset: 573},
							val:        "aggregate",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "NAMESPACE",
			pos:  position{line: 23, col: 1, offset: 587},
			expr: &zeroOrMoreExpr{
				pos: position{line: 23, col: 13, offset: 599},
				expr: &choiceExpr{
					pos: position{line: 23, col: 14, offset: 600},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 23, col: 14, offset: 600},
							name: "USING_DATABASE",
						},
						&ruleRefExpr{
							pos:  position{line: 23, col: 31, offset: 617},
							name: "FOR_DOMAIN",
						},
						&ruleRefExpr{
							pos:  position{line: 23, col: 44, offset: 630},
							name: "IN_CONTEXT",
						},
						&ruleRefExpr{
							pos:  position{line: 23, col: 57, offset: 643},
							name: "WITHIN_AGGREGATE",
						},
					},
				},
			},
		},
		{
			name: "USING_DATABASE",
			pos:  position{line: 25, col: 1, offset: 663},
			expr: &seqExpr{
				pos: position{line: 25, col: 18, offset: 680},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 25, col: 18, offset: 680},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 25, col: 20, offset: 682},
						val:        "using",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 29, offset: 691},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 25, col: 31, offset: 693},
						val:        "database",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 43, offset: 705},
						name: "_",
					},
					&labeledExpr{
						pos:   position{line: 25, col: 45, offset: 707},
						label: "name",
						expr: &ruleRefExpr{
							pos:  position{line: 25, col: 50, offset: 712},
							name: "QUOTED_NAME",
						},
					},
				},
			},
		},
		{
			name: "FOR_DOMAIN",
			pos:  position{line: 27, col: 1, offset: 725},
			expr: &seqExpr{
				pos: position{line: 27, col: 14, offset: 738},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 27, col: 14, offset: 738},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 27, col: 16, offset: 740},
						val:        "for",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 23, offset: 747},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 27, col: 25, offset: 749},
						val:        "domain",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 35, offset: 759},
						name: "_",
					},
					&labeledExpr{
						pos:   position{line: 27, col: 37, offset: 761},
						label: "name",
						expr: &ruleRefExpr{
							pos:  position{line: 27, col: 42, offset: 766},
							name: "QUOTED_NAME",
						},
					},
				},
			},
		},
		{
			name: "IN_CONTEXT",
			pos:  position{line: 29, col: 1, offset: 779},
			expr: &seqExpr{
				pos: position{line: 29, col: 14, offset: 792},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 29, col: 14, offset: 792},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 29, col: 16, offset: 794},
						val:        "in",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 29, col: 22, offset: 800},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 29, col: 24, offset: 802},
						val:        "context",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 29, col: 35, offset: 813},
						name: "_",
					},
					&labeledExpr{
						pos:   position{line: 29, col: 37, offset: 815},
						label: "name",
						expr: &ruleRefExpr{
							pos:  position{line: 29, col: 42, offset: 820},
							name: "QUOTED_NAME",
						},
					},
				},
			},
		},
		{
			name: "WITHIN_AGGREGATE",
			pos:  position{line: 31, col: 1, offset: 833},
			expr: &seqExpr{
				pos: position{line: 31, col: 20, offset: 852},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 31, col: 20, offset: 852},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 31, col: 22, offset: 854},
						val:        "within",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 31, col: 32, offset: 864},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 31, col: 34, offset: 866},
						val:        "aggregate",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 31, col: 47, offset: 879},
						name: "_",
					},
					&labeledExpr{
						pos:   position{line: 31, col: 49, offset: 881},
						label: "name",
						expr: &ruleRefExpr{
							pos:  position{line: 31, col: 54, offset: 886},
							name: "QUOTED_NAME",
						},
					},
				},
			},
		},
		{
			name: "CREATE_CLASS",
			pos:  position{line: 33, col: 1, offset: 899},
			expr: &seqExpr{
				pos: position{line: 33, col: 16, offset: 914},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 33, col: 16, offset: 914},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 18, offset: 916},
						name: "CLASS_OPEN",
					},
					&choiceExpr{
						pos: position{line: 33, col: 30, offset: 928},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 33, col: 30, offset: 928},
								name: "CREATE_VALUE",
							},
							&ruleRefExpr{
								pos:  position{line: 33, col: 45, offset: 943},
								name: "CREATE_COMMAND",
							},
							&ruleRefExpr{
								pos:  position{line: 33, col: 62, offset: 960},
								name: "CREATE_PROJECTION",
							},
							&ruleRefExpr{
								pos:  position{line: 33, col: 82, offset: 980},
								name: "CREATE_INVARIANT",
							},
							&ruleRefExpr{
								pos:  position{line: 33, col: 101, offset: 999},
								name: "CREATE_QUERY",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 115, offset: 1013},
						name: "CLASS_CLOSE",
					},
				},
			},
		},
		{
			name: "CREATE_VALUE",
			pos:  position{line: 35, col: 1, offset: 1026},
			expr: &seqExpr{
				pos: position{line: 35, col: 16, offset: 1041},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 35, col: 16, offset: 1041},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 35, col: 19, offset: 1044},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 35, col: 19, offset: 1044},
								val:        "value",
								ignoreCase: true,
							},
							&litMatcher{
								pos:        position{line: 35, col: 30, offset: 1055},
								val:        "entity",
								ignoreCase: true,
							},
							&litMatcher{
								pos:        position{line: 35, col: 42, offset: 1067},
								val:        "event",
								ignoreCase: true,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 53, offset: 1078},
						name: "_",
					},
					&labeledExpr{
						pos:   position{line: 35, col: 55, offset: 1080},
						label: "name",
						expr: &ruleRefExpr{
							pos:  position{line: 35, col: 60, offset: 1085},
							name: "QUOTED_NAME",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 72, offset: 1097},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 74, offset: 1099},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 84, offset: 1109},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 86, offset: 1111},
						name: "VALUE_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 97, offset: 1122},
						name: "_",
					},
				},
			},
		},
		{
			name: "CREATE_COMMAND",
			pos:  position{line: 37, col: 1, offset: 1125},
			expr: &seqExpr{
				pos: position{line: 37, col: 18, offset: 1142},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 37, col: 18, offset: 1142},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 37, col: 20, offset: 1144},
						val:        "command",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 32, offset: 1156},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 34, offset: 1158},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 46, offset: 1170},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 48, offset: 1172},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 58, offset: 1182},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 60, offset: 1184},
						name: "COMMAND_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 73, offset: 1197},
						name: "_",
					},
				},
			},
		},
		{
			name: "CREATE_PROJECTION",
			pos:  position{line: 39, col: 1, offset: 1200},
			expr: &seqExpr{
				pos: position{line: 39, col: 21, offset: 1220},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 39, col: 21, offset: 1220},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 39, col: 24, offset: 1223},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 39, col: 24, offset: 1223},
								val:        "aggregate",
								ignoreCase: true,
							},
							&litMatcher{
								pos:        position{line: 39, col: 39, offset: 1238},
								val:        "domain",
								ignoreCase: true,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 50, offset: 1249},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 39, col: 52, offset: 1251},
						val:        "projection",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 67, offset: 1266},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 69, offset: 1268},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 81, offset: 1280},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 83, offset: 1282},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 93, offset: 1292},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 95, offset: 1294},
						name: "PROJECTION_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 111, offset: 1310},
						name: "_",
					},
				},
			},
		},
		{
			name: "CREATE_INVARIANT",
			pos:  position{line: 41, col: 1, offset: 1313},
			expr: &seqExpr{
				pos: position{line: 41, col: 20, offset: 1332},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 41, col: 20, offset: 1332},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 41, col: 23, offset: 1335},
						val:        "invariant",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 37, offset: 1349},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 39, offset: 1351},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 51, offset: 1363},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 41, col: 53, offset: 1365},
						val:        "on",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 59, offset: 1371},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 61, offset: 1373},
						name: "CLASS_REF_QUOTES",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 78, offset: 1390},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 80, offset: 1392},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 90, offset: 1402},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 92, offset: 1404},
						name: "INVARIANT_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 107, offset: 1419},
						name: "_",
					},
				},
			},
		},
		{
			name: "CREATE_QUERY",
			pos:  position{line: 43, col: 1, offset: 1422},
			expr: &seqExpr{
				pos: position{line: 43, col: 16, offset: 1437},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 43, col: 16, offset: 1437},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 43, col: 19, offset: 1440},
						val:        "query",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 29, offset: 1450},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 31, offset: 1452},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 43, offset: 1464},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 43, col: 45, offset: 1466},
						val:        "on",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 51, offset: 1472},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 53, offset: 1474},
						name: "CLASS_REF_QUOTES",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 70, offset: 1491},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 72, offset: 1493},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 82, offset: 1503},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 84, offset: 1505},
						name: "QUERY_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 95, offset: 1516},
						name: "_",
					},
				},
			},
		},
		{
			name: "CLASS_OPEN",
			pos:  position{line: 45, col: 1, offset: 1519},
			expr: &litMatcher{
				pos:        position{line: 45, col: 14, offset: 1532},
				val:        "<|",
				ignoreCase: false,
			},
		},
		{
			name: "CLASS_CLOSE",
			pos:  position{line: 47, col: 1, offset: 1538},
			expr: &litMatcher{
				pos:        position{line: 47, col: 15, offset: 1552},
				val:        "|>",
				ignoreCase: false,
			},
		},
		{
			name: "CLASS_COMPONENT_TEST",
			pos:  position{line: 53, col: 1, offset: 1734},
			expr: &seqExpr{
				pos: position{line: 53, col: 24, offset: 1757},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 53, col: 24, offset: 1757},
						expr: &choiceExpr{
							pos: position{line: 53, col: 25, offset: 1758},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 53, col: 25, offset: 1758},
									name: "WHEN",
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 32, offset: 1765},
									name: "COMMAND_HANDLER",
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 50, offset: 1783},
									name: "PROPERTIES",
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 63, offset: 1796},
									name: "CHECK",
								},
								&ruleRefExpr{
									pos:  position{line: 53, col: 71, offset: 1804},
									name: "FUNCTION",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 53, col: 82, offset: 1815},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "QUERY_BODY",
			pos:  position{line: 55, col: 1, offset: 1820},
			expr: &zeroOrMoreExpr{
				pos: position{line: 55, col: 14, offset: 1833},
				expr: &choiceExpr{
					pos: position{line: 55, col: 15, offset: 1834},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 55, col: 15, offset: 1834},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 55, col: 28, offset: 1847},
							name: "QUERY_HANDLER",
						},
					},
				},
			},
		},
		{
			name: "INVARIANT_BODY",
			pos:  position{line: 57, col: 1, offset: 1864},
			expr: &zeroOrMoreExpr{
				pos: position{line: 57, col: 18, offset: 1881},
				expr: &choiceExpr{
					pos: position{line: 57, col: 19, offset: 1882},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 57, col: 19, offset: 1882},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 57, col: 32, offset: 1895},
							name: "CHECK",
						},
					},
				},
			},
		},
		{
			name: "PROJECTION_BODY",
			pos:  position{line: 59, col: 1, offset: 1904},
			expr: &zeroOrMoreExpr{
				pos: position{line: 59, col: 19, offset: 1922},
				expr: &choiceExpr{
					pos: position{line: 59, col: 20, offset: 1923},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 59, col: 20, offset: 1923},
							name: "WHEN",
						},
						&ruleRefExpr{
							pos:  position{line: 59, col: 27, offset: 1930},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 59, col: 40, offset: 1943},
							name: "CHECK",
						},
						&ruleRefExpr{
							pos:  position{line: 59, col: 48, offset: 1951},
							name: "FUNCTION",
						},
					},
				},
			},
		},
		{
			name: "WHEN",
			pos:  position{line: 61, col: 1, offset: 1963},
			expr: &seqExpr{
				pos: position{line: 61, col: 8, offset: 1970},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 61, col: 8, offset: 1970},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 61, col: 10, offset: 1972},
						val:        "when",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 61, col: 18, offset: 1980},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 61, col: 20, offset: 1982},
						val:        "event",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 61, col: 29, offset: 1991},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 61, col: 31, offset: 1993},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 61, col: 43, offset: 2005},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 61, col: 45, offset: 2007},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 61, col: 49, offset: 2011},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 61, col: 51, offset: 2013},
						expr: &ruleRefExpr{
							pos:  position{line: 61, col: 51, offset: 2013},
							name: "STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 61, col: 68, offset: 2030},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 61, col: 70, offset: 2032},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 61, col: 74, offset: 2036},
						name: "_",
					},
				},
			},
		},
		{
			name: "COMMAND_BODY",
			pos:  position{line: 63, col: 1, offset: 2039},
			expr: &zeroOrMoreExpr{
				pos: position{line: 63, col: 16, offset: 2054},
				expr: &choiceExpr{
					pos: position{line: 63, col: 17, offset: 2055},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 63, col: 17, offset: 2055},
							name: "COMMAND_HANDLER",
						},
						&ruleRefExpr{
							pos:  position{line: 63, col: 35, offset: 2073},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 63, col: 48, offset: 2086},
							name: "CHECK",
						},
						&ruleRefExpr{
							pos:  position{line: 63, col: 56, offset: 2094},
							name: "FUNCTION",
						},
					},
				},
			},
		},
		{
			name: "COMMAND_HANDLER",
			pos:  position{line: 65, col: 1, offset: 2106},
			expr: &seqExpr{
				pos: position{line: 65, col: 19, offset: 2124},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 65, col: 19, offset: 2124},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 65, col: 21, offset: 2126},
						val:        "handler",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 65, col: 32, offset: 2137},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 65, col: 34, offset: 2139},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 65, col: 38, offset: 2143},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 65, col: 40, offset: 2145},
						expr: &ruleRefExpr{
							pos:  position{line: 65, col: 40, offset: 2145},
							name: "COMMAND_STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 65, col: 65, offset: 2170},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 65, col: 67, offset: 2172},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 65, col: 71, offset: 2176},
						name: "_",
					},
				},
			},
		},
		{
			name: "QUERY_HANDLER",
			pos:  position{line: 67, col: 1, offset: 2179},
			expr: &seqExpr{
				pos: position{line: 67, col: 17, offset: 2195},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 67, col: 17, offset: 2195},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 67, col: 19, offset: 2197},
						val:        "handler",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 67, col: 30, offset: 2208},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 67, col: 32, offset: 2210},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 67, col: 36, offset: 2214},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 67, col: 38, offset: 2216},
						expr: &ruleRefExpr{
							pos:  position{line: 67, col: 38, offset: 2216},
							name: "STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 67, col: 55, offset: 2233},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 67, col: 57, offset: 2235},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 67, col: 61, offset: 2239},
						name: "_",
					},
				},
			},
		},
		{
			name: "COMMAND_STATEMENT_BLOCK",
			pos:  position{line: 69, col: 1, offset: 2242},
			expr: &seqExpr{
				pos: position{line: 69, col: 27, offset: 2268},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 69, col: 27, offset: 2268},
						name: "_",
					},
					&oneOrMoreExpr{
						pos: position{line: 69, col: 29, offset: 2270},
						expr: &ruleRefExpr{
							pos:  position{line: 69, col: 30, offset: 2271},
							name: "COMMAND_STATEMENT",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 69, col: 50, offset: 2291},
						name: "_",
					},
				},
			},
		},
		{
			name: "COMMAND_STATEMENT",
			pos:  position{line: 71, col: 1, offset: 2294},
			expr: &choiceExpr{
				pos: position{line: 71, col: 21, offset: 2314},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 71, col: 21, offset: 2314},
						name: "STATEMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 71, col: 33, offset: 2326},
						name: "ASSERT",
					},
					&ruleRefExpr{
						pos:  position{line: 71, col: 42, offset: 2335},
						name: "APPLY",
					},
				},
			},
		},
		{
			name: "ASSERT",
			pos:  position{line: 73, col: 1, offset: 2342},
			expr: &seqExpr{
				pos: position{line: 73, col: 10, offset: 2351},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 73, col: 10, offset: 2351},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 73, col: 12, offset: 2353},
						val:        "assert",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 73, col: 22, offset: 2363},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 73, col: 24, offset: 2365},
						val:        "invariant",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 73, col: 37, offset: 2378},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 73, col: 39, offset: 2380},
						expr: &litMatcher{
							pos:        position{line: 73, col: 40, offset: 2381},
							val:        "not",
							ignoreCase: true,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 73, col: 49, offset: 2390},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 73, col: 51, offset: 2392},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 73, col: 63, offset: 2404},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 73, col: 65, offset: 2406},
						expr: &ruleRefExpr{
							pos:  position{line: 73, col: 65, offset: 2406},
							name: "ARGUMENTS",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 73, col: 76, offset: 2417},
						name: "SEMI",
					},
				},
			},
		},
		{
			name: "APPLY",
			pos:  position{line: 75, col: 1, offset: 2423},
			expr: &seqExpr{
				pos: position{line: 75, col: 9, offset: 2431},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 75, col: 9, offset: 2431},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 75, col: 11, offset: 2433},
						val:        "apply",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 20, offset: 2442},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 75, col: 22, offset: 2444},
						val:        "event",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 31, offset: 2453},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 33, offset: 2455},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 45, offset: 2467},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 75, col: 47, offset: 2469},
						expr: &ruleRefExpr{
							pos:  position{line: 75, col: 47, offset: 2469},
							name: "ARGUMENTS",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 58, offset: 2480},
						name: "SEMI",
					},
				},
			},
		},
		{
			name: "VALUE_BODY",
			pos:  position{line: 77, col: 1, offset: 2486},
			expr: &zeroOrMoreExpr{
				pos: position{line: 77, col: 14, offset: 2499},
				expr: &ruleRefExpr{
					pos:  position{line: 77, col: 15, offset: 2500},
					name: "VALUE_COMPONENTS",
				},
			},
		},
		{
			name: "VALUE_COMPONENTS",
			pos:  position{line: 79, col: 1, offset: 2520},
			expr: &choiceExpr{
				pos: position{line: 79, col: 20, offset: 2539},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 79, col: 20, offset: 2539},
						name: "PROPERTIES",
					},
					&ruleRefExpr{
						pos:  position{line: 79, col: 33, offset: 2552},
						name: "CHECK",
					},
					&ruleRefExpr{
						pos:  position{line: 79, col: 41, offset: 2560},
						name: "FUNCTION",
					},
				},
			},
		},
		{
			name: "PROPERTIES",
			pos:  position{line: 81, col: 1, offset: 2570},
			expr: &seqExpr{
				pos: position{line: 81, col: 14, offset: 2583},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 81, col: 14, offset: 2583},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 81, col: 16, offset: 2585},
						val:        "properties",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 81, col: 30, offset: 2599},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 81, col: 32, offset: 2601},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 81, col: 36, offset: 2605},
						name: "PROPERTY_LIST",
					},
					&litMatcher{
						pos:        position{line: 81, col: 50, offset: 2619},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 81, col: 54, offset: 2623},
						name: "_",
					},
				},
			},
		},
		{
			name: "PROPERTY_LIST",
			pos:  position{line: 83, col: 1, offset: 2626},
			expr: &zeroOrMoreExpr{
				pos: position{line: 83, col: 17, offset: 2642},
				expr: &ruleRefExpr{
					pos:  position{line: 83, col: 18, offset: 2643},
					name: "PROPERTY",
				},
			},
		},
		{
			name: "PROPERTY",
			pos:  position{line: 85, col: 1, offset: 2655},
			expr: &seqExpr{
				pos: position{line: 85, col: 12, offset: 2666},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 85, col: 12, offset: 2666},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 85, col: 14, offset: 2668},
						name: "TYPE",
					},
					&ruleRefExpr{
						pos:  position{line: 85, col: 19, offset: 2673},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 85, col: 21, offset: 2675},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 85, col: 32, offset: 2686},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 85, col: 35, offset: 2689},
						expr: &seqExpr{
							pos: position{line: 85, col: 36, offset: 2690},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 85, col: 36, offset: 2690},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 85, col: 40, offset: 2694},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 85, col: 42, offset: 2696},
									name: "EXPRESSION",
								},
								&ruleRefExpr{
									pos:  position{line: 85, col: 53, offset: 2707},
									name: "_",
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 85, col: 57, offset: 2711},
						val:        ";",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 85, col: 61, offset: 2715},
						name: "_",
					},
				},
			},
		},
		{
			name: "CHECK",
			pos:  position{line: 87, col: 1, offset: 2718},
			expr: &seqExpr{
				pos: position{line: 87, col: 9, offset: 2726},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 87, col: 9, offset: 2726},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 87, col: 11, offset: 2728},
						val:        "check",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 87, col: 20, offset: 2737},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 87, col: 22, offset: 2739},
						val:        "(",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 87, col: 26, offset: 2743},
						expr: &ruleRefExpr{
							pos:  position{line: 87, col: 26, offset: 2743},
							name: "STATEMENT_BLOCK",
						},
					},
					&litMatcher{
						pos:        position{line: 87, col: 43, offset: 2760},
						val:        ")",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 87, col: 47, offset: 2764},
						name: "_",
					},
				},
			},
		},
		{
			name: "FUNCTION",
			pos:  position{line: 89, col: 1, offset: 2767},
			expr: &seqExpr{
				pos: position{line: 89, col: 12, offset: 2778},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 89, col: 12, offset: 2778},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 89, col: 14, offset: 2780},
						val:        "function",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 89, col: 26, offset: 2792},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 89, col: 28, offset: 2794},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 89, col: 39, offset: 2805},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 89, col: 41, offset: 2807},
						name: "PARAMETERS",
					},
					&ruleRefExpr{
						pos:  position{line: 89, col: 53, offset: 2819},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 89, col: 55, offset: 2821},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 89, col: 59, offset: 2825},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 89, col: 61, offset: 2827},
						expr: &ruleRefExpr{
							pos:  position{line: 89, col: 61, offset: 2827},
							name: "STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 89, col: 78, offset: 2844},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 89, col: 80, offset: 2846},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 89, col: 84, offset: 2850},
						name: "_",
					},
				},
			},
		},
		{
			name: "PARAMETERS",
			pos:  position{line: 91, col: 1, offset: 2853},
			expr: &seqExpr{
				pos: position{line: 91, col: 14, offset: 2866},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 91, col: 14, offset: 2866},
						val:        "(",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 91, col: 18, offset: 2870},
						name: "PARAMETER_LIST",
					},
					&litMatcher{
						pos:        position{line: 91, col: 33, offset: 2885},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "PARAMETER_LIST",
			pos:  position{line: 93, col: 1, offset: 2890},
			expr: &seqExpr{
				pos: position{line: 93, col: 18, offset: 2907},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 93, col: 18, offset: 2907},
						name: "_",
					},
					&zeroOrMoreExpr{
						pos: position{line: 93, col: 20, offset: 2909},
						expr: &seqExpr{
							pos: position{line: 93, col: 21, offset: 2910},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 93, col: 21, offset: 2910},
									name: "PARAMETER",
								},
								&litMatcher{
									pos:        position{line: 93, col: 31, offset: 2920},
									val:        ",",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 93, col: 35, offset: 2924},
									name: "_",
								},
							},
						},
					},
					&zeroOrOneExpr{
						pos: position{line: 93, col: 40, offset: 2929},
						expr: &ruleRefExpr{
							pos:  position{line: 93, col: 40, offset: 2929},
							name: "PARAMETER",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 93, col: 51, offset: 2940},
						name: "_",
					},
				},
			},
		},
		{
			name: "PARAMETER",
			pos:  position{line: 95, col: 1, offset: 2943},
			expr: &seqExpr{
				pos: position{line: 95, col: 13, offset: 2955},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 95, col: 13, offset: 2955},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 95, col: 15, offset: 2957},
						name: "CLASS_REF",
					},
					&ruleRefExpr{
						pos:  position{line: 95, col: 25, offset: 2967},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 95, col: 27, offset: 2969},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 95, col: 38, offset: 2980},
						name: "_",
					},
				},
			},
		},
		{
			name: "STATEMENT_BLOCK",
			pos:  position{line: 100, col: 1, offset: 3152},
			expr: &seqExpr{
				pos: position{line: 100, col: 19, offset: 3170},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 100, col: 19, offset: 3170},
						name: "_",
					},
					&oneOrMoreExpr{
						pos: position{line: 100, col: 21, offset: 3172},
						expr: &ruleRefExpr{
							pos:  position{line: 100, col: 22, offset: 3173},
							name: "STATEMENT",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 100, col: 34, offset: 3185},
						name: "_",
					},
				},
			},
		},
		{
			name: "STATEMENT",
			pos:  position{line: 102, col: 1, offset: 3188},
			expr: &choiceExpr{
				pos: position{line: 102, col: 13, offset: 3200},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 102, col: 13, offset: 3200},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 102, col: 13, offset: 3200},
								name: "RETURN",
							},
							&ruleRefExpr{
								pos:  position{line: 102, col: 20, offset: 3207},
								name: "SEMI",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 102, col: 27, offset: 3214},
						name: "IF",
					},
					&ruleRefExpr{
						pos:  position{line: 102, col: 32, offset: 3219},
						name: "FOREACH",
					},
					&seqExpr{
						pos: position{line: 102, col: 42, offset: 3229},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 102, col: 42, offset: 3229},
								name: "EXPRESSION",
							},
							&ruleRefExpr{
								pos:  position{line: 102, col: 53, offset: 3240},
								name: "SEMI",
							},
						},
					},
				},
			},
		},
		{
			name: "IF",
			pos:  position{line: 104, col: 1, offset: 3246},
			expr: &seqExpr{
				pos: position{line: 104, col: 6, offset: 3251},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 104, col: 6, offset: 3251},
						val:        "if",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 104, col: 11, offset: 3256},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 104, col: 13, offset: 3258},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 104, col: 24, offset: 3269},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 104, col: 26, offset: 3271},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 104, col: 30, offset: 3275},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 104, col: 32, offset: 3277},
						expr: &ruleRefExpr{
							pos:  position{line: 104, col: 32, offset: 3277},
							name: "STATEMENT_BLOCK",
						},
					},
					&litMatcher{
						pos:        position{line: 104, col: 49, offset: 3294},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 104, col: 53, offset: 3298},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 104, col: 55, offset: 3300},
						expr: &seqExpr{
							pos: position{line: 104, col: 56, offset: 3301},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 104, col: 56, offset: 3301},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 104, col: 63, offset: 3308},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 104, col: 65, offset: 3310},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 104, col: 69, offset: 3314},
									name: "_",
								},
								&zeroOrOneExpr{
									pos: position{line: 104, col: 71, offset: 3316},
									expr: &ruleRefExpr{
										pos:  position{line: 104, col: 71, offset: 3316},
										name: "STATEMENT_BLOCK",
									},
								},
								&litMatcher{
									pos:        position{line: 104, col: 88, offset: 3333},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 104, col: 92, offset: 3337},
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
			pos:  position{line: 106, col: 1, offset: 3342},
			expr: &seqExpr{
				pos: position{line: 106, col: 11, offset: 3352},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 106, col: 11, offset: 3352},
						val:        "foreach",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 106, col: 21, offset: 3362},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 106, col: 23, offset: 3364},
						val:        "(",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 106, col: 27, offset: 3368},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 106, col: 29, offset: 3370},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 106, col: 40, offset: 3381},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 106, col: 42, offset: 3383},
						val:        "as",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 106, col: 47, offset: 3388},
						expr: &seqExpr{
							pos: position{line: 106, col: 48, offset: 3389},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 106, col: 48, offset: 3389},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 106, col: 50, offset: 3391},
									name: "IDENTIFIER",
								},
								&ruleRefExpr{
									pos:  position{line: 106, col: 61, offset: 3402},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 106, col: 63, offset: 3404},
									val:        "=>",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 106, col: 70, offset: 3411},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 106, col: 72, offset: 3413},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 106, col: 83, offset: 3424},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 106, col: 85, offset: 3426},
						val:        ")",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 106, col: 89, offset: 3430},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 106, col: 91, offset: 3432},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 106, col: 95, offset: 3436},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 106, col: 97, offset: 3438},
						expr: &ruleRefExpr{
							pos:  position{line: 106, col: 97, offset: 3438},
							name: "STATEMENT_BLOCK",
						},
					},
					&litMatcher{
						pos:        position{line: 106, col: 114, offset: 3455},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "RETURN",
			pos:  position{line: 108, col: 1, offset: 3460},
			expr: &seqExpr{
				pos: position{line: 108, col: 10, offset: 3469},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 108, col: 10, offset: 3469},
						val:        "return",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 108, col: 19, offset: 3478},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 108, col: 21, offset: 3480},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "EXRESSION_TEST",
			pos:  position{line: 114, col: 1, offset: 3663},
			expr: &seqExpr{
				pos: position{line: 114, col: 18, offset: 3680},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 114, col: 18, offset: 3680},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 114, col: 29, offset: 3691},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "EXPRESSION",
			pos:  position{line: 116, col: 1, offset: 3696},
			expr: &choiceExpr{
				pos: position{line: 116, col: 14, offset: 3709},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 116, col: 14, offset: 3709},
						name: "QUERY",
					},
					&ruleRefExpr{
						pos:  position{line: 116, col: 22, offset: 3717},
						name: "ARITHMETIC",
					},
					&ruleRefExpr{
						pos:  position{line: 116, col: 35, offset: 3730},
						name: "COMPARISON",
					},
					&ruleRefExpr{
						pos:  position{line: 116, col: 48, offset: 3743},
						name: "ASSIGNMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 116, col: 60, offset: 3755},
						name: "LOGICAL",
					},
					&ruleRefExpr{
						pos:  position{line: 116, col: 70, offset: 3765},
						name: "ATOMIC",
					},
				},
			},
		},
		{
			name: "ATOMIC",
			pos:  position{line: 118, col: 1, offset: 3773},
			expr: &choiceExpr{
				pos: position{line: 118, col: 10, offset: 3782},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 118, col: 10, offset: 3782},
						name: "PARENTHESIS",
					},
					&ruleRefExpr{
						pos:  position{line: 118, col: 24, offset: 3796},
						name: "NEW",
					},
					&ruleRefExpr{
						pos:  position{line: 118, col: 30, offset: 3802},
						name: "METHODCALL",
					},
					&ruleRefExpr{
						pos:  position{line: 118, col: 43, offset: 3815},
						name: "OBJECTACCESS",
					},
					&ruleRefExpr{
						pos:  position{line: 118, col: 58, offset: 3830},
						name: "ARRAY",
					},
					&ruleRefExpr{
						pos:  position{line: 118, col: 66, offset: 3838},
						name: "LITERAL",
					},
					&ruleRefExpr{
						pos:  position{line: 118, col: 76, offset: 3848},
						name: "UNARY",
					},
				},
			},
		},
		{
			name: "LITERAL",
			pos:  position{line: 120, col: 1, offset: 3855},
			expr: &choiceExpr{
				pos: position{line: 120, col: 11, offset: 3865},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 120, col: 11, offset: 3865},
						name: "STRING",
					},
					&ruleRefExpr{
						pos:  position{line: 120, col: 20, offset: 3874},
						name: "FLOAT",
					},
					&ruleRefExpr{
						pos:  position{line: 120, col: 28, offset: 3882},
						name: "BOOLEAN",
					},
					&ruleRefExpr{
						pos:  position{line: 120, col: 38, offset: 3892},
						name: "NULL",
					},
					&ruleRefExpr{
						pos:  position{line: 120, col: 45, offset: 3899},
						name: "INT",
					},
				},
			},
		},
		{
			name: "NEW",
			pos:  position{line: 122, col: 1, offset: 3904},
			expr: &seqExpr{
				pos: position{line: 122, col: 7, offset: 3910},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 122, col: 7, offset: 3910},
						name: "CLASS_REF_QUOTES",
					},
					&ruleRefExpr{
						pos:  position{line: 122, col: 24, offset: 3927},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 122, col: 26, offset: 3929},
						expr: &ruleRefExpr{
							pos:  position{line: 122, col: 26, offset: 3929},
							name: "ARGUMENTS",
						},
					},
				},
			},
		},
		{
			name: "BOOLEAN",
			pos:  position{line: 124, col: 1, offset: 3941},
			expr: &choiceExpr{
				pos: position{line: 124, col: 12, offset: 3952},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 124, col: 12, offset: 3952},
						val:        "true",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 124, col: 19, offset: 3959},
						val:        "false",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "NULL",
			pos:  position{line: 126, col: 1, offset: 3968},
			expr: &litMatcher{
				pos:        position{line: 126, col: 8, offset: 3975},
				val:        "null",
				ignoreCase: false,
			},
		},
		{
			name: "ARRAY",
			pos:  position{line: 128, col: 1, offset: 3983},
			expr: &seqExpr{
				pos: position{line: 128, col: 9, offset: 3991},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 128, col: 9, offset: 3991},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 128, col: 11, offset: 3993},
						val:        "[",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 128, col: 15, offset: 3997},
						expr: &ruleRefExpr{
							pos:  position{line: 128, col: 15, offset: 3997},
							name: "ARGUMENTLIST",
						},
					},
					&litMatcher{
						pos:        position{line: 128, col: 29, offset: 4011},
						val:        "]",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 128, col: 33, offset: 4015},
						name: "_",
					},
				},
			},
		},
		{
			name: "STRING",
			pos:  position{line: 130, col: 1, offset: 4018},
			expr: &seqExpr{
				pos: position{line: 130, col: 10, offset: 4027},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 130, col: 10, offset: 4027},
						val:        "\"",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 130, col: 15, offset: 4032},
						expr: &charClassMatcher{
							pos:        position{line: 130, col: 15, offset: 4032},
							val:        "[a-zA-Z0-9]",
							ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&litMatcher{
						pos:        position{line: 130, col: 28, offset: 4045},
						val:        "\"",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "INT",
			pos:  position{line: 132, col: 1, offset: 4051},
			expr: &oneOrMoreExpr{
				pos: position{line: 132, col: 7, offset: 4057},
				expr: &charClassMatcher{
					pos:        position{line: 132, col: 7, offset: 4057},
					val:        "[0-9]",
					ranges:     []rune{'0', '9'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "FLOAT",
			pos:  position{line: 134, col: 1, offset: 4065},
			expr: &seqExpr{
				pos: position{line: 134, col: 9, offset: 4073},
				exprs: []interface{}{
					&oneOrMoreExpr{
						pos: position{line: 134, col: 9, offset: 4073},
						expr: &charClassMatcher{
							pos:        position{line: 134, col: 9, offset: 4073},
							val:        "[0-9]",
							ranges:     []rune{'0', '9'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&charClassMatcher{
						pos:        position{line: 134, col: 16, offset: 4080},
						val:        "[.]",
						chars:      []rune{'.'},
						ignoreCase: false,
						inverted:   false,
					},
					&oneOrMoreExpr{
						pos: position{line: 134, col: 20, offset: 4084},
						expr: &charClassMatcher{
							pos:        position{line: 134, col: 20, offset: 4084},
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
			pos:  position{line: 136, col: 1, offset: 4092},
			expr: &seqExpr{
				pos: position{line: 136, col: 15, offset: 4106},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 136, col: 15, offset: 4106},
						val:        "(",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 136, col: 19, offset: 4110},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 136, col: 21, offset: 4112},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 136, col: 32, offset: 4123},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 136, col: 34, offset: 4125},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "UNARY",
			pos:  position{line: 138, col: 1, offset: 4130},
			expr: &choiceExpr{
				pos: position{line: 138, col: 9, offset: 4138},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 138, col: 9, offset: 4138},
						name: "INCREMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 138, col: 21, offset: 4150},
						name: "DECREMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 138, col: 33, offset: 4162},
						name: "NEGATE",
					},
					&ruleRefExpr{
						pos:  position{line: 138, col: 42, offset: 4171},
						name: "NOT",
					},
					&ruleRefExpr{
						pos:  position{line: 138, col: 48, offset: 4177},
						name: "POSITIVE",
					},
				},
			},
		},
		{
			name: "INCREMENT",
			pos:  position{line: 140, col: 1, offset: 4187},
			expr: &seqExpr{
				pos: position{line: 140, col: 13, offset: 4199},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 140, col: 13, offset: 4199},
						name: "OBJECTACCESS",
					},
					&litMatcher{
						pos:        position{line: 140, col: 26, offset: 4212},
						val:        "++",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DECREMENT",
			pos:  position{line: 142, col: 1, offset: 4218},
			expr: &seqExpr{
				pos: position{line: 142, col: 13, offset: 4230},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 142, col: 13, offset: 4230},
						name: "OBJECTACCESS",
					},
					&litMatcher{
						pos:        position{line: 142, col: 26, offset: 4243},
						val:        "--",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "NEGATE",
			pos:  position{line: 144, col: 1, offset: 4249},
			expr: &seqExpr{
				pos: position{line: 144, col: 10, offset: 4258},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 144, col: 10, offset: 4258},
						val:        "-",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 144, col: 14, offset: 4262},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "NOT",
			pos:  position{line: 146, col: 1, offset: 4276},
			expr: &seqExpr{
				pos: position{line: 146, col: 7, offset: 4282},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 146, col: 7, offset: 4282},
						val:        "!",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 146, col: 11, offset: 4286},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "POSITIVE",
			pos:  position{line: 148, col: 1, offset: 4300},
			expr: &seqExpr{
				pos: position{line: 148, col: 12, offset: 4311},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 148, col: 12, offset: 4311},
						val:        "+",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 148, col: 16, offset: 4315},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "ARITHMETIC",
			pos:  position{line: 150, col: 1, offset: 4329},
			expr: &seqExpr{
				pos: position{line: 150, col: 14, offset: 4342},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 150, col: 14, offset: 4342},
						name: "ATOMIC",
					},
					&ruleRefExpr{
						pos:  position{line: 150, col: 21, offset: 4349},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 150, col: 23, offset: 4351},
						name: "OPERATOR",
					},
					&ruleRefExpr{
						pos:  position{line: 150, col: 32, offset: 4360},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 150, col: 34, offset: 4362},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "OPERATOR",
			pos:  position{line: 152, col: 1, offset: 4374},
			expr: &choiceExpr{
				pos: position{line: 152, col: 12, offset: 4385},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 152, col: 12, offset: 4385},
						val:        "+",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 152, col: 18, offset: 4391},
						val:        "-",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 152, col: 24, offset: 4397},
						val:        "/",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 152, col: 30, offset: 4403},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 152, col: 36, offset: 4409},
						val:        "%",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "ASSIGNMENT",
			pos:  position{line: 154, col: 1, offset: 4414},
			expr: &seqExpr{
				pos: position{line: 154, col: 14, offset: 4427},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 154, col: 14, offset: 4427},
						name: "OBJECTACCESS",
					},
					&ruleRefExpr{
						pos:  position{line: 154, col: 27, offset: 4440},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 154, col: 29, offset: 4442},
						val:        "=",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 154, col: 33, offset: 4446},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 154, col: 35, offset: 4448},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "LOGICAL",
			pos:  position{line: 156, col: 1, offset: 4460},
			expr: &seqExpr{
				pos: position{line: 156, col: 11, offset: 4470},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 156, col: 11, offset: 4470},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 156, col: 22, offset: 4481},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 156, col: 25, offset: 4484},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 156, col: 25, offset: 4484},
								val:        "and",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 156, col: 33, offset: 4492},
								val:        "or",
								ignoreCase: false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 156, col: 39, offset: 4498},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 156, col: 41, offset: 4500},
						name: "ATOMIC",
					},
				},
			},
		},
		{
			name: "COMPARISON",
			pos:  position{line: 158, col: 1, offset: 4508},
			expr: &seqExpr{
				pos: position{line: 158, col: 14, offset: 4521},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 158, col: 14, offset: 4521},
						name: "ATOMIC",
					},
					&ruleRefExpr{
						pos:  position{line: 158, col: 21, offset: 4528},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 158, col: 24, offset: 4531},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 158, col: 24, offset: 4531},
								val:        "===",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 158, col: 32, offset: 4539},
								val:        "!==",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 158, col: 40, offset: 4547},
								val:        "==",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 158, col: 47, offset: 4554},
								val:        "!=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 158, col: 54, offset: 4561},
								val:        "<=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 158, col: 61, offset: 4568},
								val:        ">=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 158, col: 68, offset: 4575},
								val:        "<",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 158, col: 74, offset: 4581},
								val:        ">",
								ignoreCase: false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 158, col: 79, offset: 4586},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 158, col: 81, offset: 4588},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "QUERY",
			pos:  position{line: 160, col: 1, offset: 4600},
			expr: &seqExpr{
				pos: position{line: 160, col: 9, offset: 4608},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 160, col: 9, offset: 4608},
						val:        "run",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 160, col: 16, offset: 4615},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 160, col: 18, offset: 4617},
						val:        "query",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 160, col: 27, offset: 4626},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 160, col: 29, offset: 4628},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 160, col: 41, offset: 4640},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 160, col: 43, offset: 4642},
						expr: &ruleRefExpr{
							pos:  position{line: 160, col: 43, offset: 4642},
							name: "ARGUMENTS",
						},
					},
				},
			},
		},
		{
			name: "OBJECTACCESS",
			pos:  position{line: 162, col: 1, offset: 4654},
			expr: &seqExpr{
				pos: position{line: 162, col: 16, offset: 4669},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 162, col: 16, offset: 4669},
						expr: &seqExpr{
							pos: position{line: 162, col: 17, offset: 4670},
							exprs: []interface{}{
								&choiceExpr{
									pos: position{line: 162, col: 18, offset: 4671},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 162, col: 18, offset: 4671},
											name: "METHODCALL",
										},
										&ruleRefExpr{
											pos:  position{line: 162, col: 31, offset: 4684},
											name: "IDENTIFIER",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 162, col: 43, offset: 4696},
									val:        "->",
									ignoreCase: false,
								},
							},
						},
					},
					&choiceExpr{
						pos: position{line: 162, col: 51, offset: 4704},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 162, col: 51, offset: 4704},
								name: "METHODCALL",
							},
							&ruleRefExpr{
								pos:  position{line: 162, col: 64, offset: 4717},
								name: "IDENTIFIER",
							},
						},
					},
				},
			},
		},
		{
			name: "METHODCALL",
			pos:  position{line: 164, col: 1, offset: 4730},
			expr: &seqExpr{
				pos: position{line: 164, col: 14, offset: 4743},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 164, col: 14, offset: 4743},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 164, col: 25, offset: 4754},
						name: "ARGUMENTS",
					},
				},
			},
		},
		{
			name: "ARGUMENTS",
			pos:  position{line: 166, col: 1, offset: 4765},
			expr: &seqExpr{
				pos: position{line: 166, col: 13, offset: 4777},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 166, col: 13, offset: 4777},
						val:        "(",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 166, col: 17, offset: 4781},
						expr: &ruleRefExpr{
							pos:  position{line: 166, col: 17, offset: 4781},
							name: "ARGUMENTLIST",
						},
					},
					&litMatcher{
						pos:        position{line: 166, col: 31, offset: 4795},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "ARGUMENTLIST",
			pos:  position{line: 168, col: 1, offset: 4800},
			expr: &seqExpr{
				pos: position{line: 168, col: 17, offset: 4816},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 168, col: 17, offset: 4816},
						name: "_",
					},
					&zeroOrMoreExpr{
						pos: position{line: 168, col: 19, offset: 4818},
						expr: &seqExpr{
							pos: position{line: 168, col: 20, offset: 4819},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 168, col: 20, offset: 4819},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 168, col: 22, offset: 4821},
									name: "EXPRESSION",
								},
								&ruleRefExpr{
									pos:  position{line: 168, col: 33, offset: 4832},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 168, col: 35, offset: 4834},
									val:        ",",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 168, col: 39, offset: 4838},
									name: "_",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 168, col: 43, offset: 4842},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 168, col: 54, offset: 4853},
						name: "_",
					},
				},
			},
		},
		{
			name: "CLASS_REF_QUOTES",
			pos:  position{line: 175, col: 1, offset: 5021},
			expr: &seqExpr{
				pos: position{line: 175, col: 20, offset: 5040},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 175, col: 20, offset: 5040},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 175, col: 22, offset: 5042},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 175, col: 26, offset: 5046},
						name: "CLASS_REF",
					},
					&litMatcher{
						pos:        position{line: 175, col: 36, offset: 5056},
						val:        "'",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CLASS_REF",
			pos:  position{line: 177, col: 1, offset: 5061},
			expr: &seqExpr{
				pos: position{line: 177, col: 13, offset: 5073},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 177, col: 13, offset: 5073},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 177, col: 15, offset: 5075},
						name: "CLASS_TYPE",
					},
					&litMatcher{
						pos:        position{line: 177, col: 26, offset: 5086},
						val:        "\\",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 177, col: 31, offset: 5091},
						name: "CLASS_NAME",
					},
				},
			},
		},
		{
			name: "CLASS_TYPE",
			pos:  position{line: 179, col: 1, offset: 5103},
			expr: &seqExpr{
				pos: position{line: 179, col: 14, offset: 5116},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 179, col: 14, offset: 5116},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 179, col: 17, offset: 5119},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 179, col: 17, offset: 5119},
								val:        "value",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 179, col: 27, offset: 5129},
								val:        "entity",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 179, col: 38, offset: 5140},
								val:        "command",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 179, col: 50, offset: 5152},
								val:        "event",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 179, col: 60, offset: 5162},
								val:        "projection",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 179, col: 75, offset: 5177},
								val:        "invariant",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 179, col: 89, offset: 5191},
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
			pos:  position{line: 181, col: 1, offset: 5201},
			expr: &actionExpr{
				pos: position{line: 181, col: 14, offset: 5214},
				run: (*parser).callonCLASS_NAME1,
				expr: &seqExpr{
					pos: position{line: 181, col: 14, offset: 5214},
					exprs: []interface{}{
						&oneOrMoreExpr{
							pos: position{line: 181, col: 14, offset: 5214},
							expr: &charClassMatcher{
								pos:        position{line: 181, col: 14, offset: 5214},
								val:        "[a-zA-Z]",
								ranges:     []rune{'a', 'z', 'A', 'Z'},
								ignoreCase: false,
								inverted:   false,
							},
						},
						&zeroOrMoreExpr{
							pos: position{line: 181, col: 24, offset: 5224},
							expr: &charClassMatcher{
								pos:        position{line: 181, col: 24, offset: 5224},
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
			pos:  position{line: 185, col: 1, offset: 5276},
			expr: &actionExpr{
				pos: position{line: 185, col: 15, offset: 5290},
				run: (*parser).callonQUOTED_NAME1,
				expr: &seqExpr{
					pos: position{line: 185, col: 15, offset: 5290},
					exprs: []interface{}{
						&litMatcher{
							pos:        position{line: 185, col: 15, offset: 5290},
							val:        "'",
							ignoreCase: false,
						},
						&labeledExpr{
							pos:   position{line: 185, col: 19, offset: 5294},
							label: "name",
							expr: &ruleRefExpr{
								pos:  position{line: 185, col: 24, offset: 5299},
								name: "CLASS_NAME",
							},
						},
						&litMatcher{
							pos:        position{line: 185, col: 35, offset: 5310},
							val:        "'",
							ignoreCase: false,
						},
					},
				},
			},
		},
		{
			name: "TYPE",
			pos:  position{line: 189, col: 1, offset: 5341},
			expr: &choiceExpr{
				pos: position{line: 189, col: 8, offset: 5348},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 189, col: 8, offset: 5348},
						name: "CLASS_REF",
					},
					&ruleRefExpr{
						pos:  position{line: 189, col: 20, offset: 5360},
						name: "VALUE_TYPE",
					},
				},
			},
		},
		{
			name: "VALUE_TYPE",
			pos:  position{line: 191, col: 1, offset: 5372},
			expr: &seqExpr{
				pos: position{line: 191, col: 14, offset: 5385},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 191, col: 14, offset: 5385},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 191, col: 17, offset: 5388},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 191, col: 17, offset: 5388},
								val:        "string",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 191, col: 28, offset: 5399},
								val:        "boolean",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 191, col: 40, offset: 5411},
								val:        "float",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 191, col: 50, offset: 5421},
								val:        "map",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 191, col: 58, offset: 5429},
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
			pos:  position{line: 193, col: 1, offset: 5439},
			expr: &seqExpr{
				pos: position{line: 193, col: 21, offset: 5459},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 193, col: 21, offset: 5459},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 193, col: 23, offset: 5461},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 193, col: 27, offset: 5465},
						name: "CLASS_NAME",
					},
					&litMatcher{
						pos:        position{line: 193, col: 38, offset: 5476},
						val:        "'",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "IDENTIFIER",
			pos:  position{line: 195, col: 1, offset: 5481},
			expr: &seqExpr{
				pos: position{line: 195, col: 14, offset: 5494},
				exprs: []interface{}{
					&oneOrMoreExpr{
						pos: position{line: 195, col: 14, offset: 5494},
						expr: &charClassMatcher{
							pos:        position{line: 195, col: 14, offset: 5494},
							val:        "[a-zA-Z]",
							ranges:     []rune{'a', 'z', 'A', 'Z'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 195, col: 24, offset: 5504},
						expr: &charClassMatcher{
							pos:        position{line: 195, col: 24, offset: 5504},
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
			pos:  position{line: 197, col: 1, offset: 5519},
			expr: &zeroOrMoreExpr{
				pos: position{line: 197, col: 5, offset: 5523},
				expr: &choiceExpr{
					pos: position{line: 197, col: 7, offset: 5525},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 197, col: 7, offset: 5525},
							name: "WHITESPACE",
						},
						&ruleRefExpr{
							pos:  position{line: 197, col: 20, offset: 5538},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "SEMI",
			pos:  position{line: 199, col: 1, offset: 5546},
			expr: &seqExpr{
				pos: position{line: 199, col: 8, offset: 5553},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 199, col: 8, offset: 5553},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 199, col: 10, offset: 5555},
						val:        ";",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 199, col: 14, offset: 5559},
						name: "_",
					},
				},
			},
		},
		{
			name: "WHITESPACE",
			pos:  position{line: 201, col: 1, offset: 5562},
			expr: &charClassMatcher{
				pos:        position{line: 201, col: 14, offset: 5575},
				val:        "[ \\t\\r]",
				chars:      []rune{' ', '\t', '\r'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EOL",
			pos:  position{line: 203, col: 1, offset: 5584},
			expr: &litMatcher{
				pos:        position{line: 203, col: 7, offset: 5590},
				val:        "\n",
				ignoreCase: false,
			},
		},
		{
			name: "EOF",
			pos:  position{line: 205, col: 1, offset: 5596},
			expr: &notExpr{
				pos: position{line: 205, col: 7, offset: 5602},
				expr: &anyMatcher{
					line: 205, col: 8, offset: 5603,
				},
			},
		},
	},
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
