package psql

import (
	"context"
	"fmt"

	"github.com/Mrcampbell/pgo2/battle-service/config"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type BattleService struct {
	DB *pg.DB
}

func NewBattleService() (*BattleService, error) {

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

	createSchema(db)

	return &BattleService{
		DB: db,
	}, nil
}

func (bs *BattleService) Upsert(ctx context.Context, battle *pokemon.Battle) error {

	err := bs.DB.Update(battle)
	if err != nil {
		err = bs.DB.Insert(battle)
		if err != nil {
			return err
		}
	}
	return nil

	return nil
}

func (bs *BattleService) GetBattle(ctx context.Context, id string) (*pokemon.Battle, error) {
	b := &pokemon.Battle{Id: id}
	err := bs.DB.Select(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func createSchema(db *pg.DB) error {

	for _, model := range []interface{}{(*pokemon.Battle)(nil)} {

		// todo: remove after debug
		// db.DropTable(model, &orm.DropTableOptions{
		// 	Cascade:  true,
		// 	IfExists: true,
		// })
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
