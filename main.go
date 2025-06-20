package main

import (
	"fmt"
	"os"

	"github.com/DanielJohn17/rss-blog-aggregator/internal/config"
	"github.com/DanielJohn17/rss-blog-aggregator/internal/handlers"
)

func main() {
	cfg := config.Read()
	if cfg == nil {
		fmt.Println("Config not found!")
		os.Exit(1)
	}

	state := &handlers.State{
		Config: cfg,
	}

	commands := &handlers.Commands{}

	commands.Register("login", handlers.HandlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("No command provided.")
		os.Exit(1)
	}

	cmd := handlers.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	if err:= commands.Run(state, cmd); err != nil {
		fmt.Printf("Error executing command '%s': %v\n", cmd.Name, err)
		os.Exit(1)
	}
}
