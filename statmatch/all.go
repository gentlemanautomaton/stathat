package statmatch

import "github.com/gentlemanautomaton/stathat"

// All returns a filter that returns true when all filters return true.
func All(filters ...stathat.Filter) stathat.Filter {
	return func(stat stathat.StatItem) bool {
		for _, match := range filters {
			if !match(stat) {
				return false
			}
		}
		return true
	}
}
