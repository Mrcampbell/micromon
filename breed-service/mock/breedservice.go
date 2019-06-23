package mock

import (
	"context"

	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
)

type BreedService struct {
}

func NewBreedService() *BreedService {
	return &BreedService{}
}

func (*BreedService) GetBreedSummary(ctx context.Context, req *pokemon.GetBreedSummaryRequest) (*pokemon.GetBreedSummaryResponse, error) {
	return &pokemon.GetBreedSummaryResponse{
		Breed: &pokemon.BreedSummary{
			Id:   1,
			Name: "Bulbasaur",
		},
	}, nil
}

func (*BreedService) GetBreedDetail(ctx context.Context, req *pokemon.GetBreedDetailRequest) (*pokemon.GetBreedDetailResponse, error) {
	return &pokemon.GetBreedDetailResponse{
		Breed: &pokemon.BreedDetail{
			Id:             1,
			Name:           "Bulbasaur",
			Hp:             2,
			Attack:         3,
			Defense:        4,
			SpecialAttack:  5,
			SpecialDefense: 6,
			Speed:          7,
		},
	}, nil
}
