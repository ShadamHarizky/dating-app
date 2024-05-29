package interfaces

import (
	"net/http"

	"github.com/ShadamHarizky/dating-app/domain/entity"
	"github.com/ShadamHarizky/dating-app/infrastructure/auth"
	"github.com/ShadamHarizky/dating-app/service"
	"github.com/gin-gonic/gin"
)

// Users struct defines the dependencies that will be used
type Swipes struct {
	ss service.SwipeServiceInterface
	rd auth.AuthInterface
	tk auth.TokenInterface
}

// Users constructor
func NewSwipes(ss service.SwipeServiceInterface, rd auth.AuthInterface, tk auth.TokenInterface) *Swipes {
	return &Swipes{
		ss: ss,
		rd: rd,
		tk: tk,
	}
}

func (s *Swipes) Swipes(c *gin.Context) {
	var swipes entity.Swipes
	if err := c.ShouldBindJSON(&swipes); err != nil {
		c.JSON(http.StatusUnprocessableEntity, entity.ResponseDefaultJson{
			Status:  http.StatusUnprocessableEntity,
			Message: "Invalid json format!",
		})
		return
	}

	metadata, err := s.tk.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, entity.ResponseDefaultJson{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	err = s.ss.RecordSwipe(int(metadata.UserId), int(swipes.ProfileId), swipes.Direction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.ResponseDefaultJson{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, entity.ResponseDefaultJson{
		Status:  http.StatusOK,
		Message: "success swipe",
	})
}
