package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/alecthomas/kong"
)

func main() {
	// Capture shutdown signals
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var cli struct {
		Count  CountCmd  `kong:"cmd,help='Count the number of stats matching a query.'"`
		List   ListCmd   `kong:"cmd,help='List stats matching a query.'"`
		Delete DeleteCmd `kong:"cmd,help='Delete stats matching a query (requires confirmation).'"`
	}

	app := kong.Parse(&cli,
		kong.Description("A command line tool that interacts with the StatHat API."),
		kong.BindTo(ctx, (*context.Context)(nil)),
		kong.UsageOnError())

	if err := app.Run(ctx); err != nil {
		app.FatalIfErrorf(err)
	}
}
