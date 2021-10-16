package mysql

import (
	"wizeline/model"
	"wizeline/repository"
)

type repo struct{

}

func (r *repo) SaveUsersToDataSource(users []model.User) error {
	panic("implement me")
}

func (r *repo) GetUsers() (string, error) {
	panic("implement me")
}

func (r *repo) GetUser(userID int) (model.User, error) {
	panic("implement me")
}

func (r *repo) SaveUsers(data [][]string) (int, error) {
	panic("implement me")
}

func (r *repo) GetUsersFromDataSource(data [][]string) ([]model.User, error) {
	panic("implement me")
}

func NewMysqlRepository() repository.UserRepository {
	return &repo{}
}

