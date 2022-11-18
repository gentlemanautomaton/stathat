package stathat

// Filter returns true if it matches a stat.
type Filter func(StatItem) bool
