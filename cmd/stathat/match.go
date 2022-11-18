package main

import (
	"github.com/gentlemanautomaton/stathat"
	"github.com/gentlemanautomaton/stathat/statmatch"
)

// Match provides stat name inclusion and exclusion filters for commands.
type Match struct {
	Include []statmatch.Pattern `kong:"env='INCLUDE',name='include',help='Include stats with names matching regular expression pattern.'"`
	Exclude []statmatch.Pattern `kong:"env='EXCLUDE',name='exclude',help='Exclude stats with names matching regular expression pattern.'"`
}

// Empty returns true if the match is unspecified.
func (m Match) Empty() bool {
	return len(m.Include) == 0 && len(m.Exclude) == 0
}

// Filter returns a stathat filter for m.
func (m Match) Filter() stathat.Filter {
	var includes, excludes []stathat.Filter
	for _, pattern := range m.Include {
		includes = append(includes, statmatch.Name(pattern))
	}
	for _, pattern := range m.Exclude {
		excludes = append(excludes, statmatch.Name(pattern))
	}
	return statmatch.All(statmatch.Any(includes...), statmatch.None(excludes...))
}
