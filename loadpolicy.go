package core

import (
	"strings"
)

// matchAndExtractTokens attempts to match the pattern slice with the
// filename slice, extracting tokens as it goes.
func matchAndExtractTokens(pattern []string, filename []string) (bool, map[string]string) {
	// Exhausted both pattern and filename?
	if len(pattern) == 0 && len(filename) == 0 {
		return true, map[string]string{}
	}
	// Exhausted pattern but not filename?
	if len(pattern) == 0 {
		return false, map[string]string{}
	}
	// Exhausted filename but not pattern?
	if len(filename) == 0 {
		// We can consume '**' from the pattern
		if pattern[0] == "**" {
			return matchAndExtractTokens(pattern[1:], filename)
		}
		return false, map[string]string{}
	}
	tokenName := ""
	tokenValue := ""
	wildcard := ""
	// Is the pattern a token?
	if pattern[0][0] == '{' && pattern[0][len(pattern[0])-1] == '}' {
		// Extract the token name
		tokenName = pattern[0][1 : len(pattern[0])-1]
		// Extract the token value
		tokenValue = filename[0]
		// Treat the token as the '*' wildcard
		wildcard = "*"
	} else if pattern[0] == "*" {
		wildcard = "*"
	} else if pattern[0] == "**" {
		wildcard = "**"
	}
	if wildcard == "*" {
		remainderMatch, remainderTokens := matchAndExtractTokens(pattern[1:], filename[1:])
		if remainderMatch {
			if tokenName != "" {
				// Add our token to the remainder tokens
				remainderTokens[tokenName] = tokenValue
			}
			return true, remainderTokens
		}
		return false, map[string]string{}
	} else if wildcard == "**" {
		remainderMatch, remainderTokens := matchAndExtractTokens(pattern[1:], filename[1:])
		if remainderMatch {
			return true, remainderTokens
		}
		return matchAndExtractTokens(pattern, filename[1:])
	}
	// Are the pattern and filename the same?
	if pattern[0] == filename[0] {
		// Match the rest of the pattern and filename
		remainderMatch, remainderTokens := matchAndExtractTokens(pattern[1:], filename[1:])
		if remainderMatch {
			return true, remainderTokens
		}
		return false, map[string]string{}
	}
	return false, map[string]string{}
}

// splitPath splits the path
func splitPath(path string) []string {
	result := []string{}
	for _, part := range strings.Split(path, "/") {
		if part != "" {
			result = append(result, part)
		}
	}
	return result
}

// isSubset return if the obj a is a subset of the obj b
func isSubset(a, b interface{}) bool {
	switch aval := a.(type) {
	case map[interface{}]interface{}:
		bval, ok := b.(map[interface{}]interface{})
		if !ok {
			return false
		}
		for k, v := range aval {
			if bv, ok := bval[k]; !ok || !isSubset(v, bv) {
				return false
			}
		}
		return true
	case map[string]interface{}:
		bval, ok := b.(map[string]interface{})
		if !ok {
			return false
		}
		for k, v := range aval {
			if bv, ok := bval[k]; !ok || !isSubset(v, bv) {
				return false
			}
		}
		return true
	case []interface{}:
		bval, ok := b.([]interface{})
		if !ok {
			return false
		}
		// must be same length, and corresponding elements must be subsets
		if len(aval) != len(bval) {
			return false
		}
		for i, v := range aval {
			if !isSubset(v, bval[i]) {
				return false
			}
		}
		return true
	default:
		return a == b
	}
}

// inject injects the tokens into a string
func inject(s string, tokens map[string]string) string {
	for k, v := range tokens {
		s = strings.Replace(s, "{"+k+"}", v, -1)
	}
	return s
}

// injectTokenValues injects the token values into the object
func injectTokenValues(obj interface{}, tokens map[string]string) interface{} {
	switch o := obj.(type) {
	case map[interface{}]interface{}:
		for k, v := range o {
			// Inject into v
			v = injectTokenValues(v, tokens)
			o[k] = v
			// if k is a string, try to inject into it
			if ks, ok := k.(string); ok {
				ks = inject(ks, tokens)
				if ks != k {
					o[ks] = v
					delete(o, k)
				}
			}
		}
		return o
	case map[string]interface{}:
		for k, v := range o {
			// Inject into v
			v = injectTokenValues(v, tokens)
			o[k] = v
			// inject into k
			ks := inject(k, tokens)
			if ks != k {
				o[ks] = v
				delete(o, k)
			}
		}
		return o
	case []interface{}:
		for i, v := range o {
			// Inject into v
			v = injectTokenValues(v, tokens)
			o[i] = v
		}
		return o
	case string:
		return inject(o, tokens)
	default:
		return obj
	}
}
