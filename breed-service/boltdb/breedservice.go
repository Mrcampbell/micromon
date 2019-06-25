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

type BreedService struct {
	DB *bolt.DB
}

func NewBreedService() (*BreedService, error) {
	db, err := bolt.Open("data/db.bolt", 0600, &bolt.Options{
		Timeout: 1 * time.Second,
	})

	if err != nil {
		return nil, err
	}

	bs := &BreedService{
		DB: db,
	}

	err = bs.initialize()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return bs, nil
}

func (bs *BreedService) initialize() error {
	return bs.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}

		b1, err := encode(pokemon.BreedDetail{
			Summary: &pokemon.BreedSummary{
				Id:   "1",
				Name: "Bulbasaur",
			},
			Hp:             1,
			Attack:         2,
			Defense:        3,
			SpecialAttack:  4,
			SpecialDefense: 5,
			Speed:          6,
		})
		if err != nil {
			return err
		}
		b2, err := encode(pokemon.BreedDetail{
			Summary: &pokemon.BreedSummary{
				Id:   "2",
				Name: "Ivysaur",
			},
			Hp:             1,
			Attack:         2,
			Defense:        3,
			SpecialAttack:  4,
			SpecialDefense: 5,
			Speed:          6,
		})
		if err != nil {
			return err
		}
		err = b.Put([]byte("1"), b1)
		if err != nil {
			return err
		}
		err = b.Put([]byte("2"), b2)
		if err != nil {
			return err
		}

		return nil
	})
}

func (bs *BreedService) GetBreedSummary(ctx context.Context, req *pokemon.GetBreedSummaryRequest) (*pokemon.GetBreedSummaryResponse, error) {

	b, err := bs.readBreedByID(req.Id)
	if err != nil {
		return nil, err
	}

	summary := convertDetailToSummmary(*b)

	return &pokemon.GetBreedSummaryResponse{
		Breed: &summary,
	}, nil
}

func (bs *BreedService) GetBreedDetail(ctx context.Context, req *pokemon.GetBreedDetailRequest) (*pokemon.GetBreedDetailResponse, error) {

	b, err := bs.readBreedByID(req.Id)
	if err != nil {
		return nil, err
	}

	return &pokemon.GetBreedDetailResponse{
		Breed: b,
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
		return nil, err
	}
	return &breed, nil
}
