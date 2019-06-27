package grpc

import (
	"context"

	"github.com/Mrcampbell/pgo2/breed-move-service/psql"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
)

type BreedMoveServer struct {
	breedMoveService psql.BreedMoveService
}

func NewBreedMoveServer(bmservice psql.BreedMoveService) *BreedMoveServer {
	return &BreedMoveServer{
		breedMoveService: bmservice,
	}
}

func (bms *BreedMoveServer) GetMovesForBreed(ctx context.Context, req *pokemon.GetMovesForBreedRequest) (*pokemon.GetMovesForBreedResponse, error) {
	breedMoves, err := bms.breedMoveService.GetMovesForBreed(ctx, req.BreedId, req.VersionGroupId)
	if err != nil {
		return nil, err
	}

	return &pokemon.GetMovesForBreedResponse{
		BreedMoves: breedMoves,
	}, nil
}
