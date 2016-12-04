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
			expr: &zeroOrMoreExpr{
				pos: position{line: 13, col: 14, offset: 240},
				expr: &ruleRefExpr{
					pos:  position{line: 13, col: 15, offset: 241},
					name: "NAMESPACE_STATEMENT",
				},
			},
		},
		{
			name: "NAMESPACE_STATEMENT",
			pos:  position{line: 15, col: 1, offset: 264},
			expr: &choiceExpr{
				pos: position{line: 15, col: 23, offset: 286},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 15, col: 23, offset: 286},
						name: "BLOCK_STATEMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 15, col: 40, offset: 303},
						name: "CREATE_OBJECT",
					},
					&ruleRefExpr{
						pos:  position{line: 15, col: 56, offset: 319},
						name: "CREATE_CLASS",
					},
				},
			},
		},
		{
			name: "CREATE_OBJECT",
			pos:  position{line: 17, col: 1, offset: 333},
			expr: &seqExpr{
				pos: position{line: 17, col: 17, offset: 349},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 17, col: 17, offset: 349},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 17, col: 19, offset: 351},
						expr: &ruleRefExpr{
							pos:  position{line: 17, col: 19, offset: 351},
							name: "CREATE_NAMESPACE_OBJECT",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 44, offset: 376},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 17, col: 46, offset: 378},
						expr: &ruleRefExpr{
							pos:  position{line: 17, col: 46, offset: 378},
							name: "NAMESPACE",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 57, offset: 389},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 17, col: 59, offset: 391},
						val:        ";",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "BLOCK_STATEMENT",
			pos:  position{line: 19, col: 1, offset: 396},
			expr: &seqExpr{
				pos: position{line: 19, col: 19, offset: 414},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 19, col: 19, offset: 414},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 21, offset: 416},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 31, offset: 426},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 19, col: 33, offset: 428},
						val:        ":",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 37, offset: 432},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 19, col: 39, offset: 434},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 43, offset: 438},
						name: "STATEMENTS",
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 54, offset: 449},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 19, col: 56, offset: 451},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CREATE_NAMESPACE_OBJECT",
			pos:  position{line: 21, col: 1, offset: 456},
			expr: &seqExpr{
				pos: position{line: 21, col: 27, offset: 482},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 21, col: 27, offset: 482},
						val:        "create",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 21, col: 37, offset: 492},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 21, col: 40, offset: 495},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 21, col: 40, offset: 495},
								val:        "database",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 21, col: 53, offset: 508},
								val:        "domain",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 21, col: 64, offset: 519},
								val:        "context",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 21, col: 76, offset: 531},
								val:        "aggregate",
								ignoreCase: false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 21, col: 89, offset: 544},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 21, col: 91, offset: 546},
						name: "QUOTED_NAME",
					},
				},
			},
		},
		{
			name: "NAMESPACE",
			pos:  position{line: 23, col: 1, offset: 559},
			expr: &zeroOrMoreExpr{
				pos: position{line: 23, col: 13, offset: 571},
				expr: &choiceExpr{
					pos: position{line: 23, col: 14, offset: 572},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 23, col: 14, offset: 572},
							name: "USING_DATABASE",
						},
						&ruleRefExpr{
							pos:  position{line: 23, col: 31, offset: 589},
							name: "FOR_DOMAIN",
						},
						&ruleRefExpr{
							pos:  position{line: 23, col: 44, offset: 602},
							name: "IN_CONTEXT",
						},
						&ruleRefExpr{
							pos:  position{line: 23, col: 57, offset: 615},
							name: "WITHIN_AGGREGATE",
						},
					},
				},
			},
		},
		{
			name: "USING_DATABASE",
			pos:  position{line: 25, col: 1, offset: 635},
			expr: &seqExpr{
				pos: position{line: 25, col: 18, offset: 652},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 25, col: 18, offset: 652},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 25, col: 20, offset: 654},
						val:        "using",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 29, offset: 663},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 25, col: 31, offset: 665},
						val:        "database",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 43, offset: 677},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 45, offset: 679},
						name: "QUOTED_NAME",
					},
				},
			},
		},
		{
			name: "FOR_DOMAIN",
			pos:  position{line: 27, col: 1, offset: 692},
			expr: &seqExpr{
				pos: position{line: 27, col: 14, offset: 705},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 27, col: 14, offset: 705},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 27, col: 16, offset: 707},
						val:        "for",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 23, offset: 714},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 27, col: 25, offset: 716},
						val:        "domain",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 35, offset: 726},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 37, offset: 728},
						name: "QUOTED_NAME",
					},
				},
			},
		},
		{
			name: "IN_CONTEXT",
			pos:  position{line: 29, col: 1, offset: 741},
			expr: &seqExpr{
				pos: position{line: 29, col: 14, offset: 754},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 29, col: 14, offset: 754},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 29, col: 16, offset: 756},
						val:        "in",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 29, col: 22, offset: 762},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 29, col: 24, offset: 764},
						val:        "context",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 29, col: 35, offset: 775},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 29, col: 37, offset: 777},
						name: "QUOTED_NAME",
					},
				},
			},
		},
		{
			name: "WITHIN_AGGREGATE",
			pos:  position{line: 31, col: 1, offset: 790},
			expr: &seqExpr{
				pos: position{line: 31, col: 20, offset: 809},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 31, col: 20, offset: 809},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 31, col: 22, offset: 811},
						val:        "within",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 31, col: 32, offset: 821},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 31, col: 34, offset: 823},
						val:        "aggregate",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 31, col: 47, offset: 836},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 31, col: 49, offset: 838},
						name: "QUOTED_NAME",
					},
				},
			},
		},
		{
			name: "CREATE_CLASS",
			pos:  position{line: 33, col: 1, offset: 851},
			expr: &choiceExpr{
				pos: position{line: 33, col: 16, offset: 866},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 33, col: 16, offset: 866},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 33, col: 16, offset: 866},
								name: "_",
							},
							&ruleRefExpr{
								pos:  position{line: 33, col: 19, offset: 869},
								name: "CREATE_VALUE",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 34, offset: 884},
						name: "CREATE_COMMAND",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 51, offset: 901},
						name: "CREATE_PROJECTION",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 71, offset: 921},
						name: "CREATE_INVARIANT",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 90, offset: 940},
						name: "CREATE_QUERY",
					},
				},
			},
		},
		{
			name: "CREATE_VALUE",
			pos:  position{line: 35, col: 1, offset: 954},
			expr: &seqExpr{
				pos: position{line: 35, col: 16, offset: 969},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 35, col: 16, offset: 969},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 35, col: 18, offset: 971},
						val:        "<|",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 23, offset: 976},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 35, col: 26, offset: 979},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 35, col: 26, offset: 979},
								val:        "value",
								ignoreCase: true,
							},
							&litMatcher{
								pos:        position{line: 35, col: 37, offset: 990},
								val:        "entity",
								ignoreCase: true,
							},
							&litMatcher{
								pos:        position{line: 35, col: 49, offset: 1002},
								val:        "event",
								ignoreCase: true,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 60, offset: 1013},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 62, offset: 1015},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 74, offset: 1027},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 76, offset: 1029},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 86, offset: 1039},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 88, offset: 1041},
						name: "VALUE_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 99, offset: 1052},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 35, col: 101, offset: 1054},
						val:        "|>",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CREATE_COMMAND",
			pos:  position{line: 37, col: 1, offset: 1060},
			expr: &seqExpr{
				pos: position{line: 37, col: 18, offset: 1077},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 37, col: 18, offset: 1077},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 37, col: 20, offset: 1079},
						val:        "<|",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 25, offset: 1084},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 37, col: 27, offset: 1086},
						val:        "command",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 39, offset: 1098},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 41, offset: 1100},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 53, offset: 1112},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 55, offset: 1114},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 65, offset: 1124},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 67, offset: 1126},
						name: "COMMAND_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 80, offset: 1139},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 37, col: 82, offset: 1141},
						val:        "|>",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CREATE_PROJECTION",
			pos:  position{line: 39, col: 1, offset: 1147},
			expr: &seqExpr{
				pos: position{line: 39, col: 21, offset: 1167},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 39, col: 21, offset: 1167},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 39, col: 23, offset: 1169},
						val:        "<|",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 28, offset: 1174},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 39, col: 31, offset: 1177},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 39, col: 31, offset: 1177},
								val:        "aggregate",
								ignoreCase: true,
							},
							&litMatcher{
								pos:        position{line: 39, col: 46, offset: 1192},
								val:        "domain",
								ignoreCase: true,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 57, offset: 1203},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 39, col: 59, offset: 1205},
						val:        "projection",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 74, offset: 1220},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 76, offset: 1222},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 88, offset: 1234},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 90, offset: 1236},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 100, offset: 1246},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 102, offset: 1248},
						name: "PROJECTION_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 118, offset: 1264},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 39, col: 120, offset: 1266},
						val:        "|>",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CREATE_INVARIANT",
			pos:  position{line: 41, col: 1, offset: 1272},
			expr: &seqExpr{
				pos: position{line: 41, col: 20, offset: 1291},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 41, col: 20, offset: 1291},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 41, col: 22, offset: 1293},
						val:        "<|",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 27, offset: 1298},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 41, col: 29, offset: 1300},
						val:        "invariant",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 43, offset: 1314},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 45, offset: 1316},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 57, offset: 1328},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 41, col: 59, offset: 1330},
						val:        "on",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 65, offset: 1336},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 67, offset: 1338},
						name: "CLASS_REF_QUOTES",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 84, offset: 1355},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 86, offset: 1357},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 96, offset: 1367},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 98, offset: 1369},
						name: "INVARIANT_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 113, offset: 1384},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 41, col: 115, offset: 1386},
						val:        "|>",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CREATE_QUERY",
			pos:  position{line: 43, col: 1, offset: 1392},
			expr: &seqExpr{
				pos: position{line: 43, col: 16, offset: 1407},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 43, col: 16, offset: 1407},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 43, col: 18, offset: 1409},
						val:        "<|",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 23, offset: 1414},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 43, col: 25, offset: 1416},
						val:        "query",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 35, offset: 1426},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 37, offset: 1428},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 49, offset: 1440},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 43, col: 51, offset: 1442},
						val:        "on",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 57, offset: 1448},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 59, offset: 1450},
						name: "CLASS_REF_QUOTES",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 76, offset: 1467},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 78, offset: 1469},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 88, offset: 1479},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 90, offset: 1481},
						name: "QUERY_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 101, offset: 1492},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 43, col: 103, offset: 1494},
						val:        "|>",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CLASS_COMPONENT_TEST",
			pos:  position{line: 49, col: 1, offset: 1676},
			expr: &seqExpr{
				pos: position{line: 49, col: 24, offset: 1699},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 49, col: 24, offset: 1699},
						expr: &choiceExpr{
							pos: position{line: 49, col: 25, offset: 1700},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 49, col: 25, offset: 1700},
									name: "WHEN",
								},
								&ruleRefExpr{
									pos:  position{line: 49, col: 32, offset: 1707},
									name: "COMMAND_HANDLER",
								},
								&ruleRefExpr{
									pos:  position{line: 49, col: 50, offset: 1725},
									name: "PROPERTIES",
								},
								&ruleRefExpr{
									pos:  position{line: 49, col: 63, offset: 1738},
									name: "CHECK",
								},
								&ruleRefExpr{
									pos:  position{line: 49, col: 71, offset: 1746},
									name: "FUNCTION",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 49, col: 82, offset: 1757},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "QUERY_BODY",
			pos:  position{line: 51, col: 1, offset: 1762},
			expr: &zeroOrMoreExpr{
				pos: position{line: 51, col: 14, offset: 1775},
				expr: &choiceExpr{
					pos: position{line: 51, col: 15, offset: 1776},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 51, col: 15, offset: 1776},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 51, col: 28, offset: 1789},
							name: "QUERY_HANDLER",
						},
					},
				},
			},
		},
		{
			name: "INVARIANT_BODY",
			pos:  position{line: 53, col: 1, offset: 1806},
			expr: &zeroOrMoreExpr{
				pos: position{line: 53, col: 18, offset: 1823},
				expr: &choiceExpr{
					pos: position{line: 53, col: 19, offset: 1824},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 53, col: 19, offset: 1824},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 53, col: 32, offset: 1837},
							name: "CHECK",
						},
					},
				},
			},
		},
		{
			name: "PROJECTION_BODY",
			pos:  position{line: 55, col: 1, offset: 1846},
			expr: &zeroOrMoreExpr{
				pos: position{line: 55, col: 19, offset: 1864},
				expr: &choiceExpr{
					pos: position{line: 55, col: 20, offset: 1865},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 55, col: 20, offset: 1865},
							name: "WHEN",
						},
						&ruleRefExpr{
							pos:  position{line: 55, col: 27, offset: 1872},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 55, col: 40, offset: 1885},
							name: "CHECK",
						},
						&ruleRefExpr{
							pos:  position{line: 55, col: 48, offset: 1893},
							name: "FUNCTION",
						},
					},
				},
			},
		},
		{
			name: "WHEN",
			pos:  position{line: 57, col: 1, offset: 1905},
			expr: &seqExpr{
				pos: position{line: 57, col: 8, offset: 1912},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 57, col: 8, offset: 1912},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 57, col: 10, offset: 1914},
						val:        "when",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 57, col: 18, offset: 1922},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 57, col: 20, offset: 1924},
						val:        "event",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 57, col: 29, offset: 1933},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 57, col: 31, offset: 1935},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 57, col: 43, offset: 1947},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 57, col: 45, offset: 1949},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 57, col: 49, offset: 1953},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 57, col: 51, offset: 1955},
						expr: &ruleRefExpr{
							pos:  position{line: 57, col: 51, offset: 1955},
							name: "STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 57, col: 68, offset: 1972},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 57, col: 70, offset: 1974},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 57, col: 74, offset: 1978},
						name: "_",
					},
				},
			},
		},
		{
			name: "COMMAND_BODY",
			pos:  position{line: 59, col: 1, offset: 1981},
			expr: &zeroOrMoreExpr{
				pos: position{line: 59, col: 16, offset: 1996},
				expr: &choiceExpr{
					pos: position{line: 59, col: 17, offset: 1997},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 59, col: 17, offset: 1997},
							name: "COMMAND_HANDLER",
						},
						&ruleRefExpr{
							pos:  position{line: 59, col: 35, offset: 2015},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 59, col: 48, offset: 2028},
							name: "CHECK",
						},
						&ruleRefExpr{
							pos:  position{line: 59, col: 56, offset: 2036},
							name: "FUNCTION",
						},
					},
				},
			},
		},
		{
			name: "COMMAND_HANDLER",
			pos:  position{line: 61, col: 1, offset: 2048},
			expr: &seqExpr{
				pos: position{line: 61, col: 19, offset: 2066},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 61, col: 19, offset: 2066},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 61, col: 21, offset: 2068},
						val:        "handler",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 61, col: 32, offset: 2079},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 61, col: 34, offset: 2081},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 61, col: 38, offset: 2085},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 61, col: 40, offset: 2087},
						expr: &ruleRefExpr{
							pos:  position{line: 61, col: 40, offset: 2087},
							name: "COMMAND_STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 61, col: 65, offset: 2112},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 61, col: 67, offset: 2114},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 61, col: 71, offset: 2118},
						name: "_",
					},
				},
			},
		},
		{
			name: "QUERY_HANDLER",
			pos:  position{line: 63, col: 1, offset: 2121},
			expr: &seqExpr{
				pos: position{line: 63, col: 17, offset: 2137},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 63, col: 17, offset: 2137},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 63, col: 19, offset: 2139},
						val:        "handler",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 63, col: 30, offset: 2150},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 63, col: 32, offset: 2152},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 63, col: 36, offset: 2156},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 63, col: 38, offset: 2158},
						expr: &ruleRefExpr{
							pos:  position{line: 63, col: 38, offset: 2158},
							name: "STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 63, col: 55, offset: 2175},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 63, col: 57, offset: 2177},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 63, col: 61, offset: 2181},
						name: "_",
					},
				},
			},
		},
		{
			name: "COMMAND_STATEMENT_BLOCK",
			pos:  position{line: 65, col: 1, offset: 2184},
			expr: &seqExpr{
				pos: position{line: 65, col: 27, offset: 2210},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 65, col: 27, offset: 2210},
						name: "_",
					},
					&oneOrMoreExpr{
						pos: position{line: 65, col: 29, offset: 2212},
						expr: &ruleRefExpr{
							pos:  position{line: 65, col: 30, offset: 2213},
							name: "COMMAND_STATEMENT",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 65, col: 50, offset: 2233},
						name: "_",
					},
				},
			},
		},
		{
			name: "COMMAND_STATEMENT",
			pos:  position{line: 67, col: 1, offset: 2236},
			expr: &choiceExpr{
				pos: position{line: 67, col: 21, offset: 2256},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 67, col: 21, offset: 2256},
						name: "STATEMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 67, col: 33, offset: 2268},
						name: "ASSERT",
					},
					&ruleRefExpr{
						pos:  position{line: 67, col: 42, offset: 2277},
						name: "APPLY",
					},
				},
			},
		},
		{
			name: "ASSERT",
			pos:  position{line: 69, col: 1, offset: 2284},
			expr: &seqExpr{
				pos: position{line: 69, col: 10, offset: 2293},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 69, col: 10, offset: 2293},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 69, col: 12, offset: 2295},
						val:        "assert",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 69, col: 22, offset: 2305},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 69, col: 24, offset: 2307},
						val:        "invariant",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 69, col: 37, offset: 2320},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 69, col: 39, offset: 2322},
						expr: &litMatcher{
							pos:        position{line: 69, col: 40, offset: 2323},
							val:        "not",
							ignoreCase: true,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 69, col: 49, offset: 2332},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 69, col: 51, offset: 2334},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 69, col: 63, offset: 2346},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 69, col: 65, offset: 2348},
						expr: &ruleRefExpr{
							pos:  position{line: 69, col: 65, offset: 2348},
							name: "ARGUMENTS",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 69, col: 76, offset: 2359},
						name: "SEMI",
					},
				},
			},
		},
		{
			name: "APPLY",
			pos:  position{line: 71, col: 1, offset: 2365},
			expr: &seqExpr{
				pos: position{line: 71, col: 9, offset: 2373},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 71, col: 9, offset: 2373},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 71, col: 11, offset: 2375},
						val:        "apply",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 71, col: 20, offset: 2384},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 71, col: 22, offset: 2386},
						val:        "event",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 71, col: 31, offset: 2395},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 71, col: 33, offset: 2397},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 71, col: 45, offset: 2409},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 71, col: 47, offset: 2411},
						expr: &ruleRefExpr{
							pos:  position{line: 71, col: 47, offset: 2411},
							name: "ARGUMENTS",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 71, col: 58, offset: 2422},
						name: "SEMI",
					},
				},
			},
		},
		{
			name: "VALUE_BODY",
			pos:  position{line: 73, col: 1, offset: 2428},
			expr: &zeroOrMoreExpr{
				pos: position{line: 73, col: 14, offset: 2441},
				expr: &ruleRefExpr{
					pos:  position{line: 73, col: 15, offset: 2442},
					name: "VALUE_COMPONENTS",
				},
			},
		},
		{
			name: "VALUE_COMPONENTS",
			pos:  position{line: 75, col: 1, offset: 2462},
			expr: &choiceExpr{
				pos: position{line: 75, col: 20, offset: 2481},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 75, col: 20, offset: 2481},
						name: "PROPERTIES",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 33, offset: 2494},
						name: "CHECK",
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 41, offset: 2502},
						name: "FUNCTION",
					},
				},
			},
		},
		{
			name: "PROPERTIES",
			pos:  position{line: 77, col: 1, offset: 2512},
			expr: &seqExpr{
				pos: position{line: 77, col: 14, offset: 2525},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 77, col: 14, offset: 2525},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 77, col: 16, offset: 2527},
						val:        "properties",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 77, col: 30, offset: 2541},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 77, col: 32, offset: 2543},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 77, col: 36, offset: 2547},
						name: "PROPERTY_LIST",
					},
					&litMatcher{
						pos:        position{line: 77, col: 50, offset: 2561},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 77, col: 54, offset: 2565},
						name: "_",
					},
				},
			},
		},
		{
			name: "PROPERTY_LIST",
			pos:  position{line: 79, col: 1, offset: 2568},
			expr: &zeroOrMoreExpr{
				pos: position{line: 79, col: 17, offset: 2584},
				expr: &ruleRefExpr{
					pos:  position{line: 79, col: 18, offset: 2585},
					name: "PROPERTY",
				},
			},
		},
		{
			name: "PROPERTY",
			pos:  position{line: 81, col: 1, offset: 2597},
			expr: &seqExpr{
				pos: position{line: 81, col: 12, offset: 2608},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 81, col: 12, offset: 2608},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 81, col: 14, offset: 2610},
						name: "TYPE",
					},
					&ruleRefExpr{
						pos:  position{line: 81, col: 19, offset: 2615},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 81, col: 21, offset: 2617},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 81, col: 32, offset: 2628},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 81, col: 35, offset: 2631},
						expr: &seqExpr{
							pos: position{line: 81, col: 36, offset: 2632},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 81, col: 36, offset: 2632},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 81, col: 40, offset: 2636},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 81, col: 42, offset: 2638},
									name: "EXPRESSION",
								},
								&ruleRefExpr{
									pos:  position{line: 81, col: 53, offset: 2649},
									name: "_",
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 81, col: 57, offset: 2653},
						val:        ";",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 81, col: 61, offset: 2657},
						name: "_",
					},
				},
			},
		},
		{
			name: "CHECK",
			pos:  position{line: 83, col: 1, offset: 2660},
			expr: &seqExpr{
				pos: position{line: 83, col: 9, offset: 2668},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 83, col: 9, offset: 2668},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 83, col: 11, offset: 2670},
						val:        "check",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 83, col: 20, offset: 2679},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 83, col: 22, offset: 2681},
						val:        "(",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 83, col: 26, offset: 2685},
						expr: &ruleRefExpr{
							pos:  position{line: 83, col: 26, offset: 2685},
							name: "STATEMENT_BLOCK",
						},
					},
					&litMatcher{
						pos:        position{line: 83, col: 43, offset: 2702},
						val:        ")",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 83, col: 47, offset: 2706},
						name: "_",
					},
				},
			},
		},
		{
			name: "FUNCTION",
			pos:  position{line: 85, col: 1, offset: 2709},
			expr: &seqExpr{
				pos: position{line: 85, col: 12, offset: 2720},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 85, col: 12, offset: 2720},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 85, col: 14, offset: 2722},
						val:        "function",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 85, col: 26, offset: 2734},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 85, col: 28, offset: 2736},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 85, col: 39, offset: 2747},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 85, col: 41, offset: 2749},
						name: "PARAMETERS",
					},
					&ruleRefExpr{
						pos:  position{line: 85, col: 53, offset: 2761},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 85, col: 55, offset: 2763},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 85, col: 59, offset: 2767},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 85, col: 61, offset: 2769},
						expr: &ruleRefExpr{
							pos:  position{line: 85, col: 61, offset: 2769},
							name: "STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 85, col: 78, offset: 2786},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 85, col: 80, offset: 2788},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 85, col: 84, offset: 2792},
						name: "_",
					},
				},
			},
		},
		{
			name: "PARAMETERS",
			pos:  position{line: 87, col: 1, offset: 2795},
			expr: &seqExpr{
				pos: position{line: 87, col: 14, offset: 2808},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 87, col: 14, offset: 2808},
						val:        "(",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 87, col: 18, offset: 2812},
						name: "PARAMETER_LIST",
					},
					&litMatcher{
						pos:        position{line: 87, col: 33, offset: 2827},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "PARAMETER_LIST",
			pos:  position{line: 89, col: 1, offset: 2832},
			expr: &seqExpr{
				pos: position{line: 89, col: 18, offset: 2849},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 89, col: 18, offset: 2849},
						name: "_",
					},
					&zeroOrMoreExpr{
						pos: position{line: 89, col: 20, offset: 2851},
						expr: &seqExpr{
							pos: position{line: 89, col: 21, offset: 2852},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 89, col: 21, offset: 2852},
									name: "PARAMETER",
								},
								&litMatcher{
									pos:        position{line: 89, col: 31, offset: 2862},
									val:        ",",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 89, col: 35, offset: 2866},
									name: "_",
								},
							},
						},
					},
					&zeroOrOneExpr{
						pos: position{line: 89, col: 40, offset: 2871},
						expr: &ruleRefExpr{
							pos:  position{line: 89, col: 40, offset: 2871},
							name: "PARAMETER",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 89, col: 51, offset: 2882},
						name: "_",
					},
				},
			},
		},
		{
			name: "PARAMETER",
			pos:  position{line: 91, col: 1, offset: 2885},
			expr: &seqExpr{
				pos: position{line: 91, col: 13, offset: 2897},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 91, col: 13, offset: 2897},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 91, col: 15, offset: 2899},
						name: "CLASS_REF",
					},
					&ruleRefExpr{
						pos:  position{line: 91, col: 25, offset: 2909},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 91, col: 27, offset: 2911},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 91, col: 38, offset: 2922},
						name: "_",
					},
				},
			},
		},
		{
			name: "STATEMENT_BLOCK",
			pos:  position{line: 96, col: 1, offset: 3094},
			expr: &seqExpr{
				pos: position{line: 96, col: 19, offset: 3112},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 96, col: 19, offset: 3112},
						name: "_",
					},
					&oneOrMoreExpr{
						pos: position{line: 96, col: 21, offset: 3114},
						expr: &ruleRefExpr{
							pos:  position{line: 96, col: 22, offset: 3115},
							name: "STATEMENT",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 34, offset: 3127},
						name: "_",
					},
				},
			},
		},
		{
			name: "STATEMENT",
			pos:  position{line: 98, col: 1, offset: 3130},
			expr: &choiceExpr{
				pos: position{line: 98, col: 13, offset: 3142},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 98, col: 13, offset: 3142},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 98, col: 13, offset: 3142},
								name: "RETURN",
							},
							&ruleRefExpr{
								pos:  position{line: 98, col: 20, offset: 3149},
								name: "SEMI",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 98, col: 27, offset: 3156},
						name: "IF",
					},
					&ruleRefExpr{
						pos:  position{line: 98, col: 32, offset: 3161},
						name: "FOREACH",
					},
					&seqExpr{
						pos: position{line: 98, col: 42, offset: 3171},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 98, col: 42, offset: 3171},
								name: "EXPRESSION",
							},
							&ruleRefExpr{
								pos:  position{line: 98, col: 53, offset: 3182},
								name: "SEMI",
							},
						},
					},
				},
			},
		},
		{
			name: "IF",
			pos:  position{line: 100, col: 1, offset: 3188},
			expr: &seqExpr{
				pos: position{line: 100, col: 6, offset: 3193},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 100, col: 6, offset: 3193},
						val:        "if",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 100, col: 11, offset: 3198},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 100, col: 13, offset: 3200},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 100, col: 24, offset: 3211},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 100, col: 26, offset: 3213},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 100, col: 30, offset: 3217},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 100, col: 32, offset: 3219},
						expr: &ruleRefExpr{
							pos:  position{line: 100, col: 32, offset: 3219},
							name: "STATEMENT_BLOCK",
						},
					},
					&litMatcher{
						pos:        position{line: 100, col: 49, offset: 3236},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 100, col: 53, offset: 3240},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 100, col: 55, offset: 3242},
						expr: &seqExpr{
							pos: position{line: 100, col: 56, offset: 3243},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 100, col: 56, offset: 3243},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 63, offset: 3250},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 100, col: 65, offset: 3252},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 69, offset: 3256},
									name: "_",
								},
								&zeroOrOneExpr{
									pos: position{line: 100, col: 71, offset: 3258},
									expr: &ruleRefExpr{
										pos:  position{line: 100, col: 71, offset: 3258},
										name: "STATEMENT_BLOCK",
									},
								},
								&litMatcher{
									pos:        position{line: 100, col: 88, offset: 3275},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 92, offset: 3279},
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
			pos:  position{line: 102, col: 1, offset: 3284},
			expr: &seqExpr{
				pos: position{line: 102, col: 11, offset: 3294},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 102, col: 11, offset: 3294},
						val:        "foreach",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 102, col: 21, offset: 3304},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 102, col: 23, offset: 3306},
						val:        "(",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 102, col: 27, offset: 3310},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 102, col: 29, offset: 3312},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 102, col: 40, offset: 3323},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 102, col: 42, offset: 3325},
						val:        "as",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 102, col: 47, offset: 3330},
						expr: &seqExpr{
							pos: position{line: 102, col: 48, offset: 3331},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 102, col: 48, offset: 3331},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 102, col: 50, offset: 3333},
									name: "IDENTIFIER",
								},
								&ruleRefExpr{
									pos:  position{line: 102, col: 61, offset: 3344},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 102, col: 63, offset: 3346},
									val:        "=>",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 102, col: 70, offset: 3353},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 102, col: 72, offset: 3355},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 102, col: 83, offset: 3366},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 102, col: 85, offset: 3368},
						val:        ")",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 102, col: 89, offset: 3372},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 102, col: 91, offset: 3374},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 102, col: 95, offset: 3378},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 102, col: 97, offset: 3380},
						expr: &ruleRefExpr{
							pos:  position{line: 102, col: 97, offset: 3380},
							name: "STATEMENT_BLOCK",
						},
					},
					&litMatcher{
						pos:        position{line: 102, col: 114, offset: 3397},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "RETURN",
			pos:  position{line: 104, col: 1, offset: 3402},
			expr: &seqExpr{
				pos: position{line: 104, col: 10, offset: 3411},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 104, col: 10, offset: 3411},
						val:        "return",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 104, col: 19, offset: 3420},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 104, col: 21, offset: 3422},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "EXRESSION_TEST",
			pos:  position{line: 110, col: 1, offset: 3605},
			expr: &seqExpr{
				pos: position{line: 110, col: 18, offset: 3622},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 110, col: 18, offset: 3622},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 110, col: 29, offset: 3633},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "EXPRESSION",
			pos:  position{line: 112, col: 1, offset: 3638},
			expr: &choiceExpr{
				pos: position{line: 112, col: 14, offset: 3651},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 112, col: 14, offset: 3651},
						name: "QUERY",
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 22, offset: 3659},
						name: "ARITHMETIC",
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 35, offset: 3672},
						name: "COMPARISON",
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 48, offset: 3685},
						name: "ASSIGNMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 60, offset: 3697},
						name: "LOGICAL",
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 70, offset: 3707},
						name: "ATOMIC",
					},
				},
			},
		},
		{
			name: "ATOMIC",
			pos:  position{line: 114, col: 1, offset: 3715},
			expr: &choiceExpr{
				pos: position{line: 114, col: 10, offset: 3724},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 114, col: 10, offset: 3724},
						name: "PARENTHESIS",
					},
					&ruleRefExpr{
						pos:  position{line: 114, col: 24, offset: 3738},
						name: "NEW",
					},
					&ruleRefExpr{
						pos:  position{line: 114, col: 30, offset: 3744},
						name: "METHODCALL",
					},
					&ruleRefExpr{
						pos:  position{line: 114, col: 43, offset: 3757},
						name: "OBJECTACCESS",
					},
					&ruleRefExpr{
						pos:  position{line: 114, col: 58, offset: 3772},
						name: "ARRAY",
					},
					&ruleRefExpr{
						pos:  position{line: 114, col: 66, offset: 3780},
						name: "LITERAL",
					},
					&ruleRefExpr{
						pos:  position{line: 114, col: 76, offset: 3790},
						name: "UNARY",
					},
				},
			},
		},
		{
			name: "LITERAL",
			pos:  position{line: 116, col: 1, offset: 3797},
			expr: &choiceExpr{
				pos: position{line: 116, col: 11, offset: 3807},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 116, col: 11, offset: 3807},
						name: "STRING",
					},
					&ruleRefExpr{
						pos:  position{line: 116, col: 20, offset: 3816},
						name: "FLOAT",
					},
					&ruleRefExpr{
						pos:  position{line: 116, col: 28, offset: 3824},
						name: "BOOLEAN",
					},
					&ruleRefExpr{
						pos:  position{line: 116, col: 38, offset: 3834},
						name: "NULL",
					},
					&ruleRefExpr{
						pos:  position{line: 116, col: 45, offset: 3841},
						name: "INT",
					},
				},
			},
		},
		{
			name: "NEW",
			pos:  position{line: 118, col: 1, offset: 3846},
			expr: &seqExpr{
				pos: position{line: 118, col: 7, offset: 3852},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 118, col: 7, offset: 3852},
						name: "CLASS_REF_QUOTES",
					},
					&ruleRefExpr{
						pos:  position{line: 118, col: 24, offset: 3869},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 118, col: 26, offset: 3871},
						expr: &ruleRefExpr{
							pos:  position{line: 118, col: 26, offset: 3871},
							name: "ARGUMENTS",
						},
					},
				},
			},
		},
		{
			name: "BOOLEAN",
			pos:  position{line: 120, col: 1, offset: 3883},
			expr: &choiceExpr{
				pos: position{line: 120, col: 12, offset: 3894},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 120, col: 12, offset: 3894},
						val:        "true",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 120, col: 19, offset: 3901},
						val:        "false",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "NULL",
			pos:  position{line: 122, col: 1, offset: 3910},
			expr: &litMatcher{
				pos:        position{line: 122, col: 8, offset: 3917},
				val:        "null",
				ignoreCase: false,
			},
		},
		{
			name: "ARRAY",
			pos:  position{line: 124, col: 1, offset: 3925},
			expr: &seqExpr{
				pos: position{line: 124, col: 9, offset: 3933},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 124, col: 9, offset: 3933},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 124, col: 11, offset: 3935},
						val:        "[",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 124, col: 15, offset: 3939},
						expr: &ruleRefExpr{
							pos:  position{line: 124, col: 15, offset: 3939},
							name: "ARGUMENTLIST",
						},
					},
					&litMatcher{
						pos:        position{line: 124, col: 29, offset: 3953},
						val:        "]",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 124, col: 33, offset: 3957},
						name: "_",
					},
				},
			},
		},
		{
			name: "STRING",
			pos:  position{line: 126, col: 1, offset: 3960},
			expr: &seqExpr{
				pos: position{line: 126, col: 10, offset: 3969},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 126, col: 10, offset: 3969},
						val:        "\"",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 126, col: 15, offset: 3974},
						expr: &charClassMatcher{
							pos:        position{line: 126, col: 15, offset: 3974},
							val:        "[a-zA-Z0-9]",
							ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&litMatcher{
						pos:        position{line: 126, col: 28, offset: 3987},
						val:        "\"",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "INT",
			pos:  position{line: 128, col: 1, offset: 3993},
			expr: &oneOrMoreExpr{
				pos: position{line: 128, col: 7, offset: 3999},
				expr: &charClassMatcher{
					pos:        position{line: 128, col: 7, offset: 3999},
					val:        "[0-9]",
					ranges:     []rune{'0', '9'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "FLOAT",
			pos:  position{line: 130, col: 1, offset: 4007},
			expr: &seqExpr{
				pos: position{line: 130, col: 9, offset: 4015},
				exprs: []interface{}{
					&oneOrMoreExpr{
						pos: position{line: 130, col: 9, offset: 4015},
						expr: &charClassMatcher{
							pos:        position{line: 130, col: 9, offset: 4015},
							val:        "[0-9]",
							ranges:     []rune{'0', '9'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&charClassMatcher{
						pos:        position{line: 130, col: 16, offset: 4022},
						val:        "[.]",
						chars:      []rune{'.'},
						ignoreCase: false,
						inverted:   false,
					},
					&oneOrMoreExpr{
						pos: position{line: 130, col: 20, offset: 4026},
						expr: &charClassMatcher{
							pos:        position{line: 130, col: 20, offset: 4026},
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
			pos:  position{line: 132, col: 1, offset: 4034},
			expr: &seqExpr{
				pos: position{line: 132, col: 15, offset: 4048},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 132, col: 15, offset: 4048},
						val:        "(",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 132, col: 19, offset: 4052},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 132, col: 21, offset: 4054},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 132, col: 32, offset: 4065},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 132, col: 34, offset: 4067},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "UNARY",
			pos:  position{line: 134, col: 1, offset: 4072},
			expr: &choiceExpr{
				pos: position{line: 134, col: 9, offset: 4080},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 134, col: 9, offset: 4080},
						name: "INCREMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 134, col: 21, offset: 4092},
						name: "DECREMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 134, col: 33, offset: 4104},
						name: "NEGATE",
					},
					&ruleRefExpr{
						pos:  position{line: 134, col: 42, offset: 4113},
						name: "NOT",
					},
					&ruleRefExpr{
						pos:  position{line: 134, col: 48, offset: 4119},
						name: "POSITIVE",
					},
				},
			},
		},
		{
			name: "INCREMENT",
			pos:  position{line: 136, col: 1, offset: 4129},
			expr: &seqExpr{
				pos: position{line: 136, col: 13, offset: 4141},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 136, col: 13, offset: 4141},
						name: "OBJECTACCESS",
					},
					&litMatcher{
						pos:        position{line: 136, col: 26, offset: 4154},
						val:        "++",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DECREMENT",
			pos:  position{line: 138, col: 1, offset: 4160},
			expr: &seqExpr{
				pos: position{line: 138, col: 13, offset: 4172},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 138, col: 13, offset: 4172},
						name: "OBJECTACCESS",
					},
					&litMatcher{
						pos:        position{line: 138, col: 26, offset: 4185},
						val:        "--",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "NEGATE",
			pos:  position{line: 140, col: 1, offset: 4191},
			expr: &seqExpr{
				pos: position{line: 140, col: 10, offset: 4200},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 140, col: 10, offset: 4200},
						val:        "-",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 140, col: 14, offset: 4204},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "NOT",
			pos:  position{line: 142, col: 1, offset: 4218},
			expr: &seqExpr{
				pos: position{line: 142, col: 7, offset: 4224},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 142, col: 7, offset: 4224},
						val:        "!",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 142, col: 11, offset: 4228},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "POSITIVE",
			pos:  position{line: 144, col: 1, offset: 4242},
			expr: &seqExpr{
				pos: position{line: 144, col: 12, offset: 4253},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 144, col: 12, offset: 4253},
						val:        "+",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 144, col: 16, offset: 4257},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "ARITHMETIC",
			pos:  position{line: 146, col: 1, offset: 4271},
			expr: &seqExpr{
				pos: position{line: 146, col: 14, offset: 4284},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 146, col: 14, offset: 4284},
						name: "ATOMIC",
					},
					&ruleRefExpr{
						pos:  position{line: 146, col: 21, offset: 4291},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 146, col: 23, offset: 4293},
						name: "OPERATOR",
					},
					&ruleRefExpr{
						pos:  position{line: 146, col: 32, offset: 4302},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 146, col: 34, offset: 4304},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "OPERATOR",
			pos:  position{line: 148, col: 1, offset: 4316},
			expr: &choiceExpr{
				pos: position{line: 148, col: 12, offset: 4327},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 148, col: 12, offset: 4327},
						val:        "+",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 148, col: 18, offset: 4333},
						val:        "-",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 148, col: 24, offset: 4339},
						val:        "/",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 148, col: 30, offset: 4345},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 148, col: 36, offset: 4351},
						val:        "%",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "ASSIGNMENT",
			pos:  position{line: 150, col: 1, offset: 4356},
			expr: &seqExpr{
				pos: position{line: 150, col: 14, offset: 4369},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 150, col: 14, offset: 4369},
						name: "OBJECTACCESS",
					},
					&ruleRefExpr{
						pos:  position{line: 150, col: 27, offset: 4382},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 150, col: 29, offset: 4384},
						val:        "=",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 150, col: 33, offset: 4388},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 150, col: 35, offset: 4390},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "LOGICAL",
			pos:  position{line: 152, col: 1, offset: 4402},
			expr: &seqExpr{
				pos: position{line: 152, col: 11, offset: 4412},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 152, col: 11, offset: 4412},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 152, col: 22, offset: 4423},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 152, col: 25, offset: 4426},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 152, col: 25, offset: 4426},
								val:        "and",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 152, col: 33, offset: 4434},
								val:        "or",
								ignoreCase: false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 152, col: 39, offset: 4440},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 152, col: 41, offset: 4442},
						name: "ATOMIC",
					},
				},
			},
		},
		{
			name: "COMPARISON",
			pos:  position{line: 154, col: 1, offset: 4450},
			expr: &seqExpr{
				pos: position{line: 154, col: 14, offset: 4463},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 154, col: 14, offset: 4463},
						name: "ATOMIC",
					},
					&ruleRefExpr{
						pos:  position{line: 154, col: 21, offset: 4470},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 154, col: 24, offset: 4473},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 154, col: 24, offset: 4473},
								val:        "===",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 154, col: 32, offset: 4481},
								val:        "!==",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 154, col: 40, offset: 4489},
								val:        "==",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 154, col: 47, offset: 4496},
								val:        "!=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 154, col: 54, offset: 4503},
								val:        "<=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 154, col: 61, offset: 4510},
								val:        ">=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 154, col: 68, offset: 4517},
								val:        "<",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 154, col: 74, offset: 4523},
								val:        ">",
								ignoreCase: false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 154, col: 79, offset: 4528},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 154, col: 81, offset: 4530},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "QUERY",
			pos:  position{line: 156, col: 1, offset: 4542},
			expr: &seqExpr{
				pos: position{line: 156, col: 9, offset: 4550},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 156, col: 9, offset: 4550},
						val:        "run",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 156, col: 16, offset: 4557},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 156, col: 18, offset: 4559},
						val:        "query",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 156, col: 27, offset: 4568},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 156, col: 29, offset: 4570},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 156, col: 41, offset: 4582},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 156, col: 43, offset: 4584},
						expr: &ruleRefExpr{
							pos:  position{line: 156, col: 43, offset: 4584},
							name: "ARGUMENTS",
						},
					},
				},
			},
		},
		{
			name: "OBJECTACCESS",
			pos:  position{line: 158, col: 1, offset: 4596},
			expr: &seqExpr{
				pos: position{line: 158, col: 16, offset: 4611},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 158, col: 16, offset: 4611},
						expr: &seqExpr{
							pos: position{line: 158, col: 17, offset: 4612},
							exprs: []interface{}{
								&choiceExpr{
									pos: position{line: 158, col: 18, offset: 4613},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 158, col: 18, offset: 4613},
											name: "METHODCALL",
										},
										&ruleRefExpr{
											pos:  position{line: 158, col: 31, offset: 4626},
											name: "IDENTIFIER",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 158, col: 43, offset: 4638},
									val:        "->",
									ignoreCase: false,
								},
							},
						},
					},
					&choiceExpr{
						pos: position{line: 158, col: 51, offset: 4646},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 158, col: 51, offset: 4646},
								name: "METHODCALL",
							},
							&ruleRefExpr{
								pos:  position{line: 158, col: 64, offset: 4659},
								name: "IDENTIFIER",
							},
						},
					},
				},
			},
		},
		{
			name: "METHODCALL",
			pos:  position{line: 160, col: 1, offset: 4672},
			expr: &seqExpr{
				pos: position{line: 160, col: 14, offset: 4685},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 160, col: 14, offset: 4685},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 160, col: 25, offset: 4696},
						name: "ARGUMENTS",
					},
				},
			},
		},
		{
			name: "ARGUMENTS",
			pos:  position{line: 162, col: 1, offset: 4707},
			expr: &seqExpr{
				pos: position{line: 162, col: 13, offset: 4719},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 162, col: 13, offset: 4719},
						val:        "(",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 162, col: 17, offset: 4723},
						expr: &ruleRefExpr{
							pos:  position{line: 162, col: 17, offset: 4723},
							name: "ARGUMENTLIST",
						},
					},
					&litMatcher{
						pos:        position{line: 162, col: 31, offset: 4737},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "ARGUMENTLIST",
			pos:  position{line: 164, col: 1, offset: 4742},
			expr: &seqExpr{
				pos: position{line: 164, col: 17, offset: 4758},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 164, col: 17, offset: 4758},
						name: "_",
					},
					&zeroOrMoreExpr{
						pos: position{line: 164, col: 19, offset: 4760},
						expr: &seqExpr{
							pos: position{line: 164, col: 20, offset: 4761},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 164, col: 20, offset: 4761},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 164, col: 22, offset: 4763},
									name: "EXPRESSION",
								},
								&ruleRefExpr{
									pos:  position{line: 164, col: 33, offset: 4774},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 164, col: 35, offset: 4776},
									val:        ",",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 164, col: 39, offset: 4780},
									name: "_",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 164, col: 43, offset: 4784},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 164, col: 54, offset: 4795},
						name: "_",
					},
				},
			},
		},
		{
			name: "CLASS_REF_QUOTES",
			pos:  position{line: 171, col: 1, offset: 4963},
			expr: &seqExpr{
				pos: position{line: 171, col: 20, offset: 4982},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 171, col: 20, offset: 4982},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 171, col: 22, offset: 4984},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 171, col: 26, offset: 4988},
						name: "CLASS_REF",
					},
					&litMatcher{
						pos:        position{line: 171, col: 36, offset: 4998},
						val:        "'",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CLASS_REF",
			pos:  position{line: 173, col: 1, offset: 5003},
			expr: &seqExpr{
				pos: position{line: 173, col: 13, offset: 5015},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 173, col: 13, offset: 5015},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 173, col: 15, offset: 5017},
						name: "CLASS_TYPE",
					},
					&litMatcher{
						pos:        position{line: 173, col: 26, offset: 5028},
						val:        "\\",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 173, col: 31, offset: 5033},
						name: "CLASS_NAME",
					},
				},
			},
		},
		{
			name: "CLASS_TYPE",
			pos:  position{line: 175, col: 1, offset: 5045},
			expr: &seqExpr{
				pos: position{line: 175, col: 14, offset: 5058},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 175, col: 14, offset: 5058},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 175, col: 17, offset: 5061},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 175, col: 17, offset: 5061},
								val:        "value",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 175, col: 27, offset: 5071},
								val:        "entity",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 175, col: 38, offset: 5082},
								val:        "command",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 175, col: 50, offset: 5094},
								val:        "event",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 175, col: 60, offset: 5104},
								val:        "projection",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 175, col: 75, offset: 5119},
								val:        "invariant",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 175, col: 89, offset: 5133},
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
			pos:  position{line: 177, col: 1, offset: 5143},
			expr: &seqExpr{
				pos: position{line: 177, col: 14, offset: 5156},
				exprs: []interface{}{
					&oneOrMoreExpr{
						pos: position{line: 177, col: 14, offset: 5156},
						expr: &charClassMatcher{
							pos:        position{line: 177, col: 14, offset: 5156},
							val:        "[a-zA-Z]",
							ranges:     []rune{'a', 'z', 'A', 'Z'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 177, col: 24, offset: 5166},
						expr: &charClassMatcher{
							pos:        position{line: 177, col: 24, offset: 5166},
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
			pos:  position{line: 179, col: 1, offset: 5182},
			expr: &seqExpr{
				pos: position{line: 179, col: 15, offset: 5196},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 179, col: 15, offset: 5196},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 179, col: 19, offset: 5200},
						name: "CLASS_NAME",
					},
					&litMatcher{
						pos:        position{line: 179, col: 30, offset: 5211},
						val:        "'",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "TYPE",
			pos:  position{line: 181, col: 1, offset: 5216},
			expr: &choiceExpr{
				pos: position{line: 181, col: 8, offset: 5223},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 181, col: 8, offset: 5223},
						name: "CLASS_REF",
					},
					&ruleRefExpr{
						pos:  position{line: 181, col: 20, offset: 5235},
						name: "VALUE_TYPE",
					},
				},
			},
		},
		{
			name: "VALUE_TYPE",
			pos:  position{line: 183, col: 1, offset: 5247},
			expr: &seqExpr{
				pos: position{line: 183, col: 14, offset: 5260},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 183, col: 14, offset: 5260},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 183, col: 17, offset: 5263},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 183, col: 17, offset: 5263},
								val:        "string",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 183, col: 28, offset: 5274},
								val:        "boolean",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 183, col: 40, offset: 5286},
								val:        "float",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 183, col: 50, offset: 5296},
								val:        "map",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 183, col: 58, offset: 5304},
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
			pos:  position{line: 185, col: 1, offset: 5314},
			expr: &seqExpr{
				pos: position{line: 185, col: 21, offset: 5334},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 185, col: 21, offset: 5334},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 185, col: 23, offset: 5336},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 185, col: 27, offset: 5340},
						name: "CLASS_NAME",
					},
					&litMatcher{
						pos:        position{line: 185, col: 38, offset: 5351},
						val:        "'",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "IDENTIFIER",
			pos:  position{line: 187, col: 1, offset: 5356},
			expr: &seqExpr{
				pos: position{line: 187, col: 14, offset: 5369},
				exprs: []interface{}{
					&oneOrMoreExpr{
						pos: position{line: 187, col: 14, offset: 5369},
						expr: &charClassMatcher{
							pos:        position{line: 187, col: 14, offset: 5369},
							val:        "[a-zA-Z]",
							ranges:     []rune{'a', 'z', 'A', 'Z'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 187, col: 24, offset: 5379},
						expr: &charClassMatcher{
							pos:        position{line: 187, col: 24, offset: 5379},
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
			pos:  position{line: 189, col: 1, offset: 5394},
			expr: &zeroOrMoreExpr{
				pos: position{line: 189, col: 5, offset: 5398},
				expr: &choiceExpr{
					pos: position{line: 189, col: 7, offset: 5400},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 189, col: 7, offset: 5400},
							name: "WHITESPACE",
						},
						&ruleRefExpr{
							pos:  position{line: 189, col: 20, offset: 5413},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "SEMI",
			pos:  position{line: 191, col: 1, offset: 5421},
			expr: &seqExpr{
				pos: position{line: 191, col: 8, offset: 5428},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 191, col: 8, offset: 5428},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 191, col: 10, offset: 5430},
						val:        ";",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 191, col: 14, offset: 5434},
						name: "_",
					},
				},
			},
		},
		{
			name: "WHITESPACE",
			pos:  position{line: 193, col: 1, offset: 5437},
			expr: &charClassMatcher{
				pos:        position{line: 193, col: 14, offset: 5450},
				val:        "[ \\t\\r]",
				chars:      []rune{' ', '\t', '\r'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EOL",
			pos:  position{line: 195, col: 1, offset: 5459},
			expr: &litMatcher{
				pos:        position{line: 195, col: 7, offset: 5465},
				val:        "\n",
				ignoreCase: false,
			},
		},
		{
			name: "EOF",
			pos:  position{line: 197, col: 1, offset: 5471},
			expr: &notExpr{
				pos: position{line: 197, col: 7, offset: 5477},
				expr: &anyMatcher{
					line: 197, col: 8, offset: 5478,
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
