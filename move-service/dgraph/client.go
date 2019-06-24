package dgraph

import (
	"context"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

func newClient() (*dgo.Dgraph, error) {

	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("server:8080", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	), nil
}

func setup(ctx context.Context, c *dgo.Dgraph) error {

	err := c.Alter(ctx, &api.Operation{
		Schema: `
			name: string &index(term) .
			balance: int . 
		`,
	})

	return err
}
