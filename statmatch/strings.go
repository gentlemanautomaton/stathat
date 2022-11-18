package statmatch

// StringMatcher is a function capable of matching strings.
type StringMatcher interface {
	MatchString(string) bool
}
