package convert

import "github.com/Mrcampbell/pgo2/protorepo/pokemon"

func MoveDetailToSummary(detail pokemon.MoveDetail) pokemon.MoveSummary {
	return *detail.Summary
}
