package psql

type PokemonMoveService struct{}

// GetAllPreviousLevelUpLearnableMoveIDsForBreed gets ALL the possible moves for a pokemon to have learned in the wild.
//  Useful for when generating wild pokemon and you need a random moveset.
func (ps *PokemonMoveService) GetAllPreviousLevelUpLearnableMoveIDsForBreed(breedID string, generation, level int) ([]int, error) {
	return nil, nil
}

// GetRandomMoveIDSetForBreed Calls GetAllPreviousLevelUpLearnableMoveIDsForBreed but only returns (up to) 4 id's
func (ps *PokemonMoveService) GetRandomMoveIDSetForBreed(breedID string, generation, level int) ([]int, error) {
	return nil, nil
}
