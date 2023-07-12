package postgres

import (
	"context"
	"fmt"

	"github.com/procode2/structio/models"
	"github.com/uptrace/bun"
)

func (p *PostgresStore) CreateNewUser(user *models.User) (*models.User, error) {
	fmt.Printf("Creating user with creds %+v\n", user)
	ctx := context.Background()

	_, err := p.db.
		NewInsert().
		Model(user).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *PostgresStore) GetUserById(userId string) (*models.User, error) {
	ctx := context.Background()

	user := &models.User{}
	err := p.db.NewSelect().Model(user).Where("? = ?", bun.Ident("id"), userId).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (p *PostgresStore) GetUserByEmail(email string) (*models.User, error) {
	ctx := context.Background()

	user := &models.User{}
	err := p.db.NewSelect().Model(user).Where("? = ?", bun.Ident("email"), email).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
