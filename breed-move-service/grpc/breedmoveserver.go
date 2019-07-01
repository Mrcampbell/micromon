package grpc

import (
	"context"
	"fmt"

	"github.com/Mrcampbell/pgo2/breed-move-service/psql"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
)

var _ pokemon.BreedMoveServiceServer = &BreedMoveServer{}

type BreedMoveServer struct {
	breedMoveService psql.BreedMoveService
}

func NewBreedMoveServer(bmservice psql.BreedMoveService) *BreedMoveServer {
	return &BreedMoveServer{
		breedMoveService: bmservice,
	}
}

func (bms *BreedMoveServer) GetMovesForBreed(ctx context.Context, req *pokemon.GetMovesForBreedRequest) (*pokemon.GetMovesForBreedResponse, error) {

	fmt.Println("MOVES REQUEST")

	if req.VersionGroupId == pokemon.VersionGroup_UNKNOWN_VERSION_GROUP {
		fmt.Println("Version Group Not Provided")
		req.VersionGroupId = pokemon.VersionGroup_ULTRA_SUN_ULTRA_MOON
	}

	breedMoves, err := bms.breedMoveService.GetMovesForBreed(ctx, req.BreedId, req.VersionGroupId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &pokemon.GetMovesForBreedResponse{
		BreedMoves: breedMoves,
	}, nil
}

func (bms *BreedMoveServer) GetRandomMoveSetForBreed(ctx context.Context, req *pokemon.GetRandomMoveSetForBreedRequest) (*pokemon.GetRandomMoveSetForBreedResponse, error) {

	fmt.Println("RANDOM MOVE REQUEST")
	ms, err := bms.breedMoveService.GetRandomMoveSummarySetForBreed(ctx, req.BreedId, req.VersionGroupId, int(req.Level))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	response := pokemon.GetRandomMoveSetForBreedResponse{}

	response.MoveOne = ms[0]

	response.MoveTwo = ms[1]

	response.MoveThree = ms[2]

	response.MoveFour = ms[3]

	return &response, nil
}
