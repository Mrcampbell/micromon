package boltdb

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/Mrcampbell/pgo2/breed-service/app"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"github.com/boltdb/bolt"
)

func (bs *BreedService) load() error {
	breedList := make(map[string]pokemon.BreedDetail, 0)
	err := loadBaseFromCSV(breedList)

	if err != nil {
		return err
	}

	err = applyCSVStat(breedList)

	if err != nil {
		return err
	}

	return bs.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucket)
		if err != nil {
			return err
		}

		for _, record := range breedList {
			encoded, err := encode(record)
			if err != nil {
				return err
			}

			err = b.Put([]byte(record.Summary.Id), encoded)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func loadBaseFromCSV(breedList map[string]pokemon.BreedDetail) error {
	reader, err := getCSVReader("./data/source/pokemon.csv")
	if err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		line, err := reader.Read()
		if i == 0 {
			// skip header
			continue
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		csvBase := app.CSVPokemon{
			ID:             line[0],
			Identifier:     line[1],
			SpeciesID:      line[2],
			Height:         line[3],
			Weight:         line[4],
			BaseExperience: line[5],
			Order:          line[6],
			IsDefault:      line[7],
		}

		breedList[csvBase.ID] = pokemon.BreedDetail{
			Summary: &pokemon.BreedSummary{
				Id:   csvBase.ID,
				Name: csvBase.Identifier,
			},
			Stats: &pokemon.StatGroup{},
		}
	}

	return nil
}

func applyCSVStat(breedList map[string]pokemon.BreedDetail) error {
	reader, err := getCSVReader("./data/source/pokemon_stat.csv")
	if err != nil {
		return err
	}

	for i := 0; i < 55; i++ {
		line, err := reader.Read()
		if i == 0 {
			// skip header
			continue
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		stat := app.CSVPokemonStat{
			PokemonID: line[0],
			StatID:    line[1],
			BaseStat:  line[2],
			Effort:    line[3],
		}

		valueStr, err := strconv.Atoi(stat.BaseStat)
		if err != nil {
			return err
		}

		value := int32(valueStr)

		if stat.StatID == "1" {
			breedList[stat.PokemonID].Stats.Hp = value
		} else if stat.StatID == "2" {
			breedList[stat.PokemonID].Stats.Attack = value
		} else if stat.StatID == "3" {
			breedList[stat.PokemonID].Stats.Defense = value
		} else if stat.StatID == "4" {
			breedList[stat.PokemonID].Stats.SpecialAttack = value
		} else if stat.StatID == "5" {
			breedList[stat.PokemonID].Stats.SpecialDefense = value
		} else if stat.StatID == "6" {
			breedList[stat.PokemonID].Stats.Speed = value
		} else {
			return fmt.Errorf("Unknown stat index (%s)", stat.StatID)
		}
	}
	return nil
}

func getCSVReader(path string) (*csv.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return csv.NewReader(bufio.NewReader(file)), nil
}

// func (bs *BreedService) load() error {
// 	return bs.DB.Update(func(tx *bolt.Tx) error {
// 		b, err := tx.CreateBucketIfNotExists(bucket)
// 		if err != nil {
// 			return err
// 		}

// 		b1, err := encode(pokemon.BreedDetail{
// 			Summary: &pokemon.BreedSummary{
// 				Id:   "1",
// 				Name: "Bulbasaur",
// 			},
// 			Stats: &pokemon.StatGroup{
// 				Hp:             45,
// 				Attack:         49,
// 				Defense:        49,
// 				SpecialAttack:  65,
// 				SpecialDefense: 65,
// 				Speed:          45,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		b2, err := encode(pokemon.BreedDetail{
// 			Summary: &pokemon.BreedSummary{
// 				Id:   "2",
// 				Name: "Ivysaur",
// 			},
// 			Stats: &pokemon.StatGroup{
// 				Hp:             60,
// 				Attack:         62,
// 				Defense:        63,
// 				SpecialAttack:  80,
// 				SpecialDefense: 80,
// 				Speed:          60,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		b3, err := encode(pokemon.BreedDetail{
// 			Summary: &pokemon.BreedSummary{
// 				Id:   "25",
// 				Name: "Pikachu",
// 			},
// 			Stats: &pokemon.StatGroup{
// 				Hp:             35,
// 				Attack:         55,
// 				Defense:        30,
// 				SpecialAttack:  50,
// 				SpecialDefense: 40,
// 				Speed:          90,
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		err = b.Put([]byte("1"), b1)
// 		if err != nil {
// 			return err
// 		}
// 		err = b.Put([]byte("2"), b2)
// 		if err != nil {
// 			return err
// 		}
// 		err = b.Put([]byte("25"), b3)
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	})
// }
