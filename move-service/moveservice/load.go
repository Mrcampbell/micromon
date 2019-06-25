package moveservice

import "github.com/Mrcampbell/pgo2/protorepo/pokemon"

func (ms *MoveService) load() {
	ms.set["1"] = pokemon.MoveDetail{
		Id:   "1",
		Name: "Scratch",
	}
	ms.set["2"] = pokemon.MoveDetail{
		Id:   "2",
		Name: "Pound",
	}
}
