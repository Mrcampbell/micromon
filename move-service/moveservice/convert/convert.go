package convert

import "github.com/Mrcampbell/pgo2/protorepo/pokemon"

func MoveDetailToSummary(detail pokemon.MoveDetail) pokemon.MoveSummary {
	return pokemon.MoveSummary{
		Id: detail.Id,
	}
}
