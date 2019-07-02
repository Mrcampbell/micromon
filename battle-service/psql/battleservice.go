package psql

import (
	"context"
	"fmt"

	"github.com/Mrcampbell/pgo2/battle-service/config"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"github.com/go-pg/pg"
)

var _ pokemon.BattleServiceServer = &BattleService{}

type BattleService struct {
	DB             *pg.DB
	pokemonService pokemon.PokemonServiceClient
	moveService    pokemon.MoveServiceClient
}

func NewBattleService(pokemonService pokemon.PokemonServiceClient, moveService pokemon.MoveServiceClient) (*BattleService, error) {

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
		DB:             db,
		pokemonService: pokemonService,
		moveService:    moveService,
	}, nil
}

func (bs *BattleService) Create(ctx context.Context, req *pokemon.CreateBattleRequest) (*pokemon.CreateBattleResponse, error) {
	return &pokemon.CreateBattleResponse{
		Battle: &pokemon.Battle{
			Duration: 0,
			PokemonABattleMask: &pokemon.PokemonBattleMask{
				Pokemon: &pokemon.Pokemon{
					Id: "Fuck you.",
				},
			},
		},
	}, nil
}
func (bs *BattleService) GetBattle(ctx context.Context, req *pokemon.GetBattleRequest) (*pokemon.GetBattleResponse, error) {
	return nil, nil
}
