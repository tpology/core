package core

import (
	"reflect"
	"testing"
)

type matchAndExtractTokensTestCase struct {
	name           string
	pattern        []string
	filename       []string
	expectedMatch  bool
	expectedTokens map[string]string
}

var matchAndExtractTokensTestCases = []matchAndExtractTokensTestCase{
	// Both empty
	{
		name:           "both empty",
		pattern:        []string{},
		filename:       []string{},
		expectedMatch:  true,
		expectedTokens: map[string]string{},
	},
	// Exhausted pattern but not filename
	{
		name:           "exhausted pattern but not filename",
		pattern:        []string{},
		filename:       []string{"foo"},
		expectedMatch:  false,
		expectedTokens: map[string]string{},
	},
	// Exhausted filename but not pattern
	{
		name:           "exhausted filename but not pattern",
		pattern:        []string{"foo"},
		filename:       []string{},
		expectedMatch:  false,
		expectedTokens: map[string]string{},
	},
	// Pattern is a token
	{
		name:           "pattern is a token",
		pattern:        []string{"{foo}"},
		filename:       []string{"bar"},
		expectedMatch:  true,
		expectedTokens: map[string]string{"foo": "bar"},
	},
	// Pattern is two tokens
	{
		name:           "pattern is two tokens",
		pattern:        []string{"{foo}", "{bar}"},
		filename:       []string{"bar", "baz"},
		expectedMatch:  true,
		expectedTokens: map[string]string{"foo": "bar", "bar": "baz"},
	},
	// Pattern is two tokens separated by a non-token
	{
		name:           "pattern is two tokens separated by a non-token",
		pattern:        []string{"{foo}", "bar", "{baz}"},
		filename:       []string{"foo", "bar", "baz"},
		expectedMatch:  true,
		expectedTokens: map[string]string{"foo": "foo", "baz": "baz"},
	},
	// Pattern is a non-token, a token, and a non-token
	{
		name:           "pattern is a non-token, a token, and a non-token",
		pattern:        []string{"foo", "{bar}", "baz"},
		filename:       []string{"foo", "bar", "baz"},
		expectedMatch:  true,
		expectedTokens: map[string]string{"bar": "bar"},
	},
	// A '**' matches three levels of the filename
	{
		name:           "a '**' matches three levels of the filename",
		pattern:        []string{"**"},
		filename:       []string{"foo", "bar", "baz"},
		expectedMatch:  true,
		expectedTokens: map[string]string{},
	},
	// A '*' does not match three levels of the filename
	{
		name:           "a '*' does not match three levels of the filename",
		pattern:        []string{"*"},
		filename:       []string{"foo", "bar", "baz"},
		expectedMatch:  false,
		expectedTokens: map[string]string{},
	},
	// A '*' x '*' matches the filename
	{
		name:           "a '*' bar '*' matches the filename",
		pattern:        []string{"*", "bar", "*"},
		filename:       []string{"foo", "bar", "baz"},
		expectedMatch:  true,
		expectedTokens: map[string]string{},
	},
	// '*' '*' '*' matches the filename
	{
		name:           "'*' '*' '*' matches the filename",
		pattern:        []string{"*", "*", "*"},
		filename:       []string{"foo", "bar", "baz"},
		expectedMatch:  true,
		expectedTokens: map[string]string{},
	},
	// '**' against empty
	{
		name:           "'**' against empty",
		pattern:        []string{"**"},
		filename:       []string{},
		expectedMatch:  true,
		expectedTokens: map[string]string{},
	},
	// foo, '**' against foo
	{
		name:           "foo, '**' against foo",
		pattern:        []string{"foo", "**"},
		filename:       []string{"foo"},
		expectedMatch:  true,
		expectedTokens: map[string]string{},
	},
	// some more realistic cases
	{
		name:           "some more realistic cases",
		pattern:        []string{"resources", "{team}", "{project}", "**", "baz"},
		filename:       []string{"resources", "team1", "project1", "foo", "bar", "baz"},
		expectedMatch:  true,
		expectedTokens: map[string]string{"team": "team1", "project": "project1"},
	},
	// showing greediness of **
	{
		name:           "some more realistic cases",
		pattern:        []string{"resources", "{team}", "**", "{project}", "baz"},
		filename:       []string{"resources", "team1", "foo", "bar", "project", "baz"},
		expectedMatch:  true,
		expectedTokens: map[string]string{"team": "team1", "project": "project"},
	},
	// foo/*/bar doesn't match foo/a/b/bar
	{
		name:           "foo/*/bar doesn't match foo/a/b/bar",
		pattern:        []string{"foo", "*", "bar"},
		filename:       []string{"foo", "a", "b", "bar"},
		expectedMatch:  false,
		expectedTokens: map[string]string{},
	},
	// But foo/**/bar matches foo/a/b/bar
	{
		name:           "But foo/**/bar matches foo/a/b/bar",
		pattern:        []string{"foo", "**", "bar"},
		filename:       []string{"foo", "a", "b", "bar"},
		expectedMatch:  true,
		expectedTokens: map[string]string{},
	},
	// Ambiguous case
	{
		name:    "Ambiguous case",
		pattern: []string{"**", "{bar}", "**"},
		// ten components in filename
		filename:       []string{"foo", "bar", "baz", "qux", "quux", "corge", "grault", "garply", "waldo", "bar"},
		expectedMatch:  true,
		expectedTokens: map[string]string{"bar": "bar"},
	},
	// Another ambiguous case
	{
		name:    "Another ambiguous case",
		pattern: []string{"**", "{bar}"},
		// ten components in filename
		filename:       []string{"foo", "bar", "baz", "qux", "quux", "corge", "grault", "garply", "waldo", "bar"},
		expectedMatch:  true,
		expectedTokens: map[string]string{"bar": "bar"},
	},
	// Yet another ambiguous case
	{
		name:    "Yet another ambiguous case",
		pattern: []string{"{bar}", "**"},
		// ten components in filename
		filename:       []string{"foo", "bar", "baz", "qux", "quux", "corge", "grault", "garply", "waldo", "bar"},
		expectedMatch:  true,
		expectedTokens: map[string]string{"bar": "foo"},
	},
}

