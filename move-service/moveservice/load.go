package moveservice

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/Mrcampbell/pgo2/move-service/app"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
)

func (ms *MoveService) load() {
	loadCSVBase(ms.set)
}

func loadCSVBase(moveList map[string]pokemon.MoveDetail) error {
	reader, err := getCSVReader("./data/source/moves.csv")
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
		csvBase := app.CSVMoveBase{
			ID:                   line[0],
			Identifier:           line[1],
			GenerationID:         line[2],
			TypeID:               line[3],
			Power:                line[4],
			PP:                   line[5],
			Accuracy:             line[6],
			Priority:             line[7],
			TargetID:             line[8],
			DamageClassID:        line[9],
			EffectID:             line[10],
			EffectChance:         line[11],
			ContestTypeID:        line[12],
			ContestEffectID:      line[13],
			SuperContestEffectID: line[14],
		}

		// j, _ := json.MarshalIndent(csvBase, " ", " ")
		// fmt.Println(string(j))
		var accuracy, power, pp, priority int

		accuracy, err = strconv.Atoi(csvBase.Accuracy)
		power, err = strconv.Atoi(csvBase.Power)
		pp, err = strconv.Atoi(csvBase.PP)
		priority, err = strconv.Atoi(csvBase.Priority)
		if err != nil {
			fmt.Println(err)
		}

		damageClass := getCSVKeyToDamageClassMap()[csvBase.DamageClassID]

		moveList[csvBase.ID] = pokemon.MoveDetail{
			Summary: &pokemon.MoveSummary{
				Id:          csvBase.ID,
				DamageClass: damageClass,
				Name:        csvBase.Identifier,
				Type:        getCSVKeyToTypeMap()[csvBase.TypeID],
			},
			Accuracy: int32(accuracy),
			// Effect: // TODO: Add effect,
			Power:    int32(power),
			Pp:       int32(pp),
			Priority: int32(priority),
			// Target: csvBase.TargetID // TODO: Add Target
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

func getCSVKeyToDamageClassMap() map[string]pokemon.DamageClass {
	return map[string]pokemon.DamageClass{
		"1": pokemon.DamageClass_STATUS,
		"2": pokemon.DamageClass_PHYSICAL,
		"3": pokemon.DamageClass_SPECIAL,
	}
}
