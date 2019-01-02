package storage

import (
	"github.com/haskaalo/intribox/config"
	"github.com/haskaalo/intribox/modules/storage/remotes"
	"github.com/haskaalo/intribox/modules/storage/remotes/local"
	"github.com/rs/zerolog/log"
)

// CurrentRemote Remote specified in config
var CurrentRemote remotes.Remote

func init() {
	switch config.Storage.RemoteName {
	case "local":
		CurrentRemote = local.R{}
	default:
		log.Fatal().Str("remote", config.Storage.RemoteName).Msg("Invalid Remote Name in config")
	}
}
