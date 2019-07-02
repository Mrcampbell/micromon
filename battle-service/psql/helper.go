package psql

import "github.com/Mrcampbell/pgo2/protorepo/pokemon"

func mapPokemonToBattleMask(p pokemon.Pokemon) pokemon.PokemonBattleMask {
	return pokemon.PokemonBattleMask{
		Pokemon:     &p,
		BattleStats: &pokemon.BattleStatGroup{},
	}
}
