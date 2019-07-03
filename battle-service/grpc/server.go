package grpc

import (
	"context"
	"fmt"

	"github.com/kr/pretty"

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
		Rounds:             make([]*pokemon.BattleRound, 0),
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
	// b.State = pokemon.BattleState_WAITING_ON_BOTH_PLAYERS

	if b.State == pokemon.BattleState_WAITING_ON_BOTH_PLAYERS {
		b.Rounds = append(b.Rounds, &pokemon.BattleRound{})
	}

	lastRound := b.Rounds[len(b.Rounds)-1]

	if req.PlayerId == b.PlayerAId && b.State == pokemon.BattleState_WAITING_ON_PLAYER_B ||
		req.PlayerId == b.PlayerBId && b.State == pokemon.BattleState_WAITING_ON_PLAYER_A ||
		b.State == pokemon.BattleState_MOVE_RESULTS_QUEUED { // if a player has already taken a turn and resubmits
		fmt.Println("You already took your turn, you schmuck.")
	} else if req.PlayerId == b.PlayerAId && b.State == pokemon.BattleState_WAITING_ON_BOTH_PLAYERS { // if one player moves first, then put state waiting on the other
		lastRound.PlayerATurn = req.Turn
		b.State = pokemon.BattleState_WAITING_ON_PLAYER_B
	} else if req.PlayerId == b.PlayerBId && b.State == pokemon.BattleState_WAITING_ON_BOTH_PLAYERS {
		lastRound.PlayerBTurn = req.Turn
		b.State = pokemon.BattleState_WAITING_ON_PLAYER_A
	} else if req.PlayerId == b.PlayerAId && b.State == pokemon.BattleState_WAITING_ON_PLAYER_A { // if the player who moved was being waited for, queue it up, baby.
		lastRound.PlayerATurn = req.Turn
		b.State = pokemon.BattleState_MOVE_RESULTS_QUEUED
	} else if req.PlayerId == b.PlayerBId && b.State == pokemon.BattleState_WAITING_ON_PLAYER_B {
		lastRound.PlayerBTurn = req.Turn
		b.State = pokemon.BattleState_MOVE_RESULTS_QUEUED
	}

	if b.State == pokemon.BattleState_MOVE_RESULTS_QUEUED {
		b.Duration++
	}

	pretty.Println(b)

	err = bs.battleService.Upsert(ctx, b)
	if err != nil {
		return nil, err
	}

	return &pokemon.SubmitTurnResponse{}, nil
}
