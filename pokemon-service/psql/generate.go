package psql

import (
	"context"
	"fmt"
	"math"
	"math/rand"

	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"github.com/Mrcampbell/pgo2/shared-library/uuid"
)

func (ps *PokemonService) buildPokemon(ctx context.Context, breed pokemon.BreedDetail, p pokemon.InternalCreatePokemonRequest, vg pokemon.VersionGroup) (pokemon.Pokemon, error) {

	var iHP, iAtk, iDef, iSpecAtk, iSpecDef, iSpeed int

	if p.OverrideStatGroup != nil && p.OverrideStatGroup.Hp != 0 {
		iHP = int(p.OverrideStatGroup.Hp)
	} else {
		iHP = rand.Intn(14) + 1
	}

	if p.OverrideStatGroup != nil && p.OverrideStatGroup.Attack != 0 {
		iAtk = int(p.OverrideStatGroup.Attack)
	} else {
		iAtk = rand.Intn(14) + 1
	}

	if p.OverrideStatGroup != nil && p.OverrideStatGroup.Defense != 0 {
		iDef = int(p.OverrideStatGroup.Defense)
	} else {
		iDef = rand.Intn(14) + 1
	}

	if p.OverrideStatGroup != nil && p.OverrideStatGroup.SpecialAttack != 0 {
		iSpecAtk = int(p.OverrideStatGroup.SpecialAttack)
	} else {
		iSpecDef = rand.Intn(14) + 1
	}

	if p.OverrideStatGroup != nil && p.OverrideStatGroup.SpecialDefense != 0 {
		iSpecDef = int(p.OverrideStatGroup.SpecialDefense)
	} else {
		iSpecDef = rand.Intn(14) + 1
	}

	if p.OverrideStatGroup != nil && p.OverrideStatGroup.Speed != 0 {
		iSpeed = int(p.OverrideStatGroup.Speed)
	} else {
		iSpeed = rand.Intn(14) + 1
	}

	hp := calcStat(int(p.Level), int(breed.Stats.Hp), iHP, 0)
	atk := calcStat(int(p.Level), int(breed.Stats.Attack), iAtk, 0)
	def := calcStat(int(p.Level), int(breed.Stats.Defense), iDef, 0)
	specAtk := calcStat(int(p.Level), int(breed.Stats.SpecialAttack), iSpecAtk, 0)
	specDef := calcStat(int(p.Level), int(breed.Stats.SpecialDefense), iSpecDef, 0)
	speed := calcStat(int(p.Level), int(breed.Stats.Speed), iSpeed, 0)

	moves, err := ps.breedMoveClient.GetRandomMoveSetForBreed(ctx, &pokemon.GetRandomMoveSetForBreedRequest{
		BreedId:        breed.Summary.Id,
		Level:          p.Level,
		VersionGroupId: vg,
	})
	if err != nil {
		fmt.Println(err)
		return pokemon.Pokemon{}, err
	}

	return pokemon.Pokemon{
		Id:           uuid.PrefixedUUID("p"),
		BreedId:      breed.Summary.Id,
		BreedSummary: breed.Summary,
		MoveOne:      moves.MoveOne,
		MoveOneId:    moves.MoveOne.Id,
		MoveTwo:      moves.MoveTwo,
		MoveTwoId:    moves.MoveTwo.Id,
		MoveThree:    moves.MoveThree,
		MoveThreeId:  moves.MoveThree.Id,
		MoveFour:     moves.MoveFour,
		MoveFourId:   moves.MoveFour.Id,

		Ev: &pokemon.StatGroup{
			Hp:             int32(1),
			Attack:         int32(1),
			Defense:        int32(1),
			SpecialAttack:  int32(1),
			SpecialDefense: int32(1),
			Speed:          int32(1),
		},
		Iv: &pokemon.StatGroup{
			Hp:             int32(iHP),
			Attack:         int32(iAtk),
			Defense:        int32(iDef),
			SpecialAttack:  int32(iSpecAtk),
			SpecialDefense: int32(iSpecDef),
			Speed:          int32(iSpeed),
		},
		Stat: &pokemon.StatGroup{
			Hp:             int32(hp),
			Attack:         int32(atk),
			Defense:        int32(def),
			SpecialAttack:  int32(specAtk),
			SpecialDefense: int32(specDef),
			Speed:          int32(speed),
		},
	}, nil
}

func calcStat(level, base, iv, ev int) int {
	a := int(math.Floor(float64(ev / 4)))
	b := ((2 * base) + iv + a) * level
	c := math.Floor(float64(b/100)) + 5

	// todo: implement nature
	// return c * nature when done

	return int(c)
}

func calcHP(level, base, iv, ev int) int {
	return calcStat(level, base, iv, ev) + 5 + level
}
