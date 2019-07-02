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
	err := loadCSVBase(breedList)

	if err != nil {
		return err
	}

	err = applyCSVStat(breedList)
	err = applyCSVTypes(breedList)

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

func loadCSVBase(breedList map[string]pokemon.BreedDetail) error {
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

// todo: map effort values
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

func applyCSVTypes(breedList map[string]pokemon.BreedDetail) error {
	reader, err := getCSVReader("./data/source/pokemon_type.csv")
	if err != nil {
		return err
	}

	for i := 0; i < 14; i++ {
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
		csvT := app.CSVPokemonType{
			PokemonID: line[0],
			TypeID:    line[1],
			Slot:      line[2],
		}

		t := getCSVKeyToTypeMap()[csvT.TypeID]

		if csvT.Slot == "1" {
			breedList[csvT.PokemonID].Summary.Type_1 = t
		} else {
			breedList[csvT.PokemonID].Summary.Type_2 = t
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

func getCSVKeyToTypeMap() map[string]pokemon.Type {
	return map[string]pokemon.Type{
		"1":  pokemon.Type_NORMAL,
		"2":  pokemon.Type_FIGHTING,
		"3":  pokemon.Type_FLYING,
		"4":  pokemon.Type_POISON,
		"5":  pokemon.Type_GROUND,
		"6":  pokemon.Type_ROCK,
		"7":  pokemon.Type_BUG,
		"8":  pokemon.Type_GHOST,
		"9":  pokemon.Type_STEEL,
		"10": pokemon.Type_FIRE,
		"11": pokemon.Type_WATER,
		"12": pokemon.Type_GRASS,
		"13": pokemon.Type_ELECTRIC,
		"14": pokemon.Type_PSYCHIC,
		"15": pokemon.Type_ICE,
		"16": pokemon.Type_DRAGON,
		"17": pokemon.Type_DARK,
		"18": pokemon.Type_FAIRY,
	}
}
