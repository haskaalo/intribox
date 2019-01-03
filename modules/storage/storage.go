package storage

import (
	"github.com/rs/zerolog/log"

	"github.com/haskaalo/intribox/config"
	"github.com/haskaalo/intribox/modules/remote"
	"github.com/haskaalo/intribox/modules/remote/local"
)

// CurrentRemote Remote specified in config
var CurrentRemote remote.Remote

func init() {
	switch config.Storage.RemoteName {
	case "local":
		CurrentRemote = local.R{}
	default:
		log.Fatal().Str("remote", config.Storage.RemoteName).Msg("Invalid Remote Name in config")
	}
}
