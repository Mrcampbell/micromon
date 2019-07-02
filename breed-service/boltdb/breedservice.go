package boltdb

import (
	"context"
	"fmt"
	"time"

	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"github.com/boltdb/bolt"
)

var (
	bucket = []byte("Breeds")
)
var _ pokemon.BreedServiceServer = &BreedService{}

type BreedService struct {
	DB              *bolt.DB
	breedMoveClient pokemon.BreedMoveServiceClient
}

func NewBreedService(breedMoveClient pokemon.BreedMoveServiceClient) (*BreedService, error) {
	db, err := bolt.Open("data/db.bolt", 0600, &bolt.Options{
		Timeout: 1 * time.Second,
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	bs := &BreedService{
		DB:              db,
		breedMoveClient: breedMoveClient,
	}

	err = bs.load()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return bs, nil
}

func (bs *BreedService) GetBreedSummary(ctx context.Context, req *pokemon.GetBreedSummaryRequest) (*pokemon.GetBreedSummaryResponse, error) {

	b, err := bs.readBreedByID(req.Id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	summary := convertDetailToSummmary(*b)

	return &pokemon.GetBreedSummaryResponse{
		Summary: &summary,
	}, nil
}

func (bs *BreedService) GetBreedDetail(ctx context.Context, req *pokemon.GetBreedDetailRequest) (*pokemon.GetBreedDetailResponse, error) {
	b, err := bs.readBreedByID(req.Id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	bm, err := bs.breedMoveClient.GetMovesForBreed(ctx, &pokemon.GetMovesForBreedRequest{
		BreedId: req.Id,
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	b.BreedMoves = bm.BreedMoves

	return &pokemon.GetBreedDetailResponse{
		Detail: b,
	}, nil
}

func (bs *BreedService) readBreedByID(id string) (*pokemon.BreedDetail, error) {
	var breed pokemon.BreedDetail
	err := bs.DB.View(func(tx *bolt.Tx) error {
		var err error

		b := tx.Bucket(bucket)
		k := []byte(id)

		breed, err = decode(b.Get(k))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &breed, nil
}
