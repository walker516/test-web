package handler

import (
	"backend/domain/user"
	"backend/internal/usecase"
	"backend/pkg/errmsg"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: u}
}

func (h *UserHandler) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errmsg.RespondError(c, errmsg.NewError("ERR_BAD_REQUEST", "Invalid id format"))
	}

	user, err := h.userUsecase.GetByID(id)
	if err != nil {
		return errmsg.RespondError(c, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetAll(c echo.Context) error {
	users, err := h.userUsecase.GetAll()
	if err != nil {
		return errmsg.RespondError(c, err)
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) Create(c echo.Context) error {
	var req user.User
	if err := c.Bind(&req); err != nil {
		return errmsg.RespondError(c, errmsg.NewError("ERR_BAD_REQUEST", "Invalid request format"))
	}


	fmt.Printf("Received request body: %+v\n", req)

	if err := h.userUsecase.Create(&req); err != nil {
		return errmsg.RespondError(c, err)
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "User created"})
}

func (h *UserHandler) Update(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errmsg.RespondError(c, errmsg.NewError("ERR_BAD_REQUEST", "Invalid id format"))
	}

	var req user.User
	if err := c.Bind(&req); err != nil {
		return errmsg.RespondError(c, errmsg.NewError("ERR_BAD_REQUEST", "Invalid request format"))
	}
	req.ID = id

	if err := h.userUsecase.Update(&req); err != nil {
		return errmsg.RespondError(c, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User updated"})
}

func (h *UserHandler) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errmsg.RespondError(c, errmsg.NewError("ERR_BAD_REQUEST", "Invalid id format"))
	}

	if err := h.userUsecase.Delete(id); err != nil {
		return errmsg.RespondError(c, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "User deleted"})
}
