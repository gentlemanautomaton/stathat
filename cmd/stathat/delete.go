package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gentlemanautomaton/stathat"
)

// DeleteCmd deletes a set of stats in StatHat.
type DeleteCmd struct {
	AccessToken string `kong:"env='ACCESS_TOKEN',name='token',required,help='StatHat API access token (not EZ key).'"`
	Match
	Limit int `kong:"env='CONCURRENCY_LIMIT',name='limit',default='8',help='The maximum number of concurrent deletions.'"`
}

// Run executes the delete command.
func (cmd DeleteCmd) Run(ctx context.Context) error {
	// Prepare a client
	s := stathat.New().Token(cmd.AccessToken)

	// Ask the client for all stats
	list, err := s.StatListAll()
	if err != nil {
		return err
	}
	if err := ctx.Err(); err != nil {
		return err
	}

	// Build a list of stats that match our query
	filter := cmd.Match.Filter()
	matched := make([]stathat.StatItem, 0, len(list))
	var totalStats, matchedStats int
	for _, stat := range list {
		totalStats++
		if filter(stat) {
			matchedStats++
			matched = append(matched, stat)
		}
	}

	// Exit if there are no matches
	if len(matched) == 0 {
		fmt.Printf("No stats matched.\n")
		return nil
	}
	if err := ctx.Err(); err != nil {
		return err
	}

	// Present the list of matching stats before we delete them
	fmt.Printf("----Matching Stats (%d/%d)----\n", matchedStats, totalStats)
	for _, stat := range matched {
		if err := ctx.Err(); err != nil {
			return err
		}
		fmt.Printf("  [%s] %s\n", stat.ID, stat.Name)
	}

	// Ask the user whether we should proceed
	msg := fmt.Sprintf("Should we proceed with deleting the %d %s listed above?", matchedStats, pluralize(matchedStats, "stat", "stats"))
	if proceed, err := promptYesNo(msg); err != nil {
		return err
	} else if !proceed {
		return nil
	}

	// Prepare a slice of result channels that will collect the results of
	// concurrent API calls
	results := make([]chan deletion, len(matched))
	for i := range results {
		results[i] = make(chan deletion)
	}

	// Decorate the results
	defer fmt.Printf("----\n")

	// Proceed with deleting each stat
	pool := cmd.Pool()
	go func() {
		for i, stat := range matched {
			// Determine the result channel
			ch := results[i]

			// Wait for a work token from the pool
			pool.Acquire(ctx)

			// Perform the deletion
			go deleteStat(ctx, pool, stat, ch)
		}
	}()

	// Report results
	for i, stat := range matched {
		ch := results[i]
		deletion := <-ch
		if deletion.Err != nil {
			if deletion.Err != context.Canceled && deletion.Err != context.DeadlineExceeded {
				if deletion.Msg == "" {
					fmt.Printf("  [%s] %s: %v (%s)\n", stat.ID, stat.Name, deletion.Err, deletion.Duration)
				} else {
					fmt.Printf("  [%s] %s: %s: %v (%s)\n", stat.ID, stat.Name, deletion.Msg, deletion.Err, deletion.Duration)
				}
			}
		} else if deletion.Msg != "" {
			fmt.Printf("  [%s] %s: %s (%s)\n", stat.ID, stat.Name, deletion.Msg, deletion.Duration)
		} else {
			fmt.Printf("  [%s] %s: Deleted (%s)\n", stat.ID, stat.Name, deletion.Duration)
		}
	}

	return ctx.Err()
}

// Pool returns a work token pool configured according to cmd.
func (cmd DeleteCmd) Pool() *Pool {
	// Constrain the degree of concurrency
	limit := cmd.Limit
	if limit < 1 {
		limit = 1
	} else if limit > 64 {
		limit = 64
	}

	return NewPool(limit)
}

type deletion struct {
	Msg      string
	Err      error
	Duration time.Duration
}

func deleteStat(ctx context.Context, pool *Pool, stat stathat.StatItem, ch chan<- deletion) {
	// Release our work token when we're done
	defer pool.Release()

	// Close the channel when we're done
	defer close(ch)

	// If the context has been cancelled, closing the channel immediately will s
	if err := ctx.Err(); err != nil {
		ch <- deletion{Err: err}
		return
	}

	start := time.Now()

	// Request deletion of the stat, which can take a long time
	msg, err := stat.Delete()

	// Report the result
	ch <- deletion{
		Msg:      msg,
		Err:      err,
		Duration: time.Since(start),
	}
}
