package psql

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/Mrcampbell/pgo2/breed-move-service/config"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"github.com/go-pg/pg"
)

type BreedMoveService struct {
	DB         *pg.DB
	moveClient pokemon.MoveServiceClient
}

func NewBreedMoveService(moveClient pokemon.MoveServiceClient) (*BreedMoveService, error) {
	db := pg.Connect(&pg.Options{
		Addr:     config.DBHost + ":" + config.DBUser,
		Database: config.DBDatabase,
		User:     config.DBUser,
		Password: config.DBPassword,
		OnConnect: func(conn *pg.Conn) error {
			fmt.Println("Connected to DB")
			return nil
		},
	})

	// no need to initialize or create schema because it should exist.
	// It's a huge process right now

	return &BreedMoveService{
		DB:         db,
		moveClient: moveClient,
	}, nil
}

// GetRandomMoveIDSetForBreed Calls GetAllPreviousLevelUpLearnableMoveIDsForBreed but only returns (up to) 4 id's
func (ps *BreedMoveService) GetRandomMoveSummarySetForBreed(ctx context.Context, breedID string, vg pokemon.VersionGroup, level int) ([](*pokemon.MoveSummary), error) {

	bms, err := ps.getAllPreviousLevelUpLearnableMoveIDsForBreed(ctx, breedID, vg, level)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("BMS: ", bms)

	// shuffle them
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(bms), func(i, j int) {
		bms[i], bms[j] = bms[j], bms[i]
	})

	if len(bms) < 2 {
		bms = append(bms, &pokemon.BreedMove{
			MoveId:         "0",
			LearnMethodId:  pokemon.LearnMethod_LEVEL_UP,
			BreedId:        breedID,
			Level:          0,
			VersionGroupId: pokemon.VersionGroup_UNKNOWN_VERSION_GROUP,
		})
	}

	if len(bms) < 3 {
		bms = append(bms, &pokemon.BreedMove{
			MoveId:         "0",
			LearnMethodId:  pokemon.LearnMethod_LEVEL_UP,
			BreedId:        breedID,
			Level:          0,
			VersionGroupId: pokemon.VersionGroup_UNKNOWN_VERSION_GROUP,
		})
	}
	if len(bms) < 4 {
		bms = append(bms, &pokemon.BreedMove{
			MoveId:         "0",
			LearnMethodId:  pokemon.LearnMethod_LEVEL_UP,
			BreedId:        breedID,
			Level:          0,
			VersionGroupId: pokemon.VersionGroup_UNKNOWN_VERSION_GROUP,
		})
	}

	if len(bms) < 5 {
		bms = append(bms, &pokemon.BreedMove{
			MoveId:         "0",
			LearnMethodId:  pokemon.LearnMethod_LEVEL_UP,
			BreedId:        breedID,
			Level:          0,
			VersionGroupId: pokemon.VersionGroup_UNKNOWN_VERSION_GROUP,
		})
	}

	// TODO: remove api call for unknown move
	var summarys []*pokemon.MoveSummary
	for _, bm := range bms {
		ms, err := ps.moveClient.GetMoveSummary(ctx, &pokemon.GetMoveSummaryRequest{
			Id: bm.MoveId,
		})
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		summarys = append(summarys, ms.Move)
	}

	return summarys, nil
}

// GetAllPreviousLevelUpLearnableMoveIDsForBreed gets ALL the possible moves for a pokemon to have learned in the wild.
//  Useful for when generating wild pokemon and you need a random moveset.
func (ps *BreedMoveService) getAllPreviousLevelUpLearnableMoveIDsForBreed(ctx context.Context, breedID string, vg pokemon.VersionGroup, level int) ([]*pokemon.BreedMove, error) {

	moves, err := ps.GetMovesForBreed(ctx, breedID, vg)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// get all level up learnable moves
	levelUpLearnableMoves := getLevelUpLearnableMoves(moves, level, vg)

	return levelUpLearnableMoves, nil
}

func (ps *BreedMoveService) GetMovesForBreed(ctx context.Context, breedID string, version_group_id pokemon.VersionGroup) ([]*pokemon.BreedMove, error) {
	var bms []*pokemon.BreedMove

	vgInt32 := pokemon.VersionGroup_value[version_group_id.String()]
	vgStr := strconv.Itoa(int(vgInt32))
	err := ps.DB.ModelContext(ctx, &bms).
		Where("breed_id = ?", breedID).
		Where("version_group_id = ?", vgStr).
		Select()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return bms, nil
}

func getLevelUpLearnableMoves(bm []*pokemon.BreedMove, level int, vg pokemon.VersionGroup) []*pokemon.BreedMove {
	return Filter(bm, func(bm *pokemon.BreedMove) bool {
		return bm.VersionGroupId == vg && int(bm.Level) <= level && bm.LearnMethodId == pokemon.LearnMethod_LEVEL_UP
	})
}
