package storage

import (
	"github.com/rs/zerolog/log"

	"github.com/haskaalo/intribox/config"
	"github.com/haskaalo/intribox/storage/backend"
	"github.com/haskaalo/intribox/storage/backend/local"
)

// Remote current remote used
var Remote backend.Backend

func init() {
	switch config.Storage.RemoteName {
	case "local":
		Remote = new(local.R)
	default:
		log.Fatal().Str("remote", config.Storage.RemoteName).Msg("Invalid Remote Name in config")
	}
}
