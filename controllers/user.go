package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/adhityaf/deall-be-test/params"
	"github.com/adhityaf/deall-be-test/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: *userService,
	}
}

func (u *UserController) Login(ctx *gin.Context) {
	var req params.LoginUser
	validate := validator.New()

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	err = validate.Struct(req)
	if err != nil {
		validationMessage := ""
		for _, err := range err.(validator.ValidationErrors) {
			validationMessage = fmt.Sprintf("%s field %s %s. ", validationMessage, err.Field(), err.Tag())
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: validationMessage,
		})

		return
	}

	result := u.userService.Login(req)

	ctx.JSON(result.Status, result.Payload)
}

func (u *UserController) Register(ctx *gin.Context) {
	var req params.CreateUser
	validate := validator.New()

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	// Request only username, password and name
	// Using register endpoint to make role user
	req.Role = "user"

	err = validate.Struct(req)
	if err != nil {
		validationMessage := ""
		for _, err := range err.(validator.ValidationErrors) {
			validationMessage = fmt.Sprintf("%s field %s %s. ", validationMessage, err.Field(), err.Tag())
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: validationMessage,
		})

		return
	}

	result := u.userService.CreateUser(req)

	ctx.JSON(result.Status, result.Payload)
}

func (u *UserController) GetUserProfile(ctx *gin.Context) {
	// Get user_id from token
	userId, err := strconv.Atoi(ctx.GetString("user_id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	result := u.userService.GetUserById(userId)

	ctx.JSON(result.Status, result.Payload)
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	var req params.CreateUser
	validate := validator.New()

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	err = validate.Struct(req)
	if err != nil {
		validationMessage := ""
		for _, err := range err.(validator.ValidationErrors) {
			validationMessage = fmt.Sprintf("%s field %s %s. ", validationMessage, err.Field(), err.Tag())
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: validationMessage,
		})

		return
	}

	result := u.userService.CreateUser(req)

	ctx.JSON(result.Status, result.Payload)
}

func (u *UserController) GetAllUsers(ctx *gin.Context) {
	result := u.userService.GetAllUsers()

	ctx.JSON(result.Status, result.Payload)
}

func (u *UserController) GetUserById(ctx *gin.Context) {
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	result := u.userService.GetUserById(userId)

	ctx.JSON(result.Status, result.Payload)
}

func (u *UserController) UpdateUserById(ctx *gin.Context) {
	var req params.UpdateUser
	validate := validator.New()

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	// Get user_id from param
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	// Request from json only name and role
	req.UserId = userId

	err = validate.Struct(req)
	if err != nil {
		validationMessage := ""
		for _, err := range err.(validator.ValidationErrors) {
			validationMessage = fmt.Sprintf("%s field %s %s. ", validationMessage, err.Field(), err.Tag())
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: validationMessage,
		})

		return
	}

	result := u.userService.UpdateUser(req)

	ctx.JSON(result.Status, result.Payload)
}

func (u *UserController) DeleteUserById(ctx *gin.Context) {
	// Get user_id from param
	userId, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	result := u.userService.DeleteUser(userId)

	ctx.JSON(result.Status, result.Payload)
}
