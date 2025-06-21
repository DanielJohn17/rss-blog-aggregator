package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/DanielJohn17/rss-blog-aggregator/internal/config"
	"github.com/DanielJohn17/rss-blog-aggregator/internal/database"
	"github.com/DanielJohn17/rss-blog-aggregator/internal/handlers"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Read()
	if cfg == nil {
		fmt.Println("Config not found!")
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		os.Exit(1)
	}

	dbQueries := database.New(db)
	state := &handlers.State{
		Db:     dbQueries,
		Config: cfg,
	}

	commands := &handlers.Commands{}

	commands.Register("login", handlers.HandlerLogin)
	commands.Register("register", handlers.HandlerRegister)

	if len(os.Args) < 2 {
		fmt.Println("No command provided.")
		os.Exit(1)
	}

	cmd := handlers.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err := commands.Run(state, cmd); err != nil {
		fmt.Printf("Error executing command '%s': %v\n", cmd.Name, err)
		os.Exit(1)
	}
}
