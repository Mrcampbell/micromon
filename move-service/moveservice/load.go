package moveservice

import "github.com/Mrcampbell/pgo2/protorepo/pokemon"

func (ms *MoveService) load() {
	ms.set["0"] = pokemon.MoveDetail{
		Summary: &pokemon.MoveSummary{
			Id:   "0",
			Name: "NONE",
		},
	}
	ms.set["1"] = pokemon.MoveDetail{
		Summary: &pokemon.MoveSummary{
			Id:   "1",
			Name: "Scratch",
		},
	}
	ms.set["2"] = pokemon.MoveDetail{
		Summary: &pokemon.MoveSummary{
			Id:   "2",
			Name: "Pound",
		},
	}
	ms.set["33"] = pokemon.MoveDetail{
		Summary: &pokemon.MoveSummary{
			Id:   "33",
			Name: "Tackle",
		},
	}
}
