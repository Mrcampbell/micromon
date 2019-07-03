package grpc

import (
	"context"
	"fmt"

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

	err = bs.battleService.Upsert(ctx, battle)
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

func (bs *BattleServer) SubmitTurn(ctx context.Context, req *pokemon.SubmitTurnRequest) (*pokemon.SubmitTurnResponse, error) {
	b, err := bs.battleService.GetBattle(ctx, req.BattleId)
	if err != nil {
		return nil, err
	}
	fmt.Println("REQ: ", req)

	if req.PlayerId == b.PlayerAId && b.State == pokemon.BattleState_WAITING_ON_BOTH_PLAYERS { // if one player moves first, then put state waiting on the other
		b.State = pokemon.BattleState_WAITING_ON_PLAYER_B
	} else if req.PlayerId == b.PlayerBId && b.State == pokemon.BattleState_WAITING_ON_BOTH_PLAYERS {
		b.State = pokemon.BattleState_WAITING_ON_PLAYER_A
	} else if req.PlayerId == b.PlayerAId && b.State == pokemon.BattleState_WAITING_ON_PLAYER_A { // if the player who moved was being waited for, queue it up, baby.
		b.State = pokemon.BattleState_MOVE_RESULTS_QUEUED
	} else if req.PlayerId == b.PlayerBId && b.State == pokemon.BattleState_WAITING_ON_PLAYER_B {
		b.State = pokemon.BattleState_MOVE_RESULTS_QUEUED
	}

	fmt.Println(req)
	fmt.Println(b.State)

	err = bs.battleService.Upsert(ctx, b)
	if err != nil {
		return nil, err
	}

	return &pokemon.SubmitTurnResponse{}, nil
}
