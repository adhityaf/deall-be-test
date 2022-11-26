package services

import (
	"fmt"
	"net/http"

	"github.com/adhityaf/deall-be-test/helpers"
	"github.com/adhityaf/deall-be-test/models"
	"github.com/adhityaf/deall-be-test/params"
	"github.com/adhityaf/deall-be-test/repositories"
	"github.com/gin-gonic/gin"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: repo,
	}
}

func (u *UserService) Login(request params.LoginUser) *params.Response {
	user, err := u.userRepository.FindByUsername(request.Username)
	if err != nil {
		return &params.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error": "Not Found",
				"message": fmt.Sprintf("User %s", err.Error()),
			},
		}
	}

	ok := helpers.ComparePassword([]byte(user.Password), []byte(request.Password))
	if !ok {
		return &params.Response{
			Status: http.StatusUnauthorized,
			Payload: gin.H{
				"error": "Unauthorized",
				"message": "Password not match",
			},
		}
	}

	token := helpers.GenerateToken(user.UserId, user.Name, user.Role)

	return &params.Response{
		Status: http.StatusOK,
		Payload: gin.H{
			"message": "Login successful",
			"token":   token,
		},
	}
}

func (u *UserService) CreateUser(request params.CreateUser) *params.Response {
	password := helpers.HashPassword(request.Password)
	user := models.User{
		Username: request.Username,
		Password: password,
		Name:     request.Name,
		Role:     request.Role,
	}

	userData, err := u.userRepository.Create(&user)
	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": "Bad Request",
				"message": err.Error(),
			},
		}
	}

	return &params.Response{
		Status:  http.StatusCreated,
		Payload: userData,
	}
}

func (u *UserService) UpdateUser(request params.UpdateUser) *params.Response {
	user, err := u.userRepository.FindById(request.UserId)
	if err != nil {
		return &params.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error": "Not Found",
				"message": err.Error(),
			},
		}
	}

	user.Name = request.Name // update name
	if request.Role != ""{ // if role is requested to be update
		user.Role = request.Role
	}

	user, err = u.userRepository.Update(user)
	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"error": "Bad Request",
				"message": err.Error(),
			},
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: gin.H{
			"message": fmt.Sprintf("Success update data with id : %d", user.UserId),
		},
	}
}

func (u *UserService) DeleteUser(userId int) *params.Response {
	user, err := u.userRepository.FindById(userId)
	if err != nil {
		return &params.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error": "Not Found",
				"message": err.Error(),
			},
		}
	}

	user, err = u.userRepository.Delete(user)
	if err != nil {
		return &params.Response{
			Status: http.StatusBadRequest,
			Payload: gin.H{
				"message": err.Error(),
			},
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: gin.H{
			"message": fmt.Sprintf("Success delete data with id : %d", user.UserId),
		},
	}
}

func (u *UserService) GetUserById(userId int) *params.Response {
	user, err := u.userRepository.FindById(userId)
	if err != nil {
		return &params.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error": "Not Found",
				"message": err.Error(),
			},
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: user,
	}
}

func (u *UserService) GetAllUsers() *params.Response {
	users, err := u.userRepository.FindAllUsers()
	if err != nil {
		return &params.Response{
			Status: http.StatusNotFound,
			Payload: gin.H{
				"error": "Not Found",
				"message": err.Error(),
			},
		}
	}

	if len(*users) == 0 {
		return &params.Response{
			Status:  http.StatusOK,
			Payload: "Data Users is Empty",
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Payload: users,
	}
}
