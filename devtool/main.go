package main

import (
	"flag"
	"fmt"

	"github.com/haskaalo/intribox/models"
	"github.com/rs/zerolog/log"
)

func printAvailableCommands() {
	fmt.Println("Invalid options")
	fmt.Println("\nCommand available:")
	fmt.Println("	testuser: Create test user")
}

func main() {
	email := flag.String("email", "test@example.com", "Test user email")
	action := flag.String("action", "", "Action")

	flag.Parse()

	switch *action {
	case "testuser":
		user, err := models.CreateTestUserWithCustomEmail(*email)
		if err != nil {
			log.Fatal().AnErr("error", err).Msg("Failed to create new user")
		}

		fmt.Println("email:", user.Email)
		fmt.Println("password", models.TestUserPassword)
		return
	default:
		printAvailableCommands()
	}
}
