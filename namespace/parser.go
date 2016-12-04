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
			name: "STATEMENTS",
			pos:  position{line: 11, col: 1, offset: 202},
			expr: &zeroOrOneExpr{
				pos: position{line: 11, col: 14, offset: 215},
				expr: &ruleRefExpr{
					pos:  position{line: 11, col: 15, offset: 216},
					name: "NAMESPACE_STATEMENT",
				},
			},
		},
		{
			name: "NAMESPACE_STATEMENT",
			pos:  position{line: 13, col: 1, offset: 239},
			expr: &choiceExpr{
				pos: position{line: 13, col: 23, offset: 261},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 13, col: 23, offset: 261},
						name: "BLOCK_STATEMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 13, col: 40, offset: 278},
						name: "CREATE_OBJECT",
					},
					&ruleRefExpr{
						pos:  position{line: 13, col: 56, offset: 294},
						name: "CREATE_CLASS",
					},
				},
			},
		},
		{
			name: "CREATE_OBJECT",
			pos:  position{line: 15, col: 1, offset: 308},
			expr: &seqExpr{
				pos: position{line: 15, col: 17, offset: 324},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 15, col: 17, offset: 324},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 15, col: 19, offset: 326},
						expr: &ruleRefExpr{
							pos:  position{line: 15, col: 19, offset: 326},
							name: "CREATE_NAMESPACE_OBJECT",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 15, col: 44, offset: 351},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 15, col: 46, offset: 353},
						expr: &ruleRefExpr{
							pos:  position{line: 15, col: 46, offset: 353},
							name: "NAMESPACE",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 15, col: 57, offset: 364},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 15, col: 59, offset: 366},
						val:        ";",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "BLOCK_STATEMENT",
			pos:  position{line: 17, col: 1, offset: 371},
			expr: &seqExpr{
				pos: position{line: 17, col: 19, offset: 389},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 17, col: 19, offset: 389},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 21, offset: 391},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 31, offset: 401},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 17, col: 33, offset: 403},
						val:        ":",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 37, offset: 407},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 17, col: 39, offset: 409},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 43, offset: 413},
						name: "NAMESPACE_STATEMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 17, col: 63, offset: 433},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 17, col: 65, offset: 435},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CREATE_NAMESPACE_OBJECT",
			pos:  position{line: 19, col: 1, offset: 440},
			expr: &seqExpr{
				pos: position{line: 19, col: 27, offset: 466},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 19, col: 27, offset: 466},
						val:        "create",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 37, offset: 476},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 19, col: 40, offset: 479},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 19, col: 40, offset: 479},
								val:        "database",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 19, col: 53, offset: 492},
								val:        "domain",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 19, col: 64, offset: 503},
								val:        "context",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 19, col: 76, offset: 515},
								val:        "aggregate",
								ignoreCase: false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 89, offset: 528},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 91, offset: 530},
						name: "QUOTED_NAME",
					},
				},
			},
		},
		{
			name: "NAMESPACE",
			pos:  position{line: 21, col: 1, offset: 543},
			expr: &zeroOrMoreExpr{
				pos: position{line: 21, col: 13, offset: 555},
				expr: &choiceExpr{
					pos: position{line: 21, col: 14, offset: 556},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 21, col: 14, offset: 556},
							name: "USING_DATABASE",
						},
						&ruleRefExpr{
							pos:  position{line: 21, col: 31, offset: 573},
							name: "FOR_DOMAIN",
						},
						&ruleRefExpr{
							pos:  position{line: 21, col: 44, offset: 586},
							name: "IN_CONTEXT",
						},
						&ruleRefExpr{
							pos:  position{line: 21, col: 57, offset: 599},
							name: "WITHIN_AGGREGATE",
						},
					},
				},
			},
		},
		{
			name: "USING_DATABASE",
			pos:  position{line: 23, col: 1, offset: 619},
			expr: &seqExpr{
				pos: position{line: 23, col: 18, offset: 636},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 23, col: 18, offset: 636},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 23, col: 20, offset: 638},
						val:        "using",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 29, offset: 647},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 23, col: 31, offset: 649},
						val:        "database",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 43, offset: 661},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 45, offset: 663},
						name: "QUOTED_NAME",
					},
				},
			},
		},
		{
			name: "FOR_DOMAIN",
			pos:  position{line: 25, col: 1, offset: 676},
			expr: &seqExpr{
				pos: position{line: 25, col: 14, offset: 689},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 25, col: 14, offset: 689},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 25, col: 16, offset: 691},
						val:        "for",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 23, offset: 698},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 25, col: 25, offset: 700},
						val:        "domain",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 35, offset: 710},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 37, offset: 712},
						name: "QUOTED_NAME",
					},
				},
			},
		},
		{
			name: "IN_CONTEXT",
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
						val:        "in",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 22, offset: 746},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 27, col: 24, offset: 748},
						val:        "context",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 35, offset: 759},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 37, offset: 761},
						name: "QUOTED_NAME",
					},
				},
			},
		},
		{
			name: "WITHIN_AGGREGATE",
			pos:  position{line: 29, col: 1, offset: 774},
			expr: &seqExpr{
				pos: position{line: 29, col: 20, offset: 793},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 29, col: 20, offset: 793},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 29, col: 22, offset: 795},
						val:        "within",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 29, col: 32, offset: 805},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 29, col: 34, offset: 807},
						val:        "aggregate",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 29, col: 47, offset: 820},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 29, col: 49, offset: 822},
						name: "QUOTED_NAME",
					},
				},
			},
		},
		{
			name: "CREATE_CLASS",
			pos:  position{line: 31, col: 1, offset: 835},
			expr: &choiceExpr{
				pos: position{line: 31, col: 16, offset: 850},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 31, col: 16, offset: 850},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 31, col: 16, offset: 850},
								name: "_",
							},
							&ruleRefExpr{
								pos:  position{line: 31, col: 19, offset: 853},
								name: "CREATE_VALUE",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 31, col: 34, offset: 868},
						name: "CREATE_COMMAND",
					},
					&ruleRefExpr{
						pos:  position{line: 31, col: 51, offset: 885},
						name: "CREATE_PROJECTION",
					},
					&ruleRefExpr{
						pos:  position{line: 31, col: 71, offset: 905},
						name: "CREATE_INVARIANT",
					},
					&ruleRefExpr{
						pos:  position{line: 31, col: 90, offset: 924},
						name: "CREATE_QUERY",
					},
				},
			},
		},
		{
			name: "CREATE_VALUE",
			pos:  position{line: 33, col: 1, offset: 938},
			expr: &seqExpr{
				pos: position{line: 33, col: 16, offset: 953},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 33, col: 16, offset: 953},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 33, col: 18, offset: 955},
						val:        "<|",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 23, offset: 960},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 33, col: 26, offset: 963},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 33, col: 26, offset: 963},
								val:        "value",
								ignoreCase: true,
							},
							&litMatcher{
								pos:        position{line: 33, col: 37, offset: 974},
								val:        "entity",
								ignoreCase: true,
							},
							&litMatcher{
								pos:        position{line: 33, col: 49, offset: 986},
								val:        "event",
								ignoreCase: true,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 60, offset: 997},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 62, offset: 999},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 74, offset: 1011},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 76, offset: 1013},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 86, offset: 1023},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 88, offset: 1025},
						name: "VALUE_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 99, offset: 1036},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 33, col: 101, offset: 1038},
						val:        "|>",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CREATE_COMMAND",
			pos:  position{line: 35, col: 1, offset: 1044},
			expr: &seqExpr{
				pos: position{line: 35, col: 18, offset: 1061},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 35, col: 18, offset: 1061},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 35, col: 20, offset: 1063},
						val:        "<|",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 25, offset: 1068},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 35, col: 27, offset: 1070},
						val:        "command",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 39, offset: 1082},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 41, offset: 1084},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 53, offset: 1096},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 55, offset: 1098},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 65, offset: 1108},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 67, offset: 1110},
						name: "COMMAND_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 35, col: 80, offset: 1123},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 35, col: 82, offset: 1125},
						val:        "|>",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CREATE_PROJECTION",
			pos:  position{line: 37, col: 1, offset: 1131},
			expr: &seqExpr{
				pos: position{line: 37, col: 21, offset: 1151},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 37, col: 21, offset: 1151},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 37, col: 23, offset: 1153},
						val:        "<|",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 28, offset: 1158},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 37, col: 31, offset: 1161},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 37, col: 31, offset: 1161},
								val:        "aggregate",
								ignoreCase: true,
							},
							&litMatcher{
								pos:        position{line: 37, col: 46, offset: 1176},
								val:        "domain",
								ignoreCase: true,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 57, offset: 1187},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 37, col: 59, offset: 1189},
						val:        "projection",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 74, offset: 1204},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 76, offset: 1206},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 88, offset: 1218},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 90, offset: 1220},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 100, offset: 1230},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 102, offset: 1232},
						name: "PROJECTION_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 118, offset: 1248},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 37, col: 120, offset: 1250},
						val:        "|>",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CREATE_INVARIANT",
			pos:  position{line: 39, col: 1, offset: 1256},
			expr: &seqExpr{
				pos: position{line: 39, col: 20, offset: 1275},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 39, col: 20, offset: 1275},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 39, col: 22, offset: 1277},
						val:        "<|",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 27, offset: 1282},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 39, col: 29, offset: 1284},
						val:        "invariant",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 43, offset: 1298},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 45, offset: 1300},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 57, offset: 1312},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 39, col: 59, offset: 1314},
						val:        "on",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 65, offset: 1320},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 67, offset: 1322},
						name: "CLASS_REF_QUOTES",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 84, offset: 1339},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 86, offset: 1341},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 96, offset: 1351},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 98, offset: 1353},
						name: "INVARIANT_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 113, offset: 1368},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 39, col: 115, offset: 1370},
						val:        "|>",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CREATE_QUERY",
			pos:  position{line: 41, col: 1, offset: 1376},
			expr: &seqExpr{
				pos: position{line: 41, col: 16, offset: 1391},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 41, col: 16, offset: 1391},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 41, col: 18, offset: 1393},
						val:        "<|",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 23, offset: 1398},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 41, col: 25, offset: 1400},
						val:        "query",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 35, offset: 1410},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 37, offset: 1412},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 49, offset: 1424},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 41, col: 51, offset: 1426},
						val:        "on",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 57, offset: 1432},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 59, offset: 1434},
						name: "CLASS_REF_QUOTES",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 76, offset: 1451},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 78, offset: 1453},
						name: "NAMESPACE",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 88, offset: 1463},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 90, offset: 1465},
						name: "QUERY_BODY",
					},
					&ruleRefExpr{
						pos:  position{line: 41, col: 101, offset: 1476},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 41, col: 103, offset: 1478},
						val:        "|>",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CLASS_COMPONENT_TEST",
			pos:  position{line: 47, col: 1, offset: 1660},
			expr: &seqExpr{
				pos: position{line: 47, col: 24, offset: 1683},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 47, col: 24, offset: 1683},
						expr: &choiceExpr{
							pos: position{line: 47, col: 25, offset: 1684},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 47, col: 25, offset: 1684},
									name: "WHEN",
								},
								&ruleRefExpr{
									pos:  position{line: 47, col: 32, offset: 1691},
									name: "COMMAND_HANDLER",
								},
								&ruleRefExpr{
									pos:  position{line: 47, col: 50, offset: 1709},
									name: "PROPERTIES",
								},
								&ruleRefExpr{
									pos:  position{line: 47, col: 63, offset: 1722},
									name: "CHECK",
								},
								&ruleRefExpr{
									pos:  position{line: 47, col: 71, offset: 1730},
									name: "FUNCTION",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 47, col: 82, offset: 1741},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "QUERY_BODY",
			pos:  position{line: 49, col: 1, offset: 1746},
			expr: &zeroOrMoreExpr{
				pos: position{line: 49, col: 14, offset: 1759},
				expr: &choiceExpr{
					pos: position{line: 49, col: 15, offset: 1760},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 49, col: 15, offset: 1760},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 49, col: 28, offset: 1773},
							name: "QUERY_HANDLER",
						},
					},
				},
			},
		},
		{
			name: "INVARIANT_BODY",
			pos:  position{line: 51, col: 1, offset: 1790},
			expr: &zeroOrMoreExpr{
				pos: position{line: 51, col: 18, offset: 1807},
				expr: &choiceExpr{
					pos: position{line: 51, col: 19, offset: 1808},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 51, col: 19, offset: 1808},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 51, col: 32, offset: 1821},
							name: "CHECK",
						},
					},
				},
			},
		},
		{
			name: "PROJECTION_BODY",
			pos:  position{line: 53, col: 1, offset: 1830},
			expr: &zeroOrMoreExpr{
				pos: position{line: 53, col: 19, offset: 1848},
				expr: &choiceExpr{
					pos: position{line: 53, col: 20, offset: 1849},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 53, col: 20, offset: 1849},
							name: "WHEN",
						},
						&ruleRefExpr{
							pos:  position{line: 53, col: 27, offset: 1856},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 53, col: 40, offset: 1869},
							name: "CHECK",
						},
						&ruleRefExpr{
							pos:  position{line: 53, col: 48, offset: 1877},
							name: "FUNCTION",
						},
					},
				},
			},
		},
		{
			name: "WHEN",
			pos:  position{line: 55, col: 1, offset: 1889},
			expr: &seqExpr{
				pos: position{line: 55, col: 8, offset: 1896},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 55, col: 8, offset: 1896},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 55, col: 10, offset: 1898},
						val:        "when",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 55, col: 18, offset: 1906},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 55, col: 20, offset: 1908},
						val:        "event",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 55, col: 29, offset: 1917},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 55, col: 31, offset: 1919},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 55, col: 43, offset: 1931},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 55, col: 45, offset: 1933},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 55, col: 49, offset: 1937},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 55, col: 51, offset: 1939},
						expr: &ruleRefExpr{
							pos:  position{line: 55, col: 51, offset: 1939},
							name: "STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 55, col: 68, offset: 1956},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 55, col: 70, offset: 1958},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 55, col: 74, offset: 1962},
						name: "_",
					},
				},
			},
		},
		{
			name: "COMMAND_BODY",
			pos:  position{line: 57, col: 1, offset: 1965},
			expr: &zeroOrMoreExpr{
				pos: position{line: 57, col: 16, offset: 1980},
				expr: &choiceExpr{
					pos: position{line: 57, col: 17, offset: 1981},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 57, col: 17, offset: 1981},
							name: "COMMAND_HANDLER",
						},
						&ruleRefExpr{
							pos:  position{line: 57, col: 35, offset: 1999},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 57, col: 48, offset: 2012},
							name: "CHECK",
						},
						&ruleRefExpr{
							pos:  position{line: 57, col: 56, offset: 2020},
							name: "FUNCTION",
						},
					},
				},
			},
		},
		{
			name: "COMMAND_HANDLER",
			pos:  position{line: 59, col: 1, offset: 2032},
			expr: &seqExpr{
				pos: position{line: 59, col: 19, offset: 2050},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 59, col: 19, offset: 2050},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 59, col: 21, offset: 2052},
						val:        "handler",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 59, col: 32, offset: 2063},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 59, col: 34, offset: 2065},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 59, col: 38, offset: 2069},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 59, col: 40, offset: 2071},
						expr: &ruleRefExpr{
							pos:  position{line: 59, col: 40, offset: 2071},
							name: "COMMAND_STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 59, col: 65, offset: 2096},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 59, col: 67, offset: 2098},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 59, col: 71, offset: 2102},
						name: "_",
					},
				},
			},
		},
		{
			name: "QUERY_HANDLER",
			pos:  position{line: 61, col: 1, offset: 2105},
			expr: &seqExpr{
				pos: position{line: 61, col: 17, offset: 2121},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 61, col: 17, offset: 2121},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 61, col: 19, offset: 2123},
						val:        "handler",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 61, col: 30, offset: 2134},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 61, col: 32, offset: 2136},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 61, col: 36, offset: 2140},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 61, col: 38, offset: 2142},
						expr: &ruleRefExpr{
							pos:  position{line: 61, col: 38, offset: 2142},
							name: "STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 61, col: 55, offset: 2159},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 61, col: 57, offset: 2161},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 61, col: 61, offset: 2165},
						name: "_",
					},
				},
			},
		},
		{
			name: "COMMAND_STATEMENT_BLOCK",
			pos:  position{line: 63, col: 1, offset: 2168},
			expr: &seqExpr{
				pos: position{line: 63, col: 27, offset: 2194},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 63, col: 27, offset: 2194},
						name: "_",
					},
					&oneOrMoreExpr{
						pos: position{line: 63, col: 29, offset: 2196},
						expr: &ruleRefExpr{
							pos:  position{line: 63, col: 30, offset: 2197},
							name: "COMMAND_STATEMENT",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 63, col: 50, offset: 2217},
						name: "_",
					},
				},
			},
		},
		{
			name: "COMMAND_STATEMENT",
			pos:  position{line: 65, col: 1, offset: 2220},
			expr: &choiceExpr{
				pos: position{line: 65, col: 21, offset: 2240},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 65, col: 21, offset: 2240},
						name: "STATEMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 65, col: 33, offset: 2252},
						name: "ASSERT",
					},
					&ruleRefExpr{
						pos:  position{line: 65, col: 42, offset: 2261},
						name: "APPLY",
					},
				},
			},
		},
		{
			name: "ASSERT",
			pos:  position{line: 67, col: 1, offset: 2268},
			expr: &seqExpr{
				pos: position{line: 67, col: 10, offset: 2277},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 67, col: 10, offset: 2277},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 67, col: 12, offset: 2279},
						val:        "assert",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 67, col: 22, offset: 2289},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 67, col: 24, offset: 2291},
						val:        "invariant",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 67, col: 37, offset: 2304},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 67, col: 39, offset: 2306},
						expr: &litMatcher{
							pos:        position{line: 67, col: 40, offset: 2307},
							val:        "not",
							ignoreCase: true,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 67, col: 49, offset: 2316},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 67, col: 51, offset: 2318},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 67, col: 63, offset: 2330},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 67, col: 65, offset: 2332},
						expr: &ruleRefExpr{
							pos:  position{line: 67, col: 65, offset: 2332},
							name: "ARGUMENTS",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 67, col: 76, offset: 2343},
						name: "SEMI",
					},
				},
			},
		},
		{
			name: "APPLY",
			pos:  position{line: 69, col: 1, offset: 2349},
			expr: &seqExpr{
				pos: position{line: 69, col: 9, offset: 2357},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 69, col: 9, offset: 2357},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 69, col: 11, offset: 2359},
						val:        "apply",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 69, col: 20, offset: 2368},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 69, col: 22, offset: 2370},
						val:        "event",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 69, col: 31, offset: 2379},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 69, col: 33, offset: 2381},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 69, col: 45, offset: 2393},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 69, col: 47, offset: 2395},
						expr: &ruleRefExpr{
							pos:  position{line: 69, col: 47, offset: 2395},
							name: "ARGUMENTS",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 69, col: 58, offset: 2406},
						name: "SEMI",
					},
				},
			},
		},
		{
			name: "VALUE_BODY",
			pos:  position{line: 71, col: 1, offset: 2412},
			expr: &zeroOrMoreExpr{
				pos: position{line: 71, col: 14, offset: 2425},
				expr: &ruleRefExpr{
					pos:  position{line: 71, col: 15, offset: 2426},
					name: "VALUE_COMPONENTS",
				},
			},
		},
		{
			name: "VALUE_COMPONENTS",
			pos:  position{line: 73, col: 1, offset: 2446},
			expr: &choiceExpr{
				pos: position{line: 73, col: 20, offset: 2465},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 73, col: 20, offset: 2465},
						name: "PROPERTIES",
					},
					&ruleRefExpr{
						pos:  position{line: 73, col: 33, offset: 2478},
						name: "CHECK",
					},
					&ruleRefExpr{
						pos:  position{line: 73, col: 41, offset: 2486},
						name: "FUNCTION",
					},
				},
			},
		},
		{
			name: "PROPERTIES",
			pos:  position{line: 75, col: 1, offset: 2496},
			expr: &seqExpr{
				pos: position{line: 75, col: 14, offset: 2509},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 75, col: 14, offset: 2509},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 75, col: 16, offset: 2511},
						val:        "properties",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 30, offset: 2525},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 75, col: 32, offset: 2527},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 36, offset: 2531},
						name: "PROPERTY_LIST",
					},
					&litMatcher{
						pos:        position{line: 75, col: 50, offset: 2545},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 75, col: 54, offset: 2549},
						name: "_",
					},
				},
			},
		},
		{
			name: "PROPERTY_LIST",
			pos:  position{line: 77, col: 1, offset: 2552},
			expr: &zeroOrMoreExpr{
				pos: position{line: 77, col: 17, offset: 2568},
				expr: &ruleRefExpr{
					pos:  position{line: 77, col: 18, offset: 2569},
					name: "PROPERTY",
				},
			},
		},
		{
			name: "PROPERTY",
			pos:  position{line: 79, col: 1, offset: 2581},
			expr: &seqExpr{
				pos: position{line: 79, col: 12, offset: 2592},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 79, col: 12, offset: 2592},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 79, col: 14, offset: 2594},
						name: "TYPE",
					},
					&ruleRefExpr{
						pos:  position{line: 79, col: 19, offset: 2599},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 79, col: 21, offset: 2601},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 79, col: 32, offset: 2612},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 79, col: 34, offset: 2614},
						val:        ";",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 79, col: 38, offset: 2618},
						name: "_",
					},
				},
			},
		},
		{
			name: "CHECK",
			pos:  position{line: 81, col: 1, offset: 2621},
			expr: &seqExpr{
				pos: position{line: 81, col: 9, offset: 2629},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 81, col: 9, offset: 2629},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 81, col: 11, offset: 2631},
						val:        "check",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 81, col: 20, offset: 2640},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 81, col: 22, offset: 2642},
						val:        "(",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 81, col: 26, offset: 2646},
						expr: &ruleRefExpr{
							pos:  position{line: 81, col: 26, offset: 2646},
							name: "STATEMENT_BLOCK",
						},
					},
					&litMatcher{
						pos:        position{line: 81, col: 43, offset: 2663},
						val:        ")",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 81, col: 47, offset: 2667},
						name: "_",
					},
				},
			},
		},
		{
			name: "FUNCTION",
			pos:  position{line: 83, col: 1, offset: 2670},
			expr: &seqExpr{
				pos: position{line: 83, col: 12, offset: 2681},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 83, col: 12, offset: 2681},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 83, col: 14, offset: 2683},
						val:        "function",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 83, col: 26, offset: 2695},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 83, col: 28, offset: 2697},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 83, col: 39, offset: 2708},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 83, col: 41, offset: 2710},
						name: "PARAMETERS",
					},
					&ruleRefExpr{
						pos:  position{line: 83, col: 53, offset: 2722},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 83, col: 55, offset: 2724},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 83, col: 59, offset: 2728},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 83, col: 61, offset: 2730},
						expr: &ruleRefExpr{
							pos:  position{line: 83, col: 61, offset: 2730},
							name: "STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 83, col: 78, offset: 2747},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 83, col: 80, offset: 2749},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 83, col: 84, offset: 2753},
						name: "_",
					},
				},
			},
		},
		{
			name: "PARAMETERS",
			pos:  position{line: 85, col: 1, offset: 2756},
			expr: &seqExpr{
				pos: position{line: 85, col: 14, offset: 2769},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 85, col: 14, offset: 2769},
						val:        "(",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 85, col: 18, offset: 2773},
						name: "PARAMETER_LIST",
					},
					&litMatcher{
						pos:        position{line: 85, col: 33, offset: 2788},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "PARAMETER_LIST",
			pos:  position{line: 87, col: 1, offset: 2793},
			expr: &seqExpr{
				pos: position{line: 87, col: 18, offset: 2810},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 87, col: 18, offset: 2810},
						name: "_",
					},
					&zeroOrMoreExpr{
						pos: position{line: 87, col: 20, offset: 2812},
						expr: &seqExpr{
							pos: position{line: 87, col: 21, offset: 2813},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 87, col: 21, offset: 2813},
									name: "PARAMETER",
								},
								&litMatcher{
									pos:        position{line: 87, col: 31, offset: 2823},
									val:        ",",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 87, col: 35, offset: 2827},
									name: "_",
								},
							},
						},
					},
					&zeroOrOneExpr{
						pos: position{line: 87, col: 40, offset: 2832},
						expr: &ruleRefExpr{
							pos:  position{line: 87, col: 40, offset: 2832},
							name: "PARAMETER",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 87, col: 51, offset: 2843},
						name: "_",
					},
				},
			},
		},
		{
			name: "PARAMETER",
			pos:  position{line: 89, col: 1, offset: 2846},
			expr: &seqExpr{
				pos: position{line: 89, col: 13, offset: 2858},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 89, col: 13, offset: 2858},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 89, col: 15, offset: 2860},
						name: "CLASS_REF",
					},
					&ruleRefExpr{
						pos:  position{line: 89, col: 25, offset: 2870},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 89, col: 27, offset: 2872},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 89, col: 38, offset: 2883},
						name: "_",
					},
				},
			},
		},
		{
			name: "STATEMENT_BLOCK",
			pos:  position{line: 94, col: 1, offset: 3055},
			expr: &seqExpr{
				pos: position{line: 94, col: 19, offset: 3073},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 94, col: 19, offset: 3073},
						name: "_",
					},
					&oneOrMoreExpr{
						pos: position{line: 94, col: 21, offset: 3075},
						expr: &ruleRefExpr{
							pos:  position{line: 94, col: 22, offset: 3076},
							name: "STATEMENT",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 94, col: 34, offset: 3088},
						name: "_",
					},
				},
			},
		},
		{
			name: "STATEMENT",
			pos:  position{line: 96, col: 1, offset: 3091},
			expr: &choiceExpr{
				pos: position{line: 96, col: 13, offset: 3103},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 96, col: 13, offset: 3103},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 96, col: 13, offset: 3103},
								name: "RETURN",
							},
							&ruleRefExpr{
								pos:  position{line: 96, col: 20, offset: 3110},
								name: "SEMI",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 27, offset: 3117},
						name: "IF",
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 32, offset: 3122},
						name: "FOREACH",
					},
					&seqExpr{
						pos: position{line: 96, col: 42, offset: 3132},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 96, col: 42, offset: 3132},
								name: "EXPRESSION",
							},
							&ruleRefExpr{
								pos:  position{line: 96, col: 53, offset: 3143},
								name: "SEMI",
							},
						},
					},
				},
			},
		},
		{
			name: "IF",
			pos:  position{line: 98, col: 1, offset: 3149},
			expr: &seqExpr{
				pos: position{line: 98, col: 6, offset: 3154},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 98, col: 6, offset: 3154},
						val:        "if",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 98, col: 11, offset: 3159},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 98, col: 13, offset: 3161},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 98, col: 24, offset: 3172},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 98, col: 26, offset: 3174},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 98, col: 30, offset: 3178},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 98, col: 32, offset: 3180},
						expr: &ruleRefExpr{
							pos:  position{line: 98, col: 32, offset: 3180},
							name: "STATEMENT_BLOCK",
						},
					},
					&litMatcher{
						pos:        position{line: 98, col: 49, offset: 3197},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 98, col: 53, offset: 3201},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 98, col: 55, offset: 3203},
						expr: &seqExpr{
							pos: position{line: 98, col: 56, offset: 3204},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 98, col: 56, offset: 3204},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 63, offset: 3211},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 98, col: 65, offset: 3213},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 69, offset: 3217},
									name: "_",
								},
								&zeroOrOneExpr{
									pos: position{line: 98, col: 71, offset: 3219},
									expr: &ruleRefExpr{
										pos:  position{line: 98, col: 71, offset: 3219},
										name: "STATEMENT_BLOCK",
									},
								},
								&litMatcher{
									pos:        position{line: 98, col: 88, offset: 3236},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 98, col: 92, offset: 3240},
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
			pos:  position{line: 100, col: 1, offset: 3245},
			expr: &seqExpr{
				pos: position{line: 100, col: 11, offset: 3255},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 100, col: 11, offset: 3255},
						val:        "foreach",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 100, col: 21, offset: 3265},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 100, col: 23, offset: 3267},
						val:        "(",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 100, col: 27, offset: 3271},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 100, col: 29, offset: 3273},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 100, col: 40, offset: 3284},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 100, col: 42, offset: 3286},
						val:        "as",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 100, col: 47, offset: 3291},
						expr: &seqExpr{
							pos: position{line: 100, col: 48, offset: 3292},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 100, col: 48, offset: 3292},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 50, offset: 3294},
									name: "IDENTIFIER",
								},
								&ruleRefExpr{
									pos:  position{line: 100, col: 61, offset: 3305},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 100, col: 63, offset: 3307},
									val:        "=>",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 100, col: 70, offset: 3314},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 100, col: 72, offset: 3316},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 100, col: 83, offset: 3327},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 100, col: 85, offset: 3329},
						val:        ")",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 100, col: 89, offset: 3333},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 100, col: 91, offset: 3335},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 100, col: 95, offset: 3339},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 100, col: 97, offset: 3341},
						expr: &ruleRefExpr{
							pos:  position{line: 100, col: 97, offset: 3341},
							name: "STATEMENT_BLOCK",
						},
					},
					&litMatcher{
						pos:        position{line: 100, col: 114, offset: 3358},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "RETURN",
			pos:  position{line: 102, col: 1, offset: 3363},
			expr: &seqExpr{
				pos: position{line: 102, col: 10, offset: 3372},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 102, col: 10, offset: 3372},
						val:        "return",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 102, col: 19, offset: 3381},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 102, col: 21, offset: 3383},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "EXRESSION_TEST",
			pos:  position{line: 108, col: 1, offset: 3566},
			expr: &seqExpr{
				pos: position{line: 108, col: 18, offset: 3583},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 108, col: 18, offset: 3583},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 108, col: 29, offset: 3594},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "EXPRESSION",
			pos:  position{line: 110, col: 1, offset: 3599},
			expr: &choiceExpr{
				pos: position{line: 110, col: 14, offset: 3612},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 110, col: 14, offset: 3612},
						name: "QUERY",
					},
					&ruleRefExpr{
						pos:  position{line: 110, col: 22, offset: 3620},
						name: "ARITHMETIC",
					},
					&ruleRefExpr{
						pos:  position{line: 110, col: 35, offset: 3633},
						name: "COMPARISON",
					},
					&ruleRefExpr{
						pos:  position{line: 110, col: 48, offset: 3646},
						name: "ASSIGNMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 110, col: 60, offset: 3658},
						name: "LOGICAL",
					},
					&ruleRefExpr{
						pos:  position{line: 110, col: 70, offset: 3668},
						name: "ATOMIC",
					},
				},
			},
		},
		{
			name: "ATOMIC",
			pos:  position{line: 112, col: 1, offset: 3676},
			expr: &choiceExpr{
				pos: position{line: 112, col: 10, offset: 3685},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 112, col: 10, offset: 3685},
						name: "PARENTHESIS",
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 24, offset: 3699},
						name: "NEW",
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 30, offset: 3705},
						name: "METHODCALL",
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 43, offset: 3718},
						name: "OBJECTACCESS",
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 58, offset: 3733},
						name: "LITERAL",
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 68, offset: 3743},
						name: "UNARY",
					},
				},
			},
		},
		{
			name: "LITERAL",
			pos:  position{line: 114, col: 1, offset: 3750},
			expr: &choiceExpr{
				pos: position{line: 114, col: 11, offset: 3760},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 114, col: 11, offset: 3760},
						name: "STRING",
					},
					&ruleRefExpr{
						pos:  position{line: 114, col: 20, offset: 3769},
						name: "FLOAT",
					},
					&ruleRefExpr{
						pos:  position{line: 114, col: 28, offset: 3777},
						name: "BOOLEAN",
					},
					&ruleRefExpr{
						pos:  position{line: 114, col: 38, offset: 3787},
						name: "NULL",
					},
					&ruleRefExpr{
						pos:  position{line: 114, col: 45, offset: 3794},
						name: "INT",
					},
				},
			},
		},
		{
			name: "NEW",
			pos:  position{line: 116, col: 1, offset: 3799},
			expr: &seqExpr{
				pos: position{line: 116, col: 7, offset: 3805},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 116, col: 7, offset: 3805},
						name: "CLASS_REF_QUOTES",
					},
					&ruleRefExpr{
						pos:  position{line: 116, col: 24, offset: 3822},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 116, col: 26, offset: 3824},
						expr: &ruleRefExpr{
							pos:  position{line: 116, col: 26, offset: 3824},
							name: "ARGUMENTS",
						},
					},
				},
			},
		},
		{
			name: "BOOLEAN",
			pos:  position{line: 118, col: 1, offset: 3836},
			expr: &choiceExpr{
				pos: position{line: 118, col: 12, offset: 3847},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 118, col: 12, offset: 3847},
						val:        "true",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 118, col: 19, offset: 3854},
						val:        "false",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "NULL",
			pos:  position{line: 120, col: 1, offset: 3863},
			expr: &litMatcher{
				pos:        position{line: 120, col: 8, offset: 3870},
				val:        "null",
				ignoreCase: false,
			},
		},
		{
			name: "STRING",
			pos:  position{line: 122, col: 1, offset: 3878},
			expr: &seqExpr{
				pos: position{line: 122, col: 10, offset: 3887},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 122, col: 10, offset: 3887},
						val:        "\"",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 122, col: 15, offset: 3892},
						expr: &charClassMatcher{
							pos:        position{line: 122, col: 15, offset: 3892},
							val:        "[a-zA-Z0-9]",
							ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&litMatcher{
						pos:        position{line: 122, col: 28, offset: 3905},
						val:        "\"",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "INT",
			pos:  position{line: 124, col: 1, offset: 3911},
			expr: &oneOrMoreExpr{
				pos: position{line: 124, col: 7, offset: 3917},
				expr: &charClassMatcher{
					pos:        position{line: 124, col: 7, offset: 3917},
					val:        "[0-9]",
					ranges:     []rune{'0', '9'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "FLOAT",
			pos:  position{line: 126, col: 1, offset: 3925},
			expr: &seqExpr{
				pos: position{line: 126, col: 9, offset: 3933},
				exprs: []interface{}{
					&oneOrMoreExpr{
						pos: position{line: 126, col: 9, offset: 3933},
						expr: &charClassMatcher{
							pos:        position{line: 126, col: 9, offset: 3933},
							val:        "[0-9]",
							ranges:     []rune{'0', '9'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&charClassMatcher{
						pos:        position{line: 126, col: 16, offset: 3940},
						val:        "[.]",
						chars:      []rune{'.'},
						ignoreCase: false,
						inverted:   false,
					},
					&oneOrMoreExpr{
						pos: position{line: 126, col: 20, offset: 3944},
						expr: &charClassMatcher{
							pos:        position{line: 126, col: 20, offset: 3944},
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
			pos:  position{line: 128, col: 1, offset: 3952},
			expr: &seqExpr{
				pos: position{line: 128, col: 15, offset: 3966},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 128, col: 15, offset: 3966},
						val:        "(",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 128, col: 19, offset: 3970},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 128, col: 21, offset: 3972},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 128, col: 32, offset: 3983},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 128, col: 34, offset: 3985},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "UNARY",
			pos:  position{line: 130, col: 1, offset: 3990},
			expr: &choiceExpr{
				pos: position{line: 130, col: 9, offset: 3998},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 130, col: 9, offset: 3998},
						name: "INCREMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 21, offset: 4010},
						name: "DECREMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 33, offset: 4022},
						name: "NEGATE",
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 42, offset: 4031},
						name: "NOT",
					},
					&ruleRefExpr{
						pos:  position{line: 130, col: 48, offset: 4037},
						name: "POSITIVE",
					},
				},
			},
		},
		{
			name: "INCREMENT",
			pos:  position{line: 132, col: 1, offset: 4047},
			expr: &seqExpr{
				pos: position{line: 132, col: 13, offset: 4059},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 132, col: 13, offset: 4059},
						name: "OBJECTACCESS",
					},
					&litMatcher{
						pos:        position{line: 132, col: 26, offset: 4072},
						val:        "++",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DECREMENT",
			pos:  position{line: 134, col: 1, offset: 4078},
			expr: &seqExpr{
				pos: position{line: 134, col: 13, offset: 4090},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 134, col: 13, offset: 4090},
						name: "OBJECTACCESS",
					},
					&litMatcher{
						pos:        position{line: 134, col: 26, offset: 4103},
						val:        "--",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "NEGATE",
			pos:  position{line: 136, col: 1, offset: 4109},
			expr: &seqExpr{
				pos: position{line: 136, col: 10, offset: 4118},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 136, col: 10, offset: 4118},
						val:        "-",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 136, col: 14, offset: 4122},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "NOT",
			pos:  position{line: 138, col: 1, offset: 4136},
			expr: &seqExpr{
				pos: position{line: 138, col: 7, offset: 4142},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 138, col: 7, offset: 4142},
						val:        "!",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 138, col: 11, offset: 4146},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "POSITIVE",
			pos:  position{line: 140, col: 1, offset: 4160},
			expr: &seqExpr{
				pos: position{line: 140, col: 12, offset: 4171},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 140, col: 12, offset: 4171},
						val:        "+",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 140, col: 16, offset: 4175},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "ARITHMETIC",
			pos:  position{line: 142, col: 1, offset: 4189},
			expr: &seqExpr{
				pos: position{line: 142, col: 14, offset: 4202},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 142, col: 14, offset: 4202},
						name: "ATOMIC",
					},
					&ruleRefExpr{
						pos:  position{line: 142, col: 21, offset: 4209},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 142, col: 23, offset: 4211},
						name: "OPERATOR",
					},
					&ruleRefExpr{
						pos:  position{line: 142, col: 32, offset: 4220},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 142, col: 34, offset: 4222},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "OPERATOR",
			pos:  position{line: 144, col: 1, offset: 4234},
			expr: &choiceExpr{
				pos: position{line: 144, col: 12, offset: 4245},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 144, col: 12, offset: 4245},
						val:        "+",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 144, col: 18, offset: 4251},
						val:        "-",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 144, col: 24, offset: 4257},
						val:        "/",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 144, col: 30, offset: 4263},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 144, col: 36, offset: 4269},
						val:        "%",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "ASSIGNMENT",
			pos:  position{line: 146, col: 1, offset: 4274},
			expr: &seqExpr{
				pos: position{line: 146, col: 14, offset: 4287},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 146, col: 14, offset: 4287},
						name: "OBJECTACCESS",
					},
					&ruleRefExpr{
						pos:  position{line: 146, col: 27, offset: 4300},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 146, col: 29, offset: 4302},
						val:        "=",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 146, col: 33, offset: 4306},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 146, col: 35, offset: 4308},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "LOGICAL",
			pos:  position{line: 148, col: 1, offset: 4320},
			expr: &seqExpr{
				pos: position{line: 148, col: 11, offset: 4330},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 148, col: 11, offset: 4330},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 148, col: 22, offset: 4341},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 148, col: 25, offset: 4344},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 148, col: 25, offset: 4344},
								val:        "and",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 148, col: 33, offset: 4352},
								val:        "or",
								ignoreCase: false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 148, col: 39, offset: 4358},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 148, col: 41, offset: 4360},
						name: "ATOMIC",
					},
				},
			},
		},
		{
			name: "COMPARISON",
			pos:  position{line: 150, col: 1, offset: 4368},
			expr: &seqExpr{
				pos: position{line: 150, col: 14, offset: 4381},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 150, col: 14, offset: 4381},
						name: "ATOMIC",
					},
					&ruleRefExpr{
						pos:  position{line: 150, col: 21, offset: 4388},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 150, col: 24, offset: 4391},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 150, col: 24, offset: 4391},
								val:        "===",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 150, col: 32, offset: 4399},
								val:        "!==",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 150, col: 40, offset: 4407},
								val:        "==",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 150, col: 47, offset: 4414},
								val:        "!=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 150, col: 54, offset: 4421},
								val:        "<=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 150, col: 61, offset: 4428},
								val:        ">=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 150, col: 68, offset: 4435},
								val:        "<",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 150, col: 74, offset: 4441},
								val:        ">",
								ignoreCase: false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 150, col: 79, offset: 4446},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 150, col: 81, offset: 4448},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "QUERY",
			pos:  position{line: 152, col: 1, offset: 4460},
			expr: &seqExpr{
				pos: position{line: 152, col: 9, offset: 4468},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 152, col: 9, offset: 4468},
						val:        "run",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 152, col: 16, offset: 4475},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 152, col: 18, offset: 4477},
						val:        "query",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 152, col: 27, offset: 4486},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 152, col: 29, offset: 4488},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 152, col: 41, offset: 4500},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 152, col: 43, offset: 4502},
						expr: &ruleRefExpr{
							pos:  position{line: 152, col: 43, offset: 4502},
							name: "ARGUMENTS",
						},
					},
				},
			},
		},
		{
			name: "OBJECTACCESS",
			pos:  position{line: 154, col: 1, offset: 4514},
			expr: &seqExpr{
				pos: position{line: 154, col: 16, offset: 4529},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 154, col: 16, offset: 4529},
						expr: &seqExpr{
							pos: position{line: 154, col: 17, offset: 4530},
							exprs: []interface{}{
								&choiceExpr{
									pos: position{line: 154, col: 18, offset: 4531},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 154, col: 18, offset: 4531},
											name: "METHODCALL",
										},
										&ruleRefExpr{
											pos:  position{line: 154, col: 31, offset: 4544},
											name: "IDENTIFIER",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 154, col: 43, offset: 4556},
									val:        "->",
									ignoreCase: false,
								},
							},
						},
					},
					&choiceExpr{
						pos: position{line: 154, col: 51, offset: 4564},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 154, col: 51, offset: 4564},
								name: "METHODCALL",
							},
							&ruleRefExpr{
								pos:  position{line: 154, col: 64, offset: 4577},
								name: "IDENTIFIER",
							},
						},
					},
				},
			},
		},
		{
			name: "METHODCALL",
			pos:  position{line: 156, col: 1, offset: 4590},
			expr: &seqExpr{
				pos: position{line: 156, col: 14, offset: 4603},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 156, col: 14, offset: 4603},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 156, col: 25, offset: 4614},
						name: "ARGUMENTS",
					},
				},
			},
		},
		{
			name: "ARGUMENTS",
			pos:  position{line: 158, col: 1, offset: 4625},
			expr: &seqExpr{
				pos: position{line: 158, col: 13, offset: 4637},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 158, col: 13, offset: 4637},
						val:        "(",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 158, col: 17, offset: 4641},
						expr: &ruleRefExpr{
							pos:  position{line: 158, col: 17, offset: 4641},
							name: "ARGUMENTLIST",
						},
					},
					&litMatcher{
						pos:        position{line: 158, col: 31, offset: 4655},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "ARGUMENTLIST",
			pos:  position{line: 160, col: 1, offset: 4660},
			expr: &seqExpr{
				pos: position{line: 160, col: 17, offset: 4676},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 160, col: 17, offset: 4676},
						name: "_",
					},
					&zeroOrMoreExpr{
						pos: position{line: 160, col: 19, offset: 4678},
						expr: &seqExpr{
							pos: position{line: 160, col: 20, offset: 4679},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 160, col: 20, offset: 4679},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 160, col: 22, offset: 4681},
									name: "EXPRESSION",
								},
								&ruleRefExpr{
									pos:  position{line: 160, col: 33, offset: 4692},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 160, col: 35, offset: 4694},
									val:        ",",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 160, col: 39, offset: 4698},
									name: "_",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 160, col: 43, offset: 4702},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 160, col: 54, offset: 4713},
						name: "_",
					},
				},
			},
		},
		{
			name: "CLASS_REF_QUOTES",
			pos:  position{line: 167, col: 1, offset: 4881},
			expr: &seqExpr{
				pos: position{line: 167, col: 20, offset: 4900},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 167, col: 20, offset: 4900},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 167, col: 22, offset: 4902},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 167, col: 26, offset: 4906},
						name: "CLASS_REF",
					},
					&litMatcher{
						pos:        position{line: 167, col: 36, offset: 4916},
						val:        "'",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CLASS_REF",
			pos:  position{line: 169, col: 1, offset: 4921},
			expr: &seqExpr{
				pos: position{line: 169, col: 13, offset: 4933},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 169, col: 13, offset: 4933},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 169, col: 15, offset: 4935},
						name: "CLASS_TYPE",
					},
					&litMatcher{
						pos:        position{line: 169, col: 26, offset: 4946},
						val:        "\\",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 169, col: 31, offset: 4951},
						name: "CLASS_NAME",
					},
				},
			},
		},
		{
			name: "CLASS_TYPE",
			pos:  position{line: 171, col: 1, offset: 4963},
			expr: &seqExpr{
				pos: position{line: 171, col: 14, offset: 4976},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 171, col: 14, offset: 4976},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 171, col: 17, offset: 4979},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 171, col: 17, offset: 4979},
								val:        "value",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 171, col: 27, offset: 4989},
								val:        "entity",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 171, col: 38, offset: 5000},
								val:        "command",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 171, col: 50, offset: 5012},
								val:        "event",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 171, col: 60, offset: 5022},
								val:        "projection",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 171, col: 75, offset: 5037},
								val:        "invariant",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 171, col: 89, offset: 5051},
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
			pos:  position{line: 173, col: 1, offset: 5061},
			expr: &seqExpr{
				pos: position{line: 173, col: 14, offset: 5074},
				exprs: []interface{}{
					&oneOrMoreExpr{
						pos: position{line: 173, col: 14, offset: 5074},
						expr: &charClassMatcher{
							pos:        position{line: 173, col: 14, offset: 5074},
							val:        "[a-zA-Z]",
							ranges:     []rune{'a', 'z', 'A', 'Z'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 173, col: 24, offset: 5084},
						expr: &charClassMatcher{
							pos:        position{line: 173, col: 24, offset: 5084},
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
			pos:  position{line: 175, col: 1, offset: 5100},
			expr: &seqExpr{
				pos: position{line: 175, col: 15, offset: 5114},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 175, col: 15, offset: 5114},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 175, col: 19, offset: 5118},
						name: "CLASS_NAME",
					},
					&litMatcher{
						pos:        position{line: 175, col: 30, offset: 5129},
						val:        "'",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "TYPE",
			pos:  position{line: 177, col: 1, offset: 5134},
			expr: &choiceExpr{
				pos: position{line: 177, col: 8, offset: 5141},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 177, col: 8, offset: 5141},
						name: "CLASS_REF",
					},
					&ruleRefExpr{
						pos:  position{line: 177, col: 20, offset: 5153},
						name: "VALUE_TYPE",
					},
				},
			},
		},
		{
			name: "VALUE_TYPE",
			pos:  position{line: 179, col: 1, offset: 5165},
			expr: &seqExpr{
				pos: position{line: 179, col: 14, offset: 5178},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 179, col: 14, offset: 5178},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 179, col: 17, offset: 5181},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 179, col: 17, offset: 5181},
								val:        "string",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 179, col: 28, offset: 5192},
								val:        "boolean",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 179, col: 40, offset: 5204},
								val:        "float",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 179, col: 50, offset: 5214},
								val:        "map",
								ignoreCase: false,
							},
						},
					},
				},
			},
		},
		{
			name: "CLASS_IMPLIED_REF",
			pos:  position{line: 181, col: 1, offset: 5222},
			expr: &seqExpr{
				pos: position{line: 181, col: 21, offset: 5242},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 181, col: 21, offset: 5242},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 181, col: 23, offset: 5244},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 181, col: 27, offset: 5248},
						name: "CLASS_NAME",
					},
					&litMatcher{
						pos:        position{line: 181, col: 38, offset: 5259},
						val:        "'",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "IDENTIFIER",
			pos:  position{line: 183, col: 1, offset: 5264},
			expr: &seqExpr{
				pos: position{line: 183, col: 14, offset: 5277},
				exprs: []interface{}{
					&oneOrMoreExpr{
						pos: position{line: 183, col: 14, offset: 5277},
						expr: &charClassMatcher{
							pos:        position{line: 183, col: 14, offset: 5277},
							val:        "[a-zA-Z]",
							ranges:     []rune{'a', 'z', 'A', 'Z'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 183, col: 24, offset: 5287},
						expr: &charClassMatcher{
							pos:        position{line: 183, col: 24, offset: 5287},
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
			pos:  position{line: 185, col: 1, offset: 5302},
			expr: &zeroOrMoreExpr{
				pos: position{line: 185, col: 5, offset: 5306},
				expr: &choiceExpr{
					pos: position{line: 185, col: 7, offset: 5308},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 185, col: 7, offset: 5308},
							name: "WHITESPACE",
						},
						&ruleRefExpr{
							pos:  position{line: 185, col: 20, offset: 5321},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "SEMI",
			pos:  position{line: 187, col: 1, offset: 5329},
			expr: &seqExpr{
				pos: position{line: 187, col: 8, offset: 5336},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 187, col: 8, offset: 5336},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 187, col: 10, offset: 5338},
						val:        ";",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 187, col: 14, offset: 5342},
						name: "_",
					},
				},
			},
		},
		{
			name: "WHITESPACE",
			pos:  position{line: 189, col: 1, offset: 5345},
			expr: &charClassMatcher{
				pos:        position{line: 189, col: 14, offset: 5358},
				val:        "[ \\t\\r]",
				chars:      []rune{' ', '\t', '\r'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EOL",
			pos:  position{line: 191, col: 1, offset: 5367},
			expr: &litMatcher{
				pos:        position{line: 191, col: 7, offset: 5373},
				val:        "\n",
				ignoreCase: false,
			},
		},
		{
			name: "EOF",
			pos:  position{line: 193, col: 1, offset: 5379},
			expr: &notExpr{
				pos: position{line: 193, col: 7, offset: 5385},
				expr: &anyMatcher{
					line: 193, col: 8, offset: 5386,
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
