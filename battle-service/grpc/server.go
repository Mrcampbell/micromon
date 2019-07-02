package grpc

import (
	"context"

	"github.com/Mrcampbell/pgo2/battle-service/psql"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"github.com/Mrcampbell/pgo2/shared-library/uuid"
)

var _ pokemon.BattleServiceServer = &BattleServer{}

type BattleServer struct {
	battleService  psql.BattleService
	pokemonService pokemon.PokemonServiceClient
	moveService    pokemon.MoveServiceClient
}

func NewBattleServer(battleService psql.BattleService, pokemonService pokemon.PokemonServiceClient, moveService pokemon.MoveServiceClient) (*BattleServer, error) {

	return &BattleServer{
		battleService:  battleService,
		pokemonService: pokemonService,
		moveService:    moveService,
	}, nil
}

func (bs *BattleServer) Create(ctx context.Context, req *pokemon.CreateBattleRequest) (*pokemon.CreateBattleResponse, error) {
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

	err = bs.battleService.Insert(ctx, battle)
	if err != nil {
		return nil, err
	}

	return &pokemon.CreateBattleResponse{
		Battle: battle,
	}, nil
}

func (bs *BattleServer) GetBattle(ctx context.Context, req *pokemon.GetBattleRequest) (*pokemon.GetBattleResponse, error) {
	b, err := bs.battleService.GetBattle(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pokemon.GetBattleResponse{
		Battle: b,
	}, nil
}

func (bs *BattleServer) UseAttack(ctx context.Context, req *pokemon.UseAttackRequest) (*pokemon.UseAttackResponse, error) {
	return &pokemon.UseAttackResponse{}, nil
}
