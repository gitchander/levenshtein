package levenshtein

//------------------------------------------------------------------------------

// Convenience helper functions

func RuneSlices(a, b []rune) int {
	return Distance(RunePair{a, b})
}

func StringSlices(a, b []string) int {
	return Distance(StringPair{a, b})
}

func BoolSlices(a, b []bool) int {
	return Distance(BoolPair{a, b})
}

//------------------------------------------------------------------------------

func Strings(a, b string) int {
	return RuneSlices([]rune(a), []rune(b))
}

//------------------------------------------------------------------------------
