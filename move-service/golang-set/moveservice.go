package dgraph

import (
	"context"

	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
)

type MoveService struct {
	set map[int]pokemon.Move
}

func NewMoveService(ctx context.Context) (*MoveService, error) {

	var set map[int]pokemon.Move

	return &MoveService{
		set: set,
	}, nil
}

func (ms *MoveService) load() {
	ms.set[1] = pokemon.Move{
		Id:   1,
		Name: "Scratch",
	}
	ms.set[2] = pokemon.Move{
		Id:   2,
		Name: "Pound",
	}
}
