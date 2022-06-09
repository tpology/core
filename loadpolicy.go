package core

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
	// Is the pattern and filename the same?
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
