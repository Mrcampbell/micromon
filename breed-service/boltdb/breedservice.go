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
			Stats: &pokemon.StatGroup{
				Hp:             45,
				Attack:         49,
				Defense:        49,
				SpecialAttack:  65,
				SpecialDefense: 65,
				Speed:          45,
			},
		})
		if err != nil {
			return err
		}
		b2, err := encode(pokemon.BreedDetail{
			Summary: &pokemon.BreedSummary{
				Id:   "2",
				Name: "Ivysaur",
			},
			Stats: &pokemon.StatGroup{
				Hp:             60,
				Attack:         62,
				Defense:        63,
				SpecialAttack:  80,
				SpecialDefense: 80,
				Speed:          60,
			},
		})
		if err != nil {
			return err
		}

		b3, err := encode(pokemon.BreedDetail{
			Summary: &pokemon.BreedSummary{
				Id:   "25",
				Name: "Pikachu",
			},
			Stats: &pokemon.StatGroup{
				Hp:             35,
				Attack:         55,
				Defense:        30,
				SpecialAttack:  50,
				SpecialDefense: 40,
				Speed:          90,
			},
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
		err = b.Put([]byte("25"), b3)
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
		Summary: &summary,
	}, nil
}

func (bs *BreedService) GetBreedDetail(ctx context.Context, req *pokemon.GetBreedDetailRequest) (*pokemon.GetBreedDetailResponse, error) {

	b, err := bs.readBreedByID(req.Id)
	if err != nil {
		return nil, err
	}

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
		return nil, err
	}
	return &breed, nil
}
