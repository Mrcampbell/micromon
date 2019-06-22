package mock

import (
	"context"

	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
)

type PokemonService struct {
}

func NewPokemonService() *PokemonService {
	return &PokemonService{}
}

func (*PokemonService) GetPokemon(ctx context.Context, req *pokemon.GetPokemonRequest) (*pokemon.GetPokemonResponse, error) {
	return &pokemon.GetPokemonResponse{
		Pokemon: &pokemon.Pokemon{
			Id:      1,
			BreedId: 1,
		},
	}, nil
}

func (*PokemonService) ListPokemon(ctx context.Context, req *pokemon.ListPokemonRequest) (*pokemon.ListPokemonResponse, error) {
	return &pokemon.ListPokemonResponse{
		Pokemon: []*pokemon.Pokemon{
			&pokemon.Pokemon{
				Id:      1,
				BreedId: 1,
			},
			&pokemon.Pokemon{
				Id:      2,
				BreedId: 25,
			},
		},
	}, nil
}
