package main

import (
	"context"
	"fmt"

	"github.com/gentlemanautomaton/stathat"
)

// ListCmd lists a set of stats in StatHat.
type ListCmd struct {
	AccessToken string `kong:"env='ACCESS_TOKEN',name='token',required,help='StatHat API access token (not EZ key).'"`
	Match
}

// Run executes the list command.
func (cmd ListCmd) Run(ctx context.Context) error {
	s := stathat.New().Token(cmd.AccessToken)
	list, err := s.StatListAll()
	filter := cmd.Match.Filter()
	for _, stat := range list {
		if err := ctx.Err(); err != nil {
			return err
		}
		if filter(stat) {
			fmt.Printf("[%s] %s\n", stat.ID, stat.Name)
		}
	}
	return err
}
