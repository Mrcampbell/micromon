package psql

import "github.com/Mrcampbell/pgo2/protorepo/pokemon"

func Filter(bm []*pokemon.BreedMove, f func(*pokemon.BreedMove) bool) []*pokemon.BreedMove {
	bmf := make([]*pokemon.BreedMove, 0)
	for _, m := range bm {
		if f(m) {
			bmf = append(bmf, m)
		}
	}
	return bmf
}
