package handlers

import (
	"github.com/labstack/echo/v4"
	"jawncorp.com/bank/user/domain/model"
	"jawncorp.com/bank/user/domain/repository"
	"jawncorp.com/bank/user/interface/database"
	"jawncorp.com/bank/user/usecase"
	"jawncorp.com/bank/user/utils"
	"net/http"
	"time"
)

type UserHandler struct {
	Interactor usecase.UserInteractor
}

func NewUserHandler(sqlHandler database.SQLHandler) *UserHandler {
	return &UserHandler{
		Interactor: usecase.UserInteractor{
			UserRepository: &repository.UserRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

func (handler *UserHandler) GetUsers() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		leave := c.QueryParam("leave")

		resJSON, err := handler.Interactor.GetUsers(leave)

		if err != nil {
			return utils.GetErrorMessage(c, "ko", err)
		}
		return c.JSON(http.StatusOK, resJSON)
	}
}

func (handler *UserHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		resJSON, err := handler.Interactor.GetUser(c.Param("id"))

		if err != nil {
			return utils.GetErrorMessage(c, "ko", err)
		}

		return c.JSON(http.StatusOK, resJSON)
	}
}

func (handler *UserHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		i := new(model.User)

		if err = c.Bind(i); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		input := model.User{
			Name:      i.Name,
			Gender:    i.Gender,
			Phone:     i.Phone,
			Email:     i.Email,
			Birthday:  i.Birthday,
			Leave:     0,
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		if input.Name == "" {
			return c.JSON(http.StatusBadRequest, model.Response{
				Message: "No name param found",
			})
		}

		if input.Phone == "" {
			return c.JSON(http.StatusBadRequest, model.Response{
				Message: "No phone param found",
			})
		}

		if input.Email == "" {
			return c.JSON(http.StatusBadRequest, model.Response{
				Message: "No email param found",
			})
		}

		resJSON, err := handler.Interactor.CreateItem(input)

		if err != nil {
			return utils.GetErrorMessage(c, "ko", err)
		}

		return c.JSON(http.StatusOK, resJSON)
	}
}

func (handler *UserHandler) UpdateLeaveAttr() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		input := new(model.User)

		if err = c.Bind(input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		resJSON, err := handler.Interactor.UpdateLeaveAttr(*input)

		if err != nil {
			return utils.GetErrorMessage(c, "ko", err)
		}

		return c.JSON(http.StatusOK, resJSON)
	}
}
