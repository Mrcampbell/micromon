package psql

import (
	"context"
	"fmt"

	"github.com/Mrcampbell/pgo2/battle-service/config"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"github.com/Mrcampbell/pgo2/shared-library/uuid"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
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

	createSchema(db)

	return &BattleService{
		DB:             db,
		pokemonService: pokemonService,
		moveService:    moveService,
	}, nil
}

func (bs *BattleService) Create(ctx context.Context, req *pokemon.CreateBattleRequest) (*pokemon.CreateBattleResponse, error) {

	pa, err := bs.pokemonService.GetPokemon(ctx, &pokemon.GetPokemonRequest{
		Id: req.PokemonAId,
	})
	if err != nil {
		return nil, err
	}

	pb, err := bs.pokemonService.GetPokemon(ctx, &pokemon.GetPokemonRequest{
		Id: req.PokemonAId,
	})
	if err != nil {
		return nil, err
	}

	pabm := mapPokemonToBattleMask(*pa.Pokemon)
	pbbm := mapPokemonToBattleMask(*pb.Pokemon)

	battle := &pokemon.Battle{

		Id:                 uuid.PrefixedUUID("b"),
		Duration:           42,
		PlayerAId:          "a",
		PlayerBId:          "b",
		State:              pokemon.BattleState_WAITING_ON_BOTH_PLAYERS,
		PokemonABattleMask: &pabm,
		PokemonBBattleMask: &pbbm,
	}

	err = bs.DB.Insert(battle)
	if err != nil {
		return nil, err
	}

	return &pokemon.CreateBattleResponse{
		Battle: battle,
	}, nil
}

func (bs *BattleService) GetBattle(ctx context.Context, req *pokemon.GetBattleRequest) (*pokemon.GetBattleResponse, error) {
	b := &pokemon.Battle{Id: req.Id}
	err := bs.DB.Select(b)
	if err != nil {
		return nil, err
	}
	return &pokemon.GetBattleResponse{
		Battle: b,
	}, nil
}

func createSchema(db *pg.DB) error {

	for _, model := range []interface{}{(*pokemon.Battle)(nil)} {

		// todo: remove after debug
		// db.DropTable(model, &orm.DropTableOptions{
		// 	Cascade:  true,
		// 	IfExists: true,
		// })
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (bs *BattleService) UseAttack(ctx context.Context, req *pokemon.UseAttackRequest) (*pokemon.UseAttackResponse, error) {
	return &pokemon.UseAttackResponse{}, nil
}