func Test_matchAndExtractTokens(t *testing.T) {
	for _, testCase := range matchAndExtractTokensTestCases {
		match, tokens := matchAndExtractTokens(testCase.pattern, testCase.filename)
		if match != testCase.expectedMatch {
			t.Errorf("%s: match mismatch: expected %t, got %t", testCase.name, testCase.expectedMatch, match)
		}
		if len(tokens) != len(testCase.expectedTokens) {
			t.Errorf("%s: tokens length mismatch: expected %d, got %d", testCase.name, len(testCase.expectedTokens), len(tokens))
		}
		for tokenName, tokenValue := range testCase.expectedTokens {
			if tokens[tokenName] != tokenValue {
				t.Errorf("%s: token value mismatch: expected %s, got %s", testCase.name, tokenValue, tokens[tokenName])
			}
		}
	}
}

var splitPathTestCases = []struct {
	name           string
	path           string
	expectedResult []string
}{
	{
		name:           "empty path",
		path:           "",
		expectedResult: []string{},
	},
	{
		name:           "single component",
		path:           "foo",
		expectedResult: []string{"foo"},
	},
	{
		name:           "multiple components",
		path:           "foo/bar/baz",
		expectedResult: []string{"foo", "bar", "baz"},
	},
	{
		name:           "multiple components with trailing slash",
		path:           "foo/bar/baz/",
		expectedResult: []string{"foo", "bar", "baz"},
	},
	{
		name:           "multiple components with leading slash",
		path:           "/foo/bar/baz",
		expectedResult: []string{"foo", "bar", "baz"},
	},
	{
		name:           "multiple components with leading and trailing slash",
		path:           "/foo/bar/baz/",
		expectedResult: []string{"foo", "bar", "baz"},
	},
	{
		name:           "repeated slashes",
		path:           "//foo//bar//baz",
		expectedResult: []string{"foo", "bar", "baz"},
	},
}

