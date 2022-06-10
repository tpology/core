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
