package commands

import "github.com/datsun80zx/go_rss_aggregator.git/internal"

type Commands struct {
	Handlers map[string]func(*internal.State, Command) error
}