func Test_splitPathTestCases(t *testing.T) {
	for _, testCase := range splitPathTestCases {
		result := splitPath(testCase.path)
		if !reflect.DeepEqual(result, testCase.expectedResult) {
			t.Errorf("%s: result mismatch: expected %v, got %v", testCase.name, testCase.expectedResult, result)
		}
	}
}

type isSubsetTestCase struct {
	name           string
	subset         interface{}
	set            interface{}
	expectedResult bool
}

var isSubsetTestCases = []isSubsetTestCase{
	{
		name:           "empty subset",
		subset:         map[string]interface{}{},
		set:            map[string]interface{}{"foo": "bar"},
		expectedResult: true,
	},
	{
		name:           "empty set",
		subset:         map[string]interface{}{"foo": "bar"},
		set:            map[string]interface{}{},
		expectedResult: false,
	},
	{
		name:           "subset is equal to set",
		subset:         map[string]interface{}{"foo": "bar"},
		set:            map[string]interface{}{"foo": "bar"},
		expectedResult: true,
	},
	{
		name:           "subset is a subset of set",
		subset:         map[string]interface{}{"foo": "bar"},
		set:            map[string]interface{}{"foo": "bar", "baz": "qux"},
		expectedResult: true,
	},
	{
		name:           "subset is not a subset of set",
		subset:         map[string]interface{}{"foo": "bar"},
		set:            map[string]interface{}{"baz": "qux", "quux": "corge"},
		expectedResult: false,
	},
	{
		name:           "subset is a subset of set with different values",
		subset:         map[string]interface{}{"foo": "bar"},
		set:            map[string]interface{}{"foo": "baz", "baz": "qux"},
		expectedResult: false,
	},
	{
		name: "nested set",
		subset: map[string]interface{}{
			"foo": map[string]interface{}{
				"bar": "baz",
			},
		},
		set: map[string]interface{}{
			"foo": map[string]interface{}{
				"bar": "baz",
			},
			"baz": "qux",
		},
		expectedResult: true,
	},
	{
		name: "unrelated",
		subset: map[string]interface{}{
			"foo": map[string]interface{}{
				"bar": "baz",
			},
		},
		set: map[string]interface{}{
			"baz": "qux",
		},
		expectedResult: false,
	},
	{
		name: "two slices",
		subset: []interface{}{
			map[string]interface{}{
				"foo": "bar",
			},
			map[string]interface{}{
				"baz": "qux",
			},
		},
		set: []interface{}{
			map[string]interface{}{
				"foo": "bar",
			},
			map[string]interface{}{
				"baz": "qux",
			},
		},
		expectedResult: true,
	},
	{
		name: "two slices subset",
		subset: []interface{}{
			map[string]interface{}{
				"foo": "bar",
			},
			map[string]interface{}{
				"baz": "qux",
			},
		},
		set: []interface{}{
			map[string]interface{}{
				"foo": "bar",
				"baz": "qux",
			},
			map[string]interface{}{
				"baz":  "qux",
				"quux": "corge",
			},
		},
		expectedResult: true,
	},
	{
		name: "two slices different length",
		subset: []interface{}{
			map[string]interface{}{
				"foo": "bar",
			},
			map[string]interface{}{
				"baz": "qux",
			},
		},
		set: []interface{}{
			map[string]interface{}{
				"foo": "bar",
			},
			map[string]interface{}{
				"baz": "qux",
			},
			map[string]interface{}{
				"quux": "corge",
			},
		},
		expectedResult: false,
	},
	{
		name: "map vs slice",
		subset: map[string]interface{}{
			"foo": "bar",
		},
		set: []interface{}{
			map[string]interface{}{
				"foo": "bar",
			},
		},
		expectedResult: false,
	},
	{
		name: "slice vs map",
		subset: []interface{}{
			map[string]interface{}{
				"foo": "bar",
			},
		},
		set: map[string]interface{}{
			"foo": "bar",
		},
		expectedResult: false,
	},
	{
		name: "slices not subsets",
		subset: []interface{}{
			map[string]interface{}{
				"foo": "baz",
			},
		},
		set: []interface{}{
			map[string]interface{}{
				"foo": "bar",
			},
		},
		expectedResult: false,
	},
	{
		name: "map interface/interface not subsets",
		subset: map[interface{}]interface{}{
			"foo": "baz",
		},
		set: map[interface{}]interface{}{
			"foo": "bar",
		},
		expectedResult: false,
	},
	{
		name: "map interface/interface subset",
		subset: map[interface{}]interface{}{
			"foo": "baz",
		},
		set: map[interface{}]interface{}{
			"foo": "baz",
		},
		expectedResult: true,
	},
	{
		name: "map interface/interface not slice subset",
		subset: map[interface{}]interface{}{
			"foo": "baz",
		},
		set: []interface{}{
			map[string]interface{}{
				"foo": "bar",
			},
		},
		expectedResult: false,
	},
	{
		name: "three levels subset",
		subset: map[string]interface{}{
			"foo": map[string]interface{}{
				"bar": map[string]interface{}{
					"baz": "qux",
				},
			},
		},
		set: map[string]interface{}{
			"foo": map[string]interface{}{
				"bar": map[string]interface{}{
					"baz": "qux",
				},
			},
			"baz": "qux",
		},
		expectedResult: true,
	},
	{
		name: "three levels not subset",
		subset: map[string]interface{}{
			"foo": map[string]interface{}{
				"bar": map[string]interface{}{
					"baz": "quuz",
				},
			},
		},
		set: map[string]interface{}{
			"foo": map[string]interface{}{
				"bar": map[string]interface{}{
					"baz": "qux",
				},
			},
			"baz": "qux",
		},
		expectedResult: false,
	},
}

