package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/ShadamHarizky/dating-app/domain/entity"
	"github.com/ShadamHarizky/dating-app/domain/repository"
	"github.com/ShadamHarizky/dating-app/utils"
)

type UserService struct {
	us repository.UserRepository
	ss SwipeService
}

func NewUserService(us repository.UserRepository, ss SwipeService) (*UserService, error) {
	return &UserService{
		us: us,
		ss: ss,
	}, nil
}

// UserService implements the UserServiceInterface

type UserServiceInterface interface {
	SaveUser(*entity.User) (*entity.User, map[string]string)
	GetUsers(uint64) ([]entity.User, error)
	GetUser(uint64) (*entity.User, error)
	GetUserByEmailAndPassword(*entity.User) (*entity.User, map[string]string)
	PurchasePremium(uint64) (*entity.User, error)
}

func (u *UserService) SaveUser(user *entity.User) (*entity.User, map[string]string) {
	return u.us.SaveUser(user)
}

func (u *UserService) GetUser(userId uint64) (*entity.User, error) {
	return u.us.GetUser(userId)
}

func (u *UserService) GetUsers(userId uint64) ([]entity.User, error) {
	excludeUser := []string{}
	user, err := u.us.GetUser(userId)
	if err != nil {
		return []entity.User{}, err
	}

	rightSwipe, err := u.ss.GetSwipedProfiles(context.Background(), int(userId), "right")
	if err != nil {
		return []entity.User{}, err
	}

	leftSwipe, err := u.ss.GetSwipedProfiles(context.Background(), int(userId), "left")
	if err != nil {
		return []entity.User{}, err
	}

	excludeUser = append(excludeUser, rightSwipe...)
	excludeUser = append(excludeUser, leftSwipe...)

	if user.Premium {
		return u.us.GetUsers(utils.StrSliceToInt(excludeUser))
	} else {
		countRightSwipe, err := u.ss.GetSwipeCount(int(userId), "right")
		if err != nil {
			return []entity.User{}, err
		}

		countLeftSwipe, err := u.ss.GetSwipeCount(int(userId), "left")
		if err != nil {
			return []entity.User{}, err
		}

		total := countRightSwipe + countLeftSwipe

		fmt.Println("total", total)

		if total < 10 {
			return u.us.GetUsers(utils.StrSliceToInt(excludeUser))
		}

		return []entity.User{}, errors.New("your account already reached limit")

	}
}

func (u *UserService) GetUserByEmailAndPassword(user *entity.User) (*entity.User, map[string]string) {
	return u.us.GetUserByEmailAndPassword(user)
}

func (u *UserService) PurchasePremium(userId uint64) (*entity.User, error) {
	return u.us.PurchasePremium(userId)
}
