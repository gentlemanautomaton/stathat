package statmatch

import "github.com/gentlemanautomaton/stathat"

// Any returns a filter that returns true when any filter returns true.
func Any(filters ...stathat.Filter) stathat.Filter {
	return func(stat stathat.StatItem) bool {
		for _, match := range filters {
			if match(stat) {
				return true
			}
		}
		return false
	}
}
