package dgraph

import (
	"context"

	"github.com/dgraph-io/dgo"
)

type MoveService struct {
	client *dgo.Dgraph
}

func NewMoveService(ctx context.Context) (*MoveService, error) {
	client, err := newClient()
	if err != nil {
		return nil, err
	}

	err = setup(ctx, client)
	if err != nil {
		return nil, err
	}

	return &MoveService{
		client: client,
	}, nil
}
