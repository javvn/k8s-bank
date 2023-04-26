package repository

import (
	"jawncorp.com/bank/user/domain/model"
	"jawncorp.com/bank/user/interface/database"
	"net/http"
)

type UserRepositoryInterface interface {
	FindAll() (users model.Users, err error)
	Find(query interface{}, args ...interface{}) (users model.Users, err error)
	FindByID(query interface{}, args ...interface{}) (user model.User, err error)
	Create(input model.User) (out model.Response, err error)
	Update(value map[string]interface{}, query interface{}, args ...interface{}) (user model.User, err error)
}

type UserRepository struct {
	database.SQLHandler
}

func (repo *UserRepository) FindAll() (users model.Users, err error) {
	repo.SQLHandler.Scan(&users.Data, "id desc")
	return
}

func (repo *UserRepository) Find(query interface{}, args ...interface{}) (users model.Users, err error) {
	repo.SQLHandler.Where(&users.Data, query, args...)
	return
}

func (repo *UserRepository) FindByID(query interface{}, args ...interface{}) (user model.User, err error) {
	repo.SQLHandler.Where(&user, query, args...)
	return
}
func (repo *UserRepository) Create(input model.User) (out model.Response, err error) {
	_, err = repo.SQLHandler.Create(&input)

	if err != nil {
		return model.Response{
			Code:    http.StatusBadRequest,
			Message: "Create error",
		}, err
	}

	return model.Response{
		Code:    http.StatusOK,
		Message: "OK",
	}, nil

}

func (repo *UserRepository) Update(value map[string]interface{}, query interface{}, args ...interface{}) (user model.User, err error) {
	// NOTE: When update with struct, GORM will only update non-zero fields,
	// you might want to use map to update attributes or use Select to specify fields to update
	repo.SQLHandler.Update(&user, value, query, args...)

	return
}
