package main

import (
	"github.com/ikozinov/jumptunnel/config"
	"github.com/ikozinov/jumptunnel/sshtunnel"
	"github.com/rs/zerolog/log"
)

func main() {
	conf, err := config.New()
	if err != nil {
		log.Fatal().Err(err).Msg("Read config")
	}

	log.Info().Interface("config", conf).Send()

	authMethod, err := conf.ParsePrivateKey()
	if err != nil {
		log.Fatal().Err(err).Msg("Parse private key")
	}

	tunnel := sshtunnel.NewSSHTunnel(
		conf.Server(),
		authMethod,
		conf.Remote(),
		conf.Listen(),
	)

	tunnel.Log = logger{}

	if err := tunnel.Start(); err != nil {
		log.Fatal().Err(err).Msg("Tunnel start")

	}

}

type logger struct {
}

func (l logger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
