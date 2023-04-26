package usecase

import (
	"jawncorp.com/bank/user/domain/model"
	"jawncorp.com/bank/user/domain/repository"
	"jawncorp.com/bank/user/utils"
)

type UserInteractor struct {
	UserRepository repository.UserRepositoryInterface
}

func (interactor *UserInteractor) GetUsers(leave string) (users model.Users, err error) {
	var query string
	var args interface{}

	if leave == "true" {
		query = "leave = ?"
		args = true
	} else if leave == "false" {
		query = "leave = ?"
		args = false
	} else {
		query = ""
		args = ""
	}

	users, err = interactor.UserRepository.Find(query, args)

	if err != nil {
		err = utils.SetErrorMessage("10001E")
		return
	}

	return
}

func (interactor *UserInteractor) GetUser(id string) (users model.User, err error) {
	query := "id = ?"
	args := id

	users, err = interactor.UserRepository.FindByID(query, args)

	if err != nil {
		err = utils.SetErrorMessage("10001E")
		return
	}

	return
}

func (interactor *UserInteractor) CreateItem(input model.User) (response model.Response, err error) {
	response, err = interactor.UserRepository.Create(input)

	if err != nil {
		err = utils.SetErrorMessage("10001E")
		return
	}

	return
}

func (interactor *UserInteractor) UpdateLeaveAttr(input model.User) (user model.User, err error) {
	user, err = interactor.UserRepository.Update(map[string]interface{}{"Leave": input.Leave}, "id = ?", input.ID)

	if err != nil {
		err = utils.SetErrorMessage("10001E")
		return
	}

	return
}
