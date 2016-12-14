package components

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
			name: "CLASS_COMPONENT_TEST",
			pos:  position{line: 11, col: 1, offset: 210},
			expr: &seqExpr{
				pos: position{line: 11, col: 24, offset: 233},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 11, col: 24, offset: 233},
						expr: &choiceExpr{
							pos: position{line: 11, col: 25, offset: 234},
							alternatives: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 11, col: 25, offset: 234},
									name: "WHEN",
								},
								&ruleRefExpr{
									pos:  position{line: 11, col: 32, offset: 241},
									name: "COMMAND_HANDLER",
								},
								&ruleRefExpr{
									pos:  position{line: 11, col: 50, offset: 259},
									name: "PROPERTIES",
								},
								&ruleRefExpr{
									pos:  position{line: 11, col: 63, offset: 272},
									name: "CHECK",
								},
								&ruleRefExpr{
									pos:  position{line: 11, col: 71, offset: 280},
									name: "FUNCTION",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 11, col: 82, offset: 291},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "QUERY_BODY",
			pos:  position{line: 13, col: 1, offset: 296},
			expr: &zeroOrMoreExpr{
				pos: position{line: 13, col: 14, offset: 309},
				expr: &choiceExpr{
					pos: position{line: 13, col: 15, offset: 310},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 13, col: 15, offset: 310},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 13, col: 28, offset: 323},
							name: "QUERY_HANDLER",
						},
					},
				},
			},
		},
		{
			name: "INVARIANT_BODY",
			pos:  position{line: 15, col: 1, offset: 340},
			expr: &zeroOrMoreExpr{
				pos: position{line: 15, col: 18, offset: 357},
				expr: &choiceExpr{
					pos: position{line: 15, col: 19, offset: 358},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 15, col: 19, offset: 358},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 15, col: 32, offset: 371},
							name: "CHECK",
						},
					},
				},
			},
		},
		{
			name: "PROJECTION_BODY",
			pos:  position{line: 17, col: 1, offset: 380},
			expr: &zeroOrMoreExpr{
				pos: position{line: 17, col: 19, offset: 398},
				expr: &choiceExpr{
					pos: position{line: 17, col: 20, offset: 399},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 17, col: 20, offset: 399},
							name: "WHEN",
						},
						&ruleRefExpr{
							pos:  position{line: 17, col: 27, offset: 406},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 17, col: 40, offset: 419},
							name: "CHECK",
						},
						&ruleRefExpr{
							pos:  position{line: 17, col: 48, offset: 427},
							name: "FUNCTION",
						},
					},
				},
			},
		},
		{
			name: "WHEN",
			pos:  position{line: 19, col: 1, offset: 439},
			expr: &seqExpr{
				pos: position{line: 19, col: 8, offset: 446},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 19, col: 8, offset: 446},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 19, col: 10, offset: 448},
						val:        "when",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 18, offset: 456},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 19, col: 20, offset: 458},
						val:        "event",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 29, offset: 467},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 31, offset: 469},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 43, offset: 481},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 19, col: 45, offset: 483},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 49, offset: 487},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 19, col: 51, offset: 489},
						expr: &ruleRefExpr{
							pos:  position{line: 19, col: 51, offset: 489},
							name: "STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 68, offset: 506},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 19, col: 70, offset: 508},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 19, col: 74, offset: 512},
						name: "_",
					},
				},
			},
		},
		{
			name: "COMMAND_BODY",
			pos:  position{line: 21, col: 1, offset: 515},
			expr: &zeroOrMoreExpr{
				pos: position{line: 21, col: 16, offset: 530},
				expr: &choiceExpr{
					pos: position{line: 21, col: 17, offset: 531},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 21, col: 17, offset: 531},
							name: "COMMAND_HANDLER",
						},
						&ruleRefExpr{
							pos:  position{line: 21, col: 35, offset: 549},
							name: "PROPERTIES",
						},
						&ruleRefExpr{
							pos:  position{line: 21, col: 48, offset: 562},
							name: "CHECK",
						},
						&ruleRefExpr{
							pos:  position{line: 21, col: 56, offset: 570},
							name: "FUNCTION",
						},
					},
				},
			},
		},
		{
			name: "COMMAND_HANDLER",
			pos:  position{line: 23, col: 1, offset: 582},
			expr: &seqExpr{
				pos: position{line: 23, col: 19, offset: 600},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 23, col: 19, offset: 600},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 23, col: 21, offset: 602},
						val:        "handler",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 32, offset: 613},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 23, col: 34, offset: 615},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 38, offset: 619},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 23, col: 40, offset: 621},
						expr: &ruleRefExpr{
							pos:  position{line: 23, col: 40, offset: 621},
							name: "COMMAND_STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 65, offset: 646},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 23, col: 67, offset: 648},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 23, col: 71, offset: 652},
						name: "_",
					},
				},
			},
		},
		{
			name: "QUERY_HANDLER",
			pos:  position{line: 25, col: 1, offset: 655},
			expr: &seqExpr{
				pos: position{line: 25, col: 17, offset: 671},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 25, col: 17, offset: 671},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 25, col: 19, offset: 673},
						val:        "handler",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 30, offset: 684},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 25, col: 32, offset: 686},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 36, offset: 690},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 25, col: 38, offset: 692},
						expr: &ruleRefExpr{
							pos:  position{line: 25, col: 38, offset: 692},
							name: "STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 55, offset: 709},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 25, col: 57, offset: 711},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 25, col: 61, offset: 715},
						name: "_",
					},
				},
			},
		},
		{
			name: "COMMAND_STATEMENT_BLOCK",
			pos:  position{line: 27, col: 1, offset: 718},
			expr: &seqExpr{
				pos: position{line: 27, col: 27, offset: 744},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 27, col: 27, offset: 744},
						name: "_",
					},
					&oneOrMoreExpr{
						pos: position{line: 27, col: 29, offset: 746},
						expr: &ruleRefExpr{
							pos:  position{line: 27, col: 30, offset: 747},
							name: "COMMAND_STATEMENT",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 27, col: 50, offset: 767},
						name: "_",
					},
				},
			},
		},
		{
			name: "COMMAND_STATEMENT",
			pos:  position{line: 29, col: 1, offset: 770},
			expr: &choiceExpr{
				pos: position{line: 29, col: 21, offset: 790},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 29, col: 21, offset: 790},
						name: "STATEMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 29, col: 33, offset: 802},
						name: "ASSERT",
					},
					&ruleRefExpr{
						pos:  position{line: 29, col: 42, offset: 811},
						name: "APPLY",
					},
				},
			},
		},
		{
			name: "ASSERT",
			pos:  position{line: 31, col: 1, offset: 818},
			expr: &seqExpr{
				pos: position{line: 31, col: 10, offset: 827},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 31, col: 10, offset: 827},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 31, col: 12, offset: 829},
						val:        "assert",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 31, col: 22, offset: 839},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 31, col: 24, offset: 841},
						val:        "invariant",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 31, col: 37, offset: 854},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 31, col: 39, offset: 856},
						expr: &litMatcher{
							pos:        position{line: 31, col: 40, offset: 857},
							val:        "not",
							ignoreCase: true,
						},
					},
					&ruleRefExpr{
						pos:  position{line: 31, col: 49, offset: 866},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 31, col: 51, offset: 868},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 31, col: 63, offset: 880},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 31, col: 65, offset: 882},
						expr: &ruleRefExpr{
							pos:  position{line: 31, col: 65, offset: 882},
							name: "ARGUMENTS",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 31, col: 76, offset: 893},
						name: "SEMI",
					},
				},
			},
		},
		{
			name: "APPLY",
			pos:  position{line: 33, col: 1, offset: 899},
			expr: &seqExpr{
				pos: position{line: 33, col: 9, offset: 907},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 33, col: 9, offset: 907},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 33, col: 11, offset: 909},
						val:        "apply",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 20, offset: 918},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 33, col: 22, offset: 920},
						val:        "event",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 31, offset: 929},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 33, offset: 931},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 45, offset: 943},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 33, col: 47, offset: 945},
						expr: &ruleRefExpr{
							pos:  position{line: 33, col: 47, offset: 945},
							name: "ARGUMENTS",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 33, col: 58, offset: 956},
						name: "SEMI",
					},
				},
			},
		},
		{
			name: "VALUE_BODY",
			pos:  position{line: 35, col: 1, offset: 962},
			expr: &zeroOrMoreExpr{
				pos: position{line: 35, col: 14, offset: 975},
				expr: &ruleRefExpr{
					pos:  position{line: 35, col: 15, offset: 976},
					name: "VALUE_COMPONENTS",
				},
			},
		},
		{
			name: "VALUE_COMPONENTS",
			pos:  position{line: 37, col: 1, offset: 996},
			expr: &choiceExpr{
				pos: position{line: 37, col: 20, offset: 1015},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 37, col: 20, offset: 1015},
						name: "PROPERTIES",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 33, offset: 1028},
						name: "CHECK",
					},
					&ruleRefExpr{
						pos:  position{line: 37, col: 41, offset: 1036},
						name: "FUNCTION",
					},
				},
			},
		},
		{
			name: "PROPERTIES",
			pos:  position{line: 39, col: 1, offset: 1046},
			expr: &seqExpr{
				pos: position{line: 39, col: 14, offset: 1059},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 39, col: 14, offset: 1059},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 39, col: 16, offset: 1061},
						val:        "properties",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 30, offset: 1075},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 39, col: 32, offset: 1077},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 36, offset: 1081},
						name: "PROPERTY_LIST",
					},
					&litMatcher{
						pos:        position{line: 39, col: 50, offset: 1095},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 39, col: 54, offset: 1099},
						name: "_",
					},
				},
			},
		},
		{
			name: "PROPERTY_LIST",
			pos:  position{line: 41, col: 1, offset: 1102},
			expr: &zeroOrMoreExpr{
				pos: position{line: 41, col: 17, offset: 1118},
				expr: &ruleRefExpr{
					pos:  position{line: 41, col: 18, offset: 1119},
					name: "PROPERTY",
				},
			},
		},
		{
			name: "PROPERTY",
			pos:  position{line: 43, col: 1, offset: 1131},
			expr: &seqExpr{
				pos: position{line: 43, col: 12, offset: 1142},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 43, col: 12, offset: 1142},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 14, offset: 1144},
						name: "TYPE",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 19, offset: 1149},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 21, offset: 1151},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 32, offset: 1162},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 43, col: 35, offset: 1165},
						expr: &seqExpr{
							pos: position{line: 43, col: 36, offset: 1166},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 43, col: 36, offset: 1166},
									val:        "=",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 43, col: 40, offset: 1170},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 43, col: 42, offset: 1172},
									name: "EXPRESSION",
								},
								&ruleRefExpr{
									pos:  position{line: 43, col: 53, offset: 1183},
									name: "_",
								},
							},
						},
					},
					&litMatcher{
						pos:        position{line: 43, col: 57, offset: 1187},
						val:        ";",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 43, col: 61, offset: 1191},
						name: "_",
					},
				},
			},
		},
		{
			name: "CHECK",
			pos:  position{line: 45, col: 1, offset: 1194},
			expr: &seqExpr{
				pos: position{line: 45, col: 9, offset: 1202},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 45, col: 9, offset: 1202},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 45, col: 11, offset: 1204},
						val:        "check",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 45, col: 20, offset: 1213},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 45, col: 22, offset: 1215},
						val:        "(",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 45, col: 26, offset: 1219},
						expr: &ruleRefExpr{
							pos:  position{line: 45, col: 26, offset: 1219},
							name: "STATEMENT_BLOCK",
						},
					},
					&litMatcher{
						pos:        position{line: 45, col: 43, offset: 1236},
						val:        ")",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 45, col: 47, offset: 1240},
						name: "_",
					},
				},
			},
		},
		{
			name: "FUNCTION",
			pos:  position{line: 47, col: 1, offset: 1243},
			expr: &seqExpr{
				pos: position{line: 47, col: 12, offset: 1254},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 47, col: 12, offset: 1254},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 47, col: 14, offset: 1256},
						val:        "function",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 47, col: 26, offset: 1268},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 47, col: 28, offset: 1270},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 47, col: 39, offset: 1281},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 47, col: 41, offset: 1283},
						name: "PARAMETERS",
					},
					&ruleRefExpr{
						pos:  position{line: 47, col: 53, offset: 1295},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 47, col: 55, offset: 1297},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 47, col: 59, offset: 1301},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 47, col: 61, offset: 1303},
						expr: &ruleRefExpr{
							pos:  position{line: 47, col: 61, offset: 1303},
							name: "STATEMENT_BLOCK",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 47, col: 78, offset: 1320},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 47, col: 80, offset: 1322},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 47, col: 84, offset: 1326},
						name: "_",
					},
				},
			},
		},
		{
			name: "PARAMETERS",
			pos:  position{line: 49, col: 1, offset: 1329},
			expr: &seqExpr{
				pos: position{line: 49, col: 14, offset: 1342},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 49, col: 14, offset: 1342},
						val:        "(",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 49, col: 18, offset: 1346},
						name: "PARAMETER_LIST",
					},
					&litMatcher{
						pos:        position{line: 49, col: 33, offset: 1361},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "PARAMETER_LIST",
			pos:  position{line: 51, col: 1, offset: 1366},
			expr: &seqExpr{
				pos: position{line: 51, col: 18, offset: 1383},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 51, col: 18, offset: 1383},
						name: "_",
					},
					&zeroOrMoreExpr{
						pos: position{line: 51, col: 20, offset: 1385},
						expr: &seqExpr{
							pos: position{line: 51, col: 21, offset: 1386},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 51, col: 21, offset: 1386},
									name: "PARAMETER",
								},
								&litMatcher{
									pos:        position{line: 51, col: 31, offset: 1396},
									val:        ",",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 51, col: 35, offset: 1400},
									name: "_",
								},
							},
						},
					},
					&zeroOrOneExpr{
						pos: position{line: 51, col: 40, offset: 1405},
						expr: &ruleRefExpr{
							pos:  position{line: 51, col: 40, offset: 1405},
							name: "PARAMETER",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 51, col: 51, offset: 1416},
						name: "_",
					},
				},
			},
		},
		{
			name: "PARAMETER",
			pos:  position{line: 53, col: 1, offset: 1419},
			expr: &seqExpr{
				pos: position{line: 53, col: 13, offset: 1431},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 53, col: 13, offset: 1431},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 53, col: 15, offset: 1433},
						name: "CLASS_REF",
					},
					&ruleRefExpr{
						pos:  position{line: 53, col: 25, offset: 1443},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 53, col: 27, offset: 1445},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 53, col: 38, offset: 1456},
						name: "_",
					},
				},
			},
		},
		{
			name: "STATEMENT_BLOCK",
			pos:  position{line: 58, col: 1, offset: 1628},
			expr: &seqExpr{
				pos: position{line: 58, col: 19, offset: 1646},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 58, col: 19, offset: 1646},
						name: "_",
					},
					&oneOrMoreExpr{
						pos: position{line: 58, col: 21, offset: 1648},
						expr: &ruleRefExpr{
							pos:  position{line: 58, col: 22, offset: 1649},
							name: "STATEMENT",
						},
					},
					&ruleRefExpr{
						pos:  position{line: 58, col: 34, offset: 1661},
						name: "_",
					},
				},
			},
		},
		{
			name: "STATEMENT",
			pos:  position{line: 60, col: 1, offset: 1664},
			expr: &choiceExpr{
				pos: position{line: 60, col: 13, offset: 1676},
				alternatives: []interface{}{
					&seqExpr{
						pos: position{line: 60, col: 13, offset: 1676},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 60, col: 13, offset: 1676},
								name: "RETURN",
							},
							&ruleRefExpr{
								pos:  position{line: 60, col: 20, offset: 1683},
								name: "SEMI",
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 60, col: 27, offset: 1690},
						name: "IF",
					},
					&ruleRefExpr{
						pos:  position{line: 60, col: 32, offset: 1695},
						name: "FOREACH",
					},
					&seqExpr{
						pos: position{line: 60, col: 42, offset: 1705},
						exprs: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 60, col: 42, offset: 1705},
								name: "EXPRESSION",
							},
							&ruleRefExpr{
								pos:  position{line: 60, col: 53, offset: 1716},
								name: "SEMI",
							},
						},
					},
				},
			},
		},
		{
			name: "IF",
			pos:  position{line: 62, col: 1, offset: 1722},
			expr: &seqExpr{
				pos: position{line: 62, col: 6, offset: 1727},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 62, col: 6, offset: 1727},
						val:        "if",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 11, offset: 1732},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 13, offset: 1734},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 24, offset: 1745},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 62, col: 26, offset: 1747},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 30, offset: 1751},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 62, col: 32, offset: 1753},
						expr: &ruleRefExpr{
							pos:  position{line: 62, col: 32, offset: 1753},
							name: "STATEMENT_BLOCK",
						},
					},
					&litMatcher{
						pos:        position{line: 62, col: 49, offset: 1770},
						val:        "}",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 62, col: 53, offset: 1774},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 62, col: 55, offset: 1776},
						expr: &seqExpr{
							pos: position{line: 62, col: 56, offset: 1777},
							exprs: []interface{}{
								&litMatcher{
									pos:        position{line: 62, col: 56, offset: 1777},
									val:        "else",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 62, col: 63, offset: 1784},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 62, col: 65, offset: 1786},
									val:        "{",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 62, col: 69, offset: 1790},
									name: "_",
								},
								&zeroOrOneExpr{
									pos: position{line: 62, col: 71, offset: 1792},
									expr: &ruleRefExpr{
										pos:  position{line: 62, col: 71, offset: 1792},
										name: "STATEMENT_BLOCK",
									},
								},
								&litMatcher{
									pos:        position{line: 62, col: 88, offset: 1809},
									val:        "}",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 62, col: 92, offset: 1813},
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
			pos:  position{line: 64, col: 1, offset: 1818},
			expr: &seqExpr{
				pos: position{line: 64, col: 11, offset: 1828},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 64, col: 11, offset: 1828},
						val:        "foreach",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 21, offset: 1838},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 64, col: 23, offset: 1840},
						val:        "(",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 27, offset: 1844},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 29, offset: 1846},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 40, offset: 1857},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 64, col: 42, offset: 1859},
						val:        "as",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 64, col: 47, offset: 1864},
						expr: &seqExpr{
							pos: position{line: 64, col: 48, offset: 1865},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 64, col: 48, offset: 1865},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 64, col: 50, offset: 1867},
									name: "IDENTIFIER",
								},
								&ruleRefExpr{
									pos:  position{line: 64, col: 61, offset: 1878},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 64, col: 63, offset: 1880},
									val:        "=>",
									ignoreCase: false,
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 70, offset: 1887},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 72, offset: 1889},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 83, offset: 1900},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 64, col: 85, offset: 1902},
						val:        ")",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 89, offset: 1906},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 64, col: 91, offset: 1908},
						val:        "{",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 64, col: 95, offset: 1912},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 64, col: 97, offset: 1914},
						expr: &ruleRefExpr{
							pos:  position{line: 64, col: 97, offset: 1914},
							name: "STATEMENT_BLOCK",
						},
					},
					&litMatcher{
						pos:        position{line: 64, col: 114, offset: 1931},
						val:        "}",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "RETURN",
			pos:  position{line: 66, col: 1, offset: 1936},
			expr: &seqExpr{
				pos: position{line: 66, col: 10, offset: 1945},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 66, col: 10, offset: 1945},
						val:        "return",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 19, offset: 1954},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 66, col: 21, offset: 1956},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "EXRESSION_TEST",
			pos:  position{line: 72, col: 1, offset: 2139},
			expr: &seqExpr{
				pos: position{line: 72, col: 18, offset: 2156},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 72, col: 18, offset: 2156},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 72, col: 29, offset: 2167},
						name: "EOF",
					},
				},
			},
		},
		{
			name: "EXPRESSION",
			pos:  position{line: 74, col: 1, offset: 2172},
			expr: &choiceExpr{
				pos: position{line: 74, col: 14, offset: 2185},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 74, col: 14, offset: 2185},
						name: "QUERY",
					},
					&ruleRefExpr{
						pos:  position{line: 74, col: 22, offset: 2193},
						name: "ARITHMETIC",
					},
					&ruleRefExpr{
						pos:  position{line: 74, col: 35, offset: 2206},
						name: "COMPARISON",
					},
					&ruleRefExpr{
						pos:  position{line: 74, col: 48, offset: 2219},
						name: "ASSIGNMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 74, col: 60, offset: 2231},
						name: "LOGICAL",
					},
					&ruleRefExpr{
						pos:  position{line: 74, col: 70, offset: 2241},
						name: "ATOMIC",
					},
				},
			},
		},
		{
			name: "ATOMIC",
			pos:  position{line: 76, col: 1, offset: 2249},
			expr: &choiceExpr{
				pos: position{line: 76, col: 10, offset: 2258},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 76, col: 10, offset: 2258},
						name: "PARENTHESIS",
					},
					&ruleRefExpr{
						pos:  position{line: 76, col: 24, offset: 2272},
						name: "NEW",
					},
					&ruleRefExpr{
						pos:  position{line: 76, col: 30, offset: 2278},
						name: "METHODCALL",
					},
					&ruleRefExpr{
						pos:  position{line: 76, col: 43, offset: 2291},
						name: "OBJECTACCESS",
					},
					&ruleRefExpr{
						pos:  position{line: 76, col: 58, offset: 2306},
						name: "ARRAY",
					},
					&ruleRefExpr{
						pos:  position{line: 76, col: 66, offset: 2314},
						name: "LITERAL",
					},
					&ruleRefExpr{
						pos:  position{line: 76, col: 76, offset: 2324},
						name: "UNARY",
					},
				},
			},
		},
		{
			name: "LITERAL",
			pos:  position{line: 78, col: 1, offset: 2331},
			expr: &choiceExpr{
				pos: position{line: 78, col: 11, offset: 2341},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 78, col: 11, offset: 2341},
						name: "STRING",
					},
					&ruleRefExpr{
						pos:  position{line: 78, col: 20, offset: 2350},
						name: "FLOAT",
					},
					&ruleRefExpr{
						pos:  position{line: 78, col: 28, offset: 2358},
						name: "BOOLEAN",
					},
					&ruleRefExpr{
						pos:  position{line: 78, col: 38, offset: 2368},
						name: "NULL",
					},
					&ruleRefExpr{
						pos:  position{line: 78, col: 45, offset: 2375},
						name: "INT",
					},
				},
			},
		},
		{
			name: "NEW",
			pos:  position{line: 80, col: 1, offset: 2380},
			expr: &seqExpr{
				pos: position{line: 80, col: 7, offset: 2386},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 80, col: 7, offset: 2386},
						name: "CLASS_REF_QUOTES",
					},
					&ruleRefExpr{
						pos:  position{line: 80, col: 24, offset: 2403},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 80, col: 26, offset: 2405},
						expr: &ruleRefExpr{
							pos:  position{line: 80, col: 26, offset: 2405},
							name: "ARGUMENTS",
						},
					},
				},
			},
		},
		{
			name: "BOOLEAN",
			pos:  position{line: 82, col: 1, offset: 2417},
			expr: &choiceExpr{
				pos: position{line: 82, col: 12, offset: 2428},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 82, col: 12, offset: 2428},
						val:        "true",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 82, col: 19, offset: 2435},
						val:        "false",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "NULL",
			pos:  position{line: 84, col: 1, offset: 2444},
			expr: &litMatcher{
				pos:        position{line: 84, col: 8, offset: 2451},
				val:        "null",
				ignoreCase: false,
			},
		},
		{
			name: "ARRAY",
			pos:  position{line: 86, col: 1, offset: 2459},
			expr: &seqExpr{
				pos: position{line: 86, col: 9, offset: 2467},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 86, col: 9, offset: 2467},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 86, col: 11, offset: 2469},
						val:        "[",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 86, col: 15, offset: 2473},
						expr: &ruleRefExpr{
							pos:  position{line: 86, col: 15, offset: 2473},
							name: "ARGUMENTLIST",
						},
					},
					&litMatcher{
						pos:        position{line: 86, col: 29, offset: 2487},
						val:        "]",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 86, col: 33, offset: 2491},
						name: "_",
					},
				},
			},
		},
		{
			name: "STRING",
			pos:  position{line: 88, col: 1, offset: 2494},
			expr: &seqExpr{
				pos: position{line: 88, col: 10, offset: 2503},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 88, col: 10, offset: 2503},
						val:        "\"",
						ignoreCase: false,
					},
					&zeroOrMoreExpr{
						pos: position{line: 88, col: 15, offset: 2508},
						expr: &charClassMatcher{
							pos:        position{line: 88, col: 15, offset: 2508},
							val:        "[a-zA-Z0-9]",
							ranges:     []rune{'a', 'z', 'A', 'Z', '0', '9'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&litMatcher{
						pos:        position{line: 88, col: 28, offset: 2521},
						val:        "\"",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "INT",
			pos:  position{line: 90, col: 1, offset: 2527},
			expr: &oneOrMoreExpr{
				pos: position{line: 90, col: 7, offset: 2533},
				expr: &charClassMatcher{
					pos:        position{line: 90, col: 7, offset: 2533},
					val:        "[0-9]",
					ranges:     []rune{'0', '9'},
					ignoreCase: false,
					inverted:   false,
				},
			},
		},
		{
			name: "FLOAT",
			pos:  position{line: 92, col: 1, offset: 2541},
			expr: &seqExpr{
				pos: position{line: 92, col: 9, offset: 2549},
				exprs: []interface{}{
					&oneOrMoreExpr{
						pos: position{line: 92, col: 9, offset: 2549},
						expr: &charClassMatcher{
							pos:        position{line: 92, col: 9, offset: 2549},
							val:        "[0-9]",
							ranges:     []rune{'0', '9'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&charClassMatcher{
						pos:        position{line: 92, col: 16, offset: 2556},
						val:        "[.]",
						chars:      []rune{'.'},
						ignoreCase: false,
						inverted:   false,
					},
					&oneOrMoreExpr{
						pos: position{line: 92, col: 20, offset: 2560},
						expr: &charClassMatcher{
							pos:        position{line: 92, col: 20, offset: 2560},
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
			pos:  position{line: 94, col: 1, offset: 2568},
			expr: &seqExpr{
				pos: position{line: 94, col: 15, offset: 2582},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 94, col: 15, offset: 2582},
						val:        "(",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 94, col: 19, offset: 2586},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 94, col: 21, offset: 2588},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 94, col: 32, offset: 2599},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 94, col: 34, offset: 2601},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "UNARY",
			pos:  position{line: 96, col: 1, offset: 2606},
			expr: &choiceExpr{
				pos: position{line: 96, col: 9, offset: 2614},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 96, col: 9, offset: 2614},
						name: "INCREMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 21, offset: 2626},
						name: "DECREMENT",
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 33, offset: 2638},
						name: "NEGATE",
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 42, offset: 2647},
						name: "NOT",
					},
					&ruleRefExpr{
						pos:  position{line: 96, col: 48, offset: 2653},
						name: "POSITIVE",
					},
				},
			},
		},
		{
			name: "INCREMENT",
			pos:  position{line: 98, col: 1, offset: 2663},
			expr: &seqExpr{
				pos: position{line: 98, col: 13, offset: 2675},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 98, col: 13, offset: 2675},
						name: "OBJECTACCESS",
					},
					&litMatcher{
						pos:        position{line: 98, col: 26, offset: 2688},
						val:        "++",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "DECREMENT",
			pos:  position{line: 100, col: 1, offset: 2694},
			expr: &seqExpr{
				pos: position{line: 100, col: 13, offset: 2706},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 100, col: 13, offset: 2706},
						name: "OBJECTACCESS",
					},
					&litMatcher{
						pos:        position{line: 100, col: 26, offset: 2719},
						val:        "--",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "NEGATE",
			pos:  position{line: 102, col: 1, offset: 2725},
			expr: &seqExpr{
				pos: position{line: 102, col: 10, offset: 2734},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 102, col: 10, offset: 2734},
						val:        "-",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 102, col: 14, offset: 2738},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "NOT",
			pos:  position{line: 104, col: 1, offset: 2752},
			expr: &seqExpr{
				pos: position{line: 104, col: 7, offset: 2758},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 104, col: 7, offset: 2758},
						val:        "!",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 104, col: 11, offset: 2762},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "POSITIVE",
			pos:  position{line: 106, col: 1, offset: 2776},
			expr: &seqExpr{
				pos: position{line: 106, col: 12, offset: 2787},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 106, col: 12, offset: 2787},
						val:        "+",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 106, col: 16, offset: 2791},
						name: "OBJECTACCESS",
					},
				},
			},
		},
		{
			name: "ARITHMETIC",
			pos:  position{line: 108, col: 1, offset: 2805},
			expr: &seqExpr{
				pos: position{line: 108, col: 14, offset: 2818},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 108, col: 14, offset: 2818},
						name: "ATOMIC",
					},
					&ruleRefExpr{
						pos:  position{line: 108, col: 21, offset: 2825},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 108, col: 23, offset: 2827},
						name: "OPERATOR",
					},
					&ruleRefExpr{
						pos:  position{line: 108, col: 32, offset: 2836},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 108, col: 34, offset: 2838},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "OPERATOR",
			pos:  position{line: 110, col: 1, offset: 2850},
			expr: &choiceExpr{
				pos: position{line: 110, col: 12, offset: 2861},
				alternatives: []interface{}{
					&litMatcher{
						pos:        position{line: 110, col: 12, offset: 2861},
						val:        "+",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 110, col: 18, offset: 2867},
						val:        "-",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 110, col: 24, offset: 2873},
						val:        "/",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 110, col: 30, offset: 2879},
						val:        "*",
						ignoreCase: false,
					},
					&litMatcher{
						pos:        position{line: 110, col: 36, offset: 2885},
						val:        "%",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "ASSIGNMENT",
			pos:  position{line: 112, col: 1, offset: 2890},
			expr: &seqExpr{
				pos: position{line: 112, col: 14, offset: 2903},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 112, col: 14, offset: 2903},
						name: "OBJECTACCESS",
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 27, offset: 2916},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 112, col: 29, offset: 2918},
						val:        "=",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 33, offset: 2922},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 112, col: 35, offset: 2924},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "LOGICAL",
			pos:  position{line: 114, col: 1, offset: 2936},
			expr: &seqExpr{
				pos: position{line: 114, col: 11, offset: 2946},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 114, col: 11, offset: 2946},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 114, col: 22, offset: 2957},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 114, col: 25, offset: 2960},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 114, col: 25, offset: 2960},
								val:        "and",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 114, col: 33, offset: 2968},
								val:        "or",
								ignoreCase: false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 114, col: 39, offset: 2974},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 114, col: 41, offset: 2976},
						name: "ATOMIC",
					},
				},
			},
		},
		{
			name: "COMPARISON",
			pos:  position{line: 116, col: 1, offset: 2984},
			expr: &seqExpr{
				pos: position{line: 116, col: 14, offset: 2997},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 116, col: 14, offset: 2997},
						name: "ATOMIC",
					},
					&ruleRefExpr{
						pos:  position{line: 116, col: 21, offset: 3004},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 116, col: 24, offset: 3007},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 116, col: 24, offset: 3007},
								val:        "===",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 116, col: 32, offset: 3015},
								val:        "!==",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 116, col: 40, offset: 3023},
								val:        "==",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 116, col: 47, offset: 3030},
								val:        "!=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 116, col: 54, offset: 3037},
								val:        "<=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 116, col: 61, offset: 3044},
								val:        ">=",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 116, col: 68, offset: 3051},
								val:        "<",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 116, col: 74, offset: 3057},
								val:        ">",
								ignoreCase: false,
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 116, col: 79, offset: 3062},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 116, col: 81, offset: 3064},
						name: "EXPRESSION",
					},
				},
			},
		},
		{
			name: "QUERY",
			pos:  position{line: 118, col: 1, offset: 3076},
			expr: &seqExpr{
				pos: position{line: 118, col: 9, offset: 3084},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 118, col: 9, offset: 3084},
						val:        "run",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 118, col: 16, offset: 3091},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 118, col: 18, offset: 3093},
						val:        "query",
						ignoreCase: true,
					},
					&ruleRefExpr{
						pos:  position{line: 118, col: 27, offset: 3102},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 118, col: 29, offset: 3104},
						name: "QUOTED_NAME",
					},
					&ruleRefExpr{
						pos:  position{line: 118, col: 41, offset: 3116},
						name: "_",
					},
					&zeroOrOneExpr{
						pos: position{line: 118, col: 43, offset: 3118},
						expr: &ruleRefExpr{
							pos:  position{line: 118, col: 43, offset: 3118},
							name: "ARGUMENTS",
						},
					},
				},
			},
		},
		{
			name: "OBJECTACCESS",
			pos:  position{line: 120, col: 1, offset: 3130},
			expr: &seqExpr{
				pos: position{line: 120, col: 16, offset: 3145},
				exprs: []interface{}{
					&zeroOrMoreExpr{
						pos: position{line: 120, col: 16, offset: 3145},
						expr: &seqExpr{
							pos: position{line: 120, col: 17, offset: 3146},
							exprs: []interface{}{
								&choiceExpr{
									pos: position{line: 120, col: 18, offset: 3147},
									alternatives: []interface{}{
										&ruleRefExpr{
											pos:  position{line: 120, col: 18, offset: 3147},
											name: "METHODCALL",
										},
										&ruleRefExpr{
											pos:  position{line: 120, col: 31, offset: 3160},
											name: "IDENTIFIER",
										},
									},
								},
								&litMatcher{
									pos:        position{line: 120, col: 43, offset: 3172},
									val:        "->",
									ignoreCase: false,
								},
							},
						},
					},
					&choiceExpr{
						pos: position{line: 120, col: 51, offset: 3180},
						alternatives: []interface{}{
							&ruleRefExpr{
								pos:  position{line: 120, col: 51, offset: 3180},
								name: "METHODCALL",
							},
							&ruleRefExpr{
								pos:  position{line: 120, col: 64, offset: 3193},
								name: "IDENTIFIER",
							},
						},
					},
				},
			},
		},
		{
			name: "METHODCALL",
			pos:  position{line: 122, col: 1, offset: 3206},
			expr: &seqExpr{
				pos: position{line: 122, col: 14, offset: 3219},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 122, col: 14, offset: 3219},
						name: "IDENTIFIER",
					},
					&ruleRefExpr{
						pos:  position{line: 122, col: 25, offset: 3230},
						name: "ARGUMENTS",
					},
				},
			},
		},
		{
			name: "ARGUMENTS",
			pos:  position{line: 124, col: 1, offset: 3241},
			expr: &seqExpr{
				pos: position{line: 124, col: 13, offset: 3253},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 124, col: 13, offset: 3253},
						val:        "(",
						ignoreCase: false,
					},
					&zeroOrOneExpr{
						pos: position{line: 124, col: 17, offset: 3257},
						expr: &ruleRefExpr{
							pos:  position{line: 124, col: 17, offset: 3257},
							name: "ARGUMENTLIST",
						},
					},
					&litMatcher{
						pos:        position{line: 124, col: 31, offset: 3271},
						val:        ")",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "ARGUMENTLIST",
			pos:  position{line: 126, col: 1, offset: 3276},
			expr: &seqExpr{
				pos: position{line: 126, col: 17, offset: 3292},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 126, col: 17, offset: 3292},
						name: "_",
					},
					&zeroOrMoreExpr{
						pos: position{line: 126, col: 19, offset: 3294},
						expr: &seqExpr{
							pos: position{line: 126, col: 20, offset: 3295},
							exprs: []interface{}{
								&ruleRefExpr{
									pos:  position{line: 126, col: 20, offset: 3295},
									name: "_",
								},
								&ruleRefExpr{
									pos:  position{line: 126, col: 22, offset: 3297},
									name: "EXPRESSION",
								},
								&ruleRefExpr{
									pos:  position{line: 126, col: 33, offset: 3308},
									name: "_",
								},
								&litMatcher{
									pos:        position{line: 126, col: 35, offset: 3310},
									val:        ",",
									ignoreCase: false,
								},
								&ruleRefExpr{
									pos:  position{line: 126, col: 39, offset: 3314},
									name: "_",
								},
							},
						},
					},
					&ruleRefExpr{
						pos:  position{line: 126, col: 43, offset: 3318},
						name: "EXPRESSION",
					},
					&ruleRefExpr{
						pos:  position{line: 126, col: 54, offset: 3329},
						name: "_",
					},
				},
			},
		},
		{
			name: "CLASS_REF_QUOTES",
			pos:  position{line: 133, col: 1, offset: 3497},
			expr: &seqExpr{
				pos: position{line: 133, col: 20, offset: 3516},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 133, col: 20, offset: 3516},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 133, col: 22, offset: 3518},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 133, col: 26, offset: 3522},
						name: "CLASS_REF",
					},
					&litMatcher{
						pos:        position{line: 133, col: 36, offset: 3532},
						val:        "'",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "CLASS_REF",
			pos:  position{line: 135, col: 1, offset: 3537},
			expr: &seqExpr{
				pos: position{line: 135, col: 13, offset: 3549},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 135, col: 13, offset: 3549},
						name: "_",
					},
					&ruleRefExpr{
						pos:  position{line: 135, col: 15, offset: 3551},
						name: "CLASS_TYPE",
					},
					&litMatcher{
						pos:        position{line: 135, col: 26, offset: 3562},
						val:        "\\",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 135, col: 31, offset: 3567},
						name: "CLASS_NAME",
					},
				},
			},
		},
		{
			name: "CLASS_TYPE",
			pos:  position{line: 137, col: 1, offset: 3579},
			expr: &seqExpr{
				pos: position{line: 137, col: 14, offset: 3592},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 137, col: 14, offset: 3592},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 137, col: 17, offset: 3595},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 137, col: 17, offset: 3595},
								val:        "value",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 137, col: 27, offset: 3605},
								val:        "entity",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 137, col: 38, offset: 3616},
								val:        "command",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 137, col: 50, offset: 3628},
								val:        "event",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 137, col: 60, offset: 3638},
								val:        "projection",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 137, col: 75, offset: 3653},
								val:        "invariant",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 137, col: 89, offset: 3667},
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
			pos:  position{line: 139, col: 1, offset: 3677},
			expr: &seqExpr{
				pos: position{line: 139, col: 14, offset: 3690},
				exprs: []interface{}{
					&oneOrMoreExpr{
						pos: position{line: 139, col: 14, offset: 3690},
						expr: &charClassMatcher{
							pos:        position{line: 139, col: 14, offset: 3690},
							val:        "[a-zA-Z]",
							ranges:     []rune{'a', 'z', 'A', 'Z'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 139, col: 24, offset: 3700},
						expr: &charClassMatcher{
							pos:        position{line: 139, col: 24, offset: 3700},
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
			pos:  position{line: 141, col: 1, offset: 3716},
			expr: &seqExpr{
				pos: position{line: 141, col: 15, offset: 3730},
				exprs: []interface{}{
					&litMatcher{
						pos:        position{line: 141, col: 15, offset: 3730},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 141, col: 19, offset: 3734},
						name: "CLASS_NAME",
					},
					&litMatcher{
						pos:        position{line: 141, col: 30, offset: 3745},
						val:        "'",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "TYPE",
			pos:  position{line: 143, col: 1, offset: 3750},
			expr: &choiceExpr{
				pos: position{line: 143, col: 8, offset: 3757},
				alternatives: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 143, col: 8, offset: 3757},
						name: "CLASS_REF",
					},
					&ruleRefExpr{
						pos:  position{line: 143, col: 20, offset: 3769},
						name: "VALUE_TYPE",
					},
				},
			},
		},
		{
			name: "VALUE_TYPE",
			pos:  position{line: 145, col: 1, offset: 3781},
			expr: &seqExpr{
				pos: position{line: 145, col: 14, offset: 3794},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 145, col: 14, offset: 3794},
						name: "_",
					},
					&choiceExpr{
						pos: position{line: 145, col: 17, offset: 3797},
						alternatives: []interface{}{
							&litMatcher{
								pos:        position{line: 145, col: 17, offset: 3797},
								val:        "string",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 145, col: 28, offset: 3808},
								val:        "boolean",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 145, col: 40, offset: 3820},
								val:        "float",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 145, col: 50, offset: 3830},
								val:        "map",
								ignoreCase: false,
							},
							&litMatcher{
								pos:        position{line: 145, col: 58, offset: 3838},
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
			pos:  position{line: 147, col: 1, offset: 3848},
			expr: &seqExpr{
				pos: position{line: 147, col: 21, offset: 3868},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 147, col: 21, offset: 3868},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 147, col: 23, offset: 3870},
						val:        "'",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 147, col: 27, offset: 3874},
						name: "CLASS_NAME",
					},
					&litMatcher{
						pos:        position{line: 147, col: 38, offset: 3885},
						val:        "'",
						ignoreCase: false,
					},
				},
			},
		},
		{
			name: "IDENTIFIER",
			pos:  position{line: 149, col: 1, offset: 3890},
			expr: &seqExpr{
				pos: position{line: 149, col: 14, offset: 3903},
				exprs: []interface{}{
					&oneOrMoreExpr{
						pos: position{line: 149, col: 14, offset: 3903},
						expr: &charClassMatcher{
							pos:        position{line: 149, col: 14, offset: 3903},
							val:        "[a-zA-Z]",
							ranges:     []rune{'a', 'z', 'A', 'Z'},
							ignoreCase: false,
							inverted:   false,
						},
					},
					&zeroOrMoreExpr{
						pos: position{line: 149, col: 24, offset: 3913},
						expr: &charClassMatcher{
							pos:        position{line: 149, col: 24, offset: 3913},
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
			pos:  position{line: 151, col: 1, offset: 3928},
			expr: &zeroOrMoreExpr{
				pos: position{line: 151, col: 5, offset: 3932},
				expr: &choiceExpr{
					pos: position{line: 151, col: 7, offset: 3934},
					alternatives: []interface{}{
						&ruleRefExpr{
							pos:  position{line: 151, col: 7, offset: 3934},
							name: "WHITESPACE",
						},
						&ruleRefExpr{
							pos:  position{line: 151, col: 20, offset: 3947},
							name: "EOL",
						},
					},
				},
			},
		},
		{
			name: "SEMI",
			pos:  position{line: 153, col: 1, offset: 3955},
			expr: &seqExpr{
				pos: position{line: 153, col: 8, offset: 3962},
				exprs: []interface{}{
					&ruleRefExpr{
						pos:  position{line: 153, col: 8, offset: 3962},
						name: "_",
					},
					&litMatcher{
						pos:        position{line: 153, col: 10, offset: 3964},
						val:        ";",
						ignoreCase: false,
					},
					&ruleRefExpr{
						pos:  position{line: 153, col: 14, offset: 3968},
						name: "_",
					},
				},
			},
		},
		{
			name: "WHITESPACE",
			pos:  position{line: 155, col: 1, offset: 3971},
			expr: &charClassMatcher{
				pos:        position{line: 155, col: 14, offset: 3984},
				val:        "[ \\t\\r]",
				chars:      []rune{' ', '\t', '\r'},
				ignoreCase: false,
				inverted:   false,
			},
		},
		{
			name: "EOL",
			pos:  position{line: 157, col: 1, offset: 3993},
			expr: &litMatcher{
				pos:        position{line: 157, col: 7, offset: 3999},
				val:        "\n",
				ignoreCase: false,
			},
		},
		{
			name: "EOF",
			pos:  position{line: 159, col: 1, offset: 4005},
			expr: &notExpr{
				pos: position{line: 159, col: 7, offset: 4011},
				expr: &anyMatcher{
					line: 159, col: 8, offset: 4012,
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
