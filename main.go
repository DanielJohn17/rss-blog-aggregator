package main

import (
	"fmt"
	"github.com/DanielJohn17/rss-blog-aggregator/internal/config"
)

func main() {
	config := config.Read()

	fmt.Printf("Db url: %s\n", config.DBURL)
	fmt.Printf("Current user: %s\n", config.CurrentUsername)

	config.SetUser("daniel")

	fmt.Printf("Db url: %s\n", config.DBURL)
	fmt.Printf("Current user: %s\n", config.CurrentUsername)
}
