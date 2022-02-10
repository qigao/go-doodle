package main

import (
	"fmt"
	"log"

	"gforum/db"
)

func main() {
	d, err := db.New()

	if err != nil {
		log.Fatal(fmt.Errorf("failed to connect database: %w", err))
	}

	if err := db.Seed(d); err != nil {
		log.Fatal(fmt.Errorf("failed to seed: %w", err))
	}
}
