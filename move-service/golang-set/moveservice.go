package dgraph

import (
	"context"

	mapset "github.com/deckarep/golang-set"
)

type MoveService struct {
	set mapset.Set
}

func NewMoveService(ctx context.Context) (*MoveService, error) {

	set := mapset.NewSet()

	return &MoveService{
		set: set,
	}, nil
}

func (ms *MoveService) load() {
	ms.set.Add()
}