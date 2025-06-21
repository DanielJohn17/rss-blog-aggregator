package handlers

import (
	"github.com/DanielJohn17/rss-blog-aggregator/internal/config"
	"github.com/DanielJohn17/rss-blog-aggregator/internal/database"
)

type State struct {
	Db     *database.Queries
	Config *config.Config
}
