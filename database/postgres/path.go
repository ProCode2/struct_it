package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/procode2/structio/models"
	"github.com/uptrace/bun"
)

func (p *PostgresStore) GetAllPaths(search string) ([]*models.Path, error) {
	paths := make([]*models.Path, 0)
	ctx := context.Background()
	q := p.db.NewSelect().Model(&paths)
	if search != "" {
		q = q.Where("? ILIKE ?", bun.Ident("title"), "%"+search+"%")
	}
	err := q.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return paths, nil
}

func (p *PostgresStore) CreateNewPath(path *models.Path) (*models.Path, error) {
	ctx := context.Background()
	err := p.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {

		_, err := tx.NewInsert().Model(path).Exec(ctx)
		if err != nil {
			fmt.Println(err)
			return err
		}

		for _, level := range path.Levels {
			level.PathId = path.ID
		}

		_, err = tx.NewInsert().Model(&path.Levels).Exec(ctx)
		if err != nil {
			fmt.Println(err)
			return err
		}

		for _, level := range path.Levels {

			for _, bit := range level.Bits {
				bit.LevelId = level.ID
			}
			_, err = tx.NewInsert().Model(&level.Bits).Exec(ctx)
		}

		return err
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return path, nil
}

func (p *PostgresStore) UpdatePath(path *models.Path) error {
	// TODO
	return nil
}

func (p *PostgresStore) GetPathById(pathId string) (*models.Path, error) {
	path := &models.Path{}

	ctx := context.Background()
	err := p.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {

		err := tx.NewSelect().Model(path).Where("? = ?", bun.Ident("id"), pathId).Relation("Levels").Scan(ctx)
		if err != nil {
			return err
		}

		for _, level := range path.Levels {
			err = tx.NewSelect().Model(&level.Bits).Where("? = ?", bun.Ident("level_id"), level.ID).Scan(ctx)
		}
		return err
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return path, nil
}

func (p *PostgresStore) DeletePathById(pathId string) error {
	// TODO
	return nil
}
