package psql

import (
	"fmt"

	"github.com/Mrcampbell/pgo2/battle-service/config"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"github.com/go-pg/pg"
)

type BattleService struct {
	DB             *pg.DB
	pokemonService pokemon.PokemonServiceClient
	moveService    pokemon.MoveServiceClient
}

func NewBatleService(pokemonService pokemon.PokemonServiceClient, moveService pokemon.MoveServiceClient) (*BattleService, error) {

	db := pg.Connect(&pg.Options{
		Addr:     config.DBHost + ":" + config.DBUser,
		Database: config.DBDatabase,
		User:     config.DBUser,
		Password: config.DBPassword,
		OnConnect: func(conn *pg.Conn) error {
			fmt.Println("Connected to DB")
			return nil
		},
	})

	return &BattleService{
		pokemonService: pokemonService,
		moveService:    moveService,
	}, nil
}
