package main

import (
	"fmt"

	"github.com/haskaalo/intribox/models"
	"github.com/rs/zerolog/log"
)

func main() {
	user, err := models.CreateTestUser()
	if err != nil {
		log.Fatal().AnErr("error", err).Msg("Failed to create new user")
	}

	fmt.Println("email:", user.Email)
	fmt.Println("password", models.TestUserPassword)
}