// Test_isSubset tests isSubset function
func Test_isSubset(t *testing.T) {
	for _, testCase := range isSubsetTestCases {
		result := isSubset(testCase.subset, testCase.set)
		if result != testCase.expectedResult {
			t.Errorf("%s: result mismatch: expected %t, got %t", testCase.name, testCase.expectedResult, result)
		}
	}
}

// injectTestCase is a test case for injectTokenValues
type injectTestCase struct {
	name     string
	in       interface{}
	tokens   map[string]string
	expected interface{}
}

// injectTestCases is a list of injectTestCase
var injectTestCases = []injectTestCase{
	{
		name: "no tokens",
		in: map[string]interface{}{
			"foo": "bar",
		},
		tokens: map[string]string{
			"foo": "bar",
		},
		expected: map[string]interface{}{
			"foo": "bar",
		},
	},
	{
		name: "one token in value",
		in: map[string]interface{}{
			"foo": "{foo}",
		},
		tokens: map[string]string{
			"foo": "baz",
		},
		expected: map[string]interface{}{
			"foo": "baz",
		},
	},
	{
		name: "one token in key",
		in: map[string]interface{}{
			"{foo}": "bar",
		},
		tokens: map[string]string{
			"foo": "baz",
		},
		expected: map[string]interface{}{
			"baz": "bar",
		},
	},
	{
		name: "two tokens in value",
		in: map[string]interface{}{
			"foo": "{foo} {bar}",
		},
		tokens: map[string]string{
			"foo": "baz",
			"bar": "qux",
		},
		expected: map[string]interface{}{
			"foo": "baz qux",
		},
	},
	{
		name: "two tokens in key",
		in: map[string]interface{}{
			"{foo} {bar}": "bar",
		},
		tokens: map[string]string{
			"foo": "baz",
			"bar": "qux",
		},
		expected: map[string]interface{}{
			"baz qux": "bar",
		},
	},
	{
		name: "two tokens in key and value",
		in: map[string]interface{}{
			"{foo} {bar}": "{foo} {bar}",
		},
		tokens: map[string]string{
			"foo": "baz",
			"bar": "qux",
		},
		expected: map[string]interface{}{
			"baz qux": "baz qux",
		},
	},
	{
		name: "two tokens in key and value with different order",
		in: map[string]interface{}{
			"{foo} {bar}": "{bar} {foo}",
		},
		tokens: map[string]string{
			"foo": "baz",
			"bar": "qux",
		},
		expected: map[string]interface{}{
			"baz qux": "qux baz",
		},
	},
	{
		name: "two slices",
		in: map[string]interface{}{
			"foo": []interface{}{
				"{foo}",
				"{bar}",
			},
			"bar": []interface{}{
				"{foo}",
				"{bar}",
			},
		},
		tokens: map[string]string{
			"foo": "baz",
			"bar": "qux",
		},
		expected: map[string]interface{}{
			"foo": []interface{}{
				"baz",
				"qux",
			},
			"bar": []interface{}{
				"baz",
				"qux",
			},
		},
	},
	{
		name: "two slices with different order",
		in: map[string]interface{}{
			"foo": []interface{}{
				"{foo}",
				"{bar}",
			},
			"bar": []interface{}{
				"{bar}",
				"{foo}",
			},
		},
		tokens: map[string]string{
			"foo": "baz",
			"bar": "qux",
		},
		expected: map[string]interface{}{
			"foo": []interface{}{
				"baz",
				"qux",
			},
			"bar": []interface{}{
				"qux",
				"baz",
			},
		},
	},
	{
		name: "two slices with different order and different length",
		in: map[string]interface{}{
			"foo": []interface{}{
				"{foo}",
				"{bar}",
			},
			"bar": []interface{}{
				"{bar}",
				"{foo}",
				"{baz}",
			},
		},
		tokens: map[string]string{
			"foo": "baz",
			"bar": "qux",
			"baz": "quux",
		},
		expected: map[string]interface{}{
			"foo": []interface{}{
				"baz",
				"qux",
			},
			"bar": []interface{}{
				"qux",
				"baz",
				"quux",
			},
		},
	},
	{
		name: "example of a resource with injections in the labels and annotations",
		in: map[string]interface{}{
			"labels": map[string]interface{}{
				"team":            "{team}",
				"owned-by-{team}": "true",
			},
			"annotations": map[string]interface{}{
				"team":            "{team}",
				"owned-by-{team}": "true",
			},
		},
		tokens: map[string]string{
			"team": "foo",
		},
		expected: map[string]interface{}{
			"labels": map[string]interface{}{
				"team":         "foo",
				"owned-by-foo": "true",
			},
			"annotations": map[string]interface{}{
				"team":         "foo",
				"owned-by-foo": "true",
			},
		},
	},
}

