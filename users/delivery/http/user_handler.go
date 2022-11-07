package http

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pastorilps/echo-swagger/middleware"
	"github.com/pastorilps/echo-swagger/users/domain"
	"github.com/pastorilps/echo-swagger/users/entity"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Response struct {
	Message string `json:"message"`
}

type UserHandler struct {
	AUsecase domain.UserUseCase
}

func NewUserHandler(e *echo.Echo, uc domain.UserUseCase) {
	handler := &UserHandler{
		AUsecase: uc,
	}

	// Routes
	e.GET("/", HealthCheck)
	e.GET("/v1/users", handler.GetAllUsers)
	e.GET("/v1/users/:id", handler.GetUserById)
	e.POST("/v1/users/create", handler.CreateUser)
	e.PUT("/v1/users/update/:id", handler.UpdateUser)
	e.DELETE("/v1/users/delete/:id", handler.DeleteUser)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

}

func (u *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	err = u.AUsecase.DeleteUser(int16(id))
	if err != nil {
		return c.JSON(getStatusCode(err), Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, id)
}

func (u *UserHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, Response{Message: err.Error()})
	}

	var receive entity.Receive_User
	err = c.Bind(&receive)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	receive.ID = int16(id)

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	list, err := u.AUsecase.UpdateUser(ctx, &receive)
	if err != nil {
		return c.JSON(getStatusCode(err), Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, list)
}

func (u *UserHandler) CreateUser(c echo.Context) error {
	l := new(entity.Users)
	if err := c.Bind(l); err != nil {
		return err
	}

	_, err := u.AUsecase.CreateUsers(l)
	if err != nil {
		return c.JSON(getStatusCode(err), Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, l)
}

func (u *UserHandler) GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	uId, err := u.AUsecase.GetUserById(int16(id))
	if err != nil {
		return c.JSON(getStatusCode(err), Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, uId)
}

func (u *UserHandler) GetAllUsers(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	data, err := u.AUsecase.GetAllUsers()
	if err != nil {
		return c.JSON(getStatusCode(err), Response{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Fatalln(err)

	switch err {
	case middleware.ErrInternalServerError:
		return http.StatusInternalServerError
	case middleware.ErrorNotFound:
		return http.StatusNotFound
	case middleware.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": "Server is up and running",
	})
}
