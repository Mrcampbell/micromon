package boltdb

import (
	"encoding/json"

	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
)

func encode(breed pokemon.BreedDetail) ([]byte, error) {
	enc, err := json.Marshal(breed)
	if err != nil {
		return nil, err
	}
	return enc, nil
}

func decode(data []byte) (pokemon.BreedDetail, error) {
	var breed pokemon.BreedDetail
	err := json.Unmarshal(data, &breed)
	if err != nil {
		return pokemon.BreedDetail{}, err
	}
	return breed, nil
}

func convertDetailToSummmary(b pokemon.BreedDetail) pokemon.BreedSummary {
	return *b.Summary
}
