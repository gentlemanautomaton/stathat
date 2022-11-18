package statmatch

import "github.com/gentlemanautomaton/stathat"

// None returns a filter that returns true when none of the filters returns true.
//
// None returns true if the set of filters is empty.
func None(filters ...stathat.Filter) stathat.Filter {
	return func(stat stathat.StatItem) bool {
		for _, match := range filters {
			if match(stat) {
				return false
			}
		}
		return true
	}
}
