package psql

type BreedMoveService struct{}

func NewBreedMoveService() (*BreedMoveService, error) {
	return &BreedMoveService{}, nil
}

// GetAllPreviousLevelUpLearnableMoveIDsForBreed gets ALL the possible moves for a pokemon to have learned in the wild.
//  Useful for when generating wild pokemon and you need a random moveset.
func (ps *BreedMoveService) GetAllPreviousLevelUpLearnableMoveIDsForBreed(breedID string, generation, level int) ([]int, error) {
	return nil, nil
}

// GetRandomMoveIDSetForBreed Calls GetAllPreviousLevelUpLearnableMoveIDsForBreed but only returns (up to) 4 id's
func (ps *BreedMoveService) GetRandomMoveIDSetForBreed(breedID string, generation, level int) ([]int, error) {
	return nil, nil
}
