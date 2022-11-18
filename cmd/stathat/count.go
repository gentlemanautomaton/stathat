package main

import (
	"context"
	"fmt"

	"github.com/gentlemanautomaton/stathat"
)

// CountCmd counts the number of stats in StatHat.
type CountCmd struct {
	AccessToken string `kong:"env='ACCESS_TOKEN',name='token',required,help='StatHat API access token (not EZ key).'"`
	Match
}

// Run executes the count command.
func (cmd CountCmd) Run(ctx context.Context) error {
	s := stathat.New().Token(cmd.AccessToken)
	list, err := s.StatListAll()
	filter := cmd.Match.Filter()
	var matched, total int
	for _, stat := range list {
		if err := ctx.Err(); err != nil {
			return err
		}
		total++
		if filter(stat) {
			matched++
		}
	}
	if cmd.Match.Empty() && matched == 0 {
		fmt.Printf("Total Stats: %d\n", total)
	} else {
		fmt.Printf("Matching Stats: %d/%d\n", matched, total)
	}
	return err
}
