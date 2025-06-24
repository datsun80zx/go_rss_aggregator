package internal

import (
	"github.com/datsun80zx/go_rss_aggregator.git/internal/config"
	"github.com/datsun80zx/go_rss_aggregator.git/internal/database"
)

type State struct {
	Config   *config.Config
	Database *database.Queries
}