// Test_injectTestCases tests injectTokenValues function
func Test_injectTestCases(t *testing.T) {
	for _, testCase := range injectTestCases {
		result := injectTokenValues(testCase.in, testCase.tokens)
		// Deep comparison
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("%s: result mismatch: expected %s, got %s", testCase.name, testCase.expected, result)
		}
	}
}

// stringInjectTestCase is a test case for simple inject
type stringInjectTestCase struct {
	name     string
	in       string
	tokens   map[string]string
	expected string
}

var stringInjectTestCases = []stringInjectTestCase{
	{
		name: "one token",
		in:   "foo {bar}",
		tokens: map[string]string{
			"bar": "baz",
		},
		expected: "foo baz",
	},
	{
		name: "two tokens",
		in:   "foo {bar} {qux}",
		tokens: map[string]string{
			"bar": "baz",
			"qux": "quux",
		},
		expected: "foo baz quux",
	},
	{
		name: "two tokens in different order",
		in:   "foo {bar} {qux}",
		tokens: map[string]string{
			"qux": "quux",
			"bar": "baz",
		},
		expected: "foo baz quux",
	},
	{
		name: "same token twice",
		in:   "foo {bar} {bar}",
		tokens: map[string]string{
			"bar": "baz",
		},
		expected: "foo baz baz",
	},
	{
		name: "unknown token",
		in:   "foo {bar}",
		tokens: map[string]string{
			"qux": "quux",
		},
		expected: "foo {bar}",
	},
	{
		name: "token in double curly braces",
		in:   "foo {{bar}}",
		tokens: map[string]string{
			"bar": "baz",
		},
		expected: "foo {baz}",
	},
}

// Test_stringInjectTestCases tests injectTokenValues function
func Test_stringInjectTestCases(t *testing.T) {
	for _, testCase := range stringInjectTestCases {
		result := inject(testCase.in, testCase.tokens)
		// Deep comparison
		if result != testCase.expected {
			t.Errorf("%s: result mismatch: expected %s, got %s", testCase.name, testCase.expected, result)
		}
	}
}
