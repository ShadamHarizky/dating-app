package interfaces

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ShadamHarizky/dating-app/domain/entity"
	"github.com/ShadamHarizky/dating-app/infrastructure/auth"
	"github.com/ShadamHarizky/dating-app/service"
	"github.com/gin-gonic/gin"
)

// Users struct defines the dependencies that will be used
type Users struct {
	us service.UserServiceInterface
	rd auth.AuthInterface
	tk auth.TokenInterface
}

// Users constructor
func NewUsers(us service.UserServiceInterface, rd auth.AuthInterface, tk auth.TokenInterface) *Users {
	return &Users{
		us: us,
		rd: rd,
		tk: tk,
	}
}

func (s *Users) SaveUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, entity.ResponseDefaultJson{
			Status:  http.StatusUnprocessableEntity,
			Message: "Invalid json format!",
		})
		return
	}
	//validate the request:
	validateErr := user.Validate("")
	if len(validateErr) > 0 {
		c.JSON(http.StatusUnprocessableEntity, validateErr)
		return
	}

	newUser, err := s.us.SaveUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ResponseDefaultJson{
			Status:  http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		})
		return
	}
	c.JSON(http.StatusCreated, newUser.PublicUser())
}

func (s *Users) GetUsers(c *gin.Context) {
	users := entity.Users{} //customize user
	var err error

	metadata, err := s.tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, entity.ResponseDefaultJson{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	users, err = s.us.GetUsers(metadata.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ResponseDefaultJson{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, users.PublicUsers())
}

func (s *Users) GetUser(c *gin.Context) {
	_, err := s.tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, entity.ResponseDefaultJson{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, entity.ResponseDefaultJson{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	user, err := s.us.GetUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ResponseDefaultJson{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user.PublicUser())
}

func (s *Users) PurchasePremium(c *gin.Context) {
	metadata, err := s.tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, entity.ResponseDefaultJson{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	user, err := s.us.PurchasePremium(metadata.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ResponseDefaultJson{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, entity.ResponseDefaultJson{
		Status:  http.StatusOK,
		Message: fmt.Sprint("Congratulation ", user.FirstName, " Your account already premium!"),
	})
}
