package database

import "github.com/procode2/structio/models"

type Storer interface {
	Init()

	CreateNewUser(user *models.User) (*models.User, error)

	GetUserById(userId string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)

	GetAllPaths(search string) ([]*models.Path, error)
	CreateNewPath(path *models.Path) (*models.Path, error)
	UpdatePath(path *models.Path) error
	GetPathById(pathId string) (*models.Path, error)
	DeletePathById(pathid string) error
}
