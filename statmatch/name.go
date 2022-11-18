package statmatch

import "github.com/gentlemanautomaton/stathat"

// Name returns a filter that matches stat names.
func Name(matcher StringMatcher) stathat.Filter {
	return func(stat stathat.StatItem) bool {
		return matcher.MatchString(stat.Name)
	}
}
