package statmatch

import (
	"fmt"
	"regexp"
	"strings"
)

// Pattern is a stat matching pattern based on regular expressions.
type Pattern struct {
	Expression *regexp.Regexp
}

// UnmarshalText unmarshals the given text as a pattern in p.
func (p *Pattern) UnmarshalText(text []byte) error {
	re := string(text)

	// Interpret special values as no-ops
	if re == "" || re == "_" {
		p.Expression = nil
		return nil
	}

	// Compile the pattern
	exp, err := compileRegex(re)
	if err != nil {
		return err
	}

	p.Expression = exp

	return nil
}

// MatchString returns true if the pattern matches s.
func (p Pattern) MatchString(s string) bool {
	if p.Expression == nil {
		return true
	}
	return p.Expression.MatchString(s)
}

// String returns a string representation of the pattern.
func (p Pattern) String() string {
	if p.Expression == nil {
		return "*"
	}
	return p.Expression.String()
}

func compileRegex(re string) (*regexp.Regexp, error) {
	if re == "" {
		return nil, nil
	}

	// Force case-insensitive matching
	if !strings.HasPrefix(re, "(?i)") {
		re = "(?i)" + re
	}

	c, err := regexp.Compile(re)
	if err != nil {
		return nil, fmt.Errorf("unable to compile regular expression \"%s\": %v", re, err)
	}
	return c, nil
}
