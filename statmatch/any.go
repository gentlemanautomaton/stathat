package statmatch

import "github.com/gentlemanautomaton/stathat"

// Any returns a filter that returns true when any filter returns true.
//
// Any returns true if the set of filters is empty.
func Any(filters ...stathat.Filter) stathat.Filter {
	return func(stat stathat.StatItem) bool {
		if len(filters) == 0 {
			return true
		}
		for _, match := range filters {
			if match(stat) {
				return true
			}
		}
		return false
	}
}
