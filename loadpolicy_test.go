package core

import (
	"fmt"
	"testing"
)

func breakf() {
	fmt.Println("Breaking")
}

type matchAndExtractTokensTestCase struct {
	breakf         func()
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
		if testCase.breakf != nil {
			testCase.breakf()
		}
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
